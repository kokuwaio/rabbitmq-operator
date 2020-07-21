package controllers

import (
	"context"
	"time"

	rabbitmqv1beta1 "github.com/kokuwaio/rabbitmq-operator/api/v1beta1"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Context("Inside of a new namespace", func() {
	ctx := context.TODO()
	ns := SetupTest(ctx)

	Describe("when no user exists", func() {
		It("should create a new user", func() {

			rabbitHost := rabbitConfig.Url
			rabbitUser := rabbitConfig.User
			password := rabbitConfig.Password

			rabbitClusterName := "test-cluster"
			secretName := "rabbit-secret"
			passwordKey := "password"
			rabbitUserName := "test-user"
			passwordSecretName := "test-user-secret"

			secret := &corev1.Secret{

				ObjectMeta: metav1.ObjectMeta{
					Name:      secretName,
					Namespace: ns.Name,
				},
				StringData: map[string]string{passwordKey: password},
				Type:       "Opaque",
			}
			Expect(k8sClient.Create(context.Background(), secret)).Should(Succeed())

			toCreate := &rabbitmqv1beta1.RabbitmqCluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      rabbitClusterName,
					Namespace: ns.Name,
				},
				Spec: rabbitmqv1beta1.RabbitmqClusterSpec{
					Host: rabbitHost,
					User: rabbitUser,
					SecretRef: rabbitmqv1beta1.SecretRef{
						Name:      secretName,
						Namespace: ns.Name,
						Key:       passwordKey,
					},
				},
			}

			Expect(k8sClient.Create(context.Background(), toCreate)).Should(Succeed())
			time.Sleep(time.Second * 5)

			passwordSecret := &corev1.Secret{

				ObjectMeta: metav1.ObjectMeta{
					Name:      passwordSecretName,
					Namespace: ns.Name,
				},
				StringData: map[string]string{"password": "dummy"},
				Type:       "Opaque",
			}
			Expect(k8sClient.Create(context.Background(), passwordSecret)).Should(Succeed())

			queue := &rabbitmqv1beta1.RabbitmqUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      rabbitUserName,
					Namespace: ns.Name,
				},
				Spec: rabbitmqv1beta1.RabbitmqUserSpec{
					Name: rabbitUserName,
					Tags: "user",
					PasswordSecretRef: rabbitmqv1beta1.PasswordSecretRef{
						Name:      passwordSecretName,
						Namespace: ns.Name,
					},
					ClusterRef: rabbitmqv1beta1.RabbitmqClusterRef{
						Name:      rabbitClusterName,
						Namespace: ns.Name,
					},
				},
			}

			Expect(k8sClient.Create(context.Background(), queue)).Should(Succeed())
			time.Sleep(time.Second * 5)

			client, err := rabbithole.NewClient(rabbitHost, rabbitUser, password)
			Expect(err).NotTo(HaveOccurred())

			rq, err := client.GetUser(rabbitUserName)
			Expect(err).NotTo(HaveOccurred())
			Expect(rq).ShouldNot(BeNil())
		})
	})
})
