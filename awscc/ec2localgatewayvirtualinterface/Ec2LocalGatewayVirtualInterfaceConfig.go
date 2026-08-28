// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2localgatewayvirtualinterface

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2LocalGatewayVirtualInterfaceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The local address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#local_address Ec2LocalGatewayVirtualInterface#local_address}
	LocalAddress *string `field:"required" json:"localAddress" yaml:"localAddress"`
	// The ID of the virtual interface group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#local_gateway_virtual_interface_group_id Ec2LocalGatewayVirtualInterface#local_gateway_virtual_interface_group_id}
	LocalGatewayVirtualInterfaceGroupId *string `field:"required" json:"localGatewayVirtualInterfaceGroupId" yaml:"localGatewayVirtualInterfaceGroupId"`
	// The Outpost LAG ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#outpost_lag_id Ec2LocalGatewayVirtualInterface#outpost_lag_id}
	OutpostLagId *string `field:"required" json:"outpostLagId" yaml:"outpostLagId"`
	// The peer address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#peer_address Ec2LocalGatewayVirtualInterface#peer_address}
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// The ID of the VLAN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#vlan Ec2LocalGatewayVirtualInterface#vlan}
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// The peer BGP ASN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#peer_bgp_asn Ec2LocalGatewayVirtualInterface#peer_bgp_asn}
	PeerBgpAsn *float64 `field:"optional" json:"peerBgpAsn" yaml:"peerBgpAsn"`
	// The extended 32-bit ASN of the BGP peer for use with larger ASN values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#peer_bgp_asn_extended Ec2LocalGatewayVirtualInterface#peer_bgp_asn_extended}
	PeerBgpAsnExtended *float64 `field:"optional" json:"peerBgpAsnExtended" yaml:"peerBgpAsnExtended"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_local_gateway_virtual_interface#tags Ec2LocalGatewayVirtualInterface#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

