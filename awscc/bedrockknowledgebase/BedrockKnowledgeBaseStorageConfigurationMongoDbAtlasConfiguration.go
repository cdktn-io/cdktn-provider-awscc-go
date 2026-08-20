// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfiguration struct {
	// Name of the collection within MongoDB Atlas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#collection_name BedrockKnowledgeBase#collection_name}
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon Mongo database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#credentials_secret_arn BedrockKnowledgeBase#credentials_secret_arn}
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Name of the database within MongoDB Atlas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#database_name BedrockKnowledgeBase#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// MongoDB Atlas endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#endpoint BedrockKnowledgeBase#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MongoDB Atlas endpoint service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#endpoint_service_name BedrockKnowledgeBase#endpoint_service_name}
	EndpointServiceName *string `field:"optional" json:"endpointServiceName" yaml:"endpointServiceName"`
	// Contains the names of the fields to which to map information about the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#field_mapping BedrockKnowledgeBase#field_mapping}
	FieldMapping *BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationFieldMapping `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// Name of a MongoDB Atlas text index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#text_index_name BedrockKnowledgeBase#text_index_name}
	TextIndexName *string `field:"optional" json:"textIndexName" yaml:"textIndexName"`
	// Name of a MongoDB Atlas index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#vector_index_name BedrockKnowledgeBase#vector_index_name}
	VectorIndexName *string `field:"optional" json:"vectorIndexName" yaml:"vectorIndexName"`
}

