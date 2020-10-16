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
	"k8s.io/apimachinery/pkg/types"
)

var _ = Context("Inside of a new namespace", func() {
	ctx := context.TODO()
	ns := SetupTest(ctx)

	Describe("when no user permission exists", func() {
		It("should create user permissions", func() {

			rabbitHost := rabbitConfig.Url
			rabbitUser := rabbitConfig.User
			password := rabbitConfig.Password

			rabbitClusterName := "test-cluster"
			secretName := "rabbit-secret"
			passwordKey := "password"
			rabbitUserName := "test-permission-user"
			passwordSecretName := "test-permission-user-secret"

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
						Name:        secretName,
						Namespace:   ns.Name,
						PasswordKey: passwordKey,
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

			user := &rabbitmqv1beta1.RabbitmqUser{
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
						Name: rabbitClusterName,
					},
				},
			}

			permissions := &rabbitmqv1beta1.RabbitmqPermisson{
				ObjectMeta: metav1.ObjectMeta{
					Name:      rabbitUserName,
					Namespace: ns.Name,
				},
				Spec: rabbitmqv1beta1.RabbitmqPermissonSpec{
					Vhost:     "/",
					UserName:  rabbitUserName,
					Configure: ".*",
					Write:     ".*",
					Read:      ".*",
					ClusterRef: rabbitmqv1beta1.RabbitmqClusterRef{
						Name: rabbitClusterName,
					},
				},
			}

			Expect(k8sClient.Create(context.Background(), user)).Should(Succeed())
			Expect(k8sClient.Create(context.Background(), permissions)).Should(Succeed())
			time.Sleep(time.Second * 10)

			client, err := rabbithole.NewClient(rabbitHost, rabbitUser, password)
			Expect(err).NotTo(HaveOccurred())

			urql, err := client.ListUsers()
			Expect(err).NotTo(HaveOccurred())
			Expect(urql).ShouldNot(BeNil())
			urq, err := client.GetUser(rabbitUserName)
			Expect(err).NotTo(HaveOccurred())
			Expect(urq).ShouldNot(BeNil())

			rq, err := client.GetPermissionsIn("/", rabbitUserName)
			Expect(err).NotTo(HaveOccurred())
			Expect(rq).ShouldNot(BeNil())
		})

		It("should create no permissions with wrong settings", func() {

			rabbitHost := rabbitConfig.Url
			rabbitUser := rabbitConfig.User
			password := rabbitConfig.Password

			rabbitClusterName := "test-cluster"
			secretName := "rabbit-secret"
			passwordKey := "password"
			rabbitUserName := "test-permission-user-bad"
			passwordSecretName := "test-permission-user-secret"

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
						Name:        secretName,
						Namespace:   ns.Name,
						PasswordKey: passwordKey,
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

			user := &rabbitmqv1beta1.RabbitmqUser{
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
						Name: rabbitClusterName,
					},
				},
			}

			permissions := &rabbitmqv1beta1.RabbitmqPermisson{
				ObjectMeta: metav1.ObjectMeta{
					Name:      rabbitUserName,
					Namespace: ns.Name,
				},
				Spec: rabbitmqv1beta1.RabbitmqPermissonSpec{
					Vhost:    "/",
					UserName: rabbitUserName,
					// wrong regex
					Configure: "*",
					Write:     ".*",
					Read:      ".*",
					ClusterRef: rabbitmqv1beta1.RabbitmqClusterRef{
						Name: rabbitClusterName,
					},
				},
			}

			Expect(k8sClient.Create(context.Background(), user)).Should(Succeed())
			Expect(k8sClient.Create(context.Background(), permissions)).Should(Succeed())
			time.Sleep(time.Second * 10)

			client, err := rabbithole.NewClient(rabbitHost, rabbitUser, password)
			Expect(err).NotTo(HaveOccurred())

			urql, err := client.ListUsers()
			Expect(err).NotTo(HaveOccurred())
			Expect(urql).ShouldNot(BeNil())
			urq, err := client.GetUser(rabbitUserName)
			Expect(err).NotTo(HaveOccurred())
			Expect(urq).ShouldNot(BeNil())

			_, err = client.GetPermissionsIn("/", rabbitUserName)
			Expect(err).To(HaveOccurred())

			Expect(k8sClient.Get(context.Background(), types.NamespacedName{
				Namespace: ns.Name,
				Name:      rabbitUserName,
			}, permissions)).Should(Succeed())
			Expect(permissions).NotTo(BeNil())
			Expect(permissions.Status.Error).Should(Equal("Error 400 (Error 400 from RabbitMQ: json: cannot unmarshal array into Go struct field ErrorResponse.reason of type string): "))

		})
	})
})
