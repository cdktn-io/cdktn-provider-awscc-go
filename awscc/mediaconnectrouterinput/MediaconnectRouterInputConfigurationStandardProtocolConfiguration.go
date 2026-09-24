// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationStandardProtocolConfiguration struct {
	// The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_input#rist MediaconnectRouterInput#rist}
	Rist *MediaconnectRouterInputConfigurationStandardProtocolConfigurationRist `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_input#rtp MediaconnectRouterInput#rtp}
	Rtp *MediaconnectRouterInputConfigurationStandardProtocolConfigurationRtp `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_input#srt_caller MediaconnectRouterInput#srt_caller}
	SrtCaller *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCaller `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_input#srt_listener MediaconnectRouterInput#srt_listener}
	SrtListener *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtListener `field:"optional" json:"srtListener" yaml:"srtListener"`
}

