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
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	rabbitmqv1beta1 "github.com/kokuwaio/rabbitmq-operator/api/v1beta1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment
var rabbitC testcontainers.Container
var rabbitConfig RabbitConfig

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

type RabbitConfig struct {
	Url      string
	User     string
	Password string
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")

	if os.Getenv("TEST_EXTERNAL_KUBE") != "" {
		existing := true
		config, _ := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
		testEnv = &envtest.Environment{
			CRDDirectoryPaths:  []string{filepath.Join("..", "config", "crd", "bases")},
			UseExistingCluster: &existing,
			Config:             config,
		}
	} else {
		testEnv = &envtest.Environment{
			CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
		}
	}

	var err error
	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = rabbitmqv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = rabbitmqv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = rabbitmqv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = rabbitmqv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = rabbitmqv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = rabbitmqv1beta1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	close(done)
}, 60)

func SetupTest(ctx context.Context) *corev1.Namespace {
	var stopCh chan struct{}
	ns := &corev1.Namespace{}

	BeforeEach(func() {
		stopCh = make(chan struct{})
		*ns = corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: "testns-" + randStringRunes(5)},
		}

		req := testcontainers.ContainerRequest{
			Image:        "rabbitmq:3.8-management",
			ExposedPorts: []string{"15672/tcp"},
			Env:          map[string]string{"RABBITMQ_DEFAULT_USER": "admin", "RABBITMQ_DEFAULT_PASS": "password"},
			WaitingFor:   wait.ForListeningPort("15672"),
		}
		rabbitC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		Expect(err).NotTo(HaveOccurred())

		port, err := rabbitC.MappedPort(ctx, "15672")
		Expect(err).NotTo(HaveOccurred())
		host, err := rabbitC.Host(ctx)
		Expect(err).NotTo(HaveOccurred())

		k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
		Expect(err).ToNot(HaveOccurred())
		Expect(k8sClient).ToNot(BeNil())

		rabbitConfig = RabbitConfig{
			Url:      fmt.Sprintf("http://%s:%s", host, port.Port()),
			User:     "admin",
			Password: "password",
		}

		err = k8sClient.Create(ctx, ns)
		Expect(err).NotTo(HaveOccurred(), "failed to create test namespace")

		mgr, err := ctrl.NewManager(cfg, ctrl.Options{})
		Expect(err).NotTo(HaveOccurred(), "failed to create manager")

		clusterController := &RabbitmqClusterReconciler{
			Client: mgr.GetClient(),
			Log:    logf.Log,
		}
		err = clusterController.SetupWithManager(mgr)
		Expect(err).NotTo(HaveOccurred(), "failed to setup controller")

		queueController := &RabbitmqQueueReconciler{
			Client: mgr.GetClient(),
			Log:    logf.Log,
		}
		err = queueController.SetupWithManager(mgr)
		Expect(err).NotTo(HaveOccurred(), "failed to setup queue controller")

		exchangeController := &RabbitmqExchangeReconciler{
			Client: mgr.GetClient(),
			Log:    logf.Log,
		}
		err = exchangeController.SetupWithManager(mgr)
		Expect(err).NotTo(HaveOccurred(), "failed to setup exchange controller")

		userController := &RabbitmqUserReconciler{
			Client: mgr.GetClient(),
			Log:    logf.Log,
		}
		err = userController.SetupWithManager(mgr)
		Expect(err).NotTo(HaveOccurred(), "failed to setup user controller")

		permissonController := &RabbitmqPermissionReconciler{
			Client: mgr.GetClient(),
			Log:    logf.Log,
		}
		err = permissonController.SetupWithManager(mgr)
		Expect(err).NotTo(HaveOccurred(), "failed to setup permisson controller")

		go func() {
			err := mgr.Start(stopCh)
			Expect(err).NotTo(HaveOccurred(), "failed to start manager")
		}()
	})

	AfterEach(func() {
		close(stopCh)
		/*
			client, err := rabbithole.NewClient("", "", "")
			Expect(err).NotTo(HaveOccurred(), "failed to delete test namespace")
			bindings, err := client.ListBindings()
			for _, binding := range bindings {
				_, _ = client.DeleteBinding(binding.Vhost, binding)
			}
			queues, err := client.ListQueues()
			for _, queue := range queues {
				_, _ = client.DeleteQueue(queue.Vhost, queue.Name)
			}
			exchanges, err := client.ListExchanges()
			for _, exchange := range exchanges {
				_, _ = client.DeleteExchange(exchange.Vhost, exchange.Name)
			}
		*/
		cleanupCluster(k8sClient)
		err := k8sClient.Delete(ctx, ns)
		Expect(err).NotTo(HaveOccurred(), "failed to delete test namespace")
	})
	return ns
}

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	if rabbitC != nil {
		ctx := context.Background()
		err := rabbitC.Terminate(ctx)
		Expect(err).ToNot(HaveOccurred())
	}
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})

func init() {
	rand.Seed(time.Now().UnixNano())
}

func cleanupCluster(k8sClient client.Client) {
	clusters := &rabbitmqv1beta1.RabbitmqClusterList{}
	_ = k8sClient.List(context.Background(), clusters)
	for _, c := range clusters.Items {
		_ = k8sClient.Delete(context.Background(), &c)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
