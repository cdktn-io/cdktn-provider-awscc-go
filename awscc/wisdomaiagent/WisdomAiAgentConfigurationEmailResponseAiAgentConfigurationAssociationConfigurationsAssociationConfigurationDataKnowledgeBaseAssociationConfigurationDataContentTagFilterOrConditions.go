// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataContentTagFilterOrConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#and_conditions WisdomAiAgent#and_conditions}.
	AndConditions interface{} `field:"optional" json:"andConditions" yaml:"andConditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#tag_condition WisdomAiAgent#tag_condition}.
	TagCondition *WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataContentTagFilterOrConditionsTagCondition `field:"optional" json:"tagCondition" yaml:"tagCondition"`
}

