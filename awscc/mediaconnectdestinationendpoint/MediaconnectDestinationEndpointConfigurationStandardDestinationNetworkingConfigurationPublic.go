// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfigurationPublic struct {
	// The rules that specify which CIDR blocks can connect to the destination endpoint.
	//
	// You can specify up to 10 rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_allow_rules MediaconnectDestinationEndpoint#destination_allow_rules}
	DestinationAllowRules interface{} `field:"optional" json:"destinationAllowRules" yaml:"destinationAllowRules"`
}

