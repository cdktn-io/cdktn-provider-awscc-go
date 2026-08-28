// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayattachment


type Ec2TransitGatewayAttachmentTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_attachment#key Ec2TransitGatewayAttachment#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_attachment#value Ec2TransitGatewayAttachment#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

