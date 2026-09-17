// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamallocation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamAllocationConfig struct {
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
	// Id of the IPAM Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ipam_allocation#ipam_pool_id Ec2IpamAllocation#ipam_pool_id}
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Represents an IPAM custom allocation of a single IPv4 or IPv6 CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ipam_allocation#cidr Ec2IpamAllocation#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ipam_allocation#description Ec2IpamAllocation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The desired netmask length of the allocation.
	//
	// If set, IPAM will choose a block of free space with this size and return the CIDR representing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ipam_allocation#netmask_length Ec2IpamAllocation#netmask_length}
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

