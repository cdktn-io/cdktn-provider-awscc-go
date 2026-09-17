// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetails struct {
	// **[Event-based policies only]** The actions to be performed when the event-based policy is activated.
	//
	// You can specify only one action per policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#actions DlmLifecyclePolicy#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// **[Default policies only]** Indicates whether the policy should copy tags from the source resource to the snapshot or AMI.
	//
	// If you do not specify a value, the default is `false`.
	//
	// Default: `false`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// **[Default policies only]** Specifies how often the policy should run and create snapshots or AMIs.
	//
	// The creation frequency can range from 1 to 7 days. If you do not specify a value, the default is 1.
	//
	// Default: 1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#create_interval DlmLifecyclePolicy#create_interval}
	CreateInterval *float64 `field:"optional" json:"createInterval" yaml:"createInterval"`
	// **[Default policies only]** Specifies destination Regions for snapshot or AMI copies.
	//
	// You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#cross_region_copy_targets DlmLifecyclePolicy#cross_region_copy_targets}
	CrossRegionCopyTargets interface{} `field:"optional" json:"crossRegionCopyTargets" yaml:"crossRegionCopyTargets"`
	// **[Event-based policies only]** The event that activates the event-based policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#event_source DlmLifecyclePolicy#event_source}
	EventSource *DlmLifecyclePolicyPolicyDetailsEventSource `field:"optional" json:"eventSource" yaml:"eventSource"`
	// **[Default policies only]** Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs.
	//
	// The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#exclusions DlmLifecyclePolicy#exclusions}
	Exclusions *DlmLifecyclePolicyPolicyDetailsExclusions `field:"optional" json:"exclusions" yaml:"exclusions"`
	// **[Default policies only]** Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state.
	//
	// Default: `false`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#extend_deletion DlmLifecyclePolicy#extend_deletion}
	ExtendDeletion interface{} `field:"optional" json:"extendDeletion" yaml:"extendDeletion"`
	// **[Custom snapshot and AMI policies only]** A set of optional parameters for snapshot and AMI lifecycle policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#parameters DlmLifecyclePolicy#parameters}
	Parameters *DlmLifecyclePolicyPolicyDetailsParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// The type of policy to create. Specify one of the following:.
	//
	// - `SIMPLIFIED` -- To create a default policy.
	// - `STANDARD` -- To create a custom policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#policy_language DlmLifecyclePolicy#policy_language}
	PolicyLanguage *string `field:"optional" json:"policyLanguage" yaml:"policyLanguage"`
	// The type of policy.
	//
	// Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account.
	//
	// The default is `EBS_SNAPSHOT_MANAGEMENT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#policy_type DlmLifecyclePolicy#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// **[Custom snapshot and AMI policies only]** The location of the resources to backup.
	//
	// If the source resources are located in a Region, specify `CLOUD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#resource_locations DlmLifecyclePolicy#resource_locations}
	ResourceLocations *[]*string `field:"optional" json:"resourceLocations" yaml:"resourceLocations"`
	// **[Default policies only]** Specify the type of default policy to create.
	//
	// - To create a default policy for EBS snapshots, that creates snapshots of all volumes in the Region that do not have recent backups, specify `VOLUME`.
	// - To create a default policy for EBS-backed AMIs, that creates EBS-backed AMIs from all instances in the Region that do not have recent backups, specify `INSTANCE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#resource_type DlmLifecyclePolicy#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// **[Custom snapshot policies only]** The target resource type for snapshot and AMI lifecycle policies.
	//
	// Use `VOLUME` to create snapshots of individual volumes or use `INSTANCE` to create multi-volume snapshots from the volumes for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#resource_types DlmLifecyclePolicy#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// **[Default policies only]** Specifies how long the policy should retain snapshots or AMIs before deleting them.
	//
	// The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time. If you do not specify a value, the default is 7.
	//
	// Default: 7
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#retain_interval DlmLifecyclePolicy#retain_interval}
	RetainInterval *float64 `field:"optional" json:"retainInterval" yaml:"retainInterval"`
	// **[Custom snapshot and AMI policies only]** The schedules of policy-defined actions for snapshot and AMI lifecycle policies.
	//
	// A policy can have up to four schedules -- one mandatory schedule and up to three optional schedules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#schedules DlmLifecyclePolicy#schedules}
	Schedules interface{} `field:"optional" json:"schedules" yaml:"schedules"`
	// **[Custom snapshot and AMI policies only]** The single tag that identifies targeted resources for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#target_tags DlmLifecyclePolicy#target_tags}
	TargetTags interface{} `field:"optional" json:"targetTags" yaml:"targetTags"`
}

