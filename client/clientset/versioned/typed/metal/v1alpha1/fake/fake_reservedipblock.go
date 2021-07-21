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

	v1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/metal/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReservedIPBlocks implements ReservedIPBlockInterface
type FakeReservedIPBlocks struct {
	Fake *FakeMetalV1alpha1
	ns   string
}

var reservedipblocksResource = schema.GroupVersionResource{Group: "metal.equinixmetal.kubeform.com", Version: "v1alpha1", Resource: "reservedipblocks"}

var reservedipblocksKind = schema.GroupVersionKind{Group: "metal.equinixmetal.kubeform.com", Version: "v1alpha1", Kind: "ReservedIPBlock"}

// Get takes name of the reservedIPBlock, and returns the corresponding reservedIPBlock object, and an error if there is any.
func (c *FakeReservedIPBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReservedIPBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reservedipblocksResource, c.ns, name), &v1alpha1.ReservedIPBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReservedIPBlock), err
}

// List takes label and field selectors, and returns the list of ReservedIPBlocks that match those selectors.
func (c *FakeReservedIPBlocks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReservedIPBlockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reservedipblocksResource, reservedipblocksKind, c.ns, opts), &v1alpha1.ReservedIPBlockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReservedIPBlockList{ListMeta: obj.(*v1alpha1.ReservedIPBlockList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReservedIPBlockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reservedIPBlocks.
func (c *FakeReservedIPBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reservedipblocksResource, c.ns, opts))

}

// Create takes the representation of a reservedIPBlock and creates it.  Returns the server's representation of the reservedIPBlock, and an error, if there is any.
func (c *FakeReservedIPBlocks) Create(ctx context.Context, reservedIPBlock *v1alpha1.ReservedIPBlock, opts v1.CreateOptions) (result *v1alpha1.ReservedIPBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reservedipblocksResource, c.ns, reservedIPBlock), &v1alpha1.ReservedIPBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReservedIPBlock), err
}

// Update takes the representation of a reservedIPBlock and updates it. Returns the server's representation of the reservedIPBlock, and an error, if there is any.
func (c *FakeReservedIPBlocks) Update(ctx context.Context, reservedIPBlock *v1alpha1.ReservedIPBlock, opts v1.UpdateOptions) (result *v1alpha1.ReservedIPBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reservedipblocksResource, c.ns, reservedIPBlock), &v1alpha1.ReservedIPBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReservedIPBlock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReservedIPBlocks) UpdateStatus(ctx context.Context, reservedIPBlock *v1alpha1.ReservedIPBlock, opts v1.UpdateOptions) (*v1alpha1.ReservedIPBlock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(reservedipblocksResource, "status", c.ns, reservedIPBlock), &v1alpha1.ReservedIPBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReservedIPBlock), err
}

// Delete takes name of the reservedIPBlock and deletes it. Returns an error if one occurs.
func (c *FakeReservedIPBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reservedipblocksResource, c.ns, name), &v1alpha1.ReservedIPBlock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReservedIPBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reservedipblocksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReservedIPBlockList{})
	return err
}

// Patch applies the patch and returns the patched reservedIPBlock.
func (c *FakeReservedIPBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReservedIPBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reservedipblocksResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReservedIPBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReservedIPBlock), err
}
