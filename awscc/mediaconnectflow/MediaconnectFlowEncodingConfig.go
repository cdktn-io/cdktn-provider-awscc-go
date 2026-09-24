// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowEncodingConfig struct {
	// The encoding profile to use when transcoding the NDI source to a Transport Stream.
	//
	// You can change this value while a flow is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#encoding_profile MediaconnectFlow#encoding_profile}
	EncodingProfile *string `field:"optional" json:"encodingProfile" yaml:"encodingProfile"`
	// The maximum video bitrate to use when transcoding the NDI source to a Transport Stream.
	//
	// This parameter enables you to override the default video bitrate within the encoding profile's supported range. The supported range is 10,000,000 - 50,000,000 bits per second (bps). If you do not specify a value, MediaConnect uses the default value of 20,000,000 bps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#video_max_bitrate MediaconnectFlow#video_max_bitrate}
	VideoMaxBitrate *float64 `field:"optional" json:"videoMaxBitrate" yaml:"videoMaxBitrate"`
}

