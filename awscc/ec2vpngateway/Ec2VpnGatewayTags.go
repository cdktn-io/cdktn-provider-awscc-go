// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpngateway


type Ec2VpnGatewayTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_vpn_gateway#key Ec2VpnGateway#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_vpn_gateway#value Ec2VpnGateway#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

