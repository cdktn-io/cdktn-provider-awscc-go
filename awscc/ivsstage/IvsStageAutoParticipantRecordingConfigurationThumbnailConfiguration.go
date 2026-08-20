// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsstage


type IvsStageAutoParticipantRecordingConfigurationThumbnailConfiguration struct {
	// An object representing a configuration of thumbnails for recorded video from an individual participant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_stage#participant_thumbnail_configuration IvsStage#participant_thumbnail_configuration}
	ParticipantThumbnailConfiguration *IvsStageAutoParticipantRecordingConfigurationThumbnailConfigurationParticipantThumbnailConfiguration `field:"optional" json:"participantThumbnailConfiguration" yaml:"participantThumbnailConfiguration"`
}

