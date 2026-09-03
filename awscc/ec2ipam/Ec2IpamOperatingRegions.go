// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipam


type Ec2IpamOperatingRegions struct {
	// The name of the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ec2_ipam#region_name Ec2Ipam#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

