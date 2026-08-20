// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfigurationGenerationContextTablesColumns struct {
	// Description for the attached entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#description BedrockKnowledgeBase#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Include or Exclude status for an entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#inclusion BedrockKnowledgeBase#inclusion}
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
	// Query generation column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_knowledge_base#name BedrockKnowledgeBase#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

