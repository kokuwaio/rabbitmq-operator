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

// RabbitmqUserReconciler reconciles a RabbitmqUser object
type RabbitmqUserReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Service *Service
}

// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqusers/status,verbs=get;update;patch

func (r *RabbitmqUserReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("rabbitmquser", req.NamespacedName)

	instance := &rabbitmqv1beta1.RabbitmqUser{}

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

	// information about individual queue
	_, err = rabbitClient.GetUser(instance.Spec.Name)
	if err != nil {
		_, err = rabbitClient.PutUser(instance.Spec.Name, r.transformSettings(instance.Spec.Settings))
		if err != nil {
			r.UpdateErrorState(ctx, instance, err)
			return reconcile.Result{}, err
		}
	}
	instance.Status.Status = "Success"
	instance.Status.Error = ""
	err = r.Update(ctx, instance)
	return ctrl.Result{}, nil
}

func (r *RabbitmqUserReconciler) transformSettings(settings rabbitmqv1beta1.RabbitmqUserSetting) rabbithole.UserSettings {
	return rabbithole.UserSettings{
		Name:             "",
		Tags:             "",
		Password:         "",
		PasswordHash:     "",
		HashingAlgorithm: "",
	}
}

func (r *RabbitmqUserReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqUser, err error) {
	instance.Status.Status = "Error"
	instance.Status.Error = err.Error()
	err = r.Update(context, instance)
}

func (r *RabbitmqUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqUser{}).
		Complete(r)
}
