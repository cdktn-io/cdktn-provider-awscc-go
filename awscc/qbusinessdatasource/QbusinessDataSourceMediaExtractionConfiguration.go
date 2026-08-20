// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource


type QbusinessDataSourceMediaExtractionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#audio_extraction_configuration QbusinessDataSource#audio_extraction_configuration}.
	AudioExtractionConfiguration *QbusinessDataSourceMediaExtractionConfigurationAudioExtractionConfiguration `field:"optional" json:"audioExtractionConfiguration" yaml:"audioExtractionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#image_extraction_configuration QbusinessDataSource#image_extraction_configuration}.
	ImageExtractionConfiguration *QbusinessDataSourceMediaExtractionConfigurationImageExtractionConfiguration `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#video_extraction_configuration QbusinessDataSource#video_extraction_configuration}.
	VideoExtractionConfiguration *QbusinessDataSourceMediaExtractionConfigurationVideoExtractionConfiguration `field:"optional" json:"videoExtractionConfiguration" yaml:"videoExtractionConfiguration"`
}

