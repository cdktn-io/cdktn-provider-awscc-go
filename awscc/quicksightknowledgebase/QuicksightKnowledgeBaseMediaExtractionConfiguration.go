// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightknowledgebase


type QuicksightKnowledgeBaseMediaExtractionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_knowledge_base#audio_extraction_configuration QuicksightKnowledgeBase#audio_extraction_configuration}.
	AudioExtractionConfiguration *QuicksightKnowledgeBaseMediaExtractionConfigurationAudioExtractionConfiguration `field:"optional" json:"audioExtractionConfiguration" yaml:"audioExtractionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_knowledge_base#image_extraction_configuration QuicksightKnowledgeBase#image_extraction_configuration}.
	ImageExtractionConfiguration *QuicksightKnowledgeBaseMediaExtractionConfigurationImageExtractionConfiguration `field:"optional" json:"imageExtractionConfiguration" yaml:"imageExtractionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_knowledge_base#video_extraction_configuration QuicksightKnowledgeBase#video_extraction_configuration}.
	VideoExtractionConfiguration *QuicksightKnowledgeBaseMediaExtractionConfigurationVideoExtractionConfiguration `field:"optional" json:"videoExtractionConfiguration" yaml:"videoExtractionConfiguration"`
}

