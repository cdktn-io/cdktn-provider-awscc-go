// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfigurationGenerationContextTables struct {
	// List of Redshift query generation columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#columns BedrockKnowledgeBase#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Description for the attached entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#description BedrockKnowledgeBase#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Include or Exclude status for an entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#inclusion BedrockKnowledgeBase#inclusion}
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
	// Query generation table name. Must follow three-part notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#name BedrockKnowledgeBase#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

