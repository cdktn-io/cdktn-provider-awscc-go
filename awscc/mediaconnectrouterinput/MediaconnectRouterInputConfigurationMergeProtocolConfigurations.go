// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMergeProtocolConfigurations struct {
	// The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_input#rist MediaconnectRouterInput#rist}
	Rist *MediaconnectRouterInputConfigurationMergeProtocolConfigurationsRist `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_input#rtp MediaconnectRouterInput#rtp}
	Rtp *MediaconnectRouterInputConfigurationMergeProtocolConfigurationsRtp `field:"optional" json:"rtp" yaml:"rtp"`
}

