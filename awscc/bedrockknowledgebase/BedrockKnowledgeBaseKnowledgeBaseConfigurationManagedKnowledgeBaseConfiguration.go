// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfiguration struct {
	// The ARN of the model used to create vector embeddings for the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#embedding_model_arn BedrockKnowledgeBase#embedding_model_arn}
	EmbeddingModelArn *string `field:"optional" json:"embeddingModelArn" yaml:"embeddingModelArn"`
	// The embeddings model configuration details for the vector model used in Knowledge Base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#embedding_model_configuration BedrockKnowledgeBase#embedding_model_configuration}
	EmbeddingModelConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationEmbeddingModelConfiguration `field:"optional" json:"embeddingModelConfiguration" yaml:"embeddingModelConfiguration"`
	// The type of embedding model to use for the managed knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#embedding_model_type BedrockKnowledgeBase#embedding_model_type}
	EmbeddingModelType *string `field:"optional" json:"embeddingModelType" yaml:"embeddingModelType"`
	// Contains details about the server-side encryption for the managed knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#server_side_encryption_configuration BedrockKnowledgeBase#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Configurations for supplemental data storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#supplemental_data_storage_configuration BedrockKnowledgeBase#supplemental_data_storage_configuration}
	SupplementalDataStorageConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfiguration `field:"optional" json:"supplementalDataStorageConfiguration" yaml:"supplementalDataStorageConfiguration"`
}

