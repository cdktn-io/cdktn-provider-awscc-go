// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfiguration struct {
	// The networking configuration for a publicly accessible destination endpoint, including the CIDR blocks that are allowed to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#public MediaconnectDestinationEndpoint#public}
	Public *MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfigurationPublic `field:"optional" json:"public" yaml:"public"`
	// The networking configuration for a destination endpoint in a virtual private cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#vpc MediaconnectDestinationEndpoint#vpc}
	Vpc *MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfigurationVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

