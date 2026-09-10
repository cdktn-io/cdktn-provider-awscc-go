// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamprefixlistresolver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamPrefixListResolverConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The address family of the address space in this Prefix List Resolver. Either IPv4 or IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_prefix_list_resolver#address_family Ec2IpamPrefixListResolver#address_family}
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_prefix_list_resolver#description Ec2IpamPrefixListResolver#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Id of the IPAM this Prefix List Resolver is a part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_prefix_list_resolver#ipam_id Ec2IpamPrefixListResolver#ipam_id}
	IpamId *string `field:"optional" json:"ipamId" yaml:"ipamId"`
	// Rules define the business logic for selecting CIDRs from IPAM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_prefix_list_resolver#rules Ec2IpamPrefixListResolver#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_prefix_list_resolver#tags Ec2IpamPrefixListResolver#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

