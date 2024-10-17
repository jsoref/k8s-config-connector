//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPersistenceConfig) DeepCopyInto(out *ClusterPersistenceConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.RdbConfig != nil {
		in, out := &in.RdbConfig, &out.RdbConfig
		*out = new(ClusterPersistenceConfig_RDBConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AofConfig != nil {
		in, out := &in.AofConfig, &out.AofConfig
		*out = new(ClusterPersistenceConfig_AOFConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPersistenceConfig.
func (in *ClusterPersistenceConfig) DeepCopy() *ClusterPersistenceConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterPersistenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPersistenceConfig_AOFConfig) DeepCopyInto(out *ClusterPersistenceConfig_AOFConfig) {
	*out = *in
	if in.AppendFsync != nil {
		in, out := &in.AppendFsync, &out.AppendFsync
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPersistenceConfig_AOFConfig.
func (in *ClusterPersistenceConfig_AOFConfig) DeepCopy() *ClusterPersistenceConfig_AOFConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterPersistenceConfig_AOFConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPersistenceConfig_RDBConfig) DeepCopyInto(out *ClusterPersistenceConfig_RDBConfig) {
	*out = *in
	if in.RdbSnapshotPeriod != nil {
		in, out := &in.RdbSnapshotPeriod, &out.RdbSnapshotPeriod
		*out = new(string)
		**out = **in
	}
	if in.RdbSnapshotStartTime != nil {
		in, out := &in.RdbSnapshotStartTime, &out.RdbSnapshotStartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPersistenceConfig_RDBConfig.
func (in *ClusterPersistenceConfig_RDBConfig) DeepCopy() *ClusterPersistenceConfig_RDBConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterPersistenceConfig_RDBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster_StateInfo) DeepCopyInto(out *Cluster_StateInfo) {
	*out = *in
	if in.UpdateInfo != nil {
		in, out := &in.UpdateInfo, &out.UpdateInfo
		*out = new(Cluster_StateInfo_UpdateInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster_StateInfo.
func (in *Cluster_StateInfo) DeepCopy() *Cluster_StateInfo {
	if in == nil {
		return nil
	}
	out := new(Cluster_StateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster_StateInfo_UpdateInfo) DeepCopyInto(out *Cluster_StateInfo_UpdateInfo) {
	*out = *in
	if in.TargetShardCount != nil {
		in, out := &in.TargetShardCount, &out.TargetShardCount
		*out = new(int32)
		**out = **in
	}
	if in.TargetReplicaCount != nil {
		in, out := &in.TargetReplicaCount, &out.TargetReplicaCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster_StateInfo_UpdateInfo.
func (in *Cluster_StateInfo_UpdateInfo) DeepCopy() *Cluster_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := new(Cluster_StateInfo_UpdateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryEndpoint) DeepCopyInto(out *DiscoveryEndpoint) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PscConfig != nil {
		in, out := &in.PscConfig, &out.PscConfig
		*out = new(PscConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryEndpoint.
func (in *DiscoveryEndpoint) DeepCopy() *DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := new(DiscoveryEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PscConfig) DeepCopyInto(out *PscConfig) {
	*out = *in
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PscConfig.
func (in *PscConfig) DeepCopy() *PscConfig {
	if in == nil {
		return nil
	}
	out := new(PscConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PscConfigSpec) DeepCopyInto(out *PscConfigSpec) {
	*out = *in
	if in.NetworkRef != nil {
		in, out := &in.NetworkRef, &out.NetworkRef
		*out = new(refsv1beta1.ComputeNetworkRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PscConfigSpec.
func (in *PscConfigSpec) DeepCopy() *PscConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PscConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PscConnection) DeepCopyInto(out *PscConnection) {
	*out = *in
	if in.PscConnectionID != nil {
		in, out := &in.PscConnectionID, &out.PscConnectionID
		*out = new(string)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.ForwardingRule != nil {
		in, out := &in.ForwardingRule, &out.ForwardingRule
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PscConnection.
func (in *PscConnection) DeepCopy() *PscConnection {
	if in == nil {
		return nil
	}
	out := new(PscConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCluster) DeepCopyInto(out *RedisCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCluster.
func (in *RedisCluster) DeepCopy() *RedisCluster {
	if in == nil {
		return nil
	}
	out := new(RedisCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterList) DeepCopyInto(out *RedisClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterList.
func (in *RedisClusterList) DeepCopy() *RedisClusterList {
	if in == nil {
		return nil
	}
	out := new(RedisClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterObservedState) DeepCopyInto(out *RedisClusterObservedState) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
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
	if in.SizeGb != nil {
		in, out := &in.SizeGb, &out.SizeGb
		*out = new(int32)
		**out = **in
	}
	if in.DiscoveryEndpoints != nil {
		in, out := &in.DiscoveryEndpoints, &out.DiscoveryEndpoints
		*out = make([]DiscoveryEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PscConnections != nil {
		in, out := &in.PscConnections, &out.PscConnections
		*out = make([]PscConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StateInfo != nil {
		in, out := &in.StateInfo, &out.StateInfo
		*out = new(Cluster_StateInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.PreciseSizeGb != nil {
		in, out := &in.PreciseSizeGb, &out.PreciseSizeGb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterObservedState.
func (in *RedisClusterObservedState) DeepCopy() *RedisClusterObservedState {
	if in == nil {
		return nil
	}
	out := new(RedisClusterObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterSpec) DeepCopyInto(out *RedisClusterSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.AuthorizationMode != nil {
		in, out := &in.AuthorizationMode, &out.AuthorizationMode
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionMode != nil {
		in, out := &in.TransitEncryptionMode, &out.TransitEncryptionMode
		*out = new(string)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int32)
		**out = **in
	}
	if in.PscConfigs != nil {
		in, out := &in.PscConfigs, &out.PscConfigs
		*out = make([]PscConfigSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.PersistenceConfig != nil {
		in, out := &in.PersistenceConfig, &out.PersistenceConfig
		*out = new(ClusterPersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisConfigs != nil {
		in, out := &in.RedisConfigs, &out.RedisConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicaCount != nil {
		in, out := &in.ReplicaCount, &out.ReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.ZoneDistributionConfig != nil {
		in, out := &in.ZoneDistributionConfig, &out.ZoneDistributionConfig
		*out = new(ZoneDistributionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterSpec.
func (in *RedisClusterSpec) DeepCopy() *RedisClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RedisClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisClusterStatus) DeepCopyInto(out *RedisClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
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
		*out = new(RedisClusterObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisClusterStatus.
func (in *RedisClusterStatus) DeepCopy() *RedisClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RedisClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneDistributionConfig) DeepCopyInto(out *ZoneDistributionConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneDistributionConfig.
func (in *ZoneDistributionConfig) DeepCopy() *ZoneDistributionConfig {
	if in == nil {
		return nil
	}
	out := new(ZoneDistributionConfig)
	in.DeepCopyInto(out)
	return out
}