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

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildTrigger) DeepCopyInto(out *CloudBuildTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildTrigger.
func (in *CloudBuildTrigger) DeepCopy() *CloudBuildTrigger {
	if in == nil {
		return nil
	}
	out := new(CloudBuildTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudBuildTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildTriggerList) DeepCopyInto(out *CloudBuildTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudBuildTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildTriggerList.
func (in *CloudBuildTriggerList) DeepCopy() *CloudBuildTriggerList {
	if in == nil {
		return nil
	}
	out := new(CloudBuildTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudBuildTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildTriggerSpec) DeepCopyInto(out *CloudBuildTriggerSpec) {
	*out = *in
	if in.ApprovalConfig != nil {
		in, out := &in.ApprovalConfig, &out.ApprovalConfig
		*out = new(TriggerApprovalConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.BitbucketServerTriggerConfig != nil {
		in, out := &in.BitbucketServerTriggerConfig, &out.BitbucketServerTriggerConfig
		*out = new(TriggerBitbucketServerTriggerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(TriggerBuild)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.GitFileSource != nil {
		in, out := &in.GitFileSource, &out.GitFileSource
		*out = new(TriggerGitFileSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(TriggerGithub)
		(*in).DeepCopyInto(*out)
	}
	if in.IgnoredFiles != nil {
		in, out := &in.IgnoredFiles, &out.IgnoredFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeBuildLogs != nil {
		in, out := &in.IncludeBuildLogs, &out.IncludeBuildLogs
		*out = new(string)
		**out = **in
	}
	if in.IncludedFiles != nil {
		in, out := &in.IncludedFiles, &out.IncludedFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PubsubConfig != nil {
		in, out := &in.PubsubConfig, &out.PubsubConfig
		*out = new(TriggerPubsubConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RepositoryEventConfig != nil {
		in, out := &in.RepositoryEventConfig, &out.RepositoryEventConfig
		*out = new(TriggerRepositoryEventConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.SourceToBuild != nil {
		in, out := &in.SourceToBuild, &out.SourceToBuild
		*out = new(TriggerSourceToBuild)
		(*in).DeepCopyInto(*out)
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TriggerTemplate != nil {
		in, out := &in.TriggerTemplate, &out.TriggerTemplate
		*out = new(TriggerTriggerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookConfig != nil {
		in, out := &in.WebhookConfig, &out.WebhookConfig
		*out = new(TriggerWebhookConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildTriggerSpec.
func (in *CloudBuildTriggerSpec) DeepCopy() *CloudBuildTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(CloudBuildTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildTriggerStatus) DeepCopyInto(out *CloudBuildTriggerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.TriggerId != nil {
		in, out := &in.TriggerId, &out.TriggerId
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildTriggerStatus.
func (in *CloudBuildTriggerStatus) DeepCopy() *CloudBuildTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(CloudBuildTriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPool) DeepCopyInto(out *CloudBuildWorkerPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPool.
func (in *CloudBuildWorkerPool) DeepCopy() *CloudBuildWorkerPool {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudBuildWorkerPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPoolList) DeepCopyInto(out *CloudBuildWorkerPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudBuildWorkerPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPoolList.
func (in *CloudBuildWorkerPoolList) DeepCopy() *CloudBuildWorkerPoolList {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudBuildWorkerPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPoolSpec) DeepCopyInto(out *CloudBuildWorkerPoolSpec) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	in.PrivatePoolV1Config.DeepCopyInto(&out.PrivatePoolV1Config)
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPoolSpec.
func (in *CloudBuildWorkerPoolSpec) DeepCopy() *CloudBuildWorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudBuildWorkerPoolStatus) DeepCopyInto(out *CloudBuildWorkerPoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
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
		*out = new(WorkerpoolObservedStateStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudBuildWorkerPoolStatus.
func (in *CloudBuildWorkerPoolStatus) DeepCopy() *CloudBuildWorkerPoolStatus {
	if in == nil {
		return nil
	}
	out := new(CloudBuildWorkerPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerApprovalConfig) DeepCopyInto(out *TriggerApprovalConfig) {
	*out = *in
	if in.ApprovalRequired != nil {
		in, out := &in.ApprovalRequired, &out.ApprovalRequired
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerApprovalConfig.
func (in *TriggerApprovalConfig) DeepCopy() *TriggerApprovalConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerApprovalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerArtifacts) DeepCopyInto(out *TriggerArtifacts) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = new(TriggerObjects)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerArtifacts.
func (in *TriggerArtifacts) DeepCopy() *TriggerArtifacts {
	if in == nil {
		return nil
	}
	out := new(TriggerArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAvailableSecrets) DeepCopyInto(out *TriggerAvailableSecrets) {
	*out = *in
	if in.SecretManager != nil {
		in, out := &in.SecretManager, &out.SecretManager
		*out = make([]TriggerSecretManager, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAvailableSecrets.
func (in *TriggerAvailableSecrets) DeepCopy() *TriggerAvailableSecrets {
	if in == nil {
		return nil
	}
	out := new(TriggerAvailableSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerBitbucketServerTriggerConfig) DeepCopyInto(out *TriggerBitbucketServerTriggerConfig) {
	*out = *in
	out.BitbucketServerConfigResourceRef = in.BitbucketServerConfigResourceRef
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(TriggerPullRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = new(TriggerPush)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerBitbucketServerTriggerConfig.
func (in *TriggerBitbucketServerTriggerConfig) DeepCopy() *TriggerBitbucketServerTriggerConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerBitbucketServerTriggerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerBuild) DeepCopyInto(out *TriggerBuild) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = new(TriggerArtifacts)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailableSecrets != nil {
		in, out := &in.AvailableSecrets, &out.AvailableSecrets
		*out = new(TriggerAvailableSecrets)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogsBucketRef != nil {
		in, out := &in.LogsBucketRef, &out.LogsBucketRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(TriggerOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.QueueTtl != nil {
		in, out := &in.QueueTtl, &out.QueueTtl
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = make([]TriggerSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(TriggerSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = make([]TriggerStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerBuild.
func (in *TriggerBuild) DeepCopy() *TriggerBuild {
	if in == nil {
		return nil
	}
	out := new(TriggerBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerGitFileSource) DeepCopyInto(out *TriggerGitFileSource) {
	*out = *in
	if in.BitbucketServerConfigRef != nil {
		in, out := &in.BitbucketServerConfigRef, &out.BitbucketServerConfigRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.GithubEnterpriseConfigRef != nil {
		in, out := &in.GithubEnterpriseConfigRef, &out.GithubEnterpriseConfigRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.RepositoryRef != nil {
		in, out := &in.RepositoryRef, &out.RepositoryRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(string)
		**out = **in
	}
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerGitFileSource.
func (in *TriggerGitFileSource) DeepCopy() *TriggerGitFileSource {
	if in == nil {
		return nil
	}
	out := new(TriggerGitFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerGithub) DeepCopyInto(out *TriggerGithub) {
	*out = *in
	if in.EnterpriseConfigResourceNameRef != nil {
		in, out := &in.EnterpriseConfigResourceNameRef, &out.EnterpriseConfigResourceNameRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(TriggerPullRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = new(TriggerPush)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerGithub.
func (in *TriggerGithub) DeepCopy() *TriggerGithub {
	if in == nil {
		return nil
	}
	out := new(TriggerGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerObjects) DeepCopyInto(out *TriggerObjects) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timing != nil {
		in, out := &in.Timing, &out.Timing
		*out = make([]TriggerTiming, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerObjects.
func (in *TriggerObjects) DeepCopy() *TriggerObjects {
	if in == nil {
		return nil
	}
	out := new(TriggerObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerOptions) DeepCopyInto(out *TriggerOptions) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.DynamicSubstitutions != nil {
		in, out := &in.DynamicSubstitutions, &out.DynamicSubstitutions
		*out = new(bool)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogStreamingOption != nil {
		in, out := &in.LogStreamingOption, &out.LogStreamingOption
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(string)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.RequestedVerifyOption != nil {
		in, out := &in.RequestedVerifyOption, &out.RequestedVerifyOption
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceProvenanceHash != nil {
		in, out := &in.SourceProvenanceHash, &out.SourceProvenanceHash
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubstitutionOption != nil {
		in, out := &in.SubstitutionOption, &out.SubstitutionOption
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]TriggerVolumes, len(*in))
		copy(*out, *in)
	}
	if in.WorkerPool != nil {
		in, out := &in.WorkerPool, &out.WorkerPool
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerOptions.
func (in *TriggerOptions) DeepCopy() *TriggerOptions {
	if in == nil {
		return nil
	}
	out := new(TriggerOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerPubsubConfig) DeepCopyInto(out *TriggerPubsubConfig) {
	*out = *in
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(string)
		**out = **in
	}
	out.TopicRef = in.TopicRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerPubsubConfig.
func (in *TriggerPubsubConfig) DeepCopy() *TriggerPubsubConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerPubsubConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerPullRequest) DeepCopyInto(out *TriggerPullRequest) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.CommentControl != nil {
		in, out := &in.CommentControl, &out.CommentControl
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerPullRequest.
func (in *TriggerPullRequest) DeepCopy() *TriggerPullRequest {
	if in == nil {
		return nil
	}
	out := new(TriggerPullRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerPush) DeepCopyInto(out *TriggerPush) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerPush.
func (in *TriggerPush) DeepCopy() *TriggerPush {
	if in == nil {
		return nil
	}
	out := new(TriggerPush)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerRepoSource) DeepCopyInto(out *TriggerRepoSource) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.ProjectId != nil {
		in, out := &in.ProjectId, &out.ProjectId
		*out = new(string)
		**out = **in
	}
	out.RepoRef = in.RepoRef
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagName != nil {
		in, out := &in.TagName, &out.TagName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerRepoSource.
func (in *TriggerRepoSource) DeepCopy() *TriggerRepoSource {
	if in == nil {
		return nil
	}
	out := new(TriggerRepoSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerRepositoryEventConfig) DeepCopyInto(out *TriggerRepositoryEventConfig) {
	*out = *in
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(TriggerPullRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = new(TriggerPush)
		(*in).DeepCopyInto(*out)
	}
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerRepositoryEventConfig.
func (in *TriggerRepositoryEventConfig) DeepCopy() *TriggerRepositoryEventConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerRepositoryEventConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSecret) DeepCopyInto(out *TriggerSecret) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSecret.
func (in *TriggerSecret) DeepCopy() *TriggerSecret {
	if in == nil {
		return nil
	}
	out := new(TriggerSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSecretManager) DeepCopyInto(out *TriggerSecretManager) {
	*out = *in
	out.VersionRef = in.VersionRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSecretManager.
func (in *TriggerSecretManager) DeepCopy() *TriggerSecretManager {
	if in == nil {
		return nil
	}
	out := new(TriggerSecretManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSource) DeepCopyInto(out *TriggerSource) {
	*out = *in
	if in.RepoSource != nil {
		in, out := &in.RepoSource, &out.RepoSource
		*out = new(TriggerRepoSource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageSource != nil {
		in, out := &in.StorageSource, &out.StorageSource
		*out = new(TriggerStorageSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSource.
func (in *TriggerSource) DeepCopy() *TriggerSource {
	if in == nil {
		return nil
	}
	out := new(TriggerSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSourceToBuild) DeepCopyInto(out *TriggerSourceToBuild) {
	*out = *in
	if in.BitbucketServerConfigRef != nil {
		in, out := &in.BitbucketServerConfigRef, &out.BitbucketServerConfigRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.GithubEnterpriseConfigRef != nil {
		in, out := &in.GithubEnterpriseConfigRef, &out.GithubEnterpriseConfigRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.RepositoryRef != nil {
		in, out := &in.RepositoryRef, &out.RepositoryRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSourceToBuild.
func (in *TriggerSourceToBuild) DeepCopy() *TriggerSourceToBuild {
	if in == nil {
		return nil
	}
	out := new(TriggerSourceToBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStep) DeepCopyInto(out *TriggerStep) {
	*out = *in
	if in.AllowExitCodes != nil {
		in, out := &in.AllowExitCodes, &out.AllowExitCodes
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.AllowFailure != nil {
		in, out := &in.AllowFailure, &out.AllowFailure
		*out = new(bool)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Script != nil {
		in, out := &in.Script, &out.Script
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	if in.Timing != nil {
		in, out := &in.Timing, &out.Timing
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]TriggerVolumes, len(*in))
		copy(*out, *in)
	}
	if in.WaitFor != nil {
		in, out := &in.WaitFor, &out.WaitFor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStep.
func (in *TriggerStep) DeepCopy() *TriggerStep {
	if in == nil {
		return nil
	}
	out := new(TriggerStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStorageSource) DeepCopyInto(out *TriggerStorageSource) {
	*out = *in
	out.BucketRef = in.BucketRef
	if in.Generation != nil {
		in, out := &in.Generation, &out.Generation
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStorageSource.
func (in *TriggerStorageSource) DeepCopy() *TriggerStorageSource {
	if in == nil {
		return nil
	}
	out := new(TriggerStorageSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTiming) DeepCopyInto(out *TriggerTiming) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTiming.
func (in *TriggerTiming) DeepCopy() *TriggerTiming {
	if in == nil {
		return nil
	}
	out := new(TriggerTiming)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerTriggerTemplate) DeepCopyInto(out *TriggerTriggerTemplate) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.RepoRef != nil {
		in, out := &in.RepoRef, &out.RepoRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	if in.TagName != nil {
		in, out := &in.TagName, &out.TagName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerTriggerTemplate.
func (in *TriggerTriggerTemplate) DeepCopy() *TriggerTriggerTemplate {
	if in == nil {
		return nil
	}
	out := new(TriggerTriggerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerVolumes) DeepCopyInto(out *TriggerVolumes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerVolumes.
func (in *TriggerVolumes) DeepCopy() *TriggerVolumes {
	if in == nil {
		return nil
	}
	out := new(TriggerVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerWebhookConfig) DeepCopyInto(out *TriggerWebhookConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerWebhookConfig.
func (in *TriggerWebhookConfig) DeepCopy() *TriggerWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolNetworkConfig) DeepCopyInto(out *WorkerpoolNetworkConfig) {
	*out = *in
	if in.EgressOption != nil {
		in, out := &in.EgressOption, &out.EgressOption
		*out = new(string)
		**out = **in
	}
	if in.PeeredNetworkIPRange != nil {
		in, out := &in.PeeredNetworkIPRange, &out.PeeredNetworkIPRange
		*out = new(string)
		**out = **in
	}
	if in.PeeredNetworkRef != nil {
		in, out := &in.PeeredNetworkRef, &out.PeeredNetworkRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolNetworkConfig.
func (in *WorkerpoolNetworkConfig) DeepCopy() *WorkerpoolNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolNetworkConfigStatus) DeepCopyInto(out *WorkerpoolNetworkConfigStatus) {
	*out = *in
	if in.EgressOption != nil {
		in, out := &in.EgressOption, &out.EgressOption
		*out = new(string)
		**out = **in
	}
	if in.PeeredNetwork != nil {
		in, out := &in.PeeredNetwork, &out.PeeredNetwork
		*out = new(string)
		**out = **in
	}
	if in.PeeredNetworkIPRange != nil {
		in, out := &in.PeeredNetworkIPRange, &out.PeeredNetworkIPRange
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolNetworkConfigStatus.
func (in *WorkerpoolNetworkConfigStatus) DeepCopy() *WorkerpoolNetworkConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolNetworkConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolObservedStateStatus) DeepCopyInto(out *WorkerpoolObservedStateStatus) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(WorkerpoolNetworkConfigStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	in.WorkerConfig.DeepCopyInto(&out.WorkerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolObservedStateStatus.
func (in *WorkerpoolObservedStateStatus) DeepCopy() *WorkerpoolObservedStateStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolObservedStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolPrivatePoolV1Config) DeepCopyInto(out *WorkerpoolPrivatePoolV1Config) {
	*out = *in
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(WorkerpoolNetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	in.WorkerConfig.DeepCopyInto(&out.WorkerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolPrivatePoolV1Config.
func (in *WorkerpoolPrivatePoolV1Config) DeepCopy() *WorkerpoolPrivatePoolV1Config {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolPrivatePoolV1Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolWorkerConfig) DeepCopyInto(out *WorkerpoolWorkerConfig) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolWorkerConfig.
func (in *WorkerpoolWorkerConfig) DeepCopy() *WorkerpoolWorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolWorkerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerpoolWorkerConfigStatus) DeepCopyInto(out *WorkerpoolWorkerConfigStatus) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerpoolWorkerConfigStatus.
func (in *WorkerpoolWorkerConfigStatus) DeepCopy() *WorkerpoolWorkerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerpoolWorkerConfigStatus)
	in.DeepCopyInto(out)
	return out
}
