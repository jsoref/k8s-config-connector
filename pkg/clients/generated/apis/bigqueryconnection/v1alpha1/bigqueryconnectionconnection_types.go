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

type ConnectionAccessRole struct {
	/* The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection. */
	// +optional
	IamRoleID *string `json:"iamRoleID,omitempty"`
}

type ConnectionAws struct {
	/* Authentication using Google owned service account to assume into customer's AWS IAM Role. */
	// +optional
	AccessRole *ConnectionAccessRole `json:"accessRole,omitempty"`
}

type ConnectionCloudResource struct {
}

type ConnectionCloudSql struct {
	/* Cloud SQL credential. */
	// +optional
	Credential *ConnectionCredential `json:"credential,omitempty"`

	/* Database name. */
	// +optional
	Database *string `json:"database,omitempty"`

	/* Reference to the Cloud SQL instance ID. */
	// +optional
	InstanceRef *v1alpha1.ResourceRef `json:"instanceRef,omitempty"`

	/* Type of the Cloud SQL database. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type ConnectionCredential struct {
	/* The password for the credential. */
	// +optional
	Password *string `json:"password,omitempty"`

	/* The username for the credential. */
	// +optional
	Username *string `json:"username,omitempty"`
}

type BigQueryConnectionConnectionSpec struct {
	/* Amazon Web Services (AWS) properties. */
	// +optional
	Aws *ConnectionAws `json:"aws,omitempty"`

	/* Use Cloud Resource properties. */
	// +optional
	CloudResource *ConnectionCloudResource `json:"cloudResource,omitempty"`

	/* Cloud SQL properties. */
	// +optional
	CloudSql *ConnectionCloudSql `json:"cloudSql,omitempty"`

	/* User provided description. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* User provided display name for the connection. */
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	/* Immutable. */
	Location string `json:"location"`

	/* The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* The BigQuery ConnectionID. This is a server-generated ID in the UUID format. If not provided, ConfigConnector will create a new Connection and store the UUID in `status.serviceGeneratedID` field. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ConnectionAccessRoleStatus struct {
	/* A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's AWS IAM Role. */
	// +optional
	Identity *string `json:"identity,omitempty"`
}

type ConnectionAwsStatus struct {
	// +optional
	AccessRole *ConnectionAccessRoleStatus `json:"accessRole,omitempty"`
}

type ConnectionCloudResourceStatus struct {
	/* The account ID of the service created for the purpose of this
	connection.

	The service account does not have any permissions associated with it
	when it is created. After creation, customers delegate permissions
	to the service account. When the connection is used in the context of an
	operation in BigQuery, the service account will be used to connect to the
	desired resources in GCP.

	The account ID is in the form of:
	<service-1234>@gcp-sa-bigquery-cloudresource.iam.gserviceaccount.com */
	// +optional
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

type ConnectionCloudSqlStatus struct {
	/* The account ID of the service used for the purpose of this connection.

	When the connection is used in the context of an operation in
	BigQuery, this service account will serve as the identity being used for
	connecting to the CloudSQL instance specified in this connection. */
	// +optional
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

type ConnectionObservedStateStatus struct {
	// +optional
	Aws *ConnectionAwsStatus `json:"aws,omitempty"`

	// +optional
	CloudResource *ConnectionCloudResourceStatus `json:"cloudResource,omitempty"`

	// +optional
	CloudSql *ConnectionCloudSqlStatus `json:"cloudSql,omitempty"`

	/* The description for the connection. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The display name for the connection. */
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty"`

	/* Output only. True, if credential is configured for this connection. */
	// +optional
	HasCredential *bool `json:"hasCredential,omitempty"`
}

type BigQueryConnectionConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryConnectionConnection's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* A unique specifier for the BigQueryConnectionConnection resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *ConnectionObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryconnectionconnection;gcpbigqueryconnectionconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
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
