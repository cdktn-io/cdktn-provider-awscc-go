// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recordset


type Route53RecordSetGeoLocation struct {
	// For geolocation resource record sets, a two-letter abbreviation that identifies a continent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53_record_set#continent_code Route53RecordSet#continent_code}
	ContinentCode *string `field:"optional" json:"continentCode" yaml:"continentCode"`
	// For geolocation resource record sets, the two-letter code for a country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53_record_set#country_code Route53RecordSet#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// For geolocation resource record sets, the two-letter code for a state of the United States.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53_record_set#subdivision_code Route53RecordSet#subdivision_code}
	SubdivisionCode *string `field:"optional" json:"subdivisionCode" yaml:"subdivisionCode"`
}

