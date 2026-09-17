// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowMediaStreamsAttributes struct {
	// A set of parameters that define the media stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#fmtp MediaconnectFlow#fmtp}
	Fmtp *MediaconnectFlowMediaStreamsAttributesFmtp `field:"optional" json:"fmtp" yaml:"fmtp"`
	// The audio language, in a format that is recognized by the receiver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#lang MediaconnectFlow#lang}
	Lang *string `field:"optional" json:"lang" yaml:"lang"`
}

