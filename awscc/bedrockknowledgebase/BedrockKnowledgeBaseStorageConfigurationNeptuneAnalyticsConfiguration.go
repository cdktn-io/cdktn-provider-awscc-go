// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfiguration struct {
	// A mapping of Bedrock Knowledge Base fields to Neptune Analytics fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_knowledge_base#field_mapping BedrockKnowledgeBase#field_mapping}
	FieldMapping *BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationFieldMapping `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// ARN for Neptune Analytics graph database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrock_knowledge_base#graph_arn BedrockKnowledgeBase#graph_arn}
	GraphArn *string `field:"optional" json:"graphArn" yaml:"graphArn"`
}

