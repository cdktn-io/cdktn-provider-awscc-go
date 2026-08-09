// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddress struct {
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_segment_definition#city CustomerprofilesSegmentDefinition#city}
	City *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCity `field:"optional" json:"city" yaml:"city"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_segment_definition#country CustomerprofilesSegmentDefinition#country}
	Country *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCountry `field:"optional" json:"country" yaml:"country"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_segment_definition#county CustomerprofilesSegmentDefinition#county}
	County *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressCounty `field:"optional" json:"county" yaml:"county"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_segment_definition#postal_code CustomerprofilesSegmentDefinition#postal_code}
	PostalCode *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressPostalCode `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_segment_definition#province CustomerprofilesSegmentDefinition#province}
	Province *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressProvince `field:"optional" json:"province" yaml:"province"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/customerprofiles_segment_definition#state CustomerprofilesSegmentDefinition#state}
	State *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressState `field:"optional" json:"state" yaml:"state"`
}

