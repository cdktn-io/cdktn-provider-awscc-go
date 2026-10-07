// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowmediastream

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowMediaStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN) of the flow that the media stream belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#flow_arn MediaconnectFlowMediaStream#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// A unique identifier for the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#media_stream_id MediaconnectFlowMediaStream#media_stream_id}
	MediaStreamId *float64 `field:"required" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#media_stream_name MediaconnectFlowMediaStream#media_stream_name}
	MediaStreamName *string `field:"required" json:"mediaStreamName" yaml:"mediaStreamName"`
	// The type of media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#media_stream_type MediaconnectFlowMediaStream#media_stream_type}
	MediaStreamType *string `field:"required" json:"mediaStreamType" yaml:"mediaStreamType"`
	// Attributes that are related to the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#attributes MediaconnectFlowMediaStream#attributes}
	Attributes *MediaconnectFlowMediaStreamAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// The sample rate (in Hz) for the stream.
	//
	// If the media stream type is video or ancillary data, set this value to 90000. If the media stream type is audio, set this value to either 48000 or 96000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#clock_rate MediaconnectFlowMediaStream#clock_rate}
	ClockRate *float64 `field:"optional" json:"clockRate" yaml:"clockRate"`
	// A description that can help you quickly identify what your media stream is used for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#description MediaconnectFlowMediaStream#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The key-value pairs that can be used to tag and organize the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#tags MediaconnectFlowMediaStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The resolution of the video. Required for a video media stream and rejected for other media stream types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#video_format MediaconnectFlowMediaStream#video_format}
	VideoFormat *string `field:"optional" json:"videoFormat" yaml:"videoFormat"`
}

