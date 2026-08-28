// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddress struct {
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#city CustomerprofilesSegmentDefinition#city}
	City *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressCity `field:"optional" json:"city" yaml:"city"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#country CustomerprofilesSegmentDefinition#country}
	Country *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressCountry `field:"optional" json:"country" yaml:"country"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#county CustomerprofilesSegmentDefinition#county}
	County *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressCounty `field:"optional" json:"county" yaml:"county"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#postal_code CustomerprofilesSegmentDefinition#postal_code}
	PostalCode *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressPostalCode `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#province CustomerprofilesSegmentDefinition#province}
	Province *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressProvince `field:"optional" json:"province" yaml:"province"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_segment_definition#state CustomerprofilesSegmentDefinition#state}
	State *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressState `field:"optional" json:"state" yaml:"state"`
}

