// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfigurationProvisionedConfigurationAuthConfiguration struct {
	// Redshift database user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#database_user BedrockKnowledgeBase#database_user}
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// Provisioned Redshift auth type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#type BedrockKnowledgeBase#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Arn of a SecretsManager Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#username_password_secret_arn BedrockKnowledgeBase#username_password_secret_arn}
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

