// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontdistributiontenant


type CloudfrontDistributionTenantCustomizationsGeoRestrictions struct {
	// The locations for geographic restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#locations CloudfrontDistributionTenant#locations}
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// The method that you want to use to restrict distribution of your content by country:   +  ``none``: No geographic restriction is enabled, meaning access to content is not restricted by client geo location.
	//
	// +  ``blacklist``: The ``Location`` elements specify the countries in which you don't want CloudFront to distribute your content.
	//   +  ``whitelist``: The ``Location`` elements specify the countries in which you want CloudFront to distribute your content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudfront_distribution_tenant#restriction_type CloudfrontDistributionTenant#restriction_type}
	RestrictionType *string `field:"optional" json:"restrictionType" yaml:"restrictionType"`
}

