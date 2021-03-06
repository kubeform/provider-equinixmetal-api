/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/reserved/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIpBlocks implements IpBlockInterface
type FakeIpBlocks struct {
	Fake *FakeReservedV1alpha1
	ns   string
}

var ipblocksResource = schema.GroupVersionResource{Group: "reserved.equinixmetal.kubeform.com", Version: "v1alpha1", Resource: "ipblocks"}

var ipblocksKind = schema.GroupVersionKind{Group: "reserved.equinixmetal.kubeform.com", Version: "v1alpha1", Kind: "IpBlock"}

// Get takes name of the ipBlock, and returns the corresponding ipBlock object, and an error if there is any.
func (c *FakeIpBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IpBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipblocksResource, c.ns, name), &v1alpha1.IpBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpBlock), err
}

// List takes label and field selectors, and returns the list of IpBlocks that match those selectors.
func (c *FakeIpBlocks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IpBlockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipblocksResource, ipblocksKind, c.ns, opts), &v1alpha1.IpBlockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IpBlockList{ListMeta: obj.(*v1alpha1.IpBlockList).ListMeta}
	for _, item := range obj.(*v1alpha1.IpBlockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipBlocks.
func (c *FakeIpBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipblocksResource, c.ns, opts))

}

// Create takes the representation of a ipBlock and creates it.  Returns the server's representation of the ipBlock, and an error, if there is any.
func (c *FakeIpBlocks) Create(ctx context.Context, ipBlock *v1alpha1.IpBlock, opts v1.CreateOptions) (result *v1alpha1.IpBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipblocksResource, c.ns, ipBlock), &v1alpha1.IpBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpBlock), err
}

// Update takes the representation of a ipBlock and updates it. Returns the server's representation of the ipBlock, and an error, if there is any.
func (c *FakeIpBlocks) Update(ctx context.Context, ipBlock *v1alpha1.IpBlock, opts v1.UpdateOptions) (result *v1alpha1.IpBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipblocksResource, c.ns, ipBlock), &v1alpha1.IpBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpBlock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpBlocks) UpdateStatus(ctx context.Context, ipBlock *v1alpha1.IpBlock, opts v1.UpdateOptions) (*v1alpha1.IpBlock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipblocksResource, "status", c.ns, ipBlock), &v1alpha1.IpBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpBlock), err
}

// Delete takes name of the ipBlock and deletes it. Returns an error if one occurs.
func (c *FakeIpBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipblocksResource, c.ns, name), &v1alpha1.IpBlock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipblocksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IpBlockList{})
	return err
}

// Patch applies the patch and returns the patched ipBlock.
func (c *FakeIpBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IpBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipblocksResource, c.ns, name, pt, data, subresources...), &v1alpha1.IpBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpBlock), err
}
