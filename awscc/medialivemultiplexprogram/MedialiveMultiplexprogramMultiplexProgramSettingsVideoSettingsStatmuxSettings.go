// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplexprogram


type MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettings struct {
	// Maximum statmux bitrate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/medialive_multiplexprogram#maximum_bitrate MedialiveMultiplexprogram#maximum_bitrate}
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Minimum statmux bitrate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/medialive_multiplexprogram#minimum_bitrate MedialiveMultiplexprogram#minimum_bitrate}
	MinimumBitrate *float64 `field:"optional" json:"minimumBitrate" yaml:"minimumBitrate"`
	// The purpose of the priority is to use a combination of the\nmultiplex rate control algorithm and the QVBR capability of the\nencoder to prioritize the video quality of some channels in a\nmultiplex over others.
	//
	// Channels that have a higher priority will\nget higher video quality at the expense of the video quality of\nother channels in the multiplex with lower priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/medialive_multiplexprogram#priority MedialiveMultiplexprogram#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

