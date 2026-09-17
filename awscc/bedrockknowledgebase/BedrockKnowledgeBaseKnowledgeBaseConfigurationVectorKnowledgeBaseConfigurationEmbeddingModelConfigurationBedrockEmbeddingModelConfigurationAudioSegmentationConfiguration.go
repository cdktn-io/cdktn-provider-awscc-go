// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfigurationAudioSegmentationConfiguration struct {
	// Duration in seconds to segment the multi modal media.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#fixed_length_duration BedrockKnowledgeBase#fixed_length_duration}
	FixedLengthDuration *float64 `field:"optional" json:"fixedLengthDuration" yaml:"fixedLengthDuration"`
}

