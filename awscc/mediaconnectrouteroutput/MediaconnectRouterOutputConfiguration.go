// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfiguration struct {
	// Configuration settings for connecting a router output to a MediaConnect flow source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_output#media_connect_flow MediaconnectRouterOutput#media_connect_flow}
	MediaConnectFlow *MediaconnectRouterOutputConfigurationMediaConnectFlow `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for connecting a router output to a MediaLive input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_output#media_live_input MediaconnectRouterOutput#media_live_input}
	MediaLiveInput *MediaconnectRouterOutputConfigurationMediaLiveInput `field:"optional" json:"mediaLiveInput" yaml:"mediaLiveInput"`
	// The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_router_output#standard MediaconnectRouterOutput#standard}
	Standard *MediaconnectRouterOutputConfigurationStandard `field:"optional" json:"standard" yaml:"standard"`
}

