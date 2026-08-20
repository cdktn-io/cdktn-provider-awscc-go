// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aiopsinvestigationgroup


type AiopsInvestigationGroupEncryptionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/aiops_investigation_group#encryption_configuration_type AiopsInvestigationGroup#encryption_configuration_type}.
	EncryptionConfigurationType *string `field:"optional" json:"encryptionConfigurationType" yaml:"encryptionConfigurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/aiops_investigation_group#kms_key_id AiopsInvestigationGroup#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

