// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedules struct {
	// **[Custom snapshot policies that target volumes only]** The snapshot archiving rule for the schedule.
	//
	// When you specify an archiving rule, snapshots are automatically moved from the standard tier to the archive tier once the schedule's retention threshold is met. Snapshots are then retained in the archive tier for the archive retention period that you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#archive_rule DlmLifecyclePolicy#archive_rule}
	ArchiveRule *DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRule `field:"optional" json:"archiveRule" yaml:"archiveRule"`
	// Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// The creation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#create_rule DlmLifecyclePolicy#create_rule}
	CreateRule *DlmLifecyclePolicyPolicyDetailsSchedulesCreateRule `field:"optional" json:"createRule" yaml:"createRule"`
	// Specifies a rule for copying snapshots or AMIs across Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#cross_region_copy_rules DlmLifecyclePolicy#cross_region_copy_rules}
	CrossRegionCopyRules interface{} `field:"optional" json:"crossRegionCopyRules" yaml:"crossRegionCopyRules"`
	// **[Custom AMI policies only]** The AMI deprecation rule for the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#deprecate_rule DlmLifecyclePolicy#deprecate_rule}
	DeprecateRule *DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRule `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// **[Custom snapshot policies only]** The rule for enabling fast snapshot restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#fast_restore_rule DlmLifecyclePolicy#fast_restore_rule}
	FastRestoreRule *DlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRule `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// The name of the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#name DlmLifecyclePolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The retention rule for snapshots or AMIs created by the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsSchedulesRetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
	// **[Custom snapshot policies only]** The rule for sharing snapshots with other AWS accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#share_rules DlmLifecyclePolicy#share_rules}
	ShareRules interface{} `field:"optional" json:"shareRules" yaml:"shareRules"`
	// The tags to apply to policy-created resources. These user-defined tags are in addition to the AWS-added lifecycle tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#tags_to_add DlmLifecyclePolicy#tags_to_add}
	TagsToAdd interface{} `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
	// **[AMI policies and snapshot policies that target instances only]** A collection of key/value pairs with values determined dynamically when the policy is executed.
	//
	// Keys may be any valid Amazon EC2 tag key. Values must be in one of the two following formats: `$(instance-id)` or `$(timestamp)`. Variable tags are only valid for EBS Snapshot Management -- Instance policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#variable_tags DlmLifecyclePolicy#variable_tags}
	VariableTags interface{} `field:"optional" json:"variableTags" yaml:"variableTags"`
}

