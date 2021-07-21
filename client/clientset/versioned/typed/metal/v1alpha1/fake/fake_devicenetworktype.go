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

// FakeDeviceNetworkTypes implements DeviceNetworkTypeInterface
type FakeDeviceNetworkTypes struct {
	Fake *FakeMetalV1alpha1
	ns   string
}

var devicenetworktypesResource = schema.GroupVersionResource{Group: "metal.equinixmetal.kubeform.com", Version: "v1alpha1", Resource: "devicenetworktypes"}

var devicenetworktypesKind = schema.GroupVersionKind{Group: "metal.equinixmetal.kubeform.com", Version: "v1alpha1", Kind: "DeviceNetworkType"}

// Get takes name of the deviceNetworkType, and returns the corresponding deviceNetworkType object, and an error if there is any.
func (c *FakeDeviceNetworkTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeviceNetworkType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devicenetworktypesResource, c.ns, name), &v1alpha1.DeviceNetworkType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNetworkType), err
}

// List takes label and field selectors, and returns the list of DeviceNetworkTypes that match those selectors.
func (c *FakeDeviceNetworkTypes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeviceNetworkTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devicenetworktypesResource, devicenetworktypesKind, c.ns, opts), &v1alpha1.DeviceNetworkTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeviceNetworkTypeList{ListMeta: obj.(*v1alpha1.DeviceNetworkTypeList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeviceNetworkTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deviceNetworkTypes.
func (c *FakeDeviceNetworkTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devicenetworktypesResource, c.ns, opts))

}

// Create takes the representation of a deviceNetworkType and creates it.  Returns the server's representation of the deviceNetworkType, and an error, if there is any.
func (c *FakeDeviceNetworkTypes) Create(ctx context.Context, deviceNetworkType *v1alpha1.DeviceNetworkType, opts v1.CreateOptions) (result *v1alpha1.DeviceNetworkType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devicenetworktypesResource, c.ns, deviceNetworkType), &v1alpha1.DeviceNetworkType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNetworkType), err
}

// Update takes the representation of a deviceNetworkType and updates it. Returns the server's representation of the deviceNetworkType, and an error, if there is any.
func (c *FakeDeviceNetworkTypes) Update(ctx context.Context, deviceNetworkType *v1alpha1.DeviceNetworkType, opts v1.UpdateOptions) (result *v1alpha1.DeviceNetworkType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devicenetworktypesResource, c.ns, deviceNetworkType), &v1alpha1.DeviceNetworkType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNetworkType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeviceNetworkTypes) UpdateStatus(ctx context.Context, deviceNetworkType *v1alpha1.DeviceNetworkType, opts v1.UpdateOptions) (*v1alpha1.DeviceNetworkType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devicenetworktypesResource, "status", c.ns, deviceNetworkType), &v1alpha1.DeviceNetworkType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNetworkType), err
}

// Delete takes name of the deviceNetworkType and deletes it. Returns an error if one occurs.
func (c *FakeDeviceNetworkTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devicenetworktypesResource, c.ns, name), &v1alpha1.DeviceNetworkType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeviceNetworkTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devicenetworktypesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeviceNetworkTypeList{})
	return err
}

// Patch applies the patch and returns the patched deviceNetworkType.
func (c *FakeDeviceNetworkTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceNetworkType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devicenetworktypesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DeviceNetworkType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceNetworkType), err
}
