// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationRdsConfigurationFieldMapping struct {
	// The name of the field in which Amazon Bedrock stores custom metadata about the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#custom_metadata_field BedrockKnowledgeBase#custom_metadata_field}
	CustomMetadataField *string `field:"optional" json:"customMetadataField" yaml:"customMetadataField"`
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#metadata_field BedrockKnowledgeBase#metadata_field}
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the ID for each entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#primary_key_field BedrockKnowledgeBase#primary_key_field}
	PrimaryKeyField *string `field:"optional" json:"primaryKeyField" yaml:"primaryKeyField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#text_field BedrockKnowledgeBase#text_field}
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_knowledge_base#vector_field BedrockKnowledgeBase#vector_field}
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
}

