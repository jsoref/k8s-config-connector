//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopyInto(out *BigQueryAnalyticsHubDataExchange) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchange.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopy() *BigQueryAnalyticsHubDataExchange {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryAnalyticsHubDataExchange) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeList) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryAnalyticsHubDataExchange, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeList.
func (in *BigQueryAnalyticsHubDataExchangeList) DeepCopy() *BigQueryAnalyticsHubDataExchangeList {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryAnalyticsHubDataExchangeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeObservedState) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeObservedState) {
	*out = *in
	if in.ListingCount != nil {
		in, out := &in.ListingCount, &out.ListingCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeObservedState.
func (in *BigQueryAnalyticsHubDataExchangeObservedState) DeepCopy() *BigQueryAnalyticsHubDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeSpec) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = new(string)
		**out = **in
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeSpec.
func (in *BigQueryAnalyticsHubDataExchangeSpec) DeepCopy() *BigQueryAnalyticsHubDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryAnalyticsHubDataExchangeStatus) DeepCopyInto(out *BigQueryAnalyticsHubDataExchangeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(BigQueryAnalyticsHubDataExchangeObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryAnalyticsHubDataExchangeStatus.
func (in *BigQueryAnalyticsHubDataExchangeStatus) DeepCopy() *BigQueryAnalyticsHubDataExchangeStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryAnalyticsHubDataExchangeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExchange) DeepCopyInto(out *DataExchange) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PrimaryContact != nil {
		in, out := &in.PrimaryContact, &out.PrimaryContact
		*out = new(string)
		**out = **in
	}
	if in.Documentation != nil {
		in, out := &in.Documentation, &out.Documentation
		*out = new(string)
		**out = **in
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SharingEnvironmentConfig != nil {
		in, out := &in.SharingEnvironmentConfig, &out.SharingEnvironmentConfig
		*out = new(SharingEnvironmentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DiscoveryType != nil {
		in, out := &in.DiscoveryType, &out.DiscoveryType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExchange.
func (in *DataExchange) DeepCopy() *DataExchange {
	if in == nil {
		return nil
	}
	out := new(DataExchange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharingEnvironmentConfig) DeepCopyInto(out *SharingEnvironmentConfig) {
	*out = *in
	if in.DefaultExchangeConfig != nil {
		in, out := &in.DefaultExchangeConfig, &out.DefaultExchangeConfig
		*out = new(SharingEnvironmentConfig_DefaultExchangeConfig)
		**out = **in
	}
	if in.DcrExchangeConfig != nil {
		in, out := &in.DcrExchangeConfig, &out.DcrExchangeConfig
		*out = new(SharingEnvironmentConfig_DcrExchangeConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharingEnvironmentConfig.
func (in *SharingEnvironmentConfig) DeepCopy() *SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := new(SharingEnvironmentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharingEnvironmentConfig_DcrExchangeConfig) DeepCopyInto(out *SharingEnvironmentConfig_DcrExchangeConfig) {
	*out = *in
	if in.SingleSelectedResourceSharingRestriction != nil {
		in, out := &in.SingleSelectedResourceSharingRestriction, &out.SingleSelectedResourceSharingRestriction
		*out = new(bool)
		**out = **in
	}
	if in.SingleLinkedDatasetPerCleanroom != nil {
		in, out := &in.SingleLinkedDatasetPerCleanroom, &out.SingleLinkedDatasetPerCleanroom
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharingEnvironmentConfig_DcrExchangeConfig.
func (in *SharingEnvironmentConfig_DcrExchangeConfig) DeepCopy() *SharingEnvironmentConfig_DcrExchangeConfig {
	if in == nil {
		return nil
	}
	out := new(SharingEnvironmentConfig_DcrExchangeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharingEnvironmentConfig_DefaultExchangeConfig) DeepCopyInto(out *SharingEnvironmentConfig_DefaultExchangeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharingEnvironmentConfig_DefaultExchangeConfig.
func (in *SharingEnvironmentConfig_DefaultExchangeConfig) DeepCopy() *SharingEnvironmentConfig_DefaultExchangeConfig {
	if in == nil {
		return nil
	}
	out := new(SharingEnvironmentConfig_DefaultExchangeConfig)
	in.DeepCopyInto(out)
	return out
}
