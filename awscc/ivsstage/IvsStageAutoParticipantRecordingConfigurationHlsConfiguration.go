// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstage


type IvsStageAutoParticipantRecordingConfigurationHlsConfiguration struct {
	// An object representing a configuration of participant HLS recordings for individual participant recording.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ivs_stage#participant_recording_hls_configuration IvsStage#participant_recording_hls_configuration}
	ParticipantRecordingHlsConfiguration *IvsStageAutoParticipantRecordingConfigurationHlsConfigurationParticipantRecordingHlsConfiguration `field:"optional" json:"participantRecordingHlsConfiguration" yaml:"participantRecordingHlsConfiguration"`
}

