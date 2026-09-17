// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration struct {
	// The ARN of the OpenSearch Service vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#collection_arn BedrockKnowledgeBase#collection_arn}
	CollectionArn *string `field:"optional" json:"collectionArn" yaml:"collectionArn"`
	// A mapping of Bedrock Knowledge Base fields to OpenSearch Serverless field names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#field_mapping BedrockKnowledgeBase#field_mapping}
	FieldMapping *BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMapping `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// The name of the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_knowledge_base#vector_index_name BedrockKnowledgeBase#vector_index_name}
	VectorIndexName *string `field:"optional" json:"vectorIndexName" yaml:"vectorIndexName"`
}

