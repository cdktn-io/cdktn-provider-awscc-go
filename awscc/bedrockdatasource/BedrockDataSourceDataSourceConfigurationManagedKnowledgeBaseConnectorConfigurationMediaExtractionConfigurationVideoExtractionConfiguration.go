// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration struct {
	// Indicates whether a feature is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#video_extraction_status BedrockDataSource#video_extraction_status}
	VideoExtractionStatus *string `field:"optional" json:"videoExtractionStatus" yaml:"videoExtractionStatus"`
}

