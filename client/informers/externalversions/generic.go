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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/bgp/v1alpha1"
	connectionv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/connection/v1alpha1"
	devicev1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/device/v1alpha1"
	ipv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/ip/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/organization/v1alpha1"
	portv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/port/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/project/v1alpha1"
	reservedv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/reserved/v1alpha1"
	spotv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/spot/v1alpha1"
	sshv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/ssh/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/user/v1alpha1"
	virtualv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/virtual/v1alpha1"
	vlanv1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/vlan/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-equinixmetal-api/apis/volume/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=bgp.equinixmetal.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("sessions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Bgp().V1alpha1().Sessions().Informer()}, nil

		// Group=connection.equinixmetal.kubeform.com, Version=v1alpha1
	case connectionv1alpha1.SchemeGroupVersion.WithResource("connections"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Connection().V1alpha1().Connections().Informer()}, nil

		// Group=device.equinixmetal.kubeform.com, Version=v1alpha1
	case devicev1alpha1.SchemeGroupVersion.WithResource("devices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Device().V1alpha1().Devices().Informer()}, nil
	case devicev1alpha1.SchemeGroupVersion.WithResource("networktypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Device().V1alpha1().NetworkTypes().Informer()}, nil

		// Group=ip.equinixmetal.kubeform.com, Version=v1alpha1
	case ipv1alpha1.SchemeGroupVersion.WithResource("attachments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ip().V1alpha1().Attachments().Informer()}, nil

		// Group=organization.equinixmetal.kubeform.com, Version=v1alpha1
	case organizationv1alpha1.SchemeGroupVersion.WithResource("organizations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Organization().V1alpha1().Organizations().Informer()}, nil

		// Group=port.equinixmetal.kubeform.com, Version=v1alpha1
	case portv1alpha1.SchemeGroupVersion.WithResource("vlanattachments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Port().V1alpha1().VlanAttachments().Informer()}, nil

		// Group=project.equinixmetal.kubeform.com, Version=v1alpha1
	case projectv1alpha1.SchemeGroupVersion.WithResource("apikeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Project().V1alpha1().ApiKeys().Informer()}, nil
	case projectv1alpha1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Project().V1alpha1().Projects().Informer()}, nil
	case projectv1alpha1.SchemeGroupVersion.WithResource("sshkeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Project().V1alpha1().SshKeys().Informer()}, nil

		// Group=reserved.equinixmetal.kubeform.com, Version=v1alpha1
	case reservedv1alpha1.SchemeGroupVersion.WithResource("ipblocks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Reserved().V1alpha1().IpBlocks().Informer()}, nil

		// Group=spot.equinixmetal.kubeform.com, Version=v1alpha1
	case spotv1alpha1.SchemeGroupVersion.WithResource("marketrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Spot().V1alpha1().MarketRequests().Informer()}, nil

		// Group=ssh.equinixmetal.kubeform.com, Version=v1alpha1
	case sshv1alpha1.SchemeGroupVersion.WithResource("keys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ssh().V1alpha1().Keys().Informer()}, nil

		// Group=user.equinixmetal.kubeform.com, Version=v1alpha1
	case userv1alpha1.SchemeGroupVersion.WithResource("apikeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.User().V1alpha1().ApiKeys().Informer()}, nil

		// Group=virtual.equinixmetal.kubeform.com, Version=v1alpha1
	case virtualv1alpha1.SchemeGroupVersion.WithResource("circuits"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Virtual().V1alpha1().Circuits().Informer()}, nil

		// Group=vlan.equinixmetal.kubeform.com, Version=v1alpha1
	case vlanv1alpha1.SchemeGroupVersion.WithResource("vlans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Vlan().V1alpha1().Vlans().Informer()}, nil

		// Group=volume.equinixmetal.kubeform.com, Version=v1alpha1
	case volumev1alpha1.SchemeGroupVersion.WithResource("attachments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Volume().V1alpha1().Attachments().Informer()}, nil
	case volumev1alpha1.SchemeGroupVersion.WithResource("volumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Volume().V1alpha1().Volumes().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
