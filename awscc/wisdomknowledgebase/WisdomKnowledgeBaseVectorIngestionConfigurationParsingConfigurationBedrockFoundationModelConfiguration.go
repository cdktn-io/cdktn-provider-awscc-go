// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase


type WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_knowledge_base#model_arn WisdomKnowledgeBase#model_arn}.
	ModelArn *string `field:"optional" json:"modelArn" yaml:"modelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wisdom_knowledge_base#parsing_prompt WisdomKnowledgeBase#parsing_prompt}.
	ParsingPrompt *WisdomKnowledgeBaseVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfigurationParsingPrompt `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

