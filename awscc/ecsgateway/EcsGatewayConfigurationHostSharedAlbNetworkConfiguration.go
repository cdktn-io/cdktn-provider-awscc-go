// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway


type EcsGatewayConfigurationHostSharedAlbNetworkConfiguration struct {
	// The inbound security group for the ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#inbound_security_group EcsGateway#inbound_security_group}
	InboundSecurityGroup *string `field:"optional" json:"inboundSecurityGroup" yaml:"inboundSecurityGroup"`
	// The IP address type for the ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#ip_address_type EcsGateway#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The subnets for the ALB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#subnets EcsGateway#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

