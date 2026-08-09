// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2egressonlyinternetgateway


type Ec2EgressOnlyInternetGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_egress_only_internet_gateway#key Ec2EgressOnlyInternetGateway#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_egress_only_internet_gateway#value Ec2EgressOnlyInternetGateway#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

