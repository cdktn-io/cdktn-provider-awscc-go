// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewall


type NetworkfirewallFirewallSubnetMappings struct {
	// A IPAddressType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall#ip_address_type NetworkfirewallFirewall#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// A SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/networkfirewall_firewall#subnet_id NetworkfirewallFirewall#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

