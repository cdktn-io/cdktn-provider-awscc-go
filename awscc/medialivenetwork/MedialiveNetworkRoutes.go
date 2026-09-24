// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivenetwork


type MedialiveNetworkRoutes struct {
	// Ip address cidr.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_network#cidr MedialiveNetwork#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// IP address for the route packet paths.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/medialive_network#gateway MedialiveNetwork#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}

