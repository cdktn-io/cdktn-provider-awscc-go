// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration struct {
	// The ARN of the Kinesis Data Stream sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_media_insights_pipeline_configuration#insights_target ChimeMediaInsightsPipelineConfiguration#insights_target}
	InsightsTarget *string `field:"optional" json:"insightsTarget" yaml:"insightsTarget"`
}

