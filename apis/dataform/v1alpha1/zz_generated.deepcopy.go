//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataformRepository) DeepCopyInto(out *DataformRepository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataformRepository.
func (in *DataformRepository) DeepCopy() *DataformRepository {
	if in == nil {
		return nil
	}
	out := new(DataformRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataformRepository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataformRepositoryList) DeepCopyInto(out *DataformRepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataformRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataformRepositoryList.
func (in *DataformRepositoryList) DeepCopy() *DataformRepositoryList {
	if in == nil {
		return nil
	}
	out := new(DataformRepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataformRepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataformRepositorySpec) DeepCopyInto(out *DataformRepositorySpec) {
	*out = *in
	if in.GitRemoteSettings != nil {
		in, out := &in.GitRemoteSettings, &out.GitRemoteSettings
		*out = new(RepositoryGitRemoteSettings)
		(*in).DeepCopyInto(*out)
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
	if in.WorkspaceCompilationOverrides != nil {
		in, out := &in.WorkspaceCompilationOverrides, &out.WorkspaceCompilationOverrides
		*out = new(RepositoryWorkspaceCompilationOverrides)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataformRepositorySpec.
func (in *DataformRepositorySpec) DeepCopy() *DataformRepositorySpec {
	if in == nil {
		return nil
	}
	out := new(DataformRepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataformRepositoryStatus) DeepCopyInto(out *DataformRepositoryStatus) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataformRepositoryStatus.
func (in *DataformRepositoryStatus) DeepCopy() *DataformRepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(DataformRepositoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryGitRemoteSettings) DeepCopyInto(out *RepositoryGitRemoteSettings) {
	*out = *in
	if in.TokenStatus != nil {
		in, out := &in.TokenStatus, &out.TokenStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryGitRemoteSettings.
func (in *RepositoryGitRemoteSettings) DeepCopy() *RepositoryGitRemoteSettings {
	if in == nil {
		return nil
	}
	out := new(RepositoryGitRemoteSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryWorkspaceCompilationOverrides) DeepCopyInto(out *RepositoryWorkspaceCompilationOverrides) {
	*out = *in
	if in.DefaultDatabase != nil {
		in, out := &in.DefaultDatabase, &out.DefaultDatabase
		*out = new(string)
		**out = **in
	}
	if in.SchemaSuffix != nil {
		in, out := &in.SchemaSuffix, &out.SchemaSuffix
		*out = new(string)
		**out = **in
	}
	if in.TablePrefix != nil {
		in, out := &in.TablePrefix, &out.TablePrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryWorkspaceCompilationOverrides.
func (in *RepositoryWorkspaceCompilationOverrides) DeepCopy() *RepositoryWorkspaceCompilationOverrides {
	if in == nil {
		return nil
	}
	out := new(RepositoryWorkspaceCompilationOverrides)
	in.DeepCopyInto(out)
	return out
}
