// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConnectionCloudResource struct {
	/* The account ID of the service created for the purpose of this connection. */
	// +optional
	ServiceAccountId *string `json:"serviceAccountId,omitempty"`
}

type BigQueryConnectionConnectionSpec struct {
	/* Container for connection properties for delegation of access to GCP resources. */
	// +optional
	CloudResource *ConnectionCloudResource `json:"cloudResource,omitempty"`

	/* A descriptive description for the connection. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* A descriptive name for the connection. */
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	/* Immutable. The geographic location where the connection should reside.
	Cloud SQL instance must be in the same location as the connection
	with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	Spanner Connections same as spanner region
	AWS allowed regions are aws-us-east-1
	Azure allowed regions are azure-eastus2. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The connectionId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ConnectionCloudResourceStatus struct {
	/* The account ID of the service created for the purpose of this connection. */
	// +optional
	ServiceAccountId *string `json:"serviceAccountId,omitempty"`
}

type ConnectionObservedStateStatus struct {
	/* Container for connection properties for delegation of access to GCP resources. */
	// +optional
	CloudResource *ConnectionCloudResourceStatus `json:"cloudResource,omitempty"`
}

type BigQueryConnectionConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryConnectionConnection's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* True if the connection has credential assigned. */
	// +optional
	HasCredential *bool `json:"hasCredential,omitempty"`

	/* The resource name of the connection in the form of:
	"projects/{project_id}/locations/{location_id}/connections/{connectionId}". */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	// +optional
	ObservedState *ConnectionObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryconnectionconnection;gcpbigqueryconnectionconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryConnectionConnection is the Schema for the bigqueryconnection API
// +k8s:openapi-gen=true
type BigQueryConnectionConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryConnectionConnectionSpec   `json:"spec,omitempty"`
	Status BigQueryConnectionConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryConnectionConnectionList contains a list of BigQueryConnectionConnection
type BigQueryConnectionConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryConnectionConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryConnectionConnection{}, &BigQueryConnectionConnectionList{})
}
