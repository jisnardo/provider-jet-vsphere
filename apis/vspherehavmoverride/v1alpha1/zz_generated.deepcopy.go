//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOverride) DeepCopyInto(out *VMOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOverride.
func (in *VMOverride) DeepCopy() *VMOverride {
	if in == nil {
		return nil
	}
	out := new(VMOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOverrideList) DeepCopyInto(out *VMOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VMOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOverrideList.
func (in *VMOverrideList) DeepCopy() *VMOverrideList {
	if in == nil {
		return nil
	}
	out := new(VMOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOverrideObservation) DeepCopyInto(out *VMOverrideObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOverrideObservation.
func (in *VMOverrideObservation) DeepCopy() *VMOverrideObservation {
	if in == nil {
		return nil
	}
	out := new(VMOverrideObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOverrideParameters) DeepCopyInto(out *VMOverrideParameters) {
	*out = *in
	if in.ComputeClusterID != nil {
		in, out := &in.ComputeClusterID, &out.ComputeClusterID
		*out = new(string)
		**out = **in
	}
	if in.HaDatastoreApdRecoveryAction != nil {
		in, out := &in.HaDatastoreApdRecoveryAction, &out.HaDatastoreApdRecoveryAction
		*out = new(string)
		**out = **in
	}
	if in.HaDatastoreApdResponse != nil {
		in, out := &in.HaDatastoreApdResponse, &out.HaDatastoreApdResponse
		*out = new(string)
		**out = **in
	}
	if in.HaDatastoreApdResponseDelay != nil {
		in, out := &in.HaDatastoreApdResponseDelay, &out.HaDatastoreApdResponseDelay
		*out = new(float64)
		**out = **in
	}
	if in.HaDatastorePdlResponse != nil {
		in, out := &in.HaDatastorePdlResponse, &out.HaDatastorePdlResponse
		*out = new(string)
		**out = **in
	}
	if in.HaHostIsolationResponse != nil {
		in, out := &in.HaHostIsolationResponse, &out.HaHostIsolationResponse
		*out = new(string)
		**out = **in
	}
	if in.HaVMFailureInterval != nil {
		in, out := &in.HaVMFailureInterval, &out.HaVMFailureInterval
		*out = new(float64)
		**out = **in
	}
	if in.HaVMMaximumFailureWindow != nil {
		in, out := &in.HaVMMaximumFailureWindow, &out.HaVMMaximumFailureWindow
		*out = new(float64)
		**out = **in
	}
	if in.HaVMMaximumResets != nil {
		in, out := &in.HaVMMaximumResets, &out.HaVMMaximumResets
		*out = new(float64)
		**out = **in
	}
	if in.HaVMMinimumUptime != nil {
		in, out := &in.HaVMMinimumUptime, &out.HaVMMinimumUptime
		*out = new(float64)
		**out = **in
	}
	if in.HaVMMonitoring != nil {
		in, out := &in.HaVMMonitoring, &out.HaVMMonitoring
		*out = new(string)
		**out = **in
	}
	if in.HaVMMonitoringUseClusterDefaults != nil {
		in, out := &in.HaVMMonitoringUseClusterDefaults, &out.HaVMMonitoringUseClusterDefaults
		*out = new(bool)
		**out = **in
	}
	if in.HaVMRestartPriority != nil {
		in, out := &in.HaVMRestartPriority, &out.HaVMRestartPriority
		*out = new(string)
		**out = **in
	}
	if in.HaVMRestartTimeout != nil {
		in, out := &in.HaVMRestartTimeout, &out.HaVMRestartTimeout
		*out = new(float64)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOverrideParameters.
func (in *VMOverrideParameters) DeepCopy() *VMOverrideParameters {
	if in == nil {
		return nil
	}
	out := new(VMOverrideParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOverrideSpec) DeepCopyInto(out *VMOverrideSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOverrideSpec.
func (in *VMOverrideSpec) DeepCopy() *VMOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(VMOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMOverrideStatus) DeepCopyInto(out *VMOverrideStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMOverrideStatus.
func (in *VMOverrideStatus) DeepCopy() *VMOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(VMOverrideStatus)
	in.DeepCopyInto(out)
	return out
}