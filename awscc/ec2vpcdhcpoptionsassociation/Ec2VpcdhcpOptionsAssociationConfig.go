// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcdhcpoptionsassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcdhcpOptionsAssociationConfig struct {
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
	// The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_vpcdhcp_options_association#dhcp_options_id Ec2VpcdhcpOptionsAssociation#dhcp_options_id}
	DhcpOptionsId *string `field:"required" json:"dhcpOptionsId" yaml:"dhcpOptionsId"`
	// The ID of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_vpcdhcp_options_association#vpc_id Ec2VpcdhcpOptionsAssociation#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

