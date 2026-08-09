// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicystore


type VerifiedpermissionsPolicyStoreEncryptionSettingsKmsEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/verifiedpermissions_policy_store#encryption_context VerifiedpermissionsPolicyStore#encryption_context}.
	EncryptionContext *map[string]*string `field:"optional" json:"encryptionContext" yaml:"encryptionContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/verifiedpermissions_policy_store#key VerifiedpermissionsPolicyStore#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

