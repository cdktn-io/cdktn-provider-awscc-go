// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamresourcediscovery


type Ec2IpamResourceDiscoveryOperatingRegions struct {
	// The name of the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_resource_discovery#region_name Ec2IpamResourceDiscovery#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

