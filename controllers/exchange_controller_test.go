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

	Describe("when no exchange exists", func() {
		It("should create a new queue", func() {

			rabbitHost := rabbitConfig.Url
			rabbitUser := rabbitConfig.User
			password := rabbitConfig.Password

			rabbitClusterName := "test-cluster"
			secretName := "rabbit-secret"
			passwordKey := "password"
			rabbitExchangeName := "test-exchange"

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

			queue := &rabbitmqv1beta1.RabbitmqExchange{
				ObjectMeta: metav1.ObjectMeta{
					Name:      rabbitExchangeName,
					Namespace: ns.Name,
				},
				Spec: rabbitmqv1beta1.RabbitmqExchangeSpec{
					Vhost: "/",
					Name:  rabbitExchangeName,
					ClusterRef: rabbitmqv1beta1.RabbitmqClusterRef{
						Name:      rabbitClusterName,
						Namespace: ns.Name,
					},
					Settings: rabbitmqv1beta1.RabbitmqExchangeSetting{
						Type:       "",
						Durable:    false,
						AutoDelete: false,
						Arguments:  nil,
					},
				},
			}

			Expect(k8sClient.Create(context.Background(), queue)).Should(Succeed())
			time.Sleep(time.Second * 5)

			client, err := rabbithole.NewClient(rabbitHost, rabbitUser, password)
			Expect(err).NotTo(HaveOccurred())

			rq, err := client.GetExchange("/", rabbitExchangeName)
			Expect(err).NotTo(HaveOccurred())
			Expect(rq).ShouldNot(BeNil())
		})
	})
})
