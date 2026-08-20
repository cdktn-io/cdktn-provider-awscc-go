// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadminorganizationcentralizationrule


type ObservabilityadminOrganizationCentralizationRuleRuleDestinationDestinationLogsConfigurationLogsEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_organization_centralization_rule#encryption_conflict_resolution_strategy ObservabilityadminOrganizationCentralizationRule#encryption_conflict_resolution_strategy}.
	EncryptionConflictResolutionStrategy *string `field:"optional" json:"encryptionConflictResolutionStrategy" yaml:"encryptionConflictResolutionStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_organization_centralization_rule#encryption_strategy ObservabilityadminOrganizationCentralizationRule#encryption_strategy}.
	EncryptionStrategy *string `field:"optional" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/observabilityadmin_organization_centralization_rule#kms_key_arn ObservabilityadminOrganizationCentralizationRule#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

