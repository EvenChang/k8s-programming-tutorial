package clientset

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ClientsetFun() {

	// Read kubernetes configurations and parse it.
	kubeconfig := flag.String("kubeconfig", "/Users/even/tmp/config", "kubeconfig file")
	flag.Parse()

	// According kubernetes configuration to build rest client config.
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// kubernetes.NewForConfig
	// In order to create the actual Kubernetes client set.
	// It’s called a client set because it contains multiple clients for all native Kubernetes resources.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Use the client-go library kubernetes clientset to get/create k8s resources.
	podList, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range podList.Items {
		fmt.Printf("Pod Name: %s, Namespace: %s\n", pod.ObjectMeta.Name, pod.ObjectMeta.Namespace)
	}
}
