// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamprefixlistresolver


type Ec2IpamPrefixListResolverRules struct {
	// Two of the rule types allow you to add conditions to the rules.
	//
	// (1) For IPAM Pool CIDR rules, you can specify an ipamPoolId; if not specified, the rule will apply to all IPAM Pool CIDRs in the scope.  (2) For IPAM Resource CIDR rules, you can specify resourceId, resourceOwner, resourceRegion, cidr, or resourceTag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver#conditions Ec2IpamPrefixListResolver#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// This rule will only match resources that are in this IPAM Scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver#ipam_scope_id Ec2IpamPrefixListResolver#ipam_scope_id}
	IpamScopeId *string `field:"optional" json:"ipamScopeId" yaml:"ipamScopeId"`
	// The resourceType property only applies to ipam-resource-cidr rules;
	//
	// this property specifies what type of resources this rule will apply to, such as VPCs or Subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver#resource_type Ec2IpamPrefixListResolver#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// There are three rule types: (1) Static CIDR: A fixed list of CIDRs that don't change (like a manual list replicated across Regions).
	//
	// (2) IPAM pool CIDR: CIDRs from specific IPAM pools (like all CIDRs from your IPAM production pool).  (3) IPAM resource CIDR: CIDRs for AWS resources like VPCs, subnets, and EIPs within a specific IPAM scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver#rule_type Ec2IpamPrefixListResolver#rule_type}
	RuleType *string `field:"optional" json:"ruleType" yaml:"ruleType"`
	// A fixed CIDR that doesn't change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver#static_cidr Ec2IpamPrefixListResolver#static_cidr}
	StaticCidr *string `field:"optional" json:"staticCidr" yaml:"staticCidr"`
}

