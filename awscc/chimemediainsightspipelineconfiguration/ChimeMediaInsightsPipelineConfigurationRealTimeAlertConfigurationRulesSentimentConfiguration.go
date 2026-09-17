// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesSentimentConfiguration struct {
	// The name of the sentiment rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#rule_name ChimeMediaInsightsPipelineConfiguration#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The type of sentiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#sentiment_type ChimeMediaInsightsPipelineConfiguration#sentiment_type}
	SentimentType *string `field:"optional" json:"sentimentType" yaml:"sentimentType"`
	// The analysis interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#time_period ChimeMediaInsightsPipelineConfiguration#time_period}
	TimePeriod *float64 `field:"optional" json:"timePeriod" yaml:"timePeriod"`
}

