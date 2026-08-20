// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfigurationsRedshiftConfiguration struct {
	// Redshift database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#database_name BedrockKnowledgeBase#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

