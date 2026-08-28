// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationEmailResponseAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#association_configurations WisdomAiAgent#association_configurations}.
	AssociationConfigurations interface{} `field:"optional" json:"associationConfigurations" yaml:"associationConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#email_query_reformulation_ai_prompt_id WisdomAiAgent#email_query_reformulation_ai_prompt_id}.
	EmailQueryReformulationAiPromptId *string `field:"optional" json:"emailQueryReformulationAiPromptId" yaml:"emailQueryReformulationAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#email_response_ai_prompt_id WisdomAiAgent#email_response_ai_prompt_id}.
	EmailResponseAiPromptId *string `field:"optional" json:"emailResponseAiPromptId" yaml:"emailResponseAiPromptId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#locale WisdomAiAgent#locale}.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

