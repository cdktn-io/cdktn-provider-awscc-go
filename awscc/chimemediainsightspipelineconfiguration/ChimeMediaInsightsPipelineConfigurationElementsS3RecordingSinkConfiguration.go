// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration


type ChimeMediaInsightsPipelineConfigurationElementsS3RecordingSinkConfiguration struct {
	// The default URI of the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#destination ChimeMediaInsightsPipelineConfiguration#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The recording file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#recording_file_format ChimeMediaInsightsPipelineConfiguration#recording_file_format}
	RecordingFileFormat *string `field:"optional" json:"recordingFileFormat" yaml:"recordingFileFormat"`
}

