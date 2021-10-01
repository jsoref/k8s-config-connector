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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstancegroupmanagerAutoHealingPolicies struct {
	/*  */
	// +optional
	HealthCheckRef *v1alpha1.ResourceRef `json:"healthCheckRef,omitempty"`

	/* The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances. This initial delay allows instances to initialize and run their startup scripts before the instance group determines that they are UNHEALTHY. This prevents the managed instance group from recreating its instances prematurely. This value must be from range [0, 3600]. */
	// +optional
	InitialDelaySec *int `json:"initialDelaySec,omitempty"`
}

type InstancegroupmanagerDisks struct {
	/* These stateful disks will never be deleted during autohealing, update or VM instance recreate operations. This flag is used to configure if the disk should be deleted after it is no longer used by the group, e.g. when the given instance or the whole group is deleted. Note: disks attached in READ_ONLY mode cannot be auto-deleted. Possible values: NEVER, ON_PERMANENT_INSTANCE_DELETION */
	// +optional
	AutoDelete *string `json:"autoDelete,omitempty"`
}

type InstancegroupmanagerDistributionPolicy struct {
	/* The distribution shape to which the group converges either proactively or on resize events (depending on the value set in `updatePolicy.instanceRedistributionType`). Possible values: TARGET_SHAPE_UNSPECIFIED, ANY, BALANCED, ANY_SINGLE_ZONE */
	// +optional
	TargetShape *string `json:"targetShape,omitempty"`

	/* Zones where the regional managed instance group will create and manage its instances. */
	// +optional
	Zones []InstancegroupmanagerZones `json:"zones,omitempty"`
}

type InstancegroupmanagerMaxSurge struct {
	/* Specifies a fixed number of VM instances. This must be a positive integer. */
	// +optional
	Fixed *int `json:"fixed,omitempty"`

	/* Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify `80` for 80%. */
	// +optional
	Percent *int `json:"percent,omitempty"`
}

type InstancegroupmanagerMaxUnavailable struct {
	/* Specifies a fixed number of VM instances. This must be a positive integer. */
	// +optional
	Fixed *int `json:"fixed,omitempty"`

	/* Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify `80` for 80%. */
	// +optional
	Percent *int `json:"percent,omitempty"`
}

type InstancegroupmanagerNamedPorts struct {
	/* The name for this named port. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). */
	// +optional
	Name *string `json:"name,omitempty"`

	/* The port number, which can be a value between 1 and 65535. */
	// +optional
	Port *int `json:"port,omitempty"`
}

type InstancegroupmanagerPreservedState struct {
	/* Disks created on the instances that will be preserved on instance delete, update, etc. This map is keyed with the device names of the disks. */
	// +optional
	Disks map[string]InstancegroupmanagerDisks `json:"disks,omitempty"`
}

type InstancegroupmanagerStatefulPolicy struct {
	/*  */
	// +optional
	PreservedState *InstancegroupmanagerPreservedState `json:"preservedState,omitempty"`
}

type InstancegroupmanagerTargetSize struct {
	/* [Output Only] Absolute value of VM instances calculated based on the specific mode. - If the value is `fixed`, then the `calculated` value is equal to the `fixed` value. - If the value is a `percent`, then the `calculated` value is `percent`/100 * `targetSize`. For example, the `calculated` value of a 80% of a managed instance group with 150 instances would be (80/100 * 150) = 120 VM instances. If there is a remainder, the number is rounded. */
	// +optional
	Calculated *int `json:"calculated,omitempty"`

	/* Specifies a fixed number of VM instances. This must be a positive integer. */
	// +optional
	Fixed *int `json:"fixed,omitempty"`

	/* Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify `80` for 80%. */
	// +optional
	Percent *int `json:"percent,omitempty"`
}

type InstancegroupmanagerUpdatePolicy struct {
	/* The [instance redistribution policy](/compute/docs/instance-groups/regional-migs#proactive_instance_redistribution) for regional managed instance groups. Valid values are: - `PROACTIVE` (default): The group attempts to maintain an even distribution of VM instances across zones in the region. - `NONE`: For non-autoscaled groups, proactive redistribution is disabled. */
	// +optional
	InstanceRedistributionType *string `json:"instanceRedistributionType,omitempty"`

	/* Specifies the intended number of instances to be created from the `instanceTemplate`. The final number of instances created from the template will be equal to: - If expressed as a fixed number, the minimum of either `targetSize.fixed` or `instanceGroupManager.targetSize` is used. - if expressed as a `percent`, the `targetSize` would be `(targetSize.percent/100 * InstanceGroupManager.targetSize)` If there is a remainder, the number is rounded. If unset, this version will update any remaining instances not updated by another `version`. Read [Starting a canary update](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update) for more information. */
	// +optional
	MaxSurge *InstancegroupmanagerMaxSurge `json:"maxSurge,omitempty"`

	/* Specifies the intended number of instances to be created from the `instanceTemplate`. The final number of instances created from the template will be equal to: - If expressed as a fixed number, the minimum of either `targetSize.fixed` or `instanceGroupManager.targetSize` is used. - if expressed as a `percent`, the `targetSize` would be `(targetSize.percent/100 * InstanceGroupManager.targetSize)` If there is a remainder, the number is rounded. If unset, this version will update any remaining instances not updated by another `version`. Read [Starting a canary update](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update) for more information. */
	// +optional
	MaxUnavailable *InstancegroupmanagerMaxUnavailable `json:"maxUnavailable,omitempty"`

	/* Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600]. */
	// +optional
	MinReadySec *int `json:"minReadySec,omitempty"`

	/* Minimal action to be taken on an instance. You can specify either `RESTART` to restart existing instances or `REPLACE` to delete and create new instances from the target template. If you specify a `RESTART`, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action. */
	// +optional
	MinimalAction *string `json:"minimalAction,omitempty"`

	/* Most disruptive action that is allowed to be taken on an instance. You can specify either `NONE` to forbid any actions, `REFRESH` to allow actions that do not need instance restart, `RESTART` to allow actions that can be applied without instance replacing or `REPLACE` to allow all possible actions. If the Updater determines that the minimal update action needed is more disruptive than most disruptive allowed action you specify it will not perform the update at all. */
	// +optional
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty"`

	/* What action should be used to replace instances. See minimal_action.REPLACE Possible values: SUBSTITUTE, RECREATE */
	// +optional
	ReplacementMethod *string `json:"replacementMethod,omitempty"`

	/* The type of update process. You can specify either `PROACTIVE` so that the instance group manager proactively executes actions in order to bring instances to their target versions or `OPPORTUNISTIC` so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or `recreateInstances` calls). */
	// +optional
	Type *string `json:"type,omitempty"`
}

type InstancegroupmanagerVersions struct {
	/*  */
	// +optional
	InstanceTemplateRef *v1alpha1.ResourceRef `json:"instanceTemplateRef,omitempty"`

	/* Name of the version. Unique among all versions in the scope of this managed instance group. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Specifies the intended number of instances to be created from the `instanceTemplate`. The final number of instances created from the template will be equal to: - If expressed as a fixed number, the minimum of either `targetSize.fixed` or `instanceGroupManager.targetSize` is used. - if expressed as a `percent`, the `targetSize` would be `(targetSize.percent/100 * InstanceGroupManager.targetSize)` If there is a remainder, the number is rounded. If unset, this version will update any remaining instances not updated by another `version`. Read [Starting a canary update](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update) for more information. */
	// +optional
	TargetSize *InstancegroupmanagerTargetSize `json:"targetSize,omitempty"`
}

type InstancegroupmanagerZones struct {
	/* The URL of the [zone](/compute/docs/regions-zones/#available). The zone must exist in the region where the managed instance group is located. */
	// +optional
	Zone *string `json:"zone,omitempty"`
}

type ComputeInstanceGroupManagerSpec struct {
	/* The autohealing policy for this managed instance group. You can specify only one value. */
	// +optional
	AutoHealingPolicies []InstancegroupmanagerAutoHealingPolicies `json:"autoHealingPolicies,omitempty"`

	/* The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). */
	// +optional
	BaseInstanceName *string `json:"baseInstanceName,omitempty"`

	/* An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Policy specifying the intended distribution of managed instances across zones in a regional managed instance group. */
	// +optional
	DistributionPolicy *InstancegroupmanagerDistributionPolicy `json:"distributionPolicy,omitempty"`

	/* The action to perform in case of zone failure. Only one value is supported, `NO_FAILOVER`. The default is `NO_FAILOVER`. Possible values: UNKNOWN, NO_FAILOVER */
	// +optional
	FailoverAction *string `json:"failoverAction,omitempty"`

	/*  */
	// +optional
	InstanceTemplateRef *v1alpha1.ResourceRef `json:"instanceTemplateRef,omitempty"`

	/* The location of this resource. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Named ports configured for the Instance Groups complementary to this Instance Group Manager. */
	// +optional
	NamedPorts []InstancegroupmanagerNamedPorts `json:"namedPorts,omitempty"`

	/* The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/*  */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Stateful configuration for this Instanced Group Manager */
	// +optional
	StatefulPolicy *InstancegroupmanagerStatefulPolicy `json:"statefulPolicy,omitempty"`

	/*  */
	// +optional
	TargetPools []v1alpha1.ResourceRef `json:"targetPools,omitempty"`

	/* The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number. */
	// +optional
	TargetSize *int `json:"targetSize,omitempty"`

	/* The update policy for this managed instance group. */
	// +optional
	UpdatePolicy *InstancegroupmanagerUpdatePolicy `json:"updatePolicy,omitempty"`

	/* Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an `instanceTemplate` and a `name`. Every version can appear at most once per instance group. This field overrides the top-level `instanceTemplate` field. Read more about the [relationships between these fields](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#relationship_between_versions_and_instancetemplate_properties_for_a_managed_instance_group). Exactly one `version` must leave the `targetSize` field unset. That version will be applied to all remaining instances. For more information, read about [canary updates](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update). */
	// +optional
	Versions []InstancegroupmanagerVersions `json:"versions,omitempty"`
}

type InstancegroupmanagerCurrentActionsStatus struct {
	/* [Output Only] The total number of instances in the managed instance group that are scheduled to be abandoned. Abandoning an instance removes it from the managed instance group without deleting it. */
	Abandoning int `json:"abandoning,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are scheduled to be created or are currently being created. If the group fails to create any of these instances, it tries again until it creates the instance successfully. If you have disabled creation retries, this field will not be populated; instead, the `creatingWithoutRetries` field will be populated. */
	Creating int `json:"creating,omitempty"`

	/* [Output Only] The number of instances that the managed instance group will attempt to create. The group attempts to create each instance only once. If the group fails to create any of these instances, it decreases the group's `targetSize` value accordingly. */
	CreatingWithoutRetries int `json:"creatingWithoutRetries,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are scheduled to be deleted or are currently being deleted. */
	Deleting int `json:"deleting,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are running and have no scheduled actions. */
	None int `json:"none,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are scheduled to be recreated or are currently being being recreated. Recreating an instance deletes the existing root persistent disk and creates a new disk from the image that is defined in the instance template. */
	Recreating int `json:"recreating,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are being reconfigured with properties that do not require a restart or a recreate action. For example, setting or removing target pools for the instance. */
	Refreshing int `json:"refreshing,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are scheduled to be restarted or are currently being restarted. */
	Restarting int `json:"restarting,omitempty"`

	/* [Output Only] The number of instances in the managed instance group that are being verified. See the `managedInstances[].currentAction` property in the `listManagedInstances` method documentation. */
	Verifying int `json:"verifying,omitempty"`
}

type InstancegroupmanagerMaxSurgeStatus struct {
	/* [Output Only] Absolute value of VM instances calculated based on the specific mode. - If the value is `fixed`, then the `calculated` value is equal to the `fixed` value. - If the value is a `percent`, then the `calculated` value is `percent`/100 * `targetSize`. For example, the `calculated` value of a 80% of a managed instance group with 150 instances would be (80/100 * 150) = 120 VM instances. If there is a remainder, the number is rounded. */
	Calculated int `json:"calculated,omitempty"`
}

type InstancegroupmanagerMaxUnavailableStatus struct {
	/* [Output Only] Absolute value of VM instances calculated based on the specific mode. - If the value is `fixed`, then the `calculated` value is equal to the `fixed` value. - If the value is a `percent`, then the `calculated` value is `percent`/100 * `targetSize`. For example, the `calculated` value of a 80% of a managed instance group with 150 instances would be (80/100 * 150) = 120 VM instances. If there is a remainder, the number is rounded. */
	Calculated int `json:"calculated,omitempty"`
}

type InstancegroupmanagerPerInstanceConfigsStatus struct {
	/* A bit indicating if all of the group's per-instance configs (listed in the output of a listPerInstanceConfigs API call) have status `EFFECTIVE` or there are no per-instance-configs. */
	AllEffective bool `json:"allEffective,omitempty"`
}

type InstancegroupmanagerStatefulStatus struct {
	/* [Output Only] A bit indicating whether the managed instance group has stateful configuration, that is, if you have configured any items in a stateful policy or in per-instance configs. The group might report that it has no stateful config even when there is still some preserved state on a managed instance, for example, if you have deleted all PICs but not yet applied those deletions. */
	HasStatefulConfig bool `json:"hasStatefulConfig,omitempty"`

	/* [Output Only] A bit indicating whether the managed instance group has stateful configuration, that is, if you have configured any items in a stateful policy or in per-instance configs. The group might report that it has no stateful config even when there is still some preserved state on a managed instance, for example, if you have deleted all PICs but not yet applied those deletions. This field is deprecated in favor of has_stateful_config. */
	IsStateful bool `json:"isStateful,omitempty"`

	/* [Output Only] Status of per-instance configs on the instance. */
	PerInstanceConfigs InstancegroupmanagerPerInstanceConfigsStatus `json:"perInstanceConfigs,omitempty"`
}

type InstancegroupmanagerStatusStatus struct {
	/* [Output Only] The URL of the [Autoscaler](/compute/docs/autoscaler/) that targets this instance group manager. */
	Autoscaler string `json:"autoscaler,omitempty"`

	/* [Output Only] A bit indicating whether the managed instance group is in a stable state. A stable state means that: none of the instances in the managed instance group is currently undergoing any type of change (for example, creation, restart, or deletion); no future changes are scheduled for instances in the managed instance group; and the managed instance group itself is not being modified. */
	IsStable bool `json:"isStable,omitempty"`

	/* [Output Only] Stateful status of the given Instance Group Manager. */
	Stateful InstancegroupmanagerStatefulStatus `json:"stateful,omitempty"`

	/* [Output Only] A status of consistency of Instances' versions with their target version specified by `version` field on Instance Group Manager. */
	VersionTarget InstancegroupmanagerVersionTargetStatus `json:"versionTarget,omitempty"`
}

type InstancegroupmanagerUpdatePolicyStatus struct {
	/*  */
	MaxSurge InstancegroupmanagerMaxSurgeStatus `json:"maxSurge,omitempty"`

	/*  */
	MaxUnavailable InstancegroupmanagerMaxUnavailableStatus `json:"maxUnavailable,omitempty"`
}

type InstancegroupmanagerVersionTargetStatus struct {
	/* [Output Only] A bit indicating whether version target has been reached in this managed instance group, i.e. all instances are in their target version. Instances' target version are specified by `version` field on Instance Group Manager. */
	IsReached bool `json:"isReached,omitempty"`
}

type ComputeInstanceGroupManagerStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeInstanceGroupManager's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The creation timestamp for this managed instance group in \[RFC3339\](https://www.ietf.org/rfc/rfc3339.txt) text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* [Output Only] The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions. */
	CurrentActions InstancegroupmanagerCurrentActionsStatus `json:"currentActions,omitempty"`
	/* Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error `412 conditionNotMet`. To see the latest fingerprint, make a `get()` request to retrieve an InstanceGroupManager. */
	Fingerprint string `json:"fingerprint,omitempty"`
	/* [Output Only] A unique identifier for this resource type. The server generates this identifier. */
	Id int `json:"id,omitempty"`
	/* [Output Only] The URL of the Instance Group resource. */
	InstanceGroup string `json:"instanceGroup,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* [Output Only] The URL of the [region](/compute/docs/regions-zones/#available) where the managed instance group resides (for regional resources). */
	Region string `json:"region,omitempty"`
	/* [Output Only] The URL for this managed instance group. The server defines this URL. */
	SelfLink string `json:"selfLink,omitempty"`
	/* [Output Only] The status of this managed instance group. */
	Status InstancegroupmanagerStatusStatus `json:"status,omitempty"`
	/*  */
	UpdatePolicy InstancegroupmanagerUpdatePolicyStatus `json:"updatePolicy,omitempty"`
	/* [Output Only] The URL of a [zone](/compute/docs/regions-zones/#available) where the managed instance group is located (for zonal resources). */
	Zone string `json:"zone,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstanceGroupManager is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupManagerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeInstanceGroupManagerList contains a list of ComputeInstanceGroupManager
type ComputeInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceGroupManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceGroupManager{}, &ComputeInstanceGroupManagerList{})
}
