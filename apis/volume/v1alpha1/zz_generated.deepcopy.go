//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Volume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeInitParameters) DeepCopyInto(out *VolumeInitParameters) {
	*out = *in
	if in.BaseVolumeID != nil {
		in, out := &in.BaseVolumeID, &out.BaseVolumeID
		*out = new(string)
		**out = **in
	}
	if in.BaseVolumeName != nil {
		in, out := &in.BaseVolumeName, &out.BaseVolumeName
		*out = new(string)
		**out = **in
	}
	if in.BaseVolumePool != nil {
		in, out := &in.BaseVolumePool, &out.BaseVolumePool
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.XML != nil {
		in, out := &in.XML, &out.XML
		*out = make([]XMLInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeInitParameters.
func (in *VolumeInitParameters) DeepCopy() *VolumeInitParameters {
	if in == nil {
		return nil
	}
	out := new(VolumeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeList) DeepCopyInto(out *VolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeList.
func (in *VolumeList) DeepCopy() *VolumeList {
	if in == nil {
		return nil
	}
	out := new(VolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeObservation) DeepCopyInto(out *VolumeObservation) {
	*out = *in
	if in.BaseVolumeID != nil {
		in, out := &in.BaseVolumeID, &out.BaseVolumeID
		*out = new(string)
		**out = **in
	}
	if in.BaseVolumeName != nil {
		in, out := &in.BaseVolumeName, &out.BaseVolumeName
		*out = new(string)
		**out = **in
	}
	if in.BaseVolumePool != nil {
		in, out := &in.BaseVolumePool, &out.BaseVolumePool
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.XML != nil {
		in, out := &in.XML, &out.XML
		*out = make([]XMLObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeObservation.
func (in *VolumeObservation) DeepCopy() *VolumeObservation {
	if in == nil {
		return nil
	}
	out := new(VolumeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeParameters) DeepCopyInto(out *VolumeParameters) {
	*out = *in
	if in.BaseVolumeID != nil {
		in, out := &in.BaseVolumeID, &out.BaseVolumeID
		*out = new(string)
		**out = **in
	}
	if in.BaseVolumeName != nil {
		in, out := &in.BaseVolumeName, &out.BaseVolumeName
		*out = new(string)
		**out = **in
	}
	if in.BaseVolumePool != nil {
		in, out := &in.BaseVolumePool, &out.BaseVolumePool
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.PoolRef != nil {
		in, out := &in.PoolRef, &out.PoolRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PoolSelector != nil {
		in, out := &in.PoolSelector, &out.PoolSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.XML != nil {
		in, out := &in.XML, &out.XML
		*out = make([]XMLParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeParameters.
func (in *VolumeParameters) DeepCopy() *VolumeParameters {
	if in == nil {
		return nil
	}
	out := new(VolumeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeStatus) DeepCopyInto(out *VolumeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeStatus.
func (in *VolumeStatus) DeepCopy() *VolumeStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XMLInitParameters) DeepCopyInto(out *XMLInitParameters) {
	*out = *in
	if in.Xslt != nil {
		in, out := &in.Xslt, &out.Xslt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XMLInitParameters.
func (in *XMLInitParameters) DeepCopy() *XMLInitParameters {
	if in == nil {
		return nil
	}
	out := new(XMLInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XMLObservation) DeepCopyInto(out *XMLObservation) {
	*out = *in
	if in.Xslt != nil {
		in, out := &in.Xslt, &out.Xslt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XMLObservation.
func (in *XMLObservation) DeepCopy() *XMLObservation {
	if in == nil {
		return nil
	}
	out := new(XMLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XMLParameters) DeepCopyInto(out *XMLParameters) {
	*out = *in
	if in.Xslt != nil {
		in, out := &in.Xslt, &out.Xslt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XMLParameters.
func (in *XMLParameters) DeepCopy() *XMLParameters {
	if in == nil {
		return nil
	}
	out := new(XMLParameters)
	in.DeepCopyInto(out)
	return out
}