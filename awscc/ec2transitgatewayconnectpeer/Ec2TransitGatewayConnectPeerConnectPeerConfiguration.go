// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewayconnectpeer


type Ec2TransitGatewayConnectPeerConnectPeerConfiguration struct {
	// The range of interior BGP peer IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_connect_peer#inside_cidr_blocks Ec2TransitGatewayConnectPeer#inside_cidr_blocks}
	InsideCidrBlocks *[]*string `field:"required" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// The peer IP address (GRE outer IP address) on the appliance side of the Connect peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_connect_peer#peer_address Ec2TransitGatewayConnectPeer#peer_address}
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// The Connect peer IP address on the transit gateway side of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_transit_gateway_connect_peer#transit_gateway_address Ec2TransitGatewayConnectPeer#transit_gateway_address}
	TransitGatewayAddress *string `field:"optional" json:"transitGatewayAddress" yaml:"transitGatewayAddress"`
}

