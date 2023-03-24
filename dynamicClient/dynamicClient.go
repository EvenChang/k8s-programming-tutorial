package dynamicClient

import (
	"context"
	"flag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func DynamicClientFunc() {

	kubeconfig := flag.String("kubeconfig", "/Users/even/tmp/config", "kubeconfig file")
	flag.Parse()

	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	gvr := schema.GroupVersionResource{Group: "k8s.ovn.org", Version: "v1", Resource: "vpcnetworks"}

	crObj := &unstructured.Unstructured{}
	crObj.SetUnstructuredContent(map[string]interface{}{
		"kind":       "VPCNetwork",
		"apiVersion": "k8s.ovn.org/v1",
		"metadata": map[string]interface{}{
			"name": "mycr",
		},
		"spec": map[string]interface{}{
			"cidr":         "10.144.0.0/12",
			"subnetLength": 24,
		},
	})

	// create the custom resource using the dynamic client
	_, err = dynamicClient.Resource(gvr).Create(context.Background(), crObj, metav1.CreateOptions{})

	if err != nil {
		panic(err.Error())
	}
}
