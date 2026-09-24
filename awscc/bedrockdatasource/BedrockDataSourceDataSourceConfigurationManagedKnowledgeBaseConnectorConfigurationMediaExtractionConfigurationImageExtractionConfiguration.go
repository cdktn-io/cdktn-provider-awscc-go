// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration struct {
	// Indicates whether a feature is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#image_extraction_status BedrockDataSource#image_extraction_status}
	ImageExtractionStatus *string `field:"optional" json:"imageExtractionStatus" yaml:"imageExtractionStatus"`
}

