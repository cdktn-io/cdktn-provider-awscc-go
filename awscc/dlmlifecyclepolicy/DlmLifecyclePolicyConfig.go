// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DlmLifecyclePolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// **[Default policies only]** Indicates whether the policy should copy tags from the source resource to the snapshot or AMI.
	//
	// If you do not specify a value, the default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// **[Default policies only]** Specifies how often the policy should run and create snapshots or AMIs.
	//
	// The creation frequency can range from 1 to 7 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#create_interval DlmLifecyclePolicy#create_interval}
	CreateInterval *float64 `field:"optional" json:"createInterval" yaml:"createInterval"`
	// **[Default policies only]** Specifies destination Regions for snapshot or AMI copies.
	//
	// You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#cross_region_copy_targets DlmLifecyclePolicy#cross_region_copy_targets}
	CrossRegionCopyTargets interface{} `field:"optional" json:"crossRegionCopyTargets" yaml:"crossRegionCopyTargets"`
	// **[Default policies only]** Specify the type of default policy to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#default_policy DlmLifecyclePolicy#default_policy}
	DefaultPolicy *string `field:"optional" json:"defaultPolicy" yaml:"defaultPolicy"`
	// A description of the lifecycle policy. The characters ^[0-9A-Za-z _-]+$ are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#description DlmLifecyclePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// **[Default policies only]** Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs.
	//
	// The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#exclusions DlmLifecyclePolicy#exclusions}
	Exclusions *DlmLifecyclePolicyExclusions `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#execution_role_arn DlmLifecyclePolicy#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// **[Default policies only]** Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#extend_deletion DlmLifecyclePolicy#extend_deletion}
	ExtendDeletion interface{} `field:"optional" json:"extendDeletion" yaml:"extendDeletion"`
	// The configuration details of the lifecycle policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#policy_details DlmLifecyclePolicy#policy_details}
	PolicyDetails *DlmLifecyclePolicyPolicyDetails `field:"optional" json:"policyDetails" yaml:"policyDetails"`
	// **[Default policies only]** Specifies how long the policy should retain snapshots or AMIs before deleting them.
	//
	// The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#retain_interval DlmLifecyclePolicy#retain_interval}
	RetainInterval *float64 `field:"optional" json:"retainInterval" yaml:"retainInterval"`
	// The activation state of the lifecycle policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#state DlmLifecyclePolicy#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags to apply to the lifecycle policy during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#tags DlmLifecyclePolicy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

