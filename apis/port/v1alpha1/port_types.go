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

type Port struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortSpec   `json:"spec,omitempty"`
	Status            PortStatus `json:"status,omitempty"`
}

type PortSpec struct {
	State *PortSpecResource `json:"state,omitempty" tf:"-"`

	Resource PortSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PortSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// UUID of the bond port
	// +optional
	BondID *string `json:"bondID,omitempty" tf:"bond_id"`
	// Name of the bond port
	// +optional
	BondName *string `json:"bondName,omitempty" tf:"bond_name"`
	// Flag indicating whether the port should be bonded
	Bonded *bool `json:"bonded" tf:"bonded"`
	// Flag indicating whether the port can be removed from a bond
	// +optional
	DisbondSupported *bool `json:"disbondSupported,omitempty" tf:"disbond_supported"`
	// Flag indicating whether the port is in layer2 (or layer3) mode
	// +optional
	Layer2 *bool `json:"layer2,omitempty" tf:"layer2"`
	// MAC address of the port
	// +optional
	Mac *string `json:"mac,omitempty" tf:"mac"`
	// Name of the port to look up, e.g. bond0, eth1
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// UUID of native VLAN of the port
	// +optional
	NativeVLANID *string `json:"nativeVLANID,omitempty" tf:"native_vlan_id"`
	// One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.
	// +optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type"`
	// UUID of the port to lookup
	PortID *string `json:"portID" tf:"port_id"`
	// Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached
	// +optional
	ResetOnDelete *bool `json:"resetOnDelete,omitempty" tf:"reset_on_delete"`
	// Port type
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// UUIDs VLANs to attach. To avoid jitter, use the UUID and not the VXLAN
	// +optional
	VlanIDS []string `json:"vlanIDS,omitempty" tf:"vlan_ids"`
	// VLAN VXLAN ids to attach (example: [1000])
	// +optional
	VxlanIDS []int64 `json:"vxlanIDS,omitempty" tf:"vxlan_ids"`
}

type PortStatus struct {
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

// PortList is a list of Ports
type PortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Port CRD objects
	Items []Port `json:"items,omitempty"`
}
