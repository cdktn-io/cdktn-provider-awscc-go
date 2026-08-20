// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectpublicvirtualinterface


type DirectconnectPublicVirtualInterfaceBgpPeers struct {
	// The address family for the BGP peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#address_family DirectconnectPublicVirtualInterface#address_family}
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#asn DirectconnectPublicVirtualInterface#asn}
	Asn *string `field:"required" json:"asn" yaml:"asn"`
	// The IP address assigned to the Amazon interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#amazon_address DirectconnectPublicVirtualInterface#amazon_address}
	AmazonAddress *string `field:"optional" json:"amazonAddress" yaml:"amazonAddress"`
	// The authentication key for BGP configuration.
	//
	// This string has a minimum length of 6 characters and and a maximum length of 80 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#auth_key DirectconnectPublicVirtualInterface#auth_key}
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#bgp_peer_id DirectconnectPublicVirtualInterface#bgp_peer_id}.
	BgpPeerId *string `field:"optional" json:"bgpPeerId" yaml:"bgpPeerId"`
	// The IP address assigned to the customer interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_public_virtual_interface#customer_address DirectconnectPublicVirtualInterface#customer_address}
	CustomerAddress *string `field:"optional" json:"customerAddress" yaml:"customerAddress"`
}

