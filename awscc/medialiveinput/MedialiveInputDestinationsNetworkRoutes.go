// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputDestinationsNetworkRoutes struct {
	// The CIDR of the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#cidr MedialiveInput#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// An optional gateway for the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#gateway MedialiveInput#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}

