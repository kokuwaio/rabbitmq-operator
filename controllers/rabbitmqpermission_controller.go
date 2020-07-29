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

// RabbitmqPermissionReconciler reconciles a RabbitmqPermisson object
type RabbitmqPermissionReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Service *Service
}

// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqpermissons,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rabbitmq.kokuwa.io,resources=rabbitmqpermissons/status,verbs=get;update;patch

func (r *RabbitmqPermissionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("rabbitmqpermisson", req.NamespacedName)

	instance := &rabbitmqv1beta1.RabbitmqPermisson{}

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

	permissions := rabbithole.Permissions{
		Configure: instance.Spec.Configure,
		Write:     instance.Spec.Write,
		Read:      instance.Spec.Read,
	}
	res, err := rabbitClient.UpdatePermissionsIn(instance.Spec.Vhost, instance.Spec.UserName, permissions)
	r.Log.Info(fmt.Sprintf("update permissions for user: %s", instance.Spec.UserName))
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		return reconcile.Result{}, err
	}
	if res.StatusCode > 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		err = fmt.Errorf("error updating permission: %s %s", res.Status, string(bodyBytes))
		r.UpdateErrorState(ctx, instance, err)
		return reconcile.Result{}, err
	}

	instance.Status.Status = "Success"
	instance.Status.Error = ""
	err = r.Update(ctx, instance)
	return ctrl.Result{}, nil
}

func (r *RabbitmqPermissionReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqPermisson, err error) {
	instance.Status.Status = "Error"
	if ferr, ok := err.(*rabbithole.ErrorResponse); ok {
		instance.Status.Error = ferr.Message
		r.Log.Error(err, ferr.Message)
	} else {
		instance.Status.Error = err.Error()
		r.Log.Error(err, err.Error())
	}
	err = r.Update(context, instance)
}

func (r *RabbitmqPermissionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqPermisson{}).
		Complete(r)
}
