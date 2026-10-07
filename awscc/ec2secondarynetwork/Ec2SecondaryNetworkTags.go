// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2secondarynetwork


type Ec2SecondaryNetworkTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ec2_secondary_network#key Ec2SecondaryNetwork#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ec2_secondary_network#value Ec2SecondaryNetwork#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

