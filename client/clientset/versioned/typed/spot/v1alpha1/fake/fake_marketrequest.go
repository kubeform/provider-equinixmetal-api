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

	v1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/spot/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMarketRequests implements MarketRequestInterface
type FakeMarketRequests struct {
	Fake *FakeSpotV1alpha1
	ns   string
}

var marketrequestsResource = schema.GroupVersionResource{Group: "spot.equinixmetal.kubeform.com", Version: "v1alpha1", Resource: "marketrequests"}

var marketrequestsKind = schema.GroupVersionKind{Group: "spot.equinixmetal.kubeform.com", Version: "v1alpha1", Kind: "MarketRequest"}

// Get takes name of the marketRequest, and returns the corresponding marketRequest object, and an error if there is any.
func (c *FakeMarketRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MarketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(marketrequestsResource, c.ns, name), &v1alpha1.MarketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketRequest), err
}

// List takes label and field selectors, and returns the list of MarketRequests that match those selectors.
func (c *FakeMarketRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MarketRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(marketrequestsResource, marketrequestsKind, c.ns, opts), &v1alpha1.MarketRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MarketRequestList{ListMeta: obj.(*v1alpha1.MarketRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.MarketRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested marketRequests.
func (c *FakeMarketRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(marketrequestsResource, c.ns, opts))

}

// Create takes the representation of a marketRequest and creates it.  Returns the server's representation of the marketRequest, and an error, if there is any.
func (c *FakeMarketRequests) Create(ctx context.Context, marketRequest *v1alpha1.MarketRequest, opts v1.CreateOptions) (result *v1alpha1.MarketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(marketrequestsResource, c.ns, marketRequest), &v1alpha1.MarketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketRequest), err
}

// Update takes the representation of a marketRequest and updates it. Returns the server's representation of the marketRequest, and an error, if there is any.
func (c *FakeMarketRequests) Update(ctx context.Context, marketRequest *v1alpha1.MarketRequest, opts v1.UpdateOptions) (result *v1alpha1.MarketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(marketrequestsResource, c.ns, marketRequest), &v1alpha1.MarketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMarketRequests) UpdateStatus(ctx context.Context, marketRequest *v1alpha1.MarketRequest, opts v1.UpdateOptions) (*v1alpha1.MarketRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(marketrequestsResource, "status", c.ns, marketRequest), &v1alpha1.MarketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketRequest), err
}

// Delete takes name of the marketRequest and deletes it. Returns an error if one occurs.
func (c *FakeMarketRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(marketrequestsResource, c.ns, name), &v1alpha1.MarketRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMarketRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(marketrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MarketRequestList{})
	return err
}

// Patch applies the patch and returns the patched marketRequest.
func (c *FakeMarketRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MarketRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(marketrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MarketRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketRequest), err
}
