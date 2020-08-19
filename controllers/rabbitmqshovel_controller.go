/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-logr/logr"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	rabbitmqv1beta1 "github.com/kokuwaio/rabbitmq-operator/api/v1beta1"
)

// RabbitmqShovelReconciler reconciles a RabbitmqShovel object
type RabbitmqShovelReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Service *Service
}

// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqshovels,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqshovels/status,verbs=get;update;patch

func (r *RabbitmqShovelReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("rabbitmqshovel", req.NamespacedName)

	instance := &rabbitmqv1beta1.RabbitmqShovel{}

	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			r.Log.Info("rabbitmq user deleted", "name", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	cluster := &rabbitmqv1beta1.RabbitmqCluster{}
	err = r.Get(ctx, types.NamespacedName{Name: instance.Spec.ClusterRef.Name}, cluster)
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	rabbitClient, err := r.Service.GetRabbitClient(cluster, r)
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object.
		if !containsString(instance.ObjectMeta.Finalizers, rabbitmqFinalizer) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, rabbitmqFinalizer)
			if err := r.Update(context.Background(), instance); err != nil {
				return reconcile.Result{Requeue: true}, nil
			}
		}
	} else {
		// The object is being deleted
		if containsString(instance.ObjectMeta.Finalizers, rabbitmqFinalizer) {
			// our finalizer is present, so lets handle our external dependency
			if _, err := rabbitClient.DeleteUser(instance.Spec.Name); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return reconcile.Result{}, err
			}

			// remove our finalizer from the list and update it.
			instance.ObjectMeta.Finalizers = removeString(instance.ObjectMeta.Finalizers, rabbitmqFinalizer)
			if err := r.Update(context.Background(), instance); err != nil {
				return reconcile.Result{Requeue: true}, nil
			}
		}

		// Our finalizer has finished, so the reconciler can do nothing.
		return reconcile.Result{}, nil
	}

	_, err = rabbitClient.GetShovel(instance.Spec.Vhost, instance.Spec.Name)
	if err != nil {
		res, err := rabbitClient.DeclareShovel(instance.Spec.Vhost, instance.Spec.Name, r.transformSettings(instance.Namespace, instance.Spec))
		if err != nil {
			r.UpdateErrorState(ctx, instance, err)
			return reconcile.Result{}, err
		}
		if res.StatusCode > 299 {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			err = fmt.Errorf("error updating queue: %s %s", res.Status, string(bodyBytes))
			r.UpdateErrorState(ctx, instance, err)
			return reconcile.Result{}, err
		}
	}

	instance.Status.Status = "Success"
	instance.Status.Error = ""
	err = r.Update(ctx, instance)
	return ctrl.Result{}, nil
}

func (r *RabbitmqShovelReconciler) transformSettings(namespace string, spec rabbitmqv1beta1.RabbitmqShovelSpec) rabbithole.ShovelDefinition {

	sourceURI := spec.SourceURI
	destinationURI := spec.DestinationURI
	if spec.SourcePasswordSecretRef.Name != "" {
		secret, _ := r.getPasswordSecret(namespace, &spec.SourcePasswordSecretRef)
		sourceURI = r.buildUrl(spec.SourceURI, spec.SourceUser, secret)
	}
	if spec.DestinationPasswordSecretRef.Name != "" {
		secret, _ := r.getPasswordSecret(namespace, &spec.DestinationPasswordSecretRef)
		destinationURI = r.buildUrl(spec.SourceURI, spec.SourceUser, secret)
	}

	return rabbithole.ShovelDefinition{
		AckMode:                          spec.AckMode,
		AddForwardHeaders:                spec.AddForwardHeaders,
		DeleteAfter:                      spec.DeleteAfter,
		DestinationAddForwardHeaders:     spec.DestinationAddForwardHeaders,
		DestinationAddTimestampHeader:    spec.DestinationAddTimestampHeader,
		DestinationAddress:               spec.DestinationAddress,
		DestinationApplicationProperties: "",
		DestinationExchange:              spec.DestinationExchange,
		DestinationExchangeKey:           spec.DestinationExchangeKey,
		DestinationProperties:            "",
		DestinationProtocol:              spec.DestinationProtocol,
		DestinationPublishProperties:     "",
		DestinationQueue:                 spec.DestinationQueue,
		DestinationURI:                   destinationURI,
		PrefetchCount:                    0,
		ReconnectDelay:                   0,
		SourceAddress:                    spec.SourceAddress,
		SourceDeleteAfter:                spec.SrcDeleteAfter,
		SourceExchange:                   spec.SourceExchange,
		SourceExchangeKey:                spec.SourceExchangeKey,
		SourcePrefetchCount:              0,
		SourceProtocol:                   spec.SourceProtocol,
		SourceQueue:                      spec.SourceQueue,
		SourceURI:                        sourceURI,
	}
}

func (r *RabbitmqShovelReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqShovel, err error) {
	instance.Status.Status = "Error"
	instance.Status.Error = err.Error()
	r.Log.Error(err, err.Error())
	err = r.Update(context, instance)
}

func (r *RabbitmqShovelReconciler) getPasswordSecret(namespace string, secretRef *rabbitmqv1beta1.PasswordSecretRef) (string, error) {
	ctx := context.Background()
	secret := &corev1.Secret{}
	ns := secretRef.Namespace
	if ns == "" {
		ns = namespace
	}
	err := r.Get(ctx, types.NamespacedName{Name: secretRef.Name, Namespace: ns}, secret)
	if err != nil {
		return "", err
	}
	var password string
	if secretRef.Key != "" {
		password = string(secret.Data[secretRef.Key])
	} else {
		password = string(secret.Data["password"])
	}
	return password, nil
}

func (r *RabbitmqShovelReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqShovel{}).
		Complete(r)
}

func (r *RabbitmqShovelReconciler) buildUrl(url, user, password string) string {
	split := strings.Split(url, "://")
	return fmt.Sprintf("%s://%s:%s@%s", split[0], user, password, split[1])
}
