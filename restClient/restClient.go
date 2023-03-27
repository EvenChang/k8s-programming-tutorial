package restClient

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func RestClientFunc() {

	kubeconfig := flag.String("kubeconfig", "/Users/even/tmp/config", "kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// core api -> api , CRD api -> apis, page40
	config.APIPath = "apis"
	config.GroupVersion = &schema.GroupVersion{Group: "k8s.ovn.org", Version: "v1"}
	// NegotiatedSerializer is used for obtaining encoders and decoders for multiple
	// supported media types.
	config.NegotiatedSerializer = scheme.Codecs

	fmt.Println("Init RESTClient.")

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	AddToScheme(scheme.Scheme)

	vpcNetwork := &VPCNetwork{
		TypeMeta: metav1.TypeMeta{
			Kind:       "VPCNetwork",
			APIVersion: "k8s.ovn.org/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-vpc-network",
		},
		Spec: VPCNetworkSpec{
			CIDR:         "10.244.10.10/24",
			SubnetLength: 24,
		},
	}

	result := &VPCNetworkList{}

	//GET
	if err := restClient.
		Get().
		Resource("vpcnetworks").
		Do(context.TODO()).
		Into(result); err != nil {
		panic(err)
	}
	//DELETE
	if len(result.Items) > 0 {
		restClient.
			Delete().
			Resource("vpcnetworks").
			Name(vpcNetwork.Name).
			Do(context.TODO())
	}

	vpcNetworkResult := &VPCNetwork{}
	//POST
	if err := restClient.
		Post().
		Resource("vpcnetworks").
		Body(vpcNetwork).
		Do(context.Background()).
		Into(vpcNetworkResult); err != nil {
		panic(err)
	}
	fmt.Println("Print all vpcnetworks resources:")
	for _, item := range result.Items {
		fmt.Printf("Name: %s\n", item.Name)
	}

	fmt.Println("---Added new VPC---")
	fmt.Println("APIVersion: ", vpcNetworkResult.APIVersion)
	fmt.Println("Kind: ", vpcNetworkResult.Kind)
	fmt.Println("Name: ", vpcNetworkResult.Name)
}
