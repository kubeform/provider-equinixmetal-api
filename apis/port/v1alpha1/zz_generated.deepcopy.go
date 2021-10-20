//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanAttachment) DeepCopyInto(out *VlanAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanAttachment.
func (in *VlanAttachment) DeepCopy() *VlanAttachment {
	if in == nil {
		return nil
	}
	out := new(VlanAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VlanAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanAttachmentList) DeepCopyInto(out *VlanAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VlanAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanAttachmentList.
func (in *VlanAttachmentList) DeepCopy() *VlanAttachmentList {
	if in == nil {
		return nil
	}
	out := new(VlanAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VlanAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanAttachmentSpec) DeepCopyInto(out *VlanAttachmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VlanAttachmentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanAttachmentSpec.
func (in *VlanAttachmentSpec) DeepCopy() *VlanAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(VlanAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanAttachmentSpecResource) DeepCopyInto(out *VlanAttachmentSpecResource) {
	*out = *in
	if in.DeviceID != nil {
		in, out := &in.DeviceID, &out.DeviceID
		*out = new(string)
		**out = **in
	}
	if in.ForceBond != nil {
		in, out := &in.ForceBond, &out.ForceBond
		*out = new(bool)
		**out = **in
	}
	if in.Native != nil {
		in, out := &in.Native, &out.Native
		*out = new(bool)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.PortName != nil {
		in, out := &in.PortName, &out.PortName
		*out = new(string)
		**out = **in
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(string)
		**out = **in
	}
	if in.VlanVnid != nil {
		in, out := &in.VlanVnid, &out.VlanVnid
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanAttachmentSpecResource.
func (in *VlanAttachmentSpecResource) DeepCopy() *VlanAttachmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(VlanAttachmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanAttachmentStatus) DeepCopyInto(out *VlanAttachmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanAttachmentStatus.
func (in *VlanAttachmentStatus) DeepCopy() *VlanAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(VlanAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}
