// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfigurationRulesKeywordMatchConfiguration struct {
	// The keywords or phrases to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#keywords ChimeMediaInsightsPipelineConfiguration#keywords}
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// Matches keywords on their presence or absence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#negate ChimeMediaInsightsPipelineConfiguration#negate}
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// The name of the keyword match rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#rule_name ChimeMediaInsightsPipelineConfiguration#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

