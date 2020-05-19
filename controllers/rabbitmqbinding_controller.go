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

	bindings, err := rabbitClient.ListExchangeBindingsBetween(instance.Spec.Vhost, instance.Spec.Source, instance.Spec.Destination)
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	if len(bindings) == 0 {
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
	return rabbithole.BindingInfo{
		Vhost:           settings.Vhost,
		Source:          settings.Source,
		Destination:     settings.Destination,
		DestinationType: settings.DestinationType,
		RoutingKey:      settings.RoutingKey,
		// TODO
		//Arguments: settings.Arguments,
		PropertiesKey: settings.PropertiesKey,
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
