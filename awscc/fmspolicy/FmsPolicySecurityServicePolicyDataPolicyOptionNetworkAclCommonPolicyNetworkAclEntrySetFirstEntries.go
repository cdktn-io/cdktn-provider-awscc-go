// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySetFirstEntries struct {
	// CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#cidr_block FmsPolicy#cidr_block}
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Whether the entry is an egress entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#egress FmsPolicy#egress}
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// ICMP type and code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#icmp_type_code FmsPolicy#icmp_type_code}
	IcmpTypeCode *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySetFirstEntriesIcmpTypeCode `field:"optional" json:"icmpTypeCode" yaml:"icmpTypeCode"`
	// IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#ipv_6_cidr_block FmsPolicy#ipv_6_cidr_block}
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#port_range FmsPolicy#port_range}
	PortRange *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySetFirstEntriesPortRange `field:"optional" json:"portRange" yaml:"portRange"`
	// Protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#protocol FmsPolicy#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Rule Action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fms_policy#rule_action FmsPolicy#rule_action}
	RuleAction *string `field:"optional" json:"ruleAction" yaml:"ruleAction"`
}

