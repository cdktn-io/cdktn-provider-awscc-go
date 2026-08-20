// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceNetworkInterfacesPrivateIpAddresses struct {
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_instance#primary Ec2Instance#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The private IPv4 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_instance#private_ip_address Ec2Instance#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

