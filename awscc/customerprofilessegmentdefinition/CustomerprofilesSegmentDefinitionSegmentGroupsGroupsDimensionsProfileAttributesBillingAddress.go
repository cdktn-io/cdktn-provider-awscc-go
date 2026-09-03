// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress struct {
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#city CustomerprofilesSegmentDefinition#city}
	City *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCity `field:"optional" json:"city" yaml:"city"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#country CustomerprofilesSegmentDefinition#country}
	Country *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCountry `field:"optional" json:"country" yaml:"country"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#county CustomerprofilesSegmentDefinition#county}
	County *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressCounty `field:"optional" json:"county" yaml:"county"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#postal_code CustomerprofilesSegmentDefinition#postal_code}
	PostalCode *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressPostalCode `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#province CustomerprofilesSegmentDefinition#province}
	Province *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressProvince `field:"optional" json:"province" yaml:"province"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/customerprofiles_segment_definition#state CustomerprofilesSegmentDefinition#state}
	State *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressState `field:"optional" json:"state" yaml:"state"`
}

