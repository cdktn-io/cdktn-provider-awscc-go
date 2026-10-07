// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowmediastream


type MediaconnectFlowMediaStreamAttributesFmtp struct {
	// The format of the audio channel. Can only be specified for an audio media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#channel_order MediaconnectFlowMediaStream#channel_order}
	ChannelOrder *string `field:"optional" json:"channelOrder" yaml:"channelOrder"`
	// The format used for the representation of color.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#colorimetry MediaconnectFlowMediaStream#colorimetry}
	Colorimetry *string `field:"optional" json:"colorimetry" yaml:"colorimetry"`
	// The frame rate for the video stream, in frames/second. For example: 60000/1001.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#exact_framerate MediaconnectFlowMediaStream#exact_framerate}
	ExactFramerate *string `field:"optional" json:"exactFramerate" yaml:"exactFramerate"`
	// The pixel aspect ratio (PAR) of the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#par MediaconnectFlowMediaStream#par}
	Par *string `field:"optional" json:"par" yaml:"par"`
	// The encoding range of the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#range MediaconnectFlowMediaStream#range}
	Range *string `field:"optional" json:"range" yaml:"range"`
	// The type of compression that was used to smooth the video's appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#scan_mode MediaconnectFlowMediaStream#scan_mode}
	ScanMode *string `field:"optional" json:"scanMode" yaml:"scanMode"`
	// The transfer characteristic system (TCS) that is used in the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#tcs MediaconnectFlowMediaStream#tcs}
	Tcs *string `field:"optional" json:"tcs" yaml:"tcs"`
}

