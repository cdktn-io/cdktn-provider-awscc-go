// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfiguration struct {
	// Connector-specific parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_source#connector_parameters BedrockDataSource#connector_parameters}
	ConnectorParameters *string `field:"optional" json:"connectorParameters" yaml:"connectorParameters"`
	// Configuration for deletion protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_source#deletion_protection_configuration BedrockDataSource#deletion_protection_configuration}
	DeletionProtectionConfiguration *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationDeletionProtectionConfiguration `field:"optional" json:"deletionProtectionConfiguration" yaml:"deletionProtectionConfiguration"`
	// Configuration for media extraction settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_data_source#media_extraction_configuration BedrockDataSource#media_extraction_configuration}
	MediaExtractionConfiguration *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
}

