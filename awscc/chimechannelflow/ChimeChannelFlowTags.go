// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimechannelflow


type ChimeChannelFlowTags struct {
	// The key in a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_channel_flow#key ChimeChannelFlow#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value in a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_channel_flow#value ChimeChannelFlow#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

