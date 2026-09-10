// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRist struct {
	// The destination IP address for the RIST protocol in the router output configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#destination_address MediaconnectRouterOutput#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The destination port number for the RIST protocol in the router output configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#destination_port MediaconnectRouterOutput#destination_port}
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
}

