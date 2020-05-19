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

// RabbitmqExchangeReconciler reconciles a RabbitmqExchange object
type RabbitmqExchangeReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Service *Service
}

// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqexchanges,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqexchanges/status,verbs=get;update;patch

func (r *RabbitmqExchangeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("rabbitmqexchange", req.NamespacedName)

	instance := &rabbitmqv1beta1.RabbitmqExchange{}

	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			r.Log.Info("rabbitmq exchange deleted", "name", req.NamespacedName)
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

	_, err = rabbitClient.GetExchange(instance.Spec.Vhost, instance.Spec.Name)
	if err != nil {
		_, err = rabbitClient.DeclareExchange(instance.Spec.Vhost, instance.Spec.Name, r.transformSettings(instance.Spec.Settings))
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

func (r *RabbitmqExchangeReconciler) transformSettings(settings rabbitmqv1beta1.RabbitmqExchangeSetting) rabbithole.ExchangeSettings {
	return rabbithole.ExchangeSettings{
		Type:       settings.Type,
		Durable:    settings.Durable,
		AutoDelete: settings.AutoDelete,
		// TODO transform data
		//Arguments:  settings.Arguments,
	}
}

func (r *RabbitmqExchangeReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqExchange, err error) {
	instance.Status.Status = "Error"
	instance.Status.Error = err.Error()
	err = r.Update(context, instance)
}

func (r *RabbitmqExchangeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqExchange{}).
		Complete(r)
}
