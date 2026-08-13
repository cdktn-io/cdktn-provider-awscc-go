// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationStandardProtocolConfiguration struct {
	// The configuration settings for a router output using the RIST (Reliable Internet Stream Transport) protocol, including the destination address and port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_output#rist MediaconnectRouterOutput#rist}
	Rist *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRist `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a router output using the RTP (Real-Time Transport Protocol) protocol, including the destination address and port, and forward error correction state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_output#rtp MediaconnectRouterOutput#rtp}
	Rtp *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtp `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in caller mode, including the destination address and port, minimum latency, stream ID, and encryption key configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_output#srt_caller MediaconnectRouterOutput#srt_caller}
	SrtCaller *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtCaller `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_output#srt_listener MediaconnectRouterOutput#srt_listener}
	SrtListener *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListener `field:"optional" json:"srtListener" yaml:"srtListener"`
}

