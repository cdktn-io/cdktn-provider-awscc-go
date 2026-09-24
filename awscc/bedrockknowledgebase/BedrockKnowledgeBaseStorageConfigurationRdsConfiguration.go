// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationRdsConfiguration struct {
	// The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#credentials_secret_arn BedrockKnowledgeBase#credentials_secret_arn}
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The name of your Amazon RDS database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#database_name BedrockKnowledgeBase#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Contains the names of the fields to which to map information about the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#field_mapping BedrockKnowledgeBase#field_mapping}
	FieldMapping *BedrockKnowledgeBaseStorageConfigurationRdsConfigurationFieldMapping `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// The ARN of the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#resource_arn BedrockKnowledgeBase#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The name of the table in the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#table_name BedrockKnowledgeBase#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

