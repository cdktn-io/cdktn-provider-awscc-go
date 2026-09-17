// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfiguration struct {
	// Configuration for audio extraction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#audio_extraction_configuration BedrockDataSource#audio_extraction_configuration}
	AudioExtractionConfiguration *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationAudioExtractionConfiguration `field:"optional" json:"audioExtractionConfiguration" yaml:"audioExtractionConfiguration"`
	// Configuration for image extraction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#image_extraction_configuration BedrockDataSource#image_extraction_configuration}
	ImageExtractionConfiguration *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationImageExtractionConfiguration `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
	// Configuration for video extraction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#video_extraction_configuration BedrockDataSource#video_extraction_configuration}
	VideoExtractionConfiguration *BedrockDataSourceDataSourceConfigurationManagedKnowledgeBaseConnectorConfigurationMediaExtractionConfigurationVideoExtractionConfiguration `field:"optional" json:"videoExtractionConfiguration" yaml:"videoExtractionConfiguration"`
}

