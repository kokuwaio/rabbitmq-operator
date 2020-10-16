package controllers

import (
	"context"

	rabbitmqv1beta1 "github.com/kokuwaio/rabbitmq-operator/api/v1beta1"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service struct {
}

func (s *Service) GetRabbitClient(cluster *rabbitmqv1beta1.RabbitmqCluster, reader client.Reader) (*rabbithole.Client, error) {
	ctx := context.Background()
	secret := &corev1.Secret{}
	err := reader.Get(ctx, types.NamespacedName{Name: cluster.Spec.SecretRef.Name, Namespace: cluster.Spec.SecretRef.Namespace}, secret)
	if err != nil {
		return nil, err
	}
	bytes := secret.Data[cluster.Spec.SecretRef.PasswordKey]
	user := cluster.Spec.User
	if cluster.Spec.SecretRef.UserKey != "" {
		user = string(secret.Data[cluster.Spec.SecretRef.UserKey])
	}
	return rabbithole.NewClient(cluster.Spec.Host, user, string(bytes))
}
