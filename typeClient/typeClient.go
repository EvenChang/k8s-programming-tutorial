package typeClient

import (
	"context"
	"flag"
	"fmt"

	ovnV1 "k8s-programming-tutorial/typeClient/pkg/apis/k8s.ovn.org/v1"
	client "k8s-programming-tutorial/typeClient/pkg/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func TypeClientFunc() {
	kubeconfig := flag.String("kubeconfig", "/Users/even/tmp/config", "kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := client.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Get list of VPC networks.
	vpcList, err := clientset.K8sV1().VPCNetworks().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Before Create VPC")
	for _, vpc := range vpcList.Items {
		fmt.Printf("VPC name : %s\n", vpc.Name)
	}

	vpcNetwork := &ovnV1.VPCNetwork{
		TypeMeta: metav1.TypeMeta{
			Kind:       "VPCNetwork",
			APIVersion: "k8s.ovn.org/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "typeclient",
		},
		Spec: ovnV1.VPCNetworkSpec{
			CIDR:         "10.244.10.10/24",
			SubnetLength: 24,
		},
	}

	fmt.Println("Creating VPC: typeclient")

	_, err = clientset.K8sV1().VPCNetworks().Create(context.Background(), vpcNetwork, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Get list of VPC networks.
	vpcList, err = clientset.K8sV1().VPCNetworks().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, vpc := range vpcList.Items {
		fmt.Printf("VPC name : %s\n", vpc.Name)
	}

	fmt.Println("After created VPC")
}
