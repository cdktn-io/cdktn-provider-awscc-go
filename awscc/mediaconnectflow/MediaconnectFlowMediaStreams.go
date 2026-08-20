// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowMediaStreams struct {
	// Attributes that are related to the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#attributes MediaconnectFlow#attributes}
	Attributes *MediaconnectFlowMediaStreamsAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// The sample rate for the stream. This value in measured in kHz.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#clock_rate MediaconnectFlow#clock_rate}
	ClockRate *float64 `field:"optional" json:"clockRate" yaml:"clockRate"`
	// A description that can help you quickly identify what your media stream is used for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#description MediaconnectFlow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The format type number (sometimes referred to as RTP payload type) of the media stream.
	//
	// MediaConnect assigns this value to the media stream. For ST 2110 JPEG XS outputs, you need to provide this value to the receiver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#fmt MediaconnectFlow#fmt}
	Fmt *float64 `field:"optional" json:"fmt" yaml:"fmt"`
	// A unique identifier for the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#media_stream_id MediaconnectFlow#media_stream_id}
	MediaStreamId *float64 `field:"optional" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#media_stream_name MediaconnectFlow#media_stream_name}
	MediaStreamName *string `field:"optional" json:"mediaStreamName" yaml:"mediaStreamName"`
	// The type of media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#media_stream_type MediaconnectFlow#media_stream_type}
	MediaStreamType *string `field:"optional" json:"mediaStreamType" yaml:"mediaStreamType"`
	// The resolution of the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_flow#video_format MediaconnectFlow#video_format}
	VideoFormat *string `field:"optional" json:"videoFormat" yaml:"videoFormat"`
}

