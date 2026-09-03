// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53cidrcollection


type Route53CidrCollectionLocations struct {
	// A list of CIDR blocks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53_cidr_collection#cidr_list Route53CidrCollection#cidr_list}
	CidrList *[]*string `field:"optional" json:"cidrList" yaml:"cidrList"`
	// The name of the location that is associated with the CIDR collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53_cidr_collection#location_name Route53CidrCollection#location_name}
	LocationName *string `field:"optional" json:"locationName" yaml:"locationName"`
}

