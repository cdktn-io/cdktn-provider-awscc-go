// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationStandard struct {
	// The Amazon Resource Name (ARN) of the network interface associated with the standard router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#network_interface_arn MediaconnectRouterInput#network_interface_arn}
	NetworkInterfaceArn *string `field:"optional" json:"networkInterfaceArn" yaml:"networkInterfaceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#protocol MediaconnectRouterInput#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol configuration settings for a router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#protocol_configuration MediaconnectRouterInput#protocol_configuration}
	ProtocolConfiguration *MediaconnectRouterInputConfigurationStandardProtocolConfiguration `field:"optional" json:"protocolConfiguration" yaml:"protocolConfiguration"`
}

