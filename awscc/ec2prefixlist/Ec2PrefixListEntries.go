// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2prefixlist


type Ec2PrefixListEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_prefix_list#cidr Ec2PrefixList#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_prefix_list#description Ec2PrefixList#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

