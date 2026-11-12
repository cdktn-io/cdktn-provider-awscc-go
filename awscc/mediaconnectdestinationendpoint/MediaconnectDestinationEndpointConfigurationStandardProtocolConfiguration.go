// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration struct {
	// The configuration settings for a destination endpoint that uses the RIST protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#rist MediaconnectDestinationEndpoint#rist}
	Rist *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRist `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a destination endpoint that uses the RTP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#rtp MediaconnectDestinationEndpoint#rtp}
	Rtp *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtp `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a destination endpoint that uses SRT in caller mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#srt_caller MediaconnectDestinationEndpoint#srt_caller}
	SrtCaller *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCaller `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a destination endpoint that uses SRT in listener mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#srt_listener MediaconnectDestinationEndpoint#srt_listener}
	SrtListener *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener `field:"optional" json:"srtListener" yaml:"srtListener"`
}

