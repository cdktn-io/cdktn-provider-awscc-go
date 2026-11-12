// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfiguration struct {
	// The configuration settings for a destination endpoint that represents a router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#router_input MediaconnectDestinationEndpoint#router_input}
	RouterInput *MediaconnectDestinationEndpointConfigurationRouterInput `field:"optional" json:"routerInput" yaml:"routerInput"`
	// The configuration settings for a standard destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#standard MediaconnectDestinationEndpoint#standard}
	Standard *MediaconnectDestinationEndpointConfigurationStandard `field:"optional" json:"standard" yaml:"standard"`
}

