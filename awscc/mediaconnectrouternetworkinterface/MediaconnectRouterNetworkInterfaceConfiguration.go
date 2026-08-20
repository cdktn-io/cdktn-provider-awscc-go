// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouternetworkinterface


type MediaconnectRouterNetworkInterfaceConfiguration struct {
	// The configuration settings for a public router network interface, including the list of allowed CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_network_interface#public MediaconnectRouterNetworkInterface#public}
	Public *MediaconnectRouterNetworkInterfaceConfigurationPublic `field:"optional" json:"public" yaml:"public"`
	// The configuration settings for a router network interface within a VPC, including the security group IDs and subnet ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_network_interface#vpc MediaconnectRouterNetworkInterface#vpc}
	Vpc *MediaconnectRouterNetworkInterfaceConfigurationVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

