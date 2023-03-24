package main

import (
	// "k8s-programming-tutorial/clientset"
	// "k8s-programming-tutorial/dynamicClient"
	// "k8s-programming-tutorial/restClient"
	"fmt"
	"k8s-programming-tutorial/pkg/apis/k8s.ovn.org/v1alpha1"
)

func main() {

	// clientset.ClientsetFun()
	// dynamicClient.DynamicClientFunc()
	// restClient.RestClientFunc()
	k := v1alpha1.VPCNetwork{}
	fmt.Println(k)
}
