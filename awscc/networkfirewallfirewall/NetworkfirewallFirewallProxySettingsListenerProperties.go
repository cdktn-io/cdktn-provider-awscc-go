// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewall


type NetworkfirewallFirewallProxySettingsListenerProperties struct {
	// The port the proxy listener accepts traffic on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#port NetworkfirewallFirewall#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol the proxy listener uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/networkfirewall_firewall#type NetworkfirewallFirewall#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

