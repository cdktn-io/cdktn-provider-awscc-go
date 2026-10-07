// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflowmediastream


type MediaconnectFlowMediaStreamTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediaconnect_flow_media_stream#key MediaconnectFlowMediaStream#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediaconnect_flow_media_stream#value MediaconnectFlowMediaStream#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

