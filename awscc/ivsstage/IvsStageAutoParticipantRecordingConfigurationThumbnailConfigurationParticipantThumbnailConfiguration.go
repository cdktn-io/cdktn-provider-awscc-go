// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstage


type IvsStageAutoParticipantRecordingConfigurationThumbnailConfigurationParticipantThumbnailConfiguration struct {
	// Thumbnail recording mode. Default: DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_stage#recording_mode IvsStage#recording_mode}
	RecordingMode *string `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Indicates the format in which thumbnails are recorded.
	//
	// SEQUENTIAL records all generated thumbnails in a serial manner, to the media/thumbnails/high directory. LATEST saves the latest thumbnail in media/latest_thumbnail/high/thumb.jpg and overwrites it at the interval specified by targetIntervalSeconds. You can enable both SEQUENTIAL and LATEST. Default: SEQUENTIAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_stage#storage IvsStage#storage}
	Storage *[]*string `field:"optional" json:"storage" yaml:"storage"`
	// The targeted thumbnail-generation interval in seconds. This is configurable only if recordingMode is INTERVAL. Default: 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_stage#target_interval_seconds IvsStage#target_interval_seconds}
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

