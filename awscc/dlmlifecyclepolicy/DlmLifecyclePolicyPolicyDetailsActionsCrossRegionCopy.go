// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsActionsCrossRegionCopy struct {
	// The encryption settings for the copied snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#encryption_configuration DlmLifecyclePolicy#encryption_configuration}
	EncryptionConfiguration *DlmLifecyclePolicyPolicyDetailsActionsCrossRegionCopyEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The retention rule that indicates how long the cross-Region snapshot or AMI copies are to be retained in the destination Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsActionsCrossRegionCopyRetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
	// The target Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#target DlmLifecyclePolicy#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

