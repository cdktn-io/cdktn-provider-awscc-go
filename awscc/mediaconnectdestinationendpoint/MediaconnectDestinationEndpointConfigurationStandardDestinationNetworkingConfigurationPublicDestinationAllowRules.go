// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfigurationPublicDestinationAllowRules struct {
	// The CIDR block that is allowed to connect to the destination endpoint, in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#cidr MediaconnectDestinationEndpoint#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

