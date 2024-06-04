//go:build !ignore_autogenerated

/*
Copyright 2023 The bpfman Authors.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfAppCommon) DeepCopyInto(out *BpfAppCommon) {
	*out = *in
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	if in.GlobalData != nil {
		in, out := &in.GlobalData, &out.GlobalData
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfAppCommon.
func (in *BpfAppCommon) DeepCopy() *BpfAppCommon {
	if in == nil {
		return nil
	}
	out := new(BpfAppCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgram) DeepCopyInto(out *BpfProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgram.
func (in *BpfProgram) DeepCopy() *BpfProgram {
	if in == nil {
		return nil
	}
	out := new(BpfProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BpfProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramCommon) DeepCopyInto(out *BpfProgramCommon) {
	*out = *in
	in.ByteCode.DeepCopyInto(&out.ByteCode)
	in.MapOwnerSelector.DeepCopyInto(&out.MapOwnerSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramCommon.
func (in *BpfProgramCommon) DeepCopy() *BpfProgramCommon {
	if in == nil {
		return nil
	}
	out := new(BpfProgramCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramList) DeepCopyInto(out *BpfProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BpfProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramList.
func (in *BpfProgramList) DeepCopy() *BpfProgramList {
	if in == nil {
		return nil
	}
	out := new(BpfProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BpfProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramSpec) DeepCopyInto(out *BpfProgramSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramSpec.
func (in *BpfProgramSpec) DeepCopy() *BpfProgramSpec {
	if in == nil {
		return nil
	}
	out := new(BpfProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramStatus) DeepCopyInto(out *BpfProgramStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramStatus.
func (in *BpfProgramStatus) DeepCopy() *BpfProgramStatus {
	if in == nil {
		return nil
	}
	out := new(BpfProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramStatusCommon) DeepCopyInto(out *BpfProgramStatusCommon) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramStatusCommon.
func (in *BpfProgramStatusCommon) DeepCopy() *BpfProgramStatusCommon {
	if in == nil {
		return nil
	}
	out := new(BpfProgramStatusCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BytecodeImage) DeepCopyInto(out *BytecodeImage) {
	*out = *in
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(ImagePullSecretSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BytecodeImage.
func (in *BytecodeImage) DeepCopy() *BytecodeImage {
	if in == nil {
		return nil
	}
	out := new(BytecodeImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BytecodeSelector) DeepCopyInto(out *BytecodeSelector) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(BytecodeImage)
		(*in).DeepCopyInto(*out)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BytecodeSelector.
func (in *BytecodeSelector) DeepCopy() *BytecodeSelector {
	if in == nil {
		return nil
	}
	out := new(BytecodeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSelector) DeepCopyInto(out *ContainerSelector) {
	*out = *in
	in.Pods.DeepCopyInto(&out.Pods)
	if in.ContainerNames != nil {
		in, out := &in.ContainerNames, &out.ContainerNames
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSelector.
func (in *ContainerSelector) DeepCopy() *ContainerSelector {
	if in == nil {
		return nil
	}
	out := new(ContainerSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FentryProgram) DeepCopyInto(out *FentryProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FentryProgram.
func (in *FentryProgram) DeepCopy() *FentryProgram {
	if in == nil {
		return nil
	}
	out := new(FentryProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FentryProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FentryProgramInfo) DeepCopyInto(out *FentryProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FentryProgramInfo.
func (in *FentryProgramInfo) DeepCopy() *FentryProgramInfo {
	if in == nil {
		return nil
	}
	out := new(FentryProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FentryProgramList) DeepCopyInto(out *FentryProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FentryProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FentryProgramList.
func (in *FentryProgramList) DeepCopy() *FentryProgramList {
	if in == nil {
		return nil
	}
	out := new(FentryProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FentryProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FentryProgramSpec) DeepCopyInto(out *FentryProgramSpec) {
	*out = *in
	in.FentryProgramInfo.DeepCopyInto(&out.FentryProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FentryProgramSpec.
func (in *FentryProgramSpec) DeepCopy() *FentryProgramSpec {
	if in == nil {
		return nil
	}
	out := new(FentryProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FentryProgramStatus) DeepCopyInto(out *FentryProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FentryProgramStatus.
func (in *FentryProgramStatus) DeepCopy() *FentryProgramStatus {
	if in == nil {
		return nil
	}
	out := new(FentryProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FexitProgram) DeepCopyInto(out *FexitProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FexitProgram.
func (in *FexitProgram) DeepCopy() *FexitProgram {
	if in == nil {
		return nil
	}
	out := new(FexitProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FexitProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FexitProgramInfo) DeepCopyInto(out *FexitProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FexitProgramInfo.
func (in *FexitProgramInfo) DeepCopy() *FexitProgramInfo {
	if in == nil {
		return nil
	}
	out := new(FexitProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FexitProgramList) DeepCopyInto(out *FexitProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FexitProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FexitProgramList.
func (in *FexitProgramList) DeepCopy() *FexitProgramList {
	if in == nil {
		return nil
	}
	out := new(FexitProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FexitProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FexitProgramSpec) DeepCopyInto(out *FexitProgramSpec) {
	*out = *in
	in.FexitProgramInfo.DeepCopyInto(&out.FexitProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FexitProgramSpec.
func (in *FexitProgramSpec) DeepCopy() *FexitProgramSpec {
	if in == nil {
		return nil
	}
	out := new(FexitProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FexitProgramStatus) DeepCopyInto(out *FexitProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FexitProgramStatus.
func (in *FexitProgramStatus) DeepCopy() *FexitProgramStatus {
	if in == nil {
		return nil
	}
	out := new(FexitProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullSecretSelector) DeepCopyInto(out *ImagePullSecretSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullSecretSelector.
func (in *ImagePullSecretSelector) DeepCopy() *ImagePullSecretSelector {
	if in == nil {
		return nil
	}
	out := new(ImagePullSecretSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSelector) DeepCopyInto(out *InterfaceSelector) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.PrimaryNodeInterface != nil {
		in, out := &in.PrimaryNodeInterface, &out.PrimaryNodeInterface
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSelector.
func (in *InterfaceSelector) DeepCopy() *InterfaceSelector {
	if in == nil {
		return nil
	}
	out := new(InterfaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KprobeProgram) DeepCopyInto(out *KprobeProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KprobeProgram.
func (in *KprobeProgram) DeepCopy() *KprobeProgram {
	if in == nil {
		return nil
	}
	out := new(KprobeProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KprobeProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KprobeProgramInfo) DeepCopyInto(out *KprobeProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KprobeProgramInfo.
func (in *KprobeProgramInfo) DeepCopy() *KprobeProgramInfo {
	if in == nil {
		return nil
	}
	out := new(KprobeProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KprobeProgramList) DeepCopyInto(out *KprobeProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KprobeProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KprobeProgramList.
func (in *KprobeProgramList) DeepCopy() *KprobeProgramList {
	if in == nil {
		return nil
	}
	out := new(KprobeProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KprobeProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KprobeProgramSpec) DeepCopyInto(out *KprobeProgramSpec) {
	*out = *in
	in.KprobeProgramInfo.DeepCopyInto(&out.KprobeProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KprobeProgramSpec.
func (in *KprobeProgramSpec) DeepCopy() *KprobeProgramSpec {
	if in == nil {
		return nil
	}
	out := new(KprobeProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KprobeProgramStatus) DeepCopyInto(out *KprobeProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KprobeProgramStatus.
func (in *KprobeProgramStatus) DeepCopy() *KprobeProgramStatus {
	if in == nil {
		return nil
	}
	out := new(KprobeProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgram) DeepCopyInto(out *TcProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgram.
func (in *TcProgram) DeepCopy() *TcProgram {
	if in == nil {
		return nil
	}
	out := new(TcProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TcProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramInfo) DeepCopyInto(out *TcProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
	in.InterfaceSelector.DeepCopyInto(&out.InterfaceSelector)
	if in.ProceedOn != nil {
		in, out := &in.ProceedOn, &out.ProceedOn
		*out = make([]TcProceedOnValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramInfo.
func (in *TcProgramInfo) DeepCopy() *TcProgramInfo {
	if in == nil {
		return nil
	}
	out := new(TcProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramList) DeepCopyInto(out *TcProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TcProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramList.
func (in *TcProgramList) DeepCopy() *TcProgramList {
	if in == nil {
		return nil
	}
	out := new(TcProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TcProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramSpec) DeepCopyInto(out *TcProgramSpec) {
	*out = *in
	in.TcProgramInfo.DeepCopyInto(&out.TcProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramSpec.
func (in *TcProgramSpec) DeepCopy() *TcProgramSpec {
	if in == nil {
		return nil
	}
	out := new(TcProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramStatus) DeepCopyInto(out *TcProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramStatus.
func (in *TcProgramStatus) DeepCopy() *TcProgramStatus {
	if in == nil {
		return nil
	}
	out := new(TcProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgram) DeepCopyInto(out *TracepointProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgram.
func (in *TracepointProgram) DeepCopy() *TracepointProgram {
	if in == nil {
		return nil
	}
	out := new(TracepointProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracepointProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramInfo) DeepCopyInto(out *TracepointProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramInfo.
func (in *TracepointProgramInfo) DeepCopy() *TracepointProgramInfo {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramList) DeepCopyInto(out *TracepointProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TracepointProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramList.
func (in *TracepointProgramList) DeepCopy() *TracepointProgramList {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracepointProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramSpec) DeepCopyInto(out *TracepointProgramSpec) {
	*out = *in
	in.TracepointProgramInfo.DeepCopyInto(&out.TracepointProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramSpec.
func (in *TracepointProgramSpec) DeepCopy() *TracepointProgramSpec {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramStatus) DeepCopyInto(out *TracepointProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramStatus.
func (in *TracepointProgramStatus) DeepCopy() *TracepointProgramStatus {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UprobeProgram) DeepCopyInto(out *UprobeProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UprobeProgram.
func (in *UprobeProgram) DeepCopy() *UprobeProgram {
	if in == nil {
		return nil
	}
	out := new(UprobeProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UprobeProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UprobeProgramInfo) DeepCopyInto(out *UprobeProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = new(ContainerSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UprobeProgramInfo.
func (in *UprobeProgramInfo) DeepCopy() *UprobeProgramInfo {
	if in == nil {
		return nil
	}
	out := new(UprobeProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UprobeProgramList) DeepCopyInto(out *UprobeProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UprobeProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UprobeProgramList.
func (in *UprobeProgramList) DeepCopy() *UprobeProgramList {
	if in == nil {
		return nil
	}
	out := new(UprobeProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UprobeProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UprobeProgramSpec) DeepCopyInto(out *UprobeProgramSpec) {
	*out = *in
	in.UprobeProgramInfo.DeepCopyInto(&out.UprobeProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UprobeProgramSpec.
func (in *UprobeProgramSpec) DeepCopy() *UprobeProgramSpec {
	if in == nil {
		return nil
	}
	out := new(UprobeProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UprobeProgramStatus) DeepCopyInto(out *UprobeProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UprobeProgramStatus.
func (in *UprobeProgramStatus) DeepCopy() *UprobeProgramStatus {
	if in == nil {
		return nil
	}
	out := new(UprobeProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgram) DeepCopyInto(out *XdpProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgram.
func (in *XdpProgram) DeepCopy() *XdpProgram {
	if in == nil {
		return nil
	}
	out := new(XdpProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XdpProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramInfo) DeepCopyInto(out *XdpProgramInfo) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
	in.InterfaceSelector.DeepCopyInto(&out.InterfaceSelector)
	if in.ProceedOn != nil {
		in, out := &in.ProceedOn, &out.ProceedOn
		*out = make([]XdpProceedOnValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramInfo.
func (in *XdpProgramInfo) DeepCopy() *XdpProgramInfo {
	if in == nil {
		return nil
	}
	out := new(XdpProgramInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramList) DeepCopyInto(out *XdpProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XdpProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramList.
func (in *XdpProgramList) DeepCopy() *XdpProgramList {
	if in == nil {
		return nil
	}
	out := new(XdpProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XdpProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramSpec) DeepCopyInto(out *XdpProgramSpec) {
	*out = *in
	in.XdpProgramInfo.DeepCopyInto(&out.XdpProgramInfo)
	in.BpfAppCommon.DeepCopyInto(&out.BpfAppCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramSpec.
func (in *XdpProgramSpec) DeepCopy() *XdpProgramSpec {
	if in == nil {
		return nil
	}
	out := new(XdpProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramStatus) DeepCopyInto(out *XdpProgramStatus) {
	*out = *in
	in.BpfProgramStatusCommon.DeepCopyInto(&out.BpfProgramStatusCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramStatus.
func (in *XdpProgramStatus) DeepCopy() *XdpProgramStatus {
	if in == nil {
		return nil
	}
	out := new(XdpProgramStatus)
	in.DeepCopyInto(out)
	return out
}
