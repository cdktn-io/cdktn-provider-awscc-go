// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRules struct {
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#cmk_arn DlmLifecyclePolicy#cmk_arn}
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Indicates whether to copy all user-defined tags from the source snapshot or AMI to the cross-Region copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#copy_tags DlmLifecyclePolicy#copy_tags}
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// **[Custom AMI policies only]** The AMI deprecation rule for cross-Region AMI copies created by the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#deprecate_rule DlmLifecyclePolicy#deprecate_rule}
	DeprecateRule *DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRule `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is `false` or if encryption by default is not enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#encrypted DlmLifecyclePolicy#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The retention rule that indicates how long the cross-Region snapshot or AMI copies are to be retained in the destination Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesRetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
	// **[Custom snapshot policies only]** The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#target DlmLifecyclePolicy#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// **[Custom AMI policies only]** The target Region or the Amazon Resource Name (ARN) of the target Outpost for the AMI copies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#target_region DlmLifecyclePolicy#target_region}
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

