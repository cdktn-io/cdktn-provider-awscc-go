// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipam


type Ec2IpamDefaultResourceDiscoveryOrganizationalUnitExclusions struct {
	// An AWS Organizations entity path.
	//
	// Build the path for the OU(s) using AWS Organizations IDs separated by a '/'. Include all child OUs by ending the path with '/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_ipam#organizations_entity_path Ec2Ipam#organizations_entity_path}
	OrganizationsEntityPath *string `field:"optional" json:"organizationsEntityPath" yaml:"organizationsEntityPath"`
}

