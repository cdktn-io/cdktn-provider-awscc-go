// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouternetworkinterface


type MediaconnectRouterNetworkInterfaceConfigurationPublicAllowRules struct {
	// The CIDR block that is allowed to access the public router network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_network_interface#cidr MediaconnectRouterNetworkInterface#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

