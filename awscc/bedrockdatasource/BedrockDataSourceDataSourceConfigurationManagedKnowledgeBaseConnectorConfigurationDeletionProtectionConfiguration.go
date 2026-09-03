// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationDeletionProtectionConfiguration struct {
	// Indicates whether a feature is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#deletion_protection_status BedrockDataSource#deletion_protection_status}
	DeletionProtectionStatus *string `field:"optional" json:"deletionProtectionStatus" yaml:"deletionProtectionStatus"`
	// Threshold for deletion protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_data_source#deletion_protection_threshold BedrockDataSource#deletion_protection_threshold}
	DeletionProtectionThreshold *float64 `field:"optional" json:"deletionProtectionThreshold" yaml:"deletionProtectionThreshold"`
}

