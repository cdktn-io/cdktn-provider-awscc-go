// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationKendraKnowledgeBaseConfiguration struct {
	// Arn of a Kendra index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrock_knowledge_base#kendra_index_arn BedrockKnowledgeBase#kendra_index_arn}
	KendraIndexArn *string `field:"optional" json:"kendraIndexArn" yaml:"kendraIndexArn"`
}

