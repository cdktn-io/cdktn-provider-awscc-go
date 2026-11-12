// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSmpte2110ReceiverGroupSettingsSmpte2110ReceiverGroupsSdpSettingsVideoSdp struct {
	// The index of the media stream in the SDP file for one SMPTE 2110 stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#media_index MedialiveInput#media_index}
	MediaIndex *float64 `field:"optional" json:"mediaIndex" yaml:"mediaIndex"`
	// The URL of the SDP file for one SMPTE 2110 stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#sdp_url MedialiveInput#sdp_url}
	SdpUrl *string `field:"optional" json:"sdpUrl" yaml:"sdpUrl"`
}

