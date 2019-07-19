// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ring) DeepCopyInto(out *Ring) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ring.
func (in *Ring) DeepCopy() *Ring {
	if in == nil {
		return nil
	}
	out := new(Ring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ring) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RingGroup) DeepCopyInto(out *RingGroup) {
	*out = *in
	if in.InitialUsers != nil {
		in, out := &in.InitialUsers, &out.InitialUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RingGroup.
func (in *RingGroup) DeepCopy() *RingGroup {
	if in == nil {
		return nil
	}
	out := new(RingGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RingList) DeepCopyInto(out *RingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RingList.
func (in *RingList) DeepCopy() *RingList {
	if in == nil {
		return nil
	}
	out := new(RingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RingPort) DeepCopyInto(out *RingPort) {
	*out = *in
	out.TargetPort = in.TargetPort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RingPort.
func (in *RingPort) DeepCopy() *RingPort {
	if in == nil {
		return nil
	}
	out := new(RingPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RingRouting) DeepCopyInto(out *RingRouting) {
	*out = *in
	in.Group.DeepCopyInto(&out.Group)
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]RingPort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RingRouting.
func (in *RingRouting) DeepCopy() *RingRouting {
	if in == nil {
		return nil
	}
	out := new(RingRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RingSpec) DeepCopyInto(out *RingSpec) {
	*out = *in
	in.Routing.DeepCopyInto(&out.Routing)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RingSpec.
func (in *RingSpec) DeepCopy() *RingSpec {
	if in == nil {
		return nil
	}
	out := new(RingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RingStatus) DeepCopyInto(out *RingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RingStatus.
func (in *RingStatus) DeepCopy() *RingStatus {
	if in == nil {
		return nil
	}
	out := new(RingStatus)
	in.DeepCopyInto(out)
	return out
}
