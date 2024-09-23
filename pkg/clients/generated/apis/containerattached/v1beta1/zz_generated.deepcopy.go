//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthorization) DeepCopyInto(out *ClusterAuthorization) {
	*out = *in
	if in.AdminUsers != nil {
		in, out := &in.AdminUsers, &out.AdminUsers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthorization.
func (in *ClusterAuthorization) DeepCopy() *ClusterAuthorization {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBinaryAuthorization) DeepCopyInto(out *ClusterBinaryAuthorization) {
	*out = *in
	if in.EvaluationMode != nil {
		in, out := &in.EvaluationMode, &out.EvaluationMode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBinaryAuthorization.
func (in *ClusterBinaryAuthorization) DeepCopy() *ClusterBinaryAuthorization {
	if in == nil {
		return nil
	}
	out := new(ClusterBinaryAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentConfig) DeepCopyInto(out *ClusterComponentConfig) {
	*out = *in
	if in.EnableComponents != nil {
		in, out := &in.EnableComponents, &out.EnableComponents
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentConfig.
func (in *ClusterComponentConfig) DeepCopy() *ClusterComponentConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterErrorsStatus) DeepCopyInto(out *ClusterErrorsStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterErrorsStatus.
func (in *ClusterErrorsStatus) DeepCopy() *ClusterErrorsStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterErrorsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFleet) DeepCopyInto(out *ClusterFleet) {
	*out = *in
	if in.Membership != nil {
		in, out := &in.Membership, &out.Membership
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFleet.
func (in *ClusterFleet) DeepCopy() *ClusterFleet {
	if in == nil {
		return nil
	}
	out := new(ClusterFleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLoggingConfig) DeepCopyInto(out *ClusterLoggingConfig) {
	*out = *in
	if in.ComponentConfig != nil {
		in, out := &in.ComponentConfig, &out.ComponentConfig
		*out = new(ClusterComponentConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLoggingConfig.
func (in *ClusterLoggingConfig) DeepCopy() *ClusterLoggingConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterLoggingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterManagedPrometheusConfig) DeepCopyInto(out *ClusterManagedPrometheusConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterManagedPrometheusConfig.
func (in *ClusterManagedPrometheusConfig) DeepCopy() *ClusterManagedPrometheusConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterManagedPrometheusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMonitoringConfig) DeepCopyInto(out *ClusterMonitoringConfig) {
	*out = *in
	if in.ManagedPrometheusConfig != nil {
		in, out := &in.ManagedPrometheusConfig, &out.ManagedPrometheusConfig
		*out = new(ClusterManagedPrometheusConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMonitoringConfig.
func (in *ClusterMonitoringConfig) DeepCopy() *ClusterMonitoringConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterMonitoringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservedStateStatus) DeepCopyInto(out *ClusterObservedStateStatus) {
	*out = *in
	if in.FleetMembership != nil {
		in, out := &in.FleetMembership, &out.FleetMembership
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservedStateStatus.
func (in *ClusterObservedStateStatus) DeepCopy() *ClusterObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterOidcConfig) DeepCopyInto(out *ClusterOidcConfig) {
	*out = *in
	if in.Jwks != nil {
		in, out := &in.Jwks, &out.Jwks
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterOidcConfig.
func (in *ClusterOidcConfig) DeepCopy() *ClusterOidcConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterOidcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkloadIdentityConfigStatus) DeepCopyInto(out *ClusterWorkloadIdentityConfigStatus) {
	*out = *in
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.IssuerUri != nil {
		in, out := &in.IssuerUri, &out.IssuerUri
		*out = new(string)
		**out = **in
	}
	if in.WorkloadPool != nil {
		in, out := &in.WorkloadPool, &out.WorkloadPool
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkloadIdentityConfigStatus.
func (in *ClusterWorkloadIdentityConfigStatus) DeepCopy() *ClusterWorkloadIdentityConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkloadIdentityConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAttachedCluster) DeepCopyInto(out *ContainerAttachedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAttachedCluster.
func (in *ContainerAttachedCluster) DeepCopy() *ContainerAttachedCluster {
	if in == nil {
		return nil
	}
	out := new(ContainerAttachedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerAttachedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAttachedClusterList) DeepCopyInto(out *ContainerAttachedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerAttachedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAttachedClusterList.
func (in *ContainerAttachedClusterList) DeepCopy() *ContainerAttachedClusterList {
	if in == nil {
		return nil
	}
	out := new(ContainerAttachedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerAttachedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAttachedClusterSpec) DeepCopyInto(out *ContainerAttachedClusterSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(ClusterAuthorization)
		(*in).DeepCopyInto(*out)
	}
	if in.BinaryAuthorization != nil {
		in, out := &in.BinaryAuthorization, &out.BinaryAuthorization
		*out = new(ClusterBinaryAuthorization)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	in.Fleet.DeepCopyInto(&out.Fleet)
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(ClusterLoggingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitoringConfig != nil {
		in, out := &in.MonitoringConfig, &out.MonitoringConfig
		*out = new(ClusterMonitoringConfig)
		(*in).DeepCopyInto(*out)
	}
	in.OidcConfig.DeepCopyInto(&out.OidcConfig)
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAttachedClusterSpec.
func (in *ContainerAttachedClusterSpec) DeepCopy() *ContainerAttachedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerAttachedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAttachedClusterStatus) DeepCopyInto(out *ContainerAttachedClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ClusterRegion != nil {
		in, out := &in.ClusterRegion, &out.ClusterRegion
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]ClusterErrorsStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(ClusterObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Reconciling != nil {
		in, out := &in.Reconciling, &out.Reconciling
		*out = new(bool)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WorkloadIdentityConfig != nil {
		in, out := &in.WorkloadIdentityConfig, &out.WorkloadIdentityConfig
		*out = make([]ClusterWorkloadIdentityConfigStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAttachedClusterStatus.
func (in *ContainerAttachedClusterStatus) DeepCopy() *ContainerAttachedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerAttachedClusterStatus)
	in.DeepCopyInto(out)
	return out
}
