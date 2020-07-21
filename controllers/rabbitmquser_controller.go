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
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"

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

	secret := &corev1.Secret{}
	err = r.Get(ctx, types.NamespacedName{Name: instance.Spec.PasswordSecretRef.Name, Namespace: instance.Spec.PasswordSecretRef.Namespace}, secret)
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		return reconcile.Result{}, err
	}
	var password string
	if instance.Spec.PasswordSecretRef.Key != "" {
		password = string(secret.Data[instance.Spec.PasswordSecretRef.Key])
	} else {
		password = string(secret.Data["password"])
	}
	r.Log.Info(fmt.Sprintf("create user with: %s %s", instance.Spec.Name, password))
	res, err := rabbitClient.PutUser(instance.Spec.Name, r.transformSettings(instance.Spec, password))
	if err != nil {
		r.UpdateErrorState(ctx, instance, err)
		return reconcile.Result{}, err
	}
	if res.StatusCode > 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		err = fmt.Errorf("error updating user: %s %s", res.Status, string(bodyBytes))
		r.UpdateErrorState(ctx, instance, err)
		return reconcile.Result{}, err
	}

	instance.Status.Status = "Success"
	instance.Status.Error = ""
	err = r.Update(ctx, instance)
	return ctrl.Result{}, nil
}

func (r *RabbitmqUserReconciler) transformSettings(spec rabbitmqv1beta1.RabbitmqUserSpec, password string) rabbithole.UserSettings {
	var salt = [4]byte{}
	salt, _ = generateSalt()
	hash := generateHashSha256(salt, password)

	hash = append(salt[:], []byte(hash[:])...)
	return rabbithole.UserSettings{
		Tags:             spec.Tags,
		PasswordHash:     base64.StdEncoding.EncodeToString(hash[:]),
		HashingAlgorithm: rabbithole.HashingAlgorithmSHA256,
	}
}

func (r *RabbitmqUserReconciler) UpdateErrorState(context context.Context, instance *rabbitmqv1beta1.RabbitmqUser, err error) {
	instance.Status.Status = "Error"
	instance.Status.Error = err.Error()
	r.Log.Error(err, err.Error())
	err = r.Update(context, instance)
}

func (r *RabbitmqUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rabbitmqv1beta1.RabbitmqUser{}).
		Complete(r)
}

func generateSalt() ([4]byte, error) {
	salt := [4]byte{}
	_, err := rand.Read(salt[:])
	salt = [4]byte{0, 0, 0, 0}
	return salt, err
}

func generateHashSha256(salt [4]byte, password string) []byte {
	temp_hash := sha256.Sum256(append(salt[:], []byte(password)...))
	return temp_hash[:]
}
