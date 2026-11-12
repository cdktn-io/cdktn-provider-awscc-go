// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationRouterInput struct {
	// The Amazon Resource Name (ARN) of the router input that the destination endpoint represents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#router_input_arn MediaconnectDestinationEndpoint#router_input_arn}
	RouterInputArn *string `field:"optional" json:"routerInputArn" yaml:"routerInputArn"`
}

