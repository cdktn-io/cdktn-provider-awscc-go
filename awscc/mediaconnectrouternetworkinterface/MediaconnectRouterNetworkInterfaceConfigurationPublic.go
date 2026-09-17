// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouternetworkinterface


type MediaconnectRouterNetworkInterfaceConfigurationPublic struct {
	// The list of allowed CIDR blocks for the public router network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_network_interface#allow_rules MediaconnectRouterNetworkInterface#allow_rules}
	AllowRules interface{} `field:"optional" json:"allowRules" yaml:"allowRules"`
}

