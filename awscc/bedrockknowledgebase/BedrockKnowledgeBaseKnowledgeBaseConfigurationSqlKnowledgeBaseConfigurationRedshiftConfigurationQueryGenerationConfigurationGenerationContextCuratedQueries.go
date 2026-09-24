// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfigurationGenerationContextCuratedQueries struct {
	// Question for the curated query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#natural_language BedrockKnowledgeBase#natural_language}
	NaturalLanguage *string `field:"optional" json:"naturalLanguage" yaml:"naturalLanguage"`
	// Answer for the curated query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#sql BedrockKnowledgeBase#sql}
	Sql *string `field:"optional" json:"sql" yaml:"sql"`
}

