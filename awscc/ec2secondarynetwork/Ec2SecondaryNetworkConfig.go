// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2secondarynetwork

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2SecondaryNetworkConfig struct {
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
	// The IPv4 CIDR block for the secondary network. The CIDR block size must be between /12 and /28.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ec2_secondary_network#ipv_4_cidr_block Ec2SecondaryNetwork#ipv_4_cidr_block}
	Ipv4CidrBlock *string `field:"required" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// The type of secondary network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ec2_secondary_network#network_type Ec2SecondaryNetwork#network_type}
	NetworkType *string `field:"required" json:"networkType" yaml:"networkType"`
	// The tags for the secondary network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ec2_secondary_network#tags Ec2SecondaryNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

