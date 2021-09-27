/*
Copyright 2021 The Hybridnet Authors.

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

	networkingv1 "github.com/alibaba/hybridnet/pkg/apis/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRemoteSubnets implements RemoteSubnetInterface
type FakeRemoteSubnets struct {
	Fake *FakeNetworkingV1
}

var remotesubnetsResource = schema.GroupVersionResource{Group: "networking.alibaba.com", Version: "v1", Resource: "remotesubnets"}

var remotesubnetsKind = schema.GroupVersionKind{Group: "networking.alibaba.com", Version: "v1", Kind: "RemoteSubnet"}

// Get takes name of the remoteSubnet, and returns the corresponding remoteSubnet object, and an error if there is any.
func (c *FakeRemoteSubnets) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.RemoteSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(remotesubnetsResource, name), &networkingv1.RemoteSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.RemoteSubnet), err
}

// List takes label and field selectors, and returns the list of RemoteSubnets that match those selectors.
func (c *FakeRemoteSubnets) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.RemoteSubnetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(remotesubnetsResource, remotesubnetsKind, opts), &networkingv1.RemoteSubnetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.RemoteSubnetList{ListMeta: obj.(*networkingv1.RemoteSubnetList).ListMeta}
	for _, item := range obj.(*networkingv1.RemoteSubnetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested remoteSubnets.
func (c *FakeRemoteSubnets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(remotesubnetsResource, opts))
}

// Create takes the representation of a remoteSubnet and creates it.  Returns the server's representation of the remoteSubnet, and an error, if there is any.
func (c *FakeRemoteSubnets) Create(ctx context.Context, remoteSubnet *networkingv1.RemoteSubnet, opts v1.CreateOptions) (result *networkingv1.RemoteSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(remotesubnetsResource, remoteSubnet), &networkingv1.RemoteSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.RemoteSubnet), err
}

// Update takes the representation of a remoteSubnet and updates it. Returns the server's representation of the remoteSubnet, and an error, if there is any.
func (c *FakeRemoteSubnets) Update(ctx context.Context, remoteSubnet *networkingv1.RemoteSubnet, opts v1.UpdateOptions) (result *networkingv1.RemoteSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(remotesubnetsResource, remoteSubnet), &networkingv1.RemoteSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.RemoteSubnet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRemoteSubnets) UpdateStatus(ctx context.Context, remoteSubnet *networkingv1.RemoteSubnet, opts v1.UpdateOptions) (*networkingv1.RemoteSubnet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(remotesubnetsResource, "status", remoteSubnet), &networkingv1.RemoteSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.RemoteSubnet), err
}

// Delete takes name of the remoteSubnet and deletes it. Returns an error if one occurs.
func (c *FakeRemoteSubnets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(remotesubnetsResource, name), &networkingv1.RemoteSubnet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRemoteSubnets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(remotesubnetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.RemoteSubnetList{})
	return err
}

// Patch applies the patch and returns the patched remoteSubnet.
func (c *FakeRemoteSubnets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.RemoteSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(remotesubnetsResource, name, pt, data, subresources...), &networkingv1.RemoteSubnet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.RemoteSubnet), err
}
