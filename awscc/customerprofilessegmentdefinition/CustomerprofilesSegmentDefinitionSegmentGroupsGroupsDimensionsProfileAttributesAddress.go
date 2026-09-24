// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddress struct {
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#city CustomerprofilesSegmentDefinition#city}
	City *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressCity `field:"optional" json:"city" yaml:"city"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#country CustomerprofilesSegmentDefinition#country}
	Country *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressCountry `field:"optional" json:"country" yaml:"country"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#county CustomerprofilesSegmentDefinition#county}
	County *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressCounty `field:"optional" json:"county" yaml:"county"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#postal_code CustomerprofilesSegmentDefinition#postal_code}
	PostalCode *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressPostalCode `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#province CustomerprofilesSegmentDefinition#province}
	Province *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressProvince `field:"optional" json:"province" yaml:"province"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#state CustomerprofilesSegmentDefinition#state}
	State *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressState `field:"optional" json:"state" yaml:"state"`
}

