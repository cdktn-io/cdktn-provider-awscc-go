// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcendpointservice


type Ec2VpcEndpointServiceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_vpc_endpoint_service#key Ec2VpcEndpointService#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_vpc_endpoint_service#value Ec2VpcEndpointService#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

