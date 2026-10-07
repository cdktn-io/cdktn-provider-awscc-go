// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowmediastream


type MediaconnectFlowMediaStreamAttributes struct {
	// A set of parameters that define the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#fmtp MediaconnectFlowMediaStream#fmtp}
	Fmtp *MediaconnectFlowMediaStreamAttributesFmtp `field:"optional" json:"fmtp" yaml:"fmtp"`
	// The audio language, in a format that is recognized by the receiver.
	//
	// Can only be specified for an audio media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediaconnect_flow_media_stream#lang MediaconnectFlowMediaStream#lang}
	Lang *string `field:"optional" json:"lang" yaml:"lang"`
}

