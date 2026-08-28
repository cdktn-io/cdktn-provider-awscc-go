// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase


type WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_knowledge_base#bedrock_foundation_model_configuration WisdomKnowledgeBase#bedrock_foundation_model_configuration}.
	BedrockFoundationModelConfiguration *WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_knowledge_base#parsing_strategy WisdomKnowledgeBase#parsing_strategy}.
	ParsingStrategy *string `field:"optional" json:"parsingStrategy" yaml:"parsingStrategy"`
}

