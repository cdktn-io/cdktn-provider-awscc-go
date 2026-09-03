// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamprefixlistresolver


type Ec2IpamPrefixListResolverRulesConditions struct {
	// Condition for the IPAM Resource CIDR rule type.  CIDR (like 10.24.34.0/23).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#cidr Ec2IpamPrefixListResolver#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Condition for the IPAM Pool CIDR rule type.
	//
	// If not chosen, the resolver applies to all IPAM Pool CIDRs in the scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#ipam_pool_id Ec2IpamPrefixListResolver#ipam_pool_id}
	IpamPoolId *string `field:"optional" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Equals, Not equals, or Subnet Of.  The subnet-of operation only applies to cidr conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#operation Ec2IpamPrefixListResolver#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Condition for the IPAM Resource CIDR rule type.  The unique ID of a resource (like vpc-1234567890abcdef0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#resource_id Ec2IpamPrefixListResolver#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Condition for the IPAM Resource CIDR rule type.  Resource owner (like 111122223333).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#resource_owner Ec2IpamPrefixListResolver#resource_owner}
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
	// Condition for the IPAM Resource CIDR rule type.  Resource region (like us-east-1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#resource_region Ec2IpamPrefixListResolver#resource_region}
	ResourceRegion *string `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// Condition for the IPAM Resource CIDR rule type.  Resource Tag (like dev-vpc-1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam_prefix_list_resolver#resource_tag Ec2IpamPrefixListResolver#resource_tag}
	ResourceTag *Ec2IpamPrefixListResolverRulesConditionsResourceTag `field:"optional" json:"resourceTag" yaml:"resourceTag"`
}

