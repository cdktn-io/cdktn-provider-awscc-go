// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpnconnection


type Ec2VpnConnectionTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_vpn_connection#key Ec2VpnConnection#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_vpn_connection#value Ec2VpnConnection#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

