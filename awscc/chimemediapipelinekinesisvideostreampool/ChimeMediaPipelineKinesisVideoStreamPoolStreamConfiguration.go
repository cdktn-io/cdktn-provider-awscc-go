// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediapipelinekinesisvideostreampool


type ChimeMediaPipelineKinesisVideoStreamPoolStreamConfiguration struct {
	// The AWS Region of the video stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chime_media_pipeline_kinesis_video_stream_pool#region ChimeMediaPipelineKinesisVideoStreamPool#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The amount of time that data is retained, in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/chime_media_pipeline_kinesis_video_stream_pool#data_retention_in_hours ChimeMediaPipelineKinesisVideoStreamPool#data_retention_in_hours}
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
}

