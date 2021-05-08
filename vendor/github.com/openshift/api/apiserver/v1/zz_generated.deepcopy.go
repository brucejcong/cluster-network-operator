// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRequestCount) DeepCopyInto(out *APIRequestCount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRequestCount.
func (in *APIRequestCount) DeepCopy() *APIRequestCount {
	if in == nil {
		return nil
	}
	out := new(APIRequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIRequestCount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRequestCountList) DeepCopyInto(out *APIRequestCountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIRequestCount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRequestCountList.
func (in *APIRequestCountList) DeepCopy() *APIRequestCountList {
	if in == nil {
		return nil
	}
	out := new(APIRequestCountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIRequestCountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRequestCountSpec) DeepCopyInto(out *APIRequestCountSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRequestCountSpec.
func (in *APIRequestCountSpec) DeepCopy() *APIRequestCountSpec {
	if in == nil {
		return nil
	}
	out := new(APIRequestCountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRequestCountStatus) DeepCopyInto(out *APIRequestCountStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CurrentHour.DeepCopyInto(&out.CurrentHour)
	if in.Last24h != nil {
		in, out := &in.Last24h, &out.Last24h
		*out = make([]PerResourceAPIRequestLog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRequestCountStatus.
func (in *APIRequestCountStatus) DeepCopy() *APIRequestCountStatus {
	if in == nil {
		return nil
	}
	out := new(APIRequestCountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeprecatedAPIRequest) DeepCopyInto(out *DeprecatedAPIRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeprecatedAPIRequest.
func (in *DeprecatedAPIRequest) DeepCopy() *DeprecatedAPIRequest {
	if in == nil {
		return nil
	}
	out := new(DeprecatedAPIRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeprecatedAPIRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeprecatedAPIRequestList) DeepCopyInto(out *DeprecatedAPIRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeprecatedAPIRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeprecatedAPIRequestList.
func (in *DeprecatedAPIRequestList) DeepCopy() *DeprecatedAPIRequestList {
	if in == nil {
		return nil
	}
	out := new(DeprecatedAPIRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeprecatedAPIRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeprecatedAPIRequestSpec) DeepCopyInto(out *DeprecatedAPIRequestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeprecatedAPIRequestSpec.
func (in *DeprecatedAPIRequestSpec) DeepCopy() *DeprecatedAPIRequestSpec {
	if in == nil {
		return nil
	}
	out := new(DeprecatedAPIRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeprecatedAPIRequestStatus) DeepCopyInto(out *DeprecatedAPIRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RequestsLastHour.DeepCopyInto(&out.RequestsLastHour)
	if in.RequestsLast24h != nil {
		in, out := &in.RequestsLast24h, &out.RequestsLast24h
		*out = make([]RequestLog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeprecatedAPIRequestStatus.
func (in *DeprecatedAPIRequestStatus) DeepCopy() *DeprecatedAPIRequestStatus {
	if in == nil {
		return nil
	}
	out := new(DeprecatedAPIRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeRequestLog) DeepCopyInto(out *NodeRequestLog) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]RequestUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeRequestLog.
func (in *NodeRequestLog) DeepCopy() *NodeRequestLog {
	if in == nil {
		return nil
	}
	out := new(NodeRequestLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerNodeAPIRequestLog) DeepCopyInto(out *PerNodeAPIRequestLog) {
	*out = *in
	if in.ByUser != nil {
		in, out := &in.ByUser, &out.ByUser
		*out = make([]PerUserAPIRequestCount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerNodeAPIRequestLog.
func (in *PerNodeAPIRequestLog) DeepCopy() *PerNodeAPIRequestLog {
	if in == nil {
		return nil
	}
	out := new(PerNodeAPIRequestLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerResourceAPIRequestLog) DeepCopyInto(out *PerResourceAPIRequestLog) {
	*out = *in
	if in.ByNode != nil {
		in, out := &in.ByNode, &out.ByNode
		*out = make([]PerNodeAPIRequestLog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerResourceAPIRequestLog.
func (in *PerResourceAPIRequestLog) DeepCopy() *PerResourceAPIRequestLog {
	if in == nil {
		return nil
	}
	out := new(PerResourceAPIRequestLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerUserAPIRequestCount) DeepCopyInto(out *PerUserAPIRequestCount) {
	*out = *in
	if in.ByVerb != nil {
		in, out := &in.ByVerb, &out.ByVerb
		*out = make([]PerVerbAPIRequestCount, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerUserAPIRequestCount.
func (in *PerUserAPIRequestCount) DeepCopy() *PerUserAPIRequestCount {
	if in == nil {
		return nil
	}
	out := new(PerUserAPIRequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerVerbAPIRequestCount) DeepCopyInto(out *PerVerbAPIRequestCount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerVerbAPIRequestCount.
func (in *PerVerbAPIRequestCount) DeepCopy() *PerVerbAPIRequestCount {
	if in == nil {
		return nil
	}
	out := new(PerVerbAPIRequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestCount) DeepCopyInto(out *RequestCount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestCount.
func (in *RequestCount) DeepCopy() *RequestCount {
	if in == nil {
		return nil
	}
	out := new(RequestCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestLog) DeepCopyInto(out *RequestLog) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodeRequestLog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestLog.
func (in *RequestLog) DeepCopy() *RequestLog {
	if in == nil {
		return nil
	}
	out := new(RequestLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestUser) DeepCopyInto(out *RequestUser) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make([]RequestCount, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestUser.
func (in *RequestUser) DeepCopy() *RequestUser {
	if in == nil {
		return nil
	}
	out := new(RequestUser)
	in.DeepCopyInto(out)
	return out
}