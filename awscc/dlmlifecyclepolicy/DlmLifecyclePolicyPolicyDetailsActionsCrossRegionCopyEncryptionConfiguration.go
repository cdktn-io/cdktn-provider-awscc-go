// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsActionsCrossRegionCopyEncryptionConfiguration struct {
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#cmk_arn DlmLifecyclePolicy#cmk_arn}
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is `false` or when encryption by default is not enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#encrypted DlmLifecyclePolicy#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
}

