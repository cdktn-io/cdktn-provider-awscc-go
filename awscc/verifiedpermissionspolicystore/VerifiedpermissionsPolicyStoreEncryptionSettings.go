// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicystore


type VerifiedpermissionsPolicyStoreEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/verifiedpermissions_policy_store#default VerifiedpermissionsPolicyStore#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/verifiedpermissions_policy_store#kms_encryption_settings VerifiedpermissionsPolicyStore#kms_encryption_settings}.
	KmsEncryptionSettings *VerifiedpermissionsPolicyStoreEncryptionSettingsKmsEncryptionSettings `field:"optional" json:"kmsEncryptionSettings" yaml:"kmsEncryptionSettings"`
}

