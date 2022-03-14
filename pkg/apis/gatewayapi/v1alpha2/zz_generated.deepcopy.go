//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Flux authors

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

package v1alpha2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressMatch) DeepCopyInto(out *AddressMatch) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(AddressType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressMatch.
func (in *AddressMatch) DeepCopy() *AddressMatch {
	if in == nil {
		return nil
	}
	out := new(AddressMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressRouteMatches) DeepCopyInto(out *AddressRouteMatches) {
	*out = *in
	if in.SourceAddresses != nil {
		in, out := &in.SourceAddresses, &out.SourceAddresses
		*out = make([]AddressMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DestinationAddresses != nil {
		in, out := &in.DestinationAddresses, &out.DestinationAddresses
		*out = make([]AddressMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressRouteMatches.
func (in *AddressRouteMatches) DeepCopy() *AddressRouteMatches {
	if in == nil {
		return nil
	}
	out := new(AddressRouteMatches)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendObjectReference) DeepCopyInto(out *BackendObjectReference) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(Group)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(Kind)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(Namespace)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(PortNumber)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendObjectReference.
func (in *BackendObjectReference) DeepCopy() *BackendObjectReference {
	if in == nil {
		return nil
	}
	out := new(BackendObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendRef) DeepCopyInto(out *BackendRef) {
	*out = *in
	in.BackendObjectReference.DeepCopyInto(&out.BackendObjectReference)
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendRef.
func (in *BackendRef) DeepCopy() *BackendRef {
	if in == nil {
		return nil
	}
	out := new(BackendRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonRouteSpec) DeepCopyInto(out *CommonRouteSpec) {
	*out = *in
	if in.ParentRefs != nil {
		in, out := &in.ParentRefs, &out.ParentRefs
		*out = make([]ParentReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonRouteSpec.
func (in *CommonRouteSpec) DeepCopy() *CommonRouteSpec {
	if in == nil {
		return nil
	}
	out := new(CommonRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBackendRef) DeepCopyInto(out *HTTPBackendRef) {
	*out = *in
	in.BackendRef.DeepCopyInto(&out.BackendRef)
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]HTTPRouteFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBackendRef.
func (in *HTTPBackendRef) DeepCopy() *HTTPBackendRef {
	if in == nil {
		return nil
	}
	out := new(HTTPBackendRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader.
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
	if in == nil {
		return nil
	}
	out := new(HTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPHeaderMatch) DeepCopyInto(out *HTTPHeaderMatch) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(HeaderMatchType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeaderMatch.
func (in *HTTPHeaderMatch) DeepCopy() *HTTPHeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HTTPHeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPPathMatch) DeepCopyInto(out *HTTPPathMatch) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(PathMatchType)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPPathMatch.
func (in *HTTPPathMatch) DeepCopy() *HTTPPathMatch {
	if in == nil {
		return nil
	}
	out := new(HTTPPathMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPPathModifier) DeepCopyInto(out *HTTPPathModifier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPPathModifier.
func (in *HTTPPathModifier) DeepCopy() *HTTPPathModifier {
	if in == nil {
		return nil
	}
	out := new(HTTPPathModifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPQueryParamMatch) DeepCopyInto(out *HTTPQueryParamMatch) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(QueryParamMatchType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPQueryParamMatch.
func (in *HTTPQueryParamMatch) DeepCopy() *HTTPQueryParamMatch {
	if in == nil {
		return nil
	}
	out := new(HTTPQueryParamMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRequestHeaderFilter) DeepCopyInto(out *HTTPRequestHeaderFilter) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]HTTPHeader, len(*in))
		copy(*out, *in)
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make([]HTTPHeader, len(*in))
		copy(*out, *in)
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRequestHeaderFilter.
func (in *HTTPRequestHeaderFilter) DeepCopy() *HTTPRequestHeaderFilter {
	if in == nil {
		return nil
	}
	out := new(HTTPRequestHeaderFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRequestMirrorFilter) DeepCopyInto(out *HTTPRequestMirrorFilter) {
	*out = *in
	in.BackendRef.DeepCopyInto(&out.BackendRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRequestMirrorFilter.
func (in *HTTPRequestMirrorFilter) DeepCopy() *HTTPRequestMirrorFilter {
	if in == nil {
		return nil
	}
	out := new(HTTPRequestMirrorFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRequestRedirectFilter) DeepCopyInto(out *HTTPRequestRedirectFilter) {
	*out = *in
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(PreciseHostname)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(HTTPPathModifier)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(PortNumber)
		**out = **in
	}
	if in.StatusCode != nil {
		in, out := &in.StatusCode, &out.StatusCode
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRequestRedirectFilter.
func (in *HTTPRequestRedirectFilter) DeepCopy() *HTTPRequestRedirectFilter {
	if in == nil {
		return nil
	}
	out := new(HTTPRequestRedirectFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRoute) DeepCopyInto(out *HTTPRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute.
func (in *HTTPRoute) DeepCopy() *HTTPRoute {
	if in == nil {
		return nil
	}
	out := new(HTTPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteFilter) DeepCopyInto(out *HTTPRouteFilter) {
	*out = *in
	if in.RequestHeaderModifier != nil {
		in, out := &in.RequestHeaderModifier, &out.RequestHeaderModifier
		*out = new(HTTPRequestHeaderFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestMirror != nil {
		in, out := &in.RequestMirror, &out.RequestMirror
		*out = new(HTTPRequestMirrorFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestRedirect != nil {
		in, out := &in.RequestRedirect, &out.RequestRedirect
		*out = new(HTTPRequestRedirectFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.URLRewrite != nil {
		in, out := &in.URLRewrite, &out.URLRewrite
		*out = new(HTTPURLRewriteFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtensionRef != nil {
		in, out := &in.ExtensionRef, &out.ExtensionRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteFilter.
func (in *HTTPRouteFilter) DeepCopy() *HTTPRouteFilter {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteList) DeepCopyInto(out *HTTPRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteList.
func (in *HTTPRouteList) DeepCopy() *HTTPRouteList {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteMatch) DeepCopyInto(out *HTTPRouteMatch) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(HTTPPathMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HTTPHeaderMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueryParams != nil {
		in, out := &in.QueryParams, &out.QueryParams
		*out = make([]HTTPQueryParamMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(HTTPMethod)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteMatch.
func (in *HTTPRouteMatch) DeepCopy() *HTTPRouteMatch {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteRule) DeepCopyInto(out *HTTPRouteRule) {
	*out = *in
	if in.Matches != nil {
		in, out := &in.Matches, &out.Matches
		*out = make([]HTTPRouteMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]HTTPRouteFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackendRefs != nil {
		in, out := &in.BackendRefs, &out.BackendRefs
		*out = make([]HTTPBackendRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteRule.
func (in *HTTPRouteRule) DeepCopy() *HTTPRouteRule {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteSpec) DeepCopyInto(out *HTTPRouteSpec) {
	*out = *in
	in.CommonRouteSpec.DeepCopyInto(&out.CommonRouteSpec)
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]Hostname, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]HTTPRouteRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteSpec.
func (in *HTTPRouteSpec) DeepCopy() *HTTPRouteSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteStatus) DeepCopyInto(out *HTTPRouteStatus) {
	*out = *in
	in.RouteStatus.DeepCopyInto(&out.RouteStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteStatus.
func (in *HTTPRouteStatus) DeepCopy() *HTTPRouteStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPURLRewriteFilter) DeepCopyInto(out *HTTPURLRewriteFilter) {
	*out = *in
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(Hostname)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(HTTPPathModifier)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPURLRewriteFilter.
func (in *HTTPURLRewriteFilter) DeepCopy() *HTTPURLRewriteFilter {
	if in == nil {
		return nil
	}
	out := new(HTTPURLRewriteFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParentReference) DeepCopyInto(out *ParentReference) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(Group)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(Kind)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(Namespace)
		**out = **in
	}
	if in.SectionName != nil {
		in, out := &in.SectionName, &out.SectionName
		*out = new(SectionName)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParentReference.
func (in *ParentReference) DeepCopy() *ParentReference {
	if in == nil {
		return nil
	}
	out := new(ParentReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteParentStatus) DeepCopyInto(out *RouteParentStatus) {
	*out = *in
	in.ParentRef.DeepCopyInto(&out.ParentRef)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteParentStatus.
func (in *RouteParentStatus) DeepCopy() *RouteParentStatus {
	if in == nil {
		return nil
	}
	out := new(RouteParentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteStatus) DeepCopyInto(out *RouteStatus) {
	*out = *in
	if in.Parents != nil {
		in, out := &in.Parents, &out.Parents
		*out = make([]RouteParentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteStatus.
func (in *RouteStatus) DeepCopy() *RouteStatus {
	if in == nil {
		return nil
	}
	out := new(RouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretObjectReference) DeepCopyInto(out *SecretObjectReference) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(Group)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(Kind)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(Namespace)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretObjectReference.
func (in *SecretObjectReference) DeepCopy() *SecretObjectReference {
	if in == nil {
		return nil
	}
	out := new(SecretObjectReference)
	in.DeepCopyInto(out)
	return out
}
