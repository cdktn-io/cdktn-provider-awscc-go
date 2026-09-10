// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightvpcconnection


type QuicksightVpcConnectionTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_vpc_connection#key QuicksightVpcConnection#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/quicksight_vpc_connection#value QuicksightVpcConnection#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

