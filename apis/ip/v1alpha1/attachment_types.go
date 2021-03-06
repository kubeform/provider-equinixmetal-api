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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Attachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentSpec   `json:"spec,omitempty"`
	Status            AttachmentStatus `json:"status,omitempty"`
}

type AttachmentSpec struct {
	State *AttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource AttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AttachmentSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// Address family as integer (4 or 6)
	// +optional
	AddressFamily *int64 `json:"addressFamily,omitempty" tf:"address_family"`
	// Length of CIDR prefix of the block as integer
	// +optional
	Cidr         *int64  `json:"cidr,omitempty" tf:"cidr"`
	CidrNotation *string `json:"cidrNotation" tf:"cidr_notation"`
	DeviceID     *string `json:"deviceID" tf:"device_id"`
	// +optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway"`
	// Flag indicating whether IP block is global, i.e. assignable in any location
	// +optional
	Global *bool `json:"global,omitempty" tf:"global"`
	// +optional
	Manageable *bool `json:"manageable,omitempty" tf:"manageable"`
	// +optional
	Management *bool `json:"management,omitempty" tf:"management"`
	// Mask in decimal notation, e.g. 255.255.255.0
	// +optional
	Netmask *string `json:"netmask,omitempty" tf:"netmask"`
	// Network IP address portion of the block specification
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// Flag indicating whether IP block is addressable from the Internet
	// +optional
	Public *bool `json:"public,omitempty" tf:"public"`
}

type AttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AttachmentList is a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Attachment CRD objects
	Items []Attachment `json:"items,omitempty"`
}
