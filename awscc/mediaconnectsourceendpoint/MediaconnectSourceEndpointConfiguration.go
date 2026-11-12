// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfiguration struct {
	// The configuration settings for a source endpoint that represents a router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#router_output MediaconnectSourceEndpoint#router_output}
	RouterOutput *MediaconnectSourceEndpointConfigurationRouterOutput `field:"optional" json:"routerOutput" yaml:"routerOutput"`
	// The configuration settings for a standard source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#standard MediaconnectSourceEndpoint#standard}
	Standard *MediaconnectSourceEndpointConfigurationStandard `field:"optional" json:"standard" yaml:"standard"`
}

