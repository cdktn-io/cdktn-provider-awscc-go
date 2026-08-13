// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostssite


type OutpostsSiteShippingAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#address_line_1 OutpostsSite#address_line_1}.
	AddressLine1 *string `field:"optional" json:"addressLine1" yaml:"addressLine1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#address_line_2 OutpostsSite#address_line_2}.
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#address_line_3 OutpostsSite#address_line_3}.
	AddressLine3 *string `field:"optional" json:"addressLine3" yaml:"addressLine3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#city OutpostsSite#city}.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#contact_name OutpostsSite#contact_name}.
	ContactName *string `field:"optional" json:"contactName" yaml:"contactName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#contact_phone_number OutpostsSite#contact_phone_number}.
	ContactPhoneNumber *string `field:"optional" json:"contactPhoneNumber" yaml:"contactPhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#country_code OutpostsSite#country_code}.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#district_or_county OutpostsSite#district_or_county}.
	DistrictOrCounty *string `field:"optional" json:"districtOrCounty" yaml:"districtOrCounty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#municipality OutpostsSite#municipality}.
	Municipality *string `field:"optional" json:"municipality" yaml:"municipality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#postal_code OutpostsSite#postal_code}.
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/outposts_site#state_or_region OutpostsSite#state_or_region}.
	StateOrRegion *string `field:"optional" json:"stateOrRegion" yaml:"stateOrRegion"`
}

