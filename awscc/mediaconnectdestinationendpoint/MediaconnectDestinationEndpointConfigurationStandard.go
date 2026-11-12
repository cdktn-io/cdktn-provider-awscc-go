// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandard struct {
	// The networking configuration for a standard destination endpoint.
	//
	// This is a union type, so exactly one of its members is set: public for a publicly accessible endpoint, or vpc for an endpoint in a virtual private cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_networking_configuration MediaconnectDestinationEndpoint#destination_networking_configuration}
	DestinationNetworkingConfiguration *MediaconnectDestinationEndpointConfigurationStandardDestinationNetworkingConfiguration `field:"optional" json:"destinationNetworkingConfiguration" yaml:"destinationNetworkingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_networking_type MediaconnectDestinationEndpoint#destination_networking_type}.
	DestinationNetworkingType *string `field:"optional" json:"destinationNetworkingType" yaml:"destinationNetworkingType"`
	// The name of the destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#name MediaconnectDestinationEndpoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#protocol MediaconnectDestinationEndpoint#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol configuration for a standard destination endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the protocol that the destination endpoint uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#protocol_configuration MediaconnectDestinationEndpoint#protocol_configuration}
	ProtocolConfiguration *MediaconnectDestinationEndpointConfigurationStandardProtocolConfiguration `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
}

