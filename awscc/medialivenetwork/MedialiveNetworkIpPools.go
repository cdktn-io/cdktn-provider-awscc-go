// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivenetwork


type MedialiveNetworkIpPools struct {
	// IP address cidr pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/medialive_network#cidr MedialiveNetwork#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

