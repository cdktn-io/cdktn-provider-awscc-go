// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationStandard struct {
	// The Amazon Resource Name (ARN) of the network interface associated with the standard router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#network_interface_arn MediaconnectRouterOutput#network_interface_arn}
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#protocol MediaconnectRouterOutput#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol configuration settings for a router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#protocol_configuration MediaconnectRouterOutput#protocol_configuration}
	ProtocolConfiguration *MediaconnectRouterOutputConfigurationStandardProtocolConfiguration `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
}

