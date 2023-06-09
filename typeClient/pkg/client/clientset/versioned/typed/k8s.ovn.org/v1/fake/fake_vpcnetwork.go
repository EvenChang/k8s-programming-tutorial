/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"
	v1 "k8s-programming-tutorial/typeClient/pkg/apis/k8s.ovn.org/v1"
	k8sovnorgv1 "k8s-programming-tutorial/typeClient/pkg/client/applyconfiguration/k8s.ovn.org/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVPCNetworks implements VPCNetworkInterface
type FakeVPCNetworks struct {
	Fake *FakeK8sV1
}

var vpcnetworksResource = v1.SchemeGroupVersion.WithResource("vpcnetworks")

var vpcnetworksKind = v1.SchemeGroupVersion.WithKind("VPCNetwork")

// Get takes name of the vPCNetwork, and returns the corresponding vPCNetwork object, and an error if there is any.
func (c *FakeVPCNetworks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VPCNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(vpcnetworksResource, name), &v1.VPCNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VPCNetwork), err
}

// List takes label and field selectors, and returns the list of VPCNetworks that match those selectors.
func (c *FakeVPCNetworks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VPCNetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(vpcnetworksResource, vpcnetworksKind, opts), &v1.VPCNetworkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.VPCNetworkList{ListMeta: obj.(*v1.VPCNetworkList).ListMeta}
	for _, item := range obj.(*v1.VPCNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vPCNetworks.
func (c *FakeVPCNetworks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(vpcnetworksResource, opts))
}

// Create takes the representation of a vPCNetwork and creates it.  Returns the server's representation of the vPCNetwork, and an error, if there is any.
func (c *FakeVPCNetworks) Create(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.CreateOptions) (result *v1.VPCNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(vpcnetworksResource, vPCNetwork), &v1.VPCNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VPCNetwork), err
}

// Update takes the representation of a vPCNetwork and updates it. Returns the server's representation of the vPCNetwork, and an error, if there is any.
func (c *FakeVPCNetworks) Update(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.UpdateOptions) (result *v1.VPCNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(vpcnetworksResource, vPCNetwork), &v1.VPCNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VPCNetwork), err
}

// Delete takes name of the vPCNetwork and deletes it. Returns an error if one occurs.
func (c *FakeVPCNetworks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(vpcnetworksResource, name, opts), &v1.VPCNetwork{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVPCNetworks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(vpcnetworksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.VPCNetworkList{})
	return err
}

// Patch applies the patch and returns the patched vPCNetwork.
func (c *FakeVPCNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VPCNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(vpcnetworksResource, name, pt, data, subresources...), &v1.VPCNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VPCNetwork), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied vPCNetwork.
func (c *FakeVPCNetworks) Apply(ctx context.Context, vPCNetwork *k8sovnorgv1.VPCNetworkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.VPCNetwork, err error) {
	if vPCNetwork == nil {
		return nil, fmt.Errorf("vPCNetwork provided to Apply must not be nil")
	}
	data, err := json.Marshal(vPCNetwork)
	if err != nil {
		return nil, err
	}
	name := vPCNetwork.Name
	if name == nil {
		return nil, fmt.Errorf("vPCNetwork.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(vpcnetworksResource, *name, types.ApplyPatchType, data), &v1.VPCNetwork{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VPCNetwork), err
}
