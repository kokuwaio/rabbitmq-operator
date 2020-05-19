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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	rabbitmqv1beta1 "github.com/kokuwaio/rabbitmq-operator/api/v1beta1"
)

// RabbitmqClusterReconciler reconciles a RabbitmqCluster object
type RabbitmqClusterReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Service *Service
}

// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqclusters/status,verbs=get;update;patch

func (r *RabbitmqClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	context := context.Background()
	_ = r.Log.WithValues("rabbitmqcluster", req.NamespacedName)

	// your logic here
	instance := &rabbitmqv1beta1.RabbitmqCluster{}

	err := r.Get(context, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			r.Log.Info("rabbitmq cluster deleted", "name", req.NamespacedName)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	rabbitClient, err := r.Service.GetRabbitClient(instance, r)
	if err != nil {
		r.UpdateErrorState(context, instance, err)
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	_, err = rabbitClient.Overview()
	if err != nil {
		r.UpdateErrorState(context, instance, err)
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	instance.Status.Status = "Success"
	instance.Status.Error = ""
	err = r.Update(context, instance)
	return ctrl.Result{}, nil
}

func (r *RabbitmqClusterReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqCluster, err error) {
	instance.Status.Status = "Error"
	instance.Status.Error = err.Error()
	err = r.Update(context, instance)
}

func (r *RabbitmqClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqCluster{}).
		Complete(r)
}
