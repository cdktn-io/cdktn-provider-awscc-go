// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationElements struct {
	// The element type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_media_insights_pipeline_configuration#type ChimeMediaInsightsPipelineConfiguration#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_media_insights_pipeline_configuration#amazon_transcribe_call_analytics_processor_configuration ChimeMediaInsightsPipelineConfiguration#amazon_transcribe_call_analytics_processor_configuration}.
	AmazonTranscribeCallAnalyticsProcessorConfiguration *ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfiguration `field:"optional" json:"amazonTranscribeCallAnalyticsProcessorConfiguration" yaml:"amazonTranscribeCallAnalyticsProcessorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_media_insights_pipeline_configuration#amazon_transcribe_processor_configuration ChimeMediaInsightsPipelineConfiguration#amazon_transcribe_processor_configuration}.
	AmazonTranscribeProcessorConfiguration *ChimeMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration `field:"optional" json:"amazonTranscribeProcessorConfiguration" yaml:"amazonTranscribeProcessorConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_media_insights_pipeline_configuration#kinesis_data_stream_sink_configuration ChimeMediaInsightsPipelineConfiguration#kinesis_data_stream_sink_configuration}.
	KinesisDataStreamSinkConfiguration *ChimeMediaInsightsPipelineConfigurationElementsKinesisDataStreamSinkConfiguration `field:"optional" json:"kinesisDataStreamSinkConfiguration" yaml:"kinesisDataStreamSinkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_media_insights_pipeline_configuration#s3_recording_sink_configuration ChimeMediaInsightsPipelineConfiguration#s3_recording_sink_configuration}.
	S3RecordingSinkConfiguration *ChimeMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration `field:"optional" json:"s3RecordingSinkConfiguration" yaml:"s3RecordingSinkConfiguration"`
}

