// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfiguration struct {
	// Configurations for Redshift query engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#query_engine_configuration BedrockKnowledgeBase#query_engine_configuration}
	QueryEngineConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfiguration `field:"optional" json:"queryEngineConfiguration" yaml:"queryEngineConfiguration"`
	// Configurations for generating Redshift engine queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#query_generation_configuration BedrockKnowledgeBase#query_generation_configuration}
	QueryGenerationConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfiguration `field:"optional" json:"queryGenerationConfiguration" yaml:"queryGenerationConfiguration"`
	// List of configurations for available Redshift query engine storage types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#storage_configurations BedrockKnowledgeBase#storage_configurations}
	StorageConfigurations interface{} `field:"optional" json:"storageConfigurations" yaml:"storageConfigurations"`
}

