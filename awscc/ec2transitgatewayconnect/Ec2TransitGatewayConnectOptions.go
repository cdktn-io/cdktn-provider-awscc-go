// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayconnect


type Ec2TransitGatewayConnectOptions struct {
	// The tunnel protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_transit_gateway_connect#protocol Ec2TransitGatewayConnect#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

