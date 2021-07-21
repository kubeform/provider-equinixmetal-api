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

	v1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/virtual/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCircuits implements CircuitInterface
type FakeCircuits struct {
	Fake *FakeVirtualV1alpha1
	ns   string
}

var circuitsResource = schema.GroupVersionResource{Group: "virtual.equinixmetal.kubeform.com", Version: "v1alpha1", Resource: "circuits"}

var circuitsKind = schema.GroupVersionKind{Group: "virtual.equinixmetal.kubeform.com", Version: "v1alpha1", Kind: "Circuit"}

// Get takes name of the circuit, and returns the corresponding circuit object, and an error if there is any.
func (c *FakeCircuits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Circuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(circuitsResource, c.ns, name), &v1alpha1.Circuit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Circuit), err
}

// List takes label and field selectors, and returns the list of Circuits that match those selectors.
func (c *FakeCircuits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CircuitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(circuitsResource, circuitsKind, c.ns, opts), &v1alpha1.CircuitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CircuitList{ListMeta: obj.(*v1alpha1.CircuitList).ListMeta}
	for _, item := range obj.(*v1alpha1.CircuitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested circuits.
func (c *FakeCircuits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(circuitsResource, c.ns, opts))

}

// Create takes the representation of a circuit and creates it.  Returns the server's representation of the circuit, and an error, if there is any.
func (c *FakeCircuits) Create(ctx context.Context, circuit *v1alpha1.Circuit, opts v1.CreateOptions) (result *v1alpha1.Circuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(circuitsResource, c.ns, circuit), &v1alpha1.Circuit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Circuit), err
}

// Update takes the representation of a circuit and updates it. Returns the server's representation of the circuit, and an error, if there is any.
func (c *FakeCircuits) Update(ctx context.Context, circuit *v1alpha1.Circuit, opts v1.UpdateOptions) (result *v1alpha1.Circuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(circuitsResource, c.ns, circuit), &v1alpha1.Circuit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Circuit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCircuits) UpdateStatus(ctx context.Context, circuit *v1alpha1.Circuit, opts v1.UpdateOptions) (*v1alpha1.Circuit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(circuitsResource, "status", c.ns, circuit), &v1alpha1.Circuit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Circuit), err
}

// Delete takes name of the circuit and deletes it. Returns an error if one occurs.
func (c *FakeCircuits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(circuitsResource, c.ns, name), &v1alpha1.Circuit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCircuits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(circuitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CircuitList{})
	return err
}

// Patch applies the patch and returns the patched circuit.
func (c *FakeCircuits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Circuit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(circuitsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Circuit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Circuit), err
}
