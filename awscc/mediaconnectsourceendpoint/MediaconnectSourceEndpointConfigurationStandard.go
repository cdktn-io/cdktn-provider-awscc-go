// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandard struct {
	// The name of the source endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#name MediaconnectSourceEndpoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#protocol MediaconnectSourceEndpoint#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol configuration for a standard source endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the protocol that the source endpoint uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#protocol_configuration MediaconnectSourceEndpoint#protocol_configuration}
	ProtocolConfiguration *MediaconnectSourceEndpointConfigurationStandardProtocolConfiguration `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
	// The networking configuration for a standard source endpoint.
	//
	// This is a union type, so exactly one of its members is set: public for a publicly accessible endpoint, or vpc for an endpoint in a virtual private cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#source_networking_configuration MediaconnectSourceEndpoint#source_networking_configuration}
	SourceNetworkingConfiguration *MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfiguration `field:"optional" json:"sourceNetworkingConfiguration" yaml:"sourceNetworkingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#source_networking_type MediaconnectSourceEndpoint#source_networking_type}.
	SourceNetworkingType *string `field:"optional" json:"sourceNetworkingType" yaml:"sourceNetworkingType"`
}

