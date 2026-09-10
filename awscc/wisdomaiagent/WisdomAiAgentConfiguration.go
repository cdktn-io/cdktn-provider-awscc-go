// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#answer_recommendation_ai_agent_configuration WisdomAiAgent#answer_recommendation_ai_agent_configuration}.
	AnswerRecommendationAiAgentConfiguration *WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfiguration `field:"optional" json:"answerRecommendationAiAgentConfiguration" yaml:"answerRecommendationAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#case_summarization_ai_agent_configuration WisdomAiAgent#case_summarization_ai_agent_configuration}.
	CaseSummarizationAiAgentConfiguration *WisdomAiAgentConfigurationCaseSummarizationAiAgentConfiguration `field:"optional" json:"caseSummarizationAiAgentConfiguration" yaml:"caseSummarizationAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#email_generative_answer_ai_agent_configuration WisdomAiAgent#email_generative_answer_ai_agent_configuration}.
	EmailGenerativeAnswerAiAgentConfiguration *WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfiguration `field:"optional" json:"emailGenerativeAnswerAiAgentConfiguration" yaml:"emailGenerativeAnswerAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#email_overview_ai_agent_configuration WisdomAiAgent#email_overview_ai_agent_configuration}.
	EmailOverviewAiAgentConfiguration *WisdomAiAgentConfigurationEmailOverviewAiAgentConfiguration `field:"optional" json:"emailOverviewAiAgentConfiguration" yaml:"emailOverviewAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#email_response_ai_agent_configuration WisdomAiAgent#email_response_ai_agent_configuration}.
	EmailResponseAiAgentConfiguration *WisdomAiAgentConfigurationEmailResponseAiAgentConfiguration `field:"optional" json:"emailResponseAiAgentConfiguration" yaml:"emailResponseAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#manual_search_ai_agent_configuration WisdomAiAgent#manual_search_ai_agent_configuration}.
	ManualSearchAiAgentConfiguration *WisdomAiAgentConfigurationManualSearchAiAgentConfiguration `field:"optional" json:"manualSearchAiAgentConfiguration" yaml:"manualSearchAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#note_taking_ai_agent_configuration WisdomAiAgent#note_taking_ai_agent_configuration}.
	NoteTakingAiAgentConfiguration *WisdomAiAgentConfigurationNoteTakingAiAgentConfiguration `field:"optional" json:"noteTakingAiAgentConfiguration" yaml:"noteTakingAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#orchestration_ai_agent_configuration WisdomAiAgent#orchestration_ai_agent_configuration}.
	OrchestrationAiAgentConfiguration *WisdomAiAgentConfigurationOrchestrationAiAgentConfiguration `field:"optional" json:"orchestrationAiAgentConfiguration" yaml:"orchestrationAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#self_service_ai_agent_configuration WisdomAiAgent#self_service_ai_agent_configuration}.
	SelfServiceAiAgentConfiguration *WisdomAiAgentConfigurationSelfServiceAiAgentConfiguration `field:"optional" json:"selfServiceAiAgentConfiguration" yaml:"selfServiceAiAgentConfiguration"`
}

