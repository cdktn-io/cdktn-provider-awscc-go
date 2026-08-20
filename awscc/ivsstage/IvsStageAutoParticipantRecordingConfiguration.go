// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstage


type IvsStageAutoParticipantRecordingConfiguration struct {
	// HLS configuration object for individual participant recording.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_stage#hls_configuration IvsStage#hls_configuration}
	HlsConfiguration *IvsStageAutoParticipantRecordingConfigurationHlsConfiguration `field:"optional" json:"hlsConfiguration" yaml:"hlsConfiguration"`
	// Types of media to be recorded. Default: AUDIO_VIDEO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_stage#media_types IvsStage#media_types}
	MediaTypes *[]*string `field:"optional" json:"mediaTypes" yaml:"mediaTypes"`
	// If a stage publisher disconnects and then reconnects within the specified interval, the multiple recordings will be considered a single recording and merged together.
	//
	// The default value is 0, which disables merging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_stage#recording_reconnect_window_seconds IvsStage#recording_reconnect_window_seconds}
	RecordingReconnectWindowSeconds *float64 `field:"optional" json:"recordingReconnectWindowSeconds" yaml:"recordingReconnectWindowSeconds"`
	// ARN of the StorageConfiguration resource to use for individual participant recording.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_stage#storage_configuration_arn IvsStage#storage_configuration_arn}
	StorageConfigurationArn *string `field:"optional" json:"storageConfigurationArn" yaml:"storageConfigurationArn"`
	// A complex type that allows you to enable/disable the recording of thumbnails for individual participant recording and modify the interval at which thumbnails are generated for the live session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_stage#thumbnail_configuration IvsStage#thumbnail_configuration}
	ThumbnailConfiguration *IvsStageAutoParticipantRecordingConfigurationThumbnailConfiguration `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
}

