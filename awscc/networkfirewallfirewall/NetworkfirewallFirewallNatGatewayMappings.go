// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewall


type NetworkfirewallFirewallNatGatewayMappings struct {
	// The unique identifier of a NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#nat_gateway_id NetworkfirewallFirewall#nat_gateway_id}
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
}

