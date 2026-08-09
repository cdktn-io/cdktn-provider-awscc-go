// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOption struct {
	// Network ACL common policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fms_policy#network_acl_common_policy FmsPolicy#network_acl_common_policy}
	NetworkAclCommonPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicy `field:"optional" json:"networkAclCommonPolicy" yaml:"networkAclCommonPolicy"`
	// Network firewall policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fms_policy#network_firewall_policy FmsPolicy#network_firewall_policy}
	NetworkFirewallPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// Third party firewall policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fms_policy#third_party_firewall_policy FmsPolicy#third_party_firewall_policy}
	ThirdPartyFirewallPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

