// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationPineconeConfiguration struct {
	// The endpoint URL for your index management page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#connection_string BedrockKnowledgeBase#connection_string}
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// The ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#credentials_secret_arn BedrockKnowledgeBase#credentials_secret_arn}
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Contains the names of the fields to which to map information about the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#field_mapping BedrockKnowledgeBase#field_mapping}
	FieldMapping *BedrockKnowledgeBaseStorageConfigurationPineconeConfigurationFieldMapping `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// The namespace to be used to write new data to your database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#namespace BedrockKnowledgeBase#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

