// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfiguration struct {
	// The Amazon Resource Name (ARN) of the OpenSearch domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#domain_arn BedrockKnowledgeBase#domain_arn}
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// The endpoint URL the OpenSearch domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#domain_endpoint BedrockKnowledgeBase#domain_endpoint}
	DomainEndpoint *string `field:"optional" json:"domainEndpoint" yaml:"domainEndpoint"`
	// A mapping of Bedrock Knowledge Base fields to OpenSearch Managed Cluster field names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#field_mapping BedrockKnowledgeBase#field_mapping}
	FieldMapping *BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationFieldMapping `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// The name of the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrock_knowledge_base#vector_index_name BedrockKnowledgeBase#vector_index_name}
	VectorIndexName *string `field:"optional" json:"vectorIndexName" yaml:"vectorIndexName"`
}

