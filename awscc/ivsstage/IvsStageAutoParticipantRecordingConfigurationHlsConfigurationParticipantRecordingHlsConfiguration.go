// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstage


type IvsStageAutoParticipantRecordingConfigurationHlsConfigurationParticipantRecordingHlsConfiguration struct {
	// Defines the target duration for recorded segments generated when recording a stage participant.
	//
	// Segments may have durations longer than the specified value when needed to ensure each segment begins with a keyframe. Default: 6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ivs_stage#target_segment_duration_seconds IvsStage#target_segment_duration_seconds}
	TargetSegmentDurationSeconds *float64 `field:"optional" json:"targetSegmentDurationSeconds" yaml:"targetSegmentDurationSeconds"`
}

