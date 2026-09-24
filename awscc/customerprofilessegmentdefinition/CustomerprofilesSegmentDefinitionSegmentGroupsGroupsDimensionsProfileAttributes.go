// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition


type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes struct {
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#account_number CustomerprofilesSegmentDefinition#account_number}
	AccountNumber *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumber `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// Specifies criteria for a segment using extended-length string values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#additional_information CustomerprofilesSegmentDefinition#additional_information}
	AdditionalInformation *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformation `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// The address based criteria for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#address CustomerprofilesSegmentDefinition#address}
	Address *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddress `field:"optional" json:"address" yaml:"address"`
	// One or more custom attributes to use as criteria for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#attributes CustomerprofilesSegmentDefinition#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The address based criteria for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#billing_address CustomerprofilesSegmentDefinition#billing_address}
	BillingAddress *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress `field:"optional" json:"billingAddress" yaml:"billingAddress"`
	// Specifies date based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#birth_date CustomerprofilesSegmentDefinition#birth_date}
	BirthDate *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDate `field:"optional" json:"birthDate" yaml:"birthDate"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#business_email_address CustomerprofilesSegmentDefinition#business_email_address}
	BusinessEmailAddress *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddress `field:"optional" json:"businessEmailAddress" yaml:"businessEmailAddress"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#business_name CustomerprofilesSegmentDefinition#business_name}
	BusinessName *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessName `field:"optional" json:"businessName" yaml:"businessName"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#business_phone_number CustomerprofilesSegmentDefinition#business_phone_number}
	BusinessPhoneNumber *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumber `field:"optional" json:"businessPhoneNumber" yaml:"businessPhoneNumber"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#email_address CustomerprofilesSegmentDefinition#email_address}
	EmailAddress *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddress `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#first_name CustomerprofilesSegmentDefinition#first_name}
	FirstName *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstName `field:"optional" json:"firstName" yaml:"firstName"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#gender_string CustomerprofilesSegmentDefinition#gender_string}
	GenderString *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderString `field:"optional" json:"genderString" yaml:"genderString"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#home_phone_number CustomerprofilesSegmentDefinition#home_phone_number}
	HomePhoneNumber *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumber `field:"optional" json:"homePhoneNumber" yaml:"homePhoneNumber"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#last_name CustomerprofilesSegmentDefinition#last_name}
	LastName *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastName `field:"optional" json:"lastName" yaml:"lastName"`
	// The address based criteria for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#mailing_address CustomerprofilesSegmentDefinition#mailing_address}
	MailingAddress *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddress `field:"optional" json:"mailingAddress" yaml:"mailingAddress"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#middle_name CustomerprofilesSegmentDefinition#middle_name}
	MiddleName *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleName `field:"optional" json:"middleName" yaml:"middleName"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#mobile_phone_number CustomerprofilesSegmentDefinition#mobile_phone_number}
	MobilePhoneNumber *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumber `field:"optional" json:"mobilePhoneNumber" yaml:"mobilePhoneNumber"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#party_type_string CustomerprofilesSegmentDefinition#party_type_string}
	PartyTypeString *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeString `field:"optional" json:"partyTypeString" yaml:"partyTypeString"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#personal_email_address CustomerprofilesSegmentDefinition#personal_email_address}
	PersonalEmailAddress *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddress `field:"optional" json:"personalEmailAddress" yaml:"personalEmailAddress"`
	// Specifies profile based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#phone_number CustomerprofilesSegmentDefinition#phone_number}
	PhoneNumber *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumber `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// Specifies profile type based criteria for a segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#profile_type CustomerprofilesSegmentDefinition#profile_type}
	ProfileType *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileType `field:"optional" json:"profileType" yaml:"profileType"`
	// The address based criteria for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/customerprofiles_segment_definition#shipping_address CustomerprofilesSegmentDefinition#shipping_address}
	ShippingAddress *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddress `field:"optional" json:"shippingAddress" yaml:"shippingAddress"`
}

