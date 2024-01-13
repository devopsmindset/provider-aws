//go:build !ignore_autogenerated

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasTarget) DeepCopyInto(out *AliasTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasTarget.
func (in *AliasTarget) DeepCopy() *AliasTarget {
	if in == nil {
		return nil
	}
	out := new(AliasTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.PrivateZone != nil {
		in, out := &in.PrivateZone, &out.PrivateZone
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelegationSet) DeepCopyInto(out *DelegationSet) {
	*out = *in
	if in.NameServers != nil {
		in, out := &in.NameServers, &out.NameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelegationSet.
func (in *DelegationSet) DeepCopy() *DelegationSet {
	if in == nil {
		return nil
	}
	out := new(DelegationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoLocation) DeepCopyInto(out *GeoLocation) {
	*out = *in
	if in.ContinentCode != nil {
		in, out := &in.ContinentCode, &out.ContinentCode
		*out = new(string)
		**out = **in
	}
	if in.CountryCode != nil {
		in, out := &in.CountryCode, &out.CountryCode
		*out = new(string)
		**out = **in
	}
	if in.SubdivisionCode != nil {
		in, out := &in.SubdivisionCode, &out.SubdivisionCode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoLocation.
func (in *GeoLocation) DeepCopy() *GeoLocation {
	if in == nil {
		return nil
	}
	out := new(GeoLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZone) DeepCopyInto(out *HostedZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZone.
func (in *HostedZone) DeepCopy() *HostedZone {
	if in == nil {
		return nil
	}
	out := new(HostedZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostedZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZoneList) DeepCopyInto(out *HostedZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostedZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZoneList.
func (in *HostedZoneList) DeepCopy() *HostedZoneList {
	if in == nil {
		return nil
	}
	out := new(HostedZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostedZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZoneObservation) DeepCopyInto(out *HostedZoneObservation) {
	*out = *in
	in.DelegationSet.DeepCopyInto(&out.DelegationSet)
	out.HostedZone = in.HostedZone
	if in.VPCs != nil {
		in, out := &in.VPCs, &out.VPCs
		*out = make([]VPCObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZoneObservation.
func (in *HostedZoneObservation) DeepCopy() *HostedZoneObservation {
	if in == nil {
		return nil
	}
	out := new(HostedZoneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZoneParameters) DeepCopyInto(out *HostedZoneParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(Config)
		(*in).DeepCopyInto(*out)
	}
	if in.DelegationSetID != nil {
		in, out := &in.DelegationSetID, &out.DelegationSetID
		*out = new(string)
		**out = **in
	}
	if in.VPC != nil {
		in, out := &in.VPC, &out.VPC
		*out = new(VPC)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZoneParameters.
func (in *HostedZoneParameters) DeepCopy() *HostedZoneParameters {
	if in == nil {
		return nil
	}
	out := new(HostedZoneParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZoneResponse) DeepCopyInto(out *HostedZoneResponse) {
	*out = *in
	out.LinkedService = in.LinkedService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZoneResponse.
func (in *HostedZoneResponse) DeepCopy() *HostedZoneResponse {
	if in == nil {
		return nil
	}
	out := new(HostedZoneResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZoneSpec) DeepCopyInto(out *HostedZoneSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZoneSpec.
func (in *HostedZoneSpec) DeepCopy() *HostedZoneSpec {
	if in == nil {
		return nil
	}
	out := new(HostedZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostedZoneStatus) DeepCopyInto(out *HostedZoneStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostedZoneStatus.
func (in *HostedZoneStatus) DeepCopy() *HostedZoneStatus {
	if in == nil {
		return nil
	}
	out := new(HostedZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkedService) DeepCopyInto(out *LinkedService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkedService.
func (in *LinkedService) DeepCopy() *LinkedService {
	if in == nil {
		return nil
	}
	out := new(LinkedService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecord) DeepCopyInto(out *ResourceRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecord.
func (in *ResourceRecord) DeepCopy() *ResourceRecord {
	if in == nil {
		return nil
	}
	out := new(ResourceRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSet) DeepCopyInto(out *ResourceRecordSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSet.
func (in *ResourceRecordSet) DeepCopy() *ResourceRecordSet {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceRecordSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetList) DeepCopyInto(out *ResourceRecordSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceRecordSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetList.
func (in *ResourceRecordSetList) DeepCopy() *ResourceRecordSetList {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceRecordSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetParameters) DeepCopyInto(out *ResourceRecordSetParameters) {
	*out = *in
	if in.AliasTarget != nil {
		in, out := &in.AliasTarget, &out.AliasTarget
		*out = new(AliasTarget)
		**out = **in
	}
	if in.GeoLocation != nil {
		in, out := &in.GeoLocation, &out.GeoLocation
		*out = new(GeoLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheckID != nil {
		in, out := &in.HealthCheckID, &out.HealthCheckID
		*out = new(string)
		**out = **in
	}
	if in.MultiValueAnswer != nil {
		in, out := &in.MultiValueAnswer, &out.MultiValueAnswer
		*out = new(bool)
		**out = **in
	}
	if in.ResourceRecords != nil {
		in, out := &in.ResourceRecords, &out.ResourceRecords
		*out = make([]ResourceRecord, len(*in))
		copy(*out, *in)
	}
	if in.SetIdentifier != nil {
		in, out := &in.SetIdentifier, &out.SetIdentifier
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int64)
		**out = **in
	}
	if in.TrafficPolicyInstanceID != nil {
		in, out := &in.TrafficPolicyInstanceID, &out.TrafficPolicyInstanceID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ZoneIDRef != nil {
		in, out := &in.ZoneIDRef, &out.ZoneIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneIDSelector != nil {
		in, out := &in.ZoneIDSelector, &out.ZoneIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetParameters.
func (in *ResourceRecordSetParameters) DeepCopy() *ResourceRecordSetParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetSpec) DeepCopyInto(out *ResourceRecordSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetSpec.
func (in *ResourceRecordSetSpec) DeepCopy() *ResourceRecordSetSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetStatus) DeepCopyInto(out *ResourceRecordSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetStatus.
func (in *ResourceRecordSetStatus) DeepCopy() *ResourceRecordSetStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCRegion != nil {
		in, out := &in.VPCRegion, &out.VPCRegion
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCObservation) DeepCopyInto(out *VPCObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCObservation.
func (in *VPCObservation) DeepCopy() *VPCObservation {
	if in == nil {
		return nil
	}
	out := new(VPCObservation)
	in.DeepCopyInto(out)
	return out
}
