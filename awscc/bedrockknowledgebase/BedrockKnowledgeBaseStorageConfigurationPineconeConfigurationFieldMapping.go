// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase


type BedrockKnowledgeBaseStorageConfigurationPineconeConfigurationFieldMapping struct {
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#metadata_field BedrockKnowledgeBase#metadata_field}
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_knowledge_base#text_field BedrockKnowledgeBase#text_field}
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
}

