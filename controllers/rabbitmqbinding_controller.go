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

	"github.com/go-logr/logr"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	rabbitmqv1beta1 "github.com/kokuwaio/rabbitmq-operator/api/v1beta1"
)

// RabbitmqBindingReconciler reconciles a RabbitmqBinding object
type RabbitmqBindingReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Service *Service
}

// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqbindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqbindings/status,verbs=get;update;patch

func (r *RabbitmqBindingReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("rabbitmqbinding", req.NamespacedName)

	instance := &rabbitmqv1beta1.RabbitmqBinding{}

	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			r.Log.Info("rabbitmq binding deleted", "name", req.NamespacedName)
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
			//TODO
			bindings, err := rabbitClient.ListQueueBindings(instance.Spec.Vhost, instance.Spec.Destination)
			if err != nil {
				return reconcile.Result{}, err
			}
			if binding, ok := r.checkIfBindingExists(bindings, instance.Spec); ok {
				if _, err := rabbitClient.DeleteBinding(instance.Spec.Vhost, *binding); err != nil {
					// if fail to delete the external dependency here, return with error
					// so that it can be retried
					r.Log.Info("error deleting binding", "err", err)
					return reconcile.Result{}, err
				}
			} else {
				r.Log.Info("binding does not exists remove it anyway")
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

	bindings, err := rabbitClient.ListQueueBindings(instance.Spec.Vhost, instance.Spec.Destination)
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	if _, ok := r.checkIfBindingExists(bindings, instance.Spec); !ok {

		// declare a binding
		_, err = rabbitClient.DeclareBinding(instance.Spec.Vhost, r.transformSettings(instance.Spec))
		if err != nil {
			r.UpdateErrorState(ctx, instance, err)
			// Error reading the object - requeue the request.
			return reconcile.Result{}, err
		}
	}

	instance.Status.Status = "Success"
	instance.Status.Error = ""
	err = r.Update(ctx, instance)
	return ctrl.Result{}, nil
}

func (r *RabbitmqBindingReconciler) transformSettings(settings rabbitmqv1beta1.RabbitmqBindingSpec) rabbithole.BindingInfo {
	result := make(map[string]interface{})
	if settings.Arguments != nil {
		for k, v := range settings.Arguments {
			result[k] = v
		}
	}

	return rabbithole.BindingInfo{
		Source:          settings.Source,
		Destination:     settings.Destination,
		DestinationType: settings.DestinationType,
		RoutingKey:      settings.RoutingKey,
		Arguments:       result,
		PropertiesKey:   settings.PropertiesKey,
	}
}

func (r *RabbitmqBindingReconciler) transformSettingsForDelete(settings rabbitmqv1beta1.RabbitmqBindingSpec) rabbithole.BindingInfo {

	return rabbithole.BindingInfo{
		Source:          settings.Source,
		Destination:     settings.Destination,
		DestinationType: settings.DestinationType,
		RoutingKey:      settings.RoutingKey,
		PropertiesKey:   "%23",
	}
}

func (r *RabbitmqBindingReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqBinding, err error) {
	instance.Status.Status = "Error"
	instance.Status.Error = err.Error()
	err = r.Update(context, instance)
}

func (r *RabbitmqBindingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqBinding{}).
		Complete(r)
}

func (r *RabbitmqBindingReconciler) checkIfBindingExists(bindings []rabbithole.BindingInfo, spec rabbitmqv1beta1.RabbitmqBindingSpec) (*rabbithole.BindingInfo, bool) {
	for _, binding := range bindings {
		if binding.Destination == spec.Destination && binding.Source == binding.Source && binding.RoutingKey == spec.RoutingKey {
			return &binding, true
		}
	}
	return nil, false
}
