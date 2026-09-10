// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectgateway


type MediaconnectGatewayNetworks struct {
	// A unique IP address range to use for this network.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_gateway#cidr_block MediaconnectGateway#cidr_block}
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The name of the network.
	//
	// This name is used to reference the network and must be unique among networks in this gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_gateway#name MediaconnectGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

