// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisvideosignalingchannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisvideoSignalingChannelConfig struct {
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
	// The period of time a signaling channel retains undelivered messages before they are discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisvideo_signaling_channel#message_ttl_seconds KinesisvideoSignalingChannel#message_ttl_seconds}
	MessageTtlSeconds *float64 `field:"optional" json:"messageTtlSeconds" yaml:"messageTtlSeconds"`
	// The name of the Kinesis Video Signaling Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisvideo_signaling_channel#name KinesisvideoSignalingChannel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisvideo_signaling_channel#tags KinesisvideoSignalingChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisvideo_signaling_channel#type KinesisvideoSignalingChannel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

