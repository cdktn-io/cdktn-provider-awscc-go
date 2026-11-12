// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfiguration struct {
	// Configuration settings for connecting a router output to a MediaConnect flow source.
	//
	// To connect the router output to the flow source, provide both flowArn and flowSourceArn. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change destinationTransitEncryption while the router output is connected to a flow source. To change it, first disconnect the router output, then update the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#media_connect_flow MediaconnectRouterOutput#media_connect_flow}
	MediaConnectFlow *MediaconnectRouterOutputConfigurationMediaConnectFlow `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for connecting a router output to a MediaLive input.
	//
	// To connect the router output to the MediaLive input, provide both mediaLiveInputArn and mediaLivePipelineId. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change destinationTransitEncryption while the router output is connected to a MediaLive input. To change it, first disconnect the router output, then update the encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#media_live_input MediaconnectRouterOutput#media_live_input}
	MediaLiveInput *MediaconnectRouterOutputConfigurationMediaLiveInput `field:"optional" json:"mediaLiveInput" yaml:"mediaLiveInput"`
	// Configuration settings for a router output that is represented by a source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#source_endpoint MediaconnectRouterOutput#source_endpoint}
	SourceEndpoint *MediaconnectRouterOutputConfigurationSourceEndpoint `field:"optional" json:"sourceEndpoint" yaml:"sourceEndpoint"`
	// The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_output#standard MediaconnectRouterOutput#standard}
	Standard *MediaconnectRouterOutputConfigurationStandard `field:"optional" json:"standard" yaml:"standard"`
}

