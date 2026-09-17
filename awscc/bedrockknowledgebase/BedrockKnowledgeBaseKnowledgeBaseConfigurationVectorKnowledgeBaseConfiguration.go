// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfiguration struct {
	// The ARN of the model used to create vector embeddings for the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#embedding_model_arn BedrockKnowledgeBase#embedding_model_arn}
	EmbeddingModelArn *string `field:"optional" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// The embeddings model configuration details for the vector model used in Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#embedding_model_configuration BedrockKnowledgeBase#embedding_model_configuration}
	EmbeddingModelConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfiguration `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// Configurations for supplemental data storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#supplemental_data_storage_configuration BedrockKnowledgeBase#supplemental_data_storage_configuration}
	SupplementalDataStorageConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationSupplementalDataStorageConfiguration `field:"optional" json:"supplementalDataStorageConfiguration" yaml:"supplementalDataStorageConfiguration"`
}

