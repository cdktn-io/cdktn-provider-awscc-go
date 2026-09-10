// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfiguration struct {
	// List of audio configurations for multi modal ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#audio BedrockKnowledgeBase#audio}
	Audio interface{} `field:"optional" json:"audio" yaml:"audio"`
	// The dimensions details for the vector configuration used on the Bedrock embeddings model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#dimensions BedrockKnowledgeBase#dimensions}
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The data type for the vectors when using a model to convert text into vector embeddings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#embedding_data_type BedrockKnowledgeBase#embedding_data_type}
	EmbeddingDataType *string `field:"optional" json:"embeddingDataType" yaml:"embeddingDataType"`
	// List of video configurations for multi modal ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#video BedrockKnowledgeBase#video}
	Video interface{} `field:"optional" json:"video" yaml:"video"`
}

