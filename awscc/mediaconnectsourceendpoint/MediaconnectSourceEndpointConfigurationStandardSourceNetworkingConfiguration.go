// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfiguration struct {
	// The networking configuration for a publicly accessible source endpoint, including the CIDR blocks that are allowed to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#public MediaconnectSourceEndpoint#public}
	Public *MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationPublic `field:"optional" json:"public" yaml:"public"`
	// The networking configuration for a source endpoint in a virtual private cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#vpc MediaconnectSourceEndpoint#vpc}
	Vpc *MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

