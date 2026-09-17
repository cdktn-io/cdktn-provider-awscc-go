// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#issue_detection_configuration ChimeMediaInsightsPipelineConfiguration#issue_detection_configuration}.
	IssueDetectionConfiguration *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesIssueDetectionConfiguration `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#keyword_match_configuration ChimeMediaInsightsPipelineConfiguration#keyword_match_configuration}.
	KeywordMatchConfiguration *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfiguration `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#sentiment_configuration ChimeMediaInsightsPipelineConfiguration#sentiment_configuration}.
	SentimentConfiguration *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfiguration `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
	// The type of alert rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#type ChimeMediaInsightsPipelineConfiguration#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

