// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfiguration struct {
	// The vector configuration details for the Bedrock embeddings model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#bedrock_embedding_model_configuration BedrockKnowledgeBase#bedrock_embedding_model_configuration}
	BedrockEmbeddingModelConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfiguration `field:"optional" json:"bedrockEmbeddingModelConfiguration" yaml:"bedrockEmbeddingModelConfiguration"`
}

