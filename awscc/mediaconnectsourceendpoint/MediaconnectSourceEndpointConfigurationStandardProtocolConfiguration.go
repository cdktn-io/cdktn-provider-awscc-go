// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardProtocolConfiguration struct {
	// The configuration settings for a source endpoint that uses the RIST protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#rist MediaconnectSourceEndpoint#rist}
	Rist *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRist `field:"optional" json:"rist" yaml:"rist"`
	// The configuration settings for a source endpoint that uses the RTP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#rtp MediaconnectSourceEndpoint#rtp}
	Rtp *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtp `field:"optional" json:"rtp" yaml:"rtp"`
	// The configuration settings for a source endpoint that uses SRT in caller mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#srt_caller MediaconnectSourceEndpoint#srt_caller}
	SrtCaller *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCaller `field:"optional" json:"srtCaller" yaml:"srtCaller"`
	// The configuration settings for a source endpoint that uses SRT in listener mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#srt_listener MediaconnectSourceEndpoint#srt_listener}
	SrtListener *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListener `field:"optional" json:"srtListener" yaml:"srtListener"`
}

