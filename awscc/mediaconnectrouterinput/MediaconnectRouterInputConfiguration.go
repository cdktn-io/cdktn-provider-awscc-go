// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfiguration struct {
	// Configuration settings for a router input that is represented by a destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#destination_endpoint MediaconnectRouterInput#destination_endpoint}
	DestinationEndpoint *MediaconnectRouterInputConfigurationDestinationEndpoint `field:"optional" json:"destinationEndpoint" yaml:"destinationEndpoint"`
	// Configuration settings for a failover router input that allows switching between two input sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#failover MediaconnectRouterInput#failover}
	Failover *MediaconnectRouterInputConfigurationFailover `field:"optional" json:"failover" yaml:"failover"`
	// Configuration settings for connecting a router input to a flow output.
	//
	// To connect the router input to the flow output, provide both flowArn and flowOutputArn. To disconnect it, call the update operation and omit both fields while keeping this configuration object. You can't change sourceTransitDecryption while the router input is connected to a flow output. To change it, first disconnect the router input, then update the decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#media_connect_flow MediaconnectRouterInput#media_connect_flow}
	MediaConnectFlow *MediaconnectRouterInputConfigurationMediaConnectFlow `field:"optional" json:"mediaConnectFlow" yaml:"mediaConnectFlow"`
	// Configuration settings for connecting a router input to a MediaLive channel output.
	//
	// To connect the router input to the MediaLive channel output, provide all three of mediaLiveChannelArn, mediaLivePipelineId, and mediaLiveChannelOutputName. To disconnect it, call the update operation and omit all three fields while keeping this configuration object. You can't change sourceTransitDecryption while the router input is connected to a MediaLive channel output. To change it, first disconnect the router input, then update the decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#media_live_channel MediaconnectRouterInput#media_live_channel}
	MediaLiveChannel *MediaconnectRouterInputConfigurationMediaLiveChannel `field:"optional" json:"mediaLiveChannel" yaml:"mediaLiveChannel"`
	// Configuration settings for a merge router input that combines two input sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#merge MediaconnectRouterInput#merge}
	Merge *MediaconnectRouterInputConfigurationMerge `field:"optional" json:"merge" yaml:"merge"`
	// The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_router_input#standard MediaconnectRouterInput#standard}
	Standard *MediaconnectRouterInputConfigurationStandard `field:"optional" json:"standard" yaml:"standard"`
}

