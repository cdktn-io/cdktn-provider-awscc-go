// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivenode


type MedialiveNodeSdiSourceMappings struct {
	// The card number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/medialive_node#card_number MedialiveNode#card_number}
	CardNumber *float64 `field:"optional" json:"cardNumber" yaml:"cardNumber"`
	// The channel number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/medialive_node#channel_number MedialiveNode#channel_number}
	ChannelNumber *float64 `field:"optional" json:"channelNumber" yaml:"channelNumber"`
	// The SDI source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/medialive_node#sdi_source MedialiveNode#sdi_source}
	SdiSource *string `field:"optional" json:"sdiSource" yaml:"sdiSource"`
}

