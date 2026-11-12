// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSmpte2110ReceiverGroupSettingsSmpte2110ReceiverGroupsSdpSettings struct {
	// A list of InputSdpLocations.
	//
	// Each item in the list specifies the SDP file and index for one ancillary SMPTE 2110 stream.
	// Each stream encapsulates one captions stream (out of any number you can include) or the single SCTE 35 stream that you can include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#ancillary_sdps MedialiveInput#ancillary_sdps}
	AncillarySdps interface{} `field:"optional" json:"ancillarySdps" yaml:"ancillarySdps"`
	// A list of InputSdpLocations.
	//
	// Each item in the list specifies the SDP file and index for one audio SMPTE 2110 stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#audio_sdps MedialiveInput#audio_sdps}
	AudioSdps interface{} `field:"optional" json:"audioSdps" yaml:"audioSdps"`
	// The InputSdpLocation that specifies the SDP file and index for the single video SMPTE 2110 stream for this 2110 input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#video_sdp MedialiveInput#video_sdp}
	VideoSdp *MedialiveInputSmpte2110ReceiverGroupSettingsSmpte2110ReceiverGroupsSdpSettingsVideoSdp `field:"optional" json:"videoSdp" yaml:"videoSdp"`
}

