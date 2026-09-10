// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#content_tag_filter WisdomAiAgent#content_tag_filter}.
	ContentTagFilter *WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataContentTagFilter `field:"optional" json:"contentTagFilter" yaml:"contentTagFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#max_results WisdomAiAgent#max_results}.
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/wisdom_ai_agent#override_knowledge_base_search_type WisdomAiAgent#override_knowledge_base_search_type}.
	OverrideKnowledgeBaseSearchType *string `field:"optional" json:"overrideKnowledgeBaseSearchType" yaml:"overrideKnowledgeBaseSearchType"`
}

