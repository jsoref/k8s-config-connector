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
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryPubsubConfigs) DeepCopyInto(out *RepositoryPubsubConfigs) {
	*out = *in
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1alpha1.ResourceRef)
		**out = **in
	}
	out.TopicRef = in.TopicRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryPubsubConfigs.
func (in *RepositoryPubsubConfigs) DeepCopy() *RepositoryPubsubConfigs {
	if in == nil {
		return nil
	}
	out := new(RepositoryPubsubConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRepoRepository) DeepCopyInto(out *SourceRepoRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRepoRepository.
func (in *SourceRepoRepository) DeepCopy() *SourceRepoRepository {
	if in == nil {
		return nil
	}
	out := new(SourceRepoRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceRepoRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRepoRepositoryList) DeepCopyInto(out *SourceRepoRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceRepoRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRepoRepositoryList.
func (in *SourceRepoRepositoryList) DeepCopy() *SourceRepoRepositoryList {
	if in == nil {
		return nil
	}
	out := new(SourceRepoRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceRepoRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRepoRepositorySpec) DeepCopyInto(out *SourceRepoRepositorySpec) {
	*out = *in
	if in.PubsubConfigs != nil {
		in, out := &in.PubsubConfigs, &out.PubsubConfigs
		*out = make([]RepositoryPubsubConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRepoRepositorySpec.
func (in *SourceRepoRepositorySpec) DeepCopy() *SourceRepoRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(SourceRepoRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRepoRepositoryStatus) DeepCopyInto(out *SourceRepoRepositoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRepoRepositoryStatus.
func (in *SourceRepoRepositoryStatus) DeepCopy() *SourceRepoRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(SourceRepoRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
