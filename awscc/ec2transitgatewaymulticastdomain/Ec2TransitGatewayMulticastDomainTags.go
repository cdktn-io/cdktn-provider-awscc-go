// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaymulticastdomain


type Ec2TransitGatewayMulticastDomainTags struct {
	// The key of the tag.
	//
	// Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_transit_gateway_multicast_domain#key Ec2TransitGatewayMulticastDomain#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_transit_gateway_multicast_domain#value Ec2TransitGatewayMulticastDomain#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

