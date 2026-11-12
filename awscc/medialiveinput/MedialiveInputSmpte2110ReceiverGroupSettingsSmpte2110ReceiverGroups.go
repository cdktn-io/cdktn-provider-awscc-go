// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSmpte2110ReceiverGroupSettingsSmpte2110ReceiverGroups struct {
	// The single Smpte2110ReceiverGroupSdpSettings that identify the video, audio, and ancillary streams for this receiver group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#sdp_settings MedialiveInput#sdp_settings}
	SdpSettings *MedialiveInputSmpte2110ReceiverGroupSettingsSmpte2110ReceiverGroupsSdpSettings `field:"optional" json:"sdpSettings" yaml:"sdpSettings"`
}

