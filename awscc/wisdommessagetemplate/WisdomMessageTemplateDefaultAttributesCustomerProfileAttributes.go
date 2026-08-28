// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateDefaultAttributesCustomerProfileAttributes struct {
	// A unique account number that you have given to the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#account_number WisdomMessageTemplate#account_number}
	AccountNumber *string `field:"optional" json:"accountNumber" yaml:"accountNumber"`
	// Any additional information relevant to the customer's profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#additional_information WisdomMessageTemplate#additional_information}
	AdditionalInformation *string `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// The first line of a customer address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#address_1 WisdomMessageTemplate#address_1}
	Address1 *string `field:"optional" json:"address1" yaml:"address1"`
	// The second line of a customer address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#address_2 WisdomMessageTemplate#address_2}
	Address2 *string `field:"optional" json:"address2" yaml:"address2"`
	// The third line of a customer address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#address_3 WisdomMessageTemplate#address_3}
	Address3 *string `field:"optional" json:"address3" yaml:"address3"`
	// The fourth line of a customer address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#address_4 WisdomMessageTemplate#address_4}
	Address4 *string `field:"optional" json:"address4" yaml:"address4"`
	// The first line of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_address_1 WisdomMessageTemplate#billing_address_1}
	BillingAddress1 *string `field:"optional" json:"billingAddress1" yaml:"billingAddress1"`
	// The second line of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_address_2 WisdomMessageTemplate#billing_address_2}
	BillingAddress2 *string `field:"optional" json:"billingAddress2" yaml:"billingAddress2"`
	// The third line of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_address_3 WisdomMessageTemplate#billing_address_3}
	BillingAddress3 *string `field:"optional" json:"billingAddress3" yaml:"billingAddress3"`
	// The fourth line of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_address_4 WisdomMessageTemplate#billing_address_4}
	BillingAddress4 *string `field:"optional" json:"billingAddress4" yaml:"billingAddress4"`
	// The city of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_city WisdomMessageTemplate#billing_city}
	BillingCity *string `field:"optional" json:"billingCity" yaml:"billingCity"`
	// The country of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_country WisdomMessageTemplate#billing_country}
	BillingCountry *string `field:"optional" json:"billingCountry" yaml:"billingCountry"`
	// The county of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_county WisdomMessageTemplate#billing_county}
	BillingCounty *string `field:"optional" json:"billingCounty" yaml:"billingCounty"`
	// The postal code of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_postal_code WisdomMessageTemplate#billing_postal_code}
	BillingPostalCode *string `field:"optional" json:"billingPostalCode" yaml:"billingPostalCode"`
	// The province of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_province WisdomMessageTemplate#billing_province}
	BillingProvince *string `field:"optional" json:"billingProvince" yaml:"billingProvince"`
	// The state of a customer?s billing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#billing_state WisdomMessageTemplate#billing_state}
	BillingState *string `field:"optional" json:"billingState" yaml:"billingState"`
	// The customer's birth date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#birth_date WisdomMessageTemplate#birth_date}
	BirthDate *string `field:"optional" json:"birthDate" yaml:"birthDate"`
	// The customer's business email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#business_email_address WisdomMessageTemplate#business_email_address}
	BusinessEmailAddress *string `field:"optional" json:"businessEmailAddress" yaml:"businessEmailAddress"`
	// The name of the customer's business.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#business_name WisdomMessageTemplate#business_name}
	BusinessName *string `field:"optional" json:"businessName" yaml:"businessName"`
	// The customer's business phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#business_phone_number WisdomMessageTemplate#business_phone_number}
	BusinessPhoneNumber *string `field:"optional" json:"businessPhoneNumber" yaml:"businessPhoneNumber"`
	// The city in which a customer lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#city WisdomMessageTemplate#city}
	City *string `field:"optional" json:"city" yaml:"city"`
	// The country in which a customer lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#country WisdomMessageTemplate#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// The county in which a customer lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#county WisdomMessageTemplate#county}
	County *string `field:"optional" json:"county" yaml:"county"`
	// The custom attributes that are used with the message template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#custom WisdomMessageTemplate#custom}
	Custom *map[string]*string `field:"optional" json:"custom" yaml:"custom"`
	// The customer's email address, which has not been specified as a personal or business address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#email_address WisdomMessageTemplate#email_address}
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The customer's first name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#first_name WisdomMessageTemplate#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The customer's gender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#gender WisdomMessageTemplate#gender}
	Gender *string `field:"optional" json:"gender" yaml:"gender"`
	// The customer's home phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#home_phone_number WisdomMessageTemplate#home_phone_number}
	HomePhoneNumber *string `field:"optional" json:"homePhoneNumber" yaml:"homePhoneNumber"`
	// The customer's last name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#last_name WisdomMessageTemplate#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// The first line of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_address_1 WisdomMessageTemplate#mailing_address_1}
	MailingAddress1 *string `field:"optional" json:"mailingAddress1" yaml:"mailingAddress1"`
	// The second line of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_address_2 WisdomMessageTemplate#mailing_address_2}
	MailingAddress2 *string `field:"optional" json:"mailingAddress2" yaml:"mailingAddress2"`
	// The third line of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_address_3 WisdomMessageTemplate#mailing_address_3}
	MailingAddress3 *string `field:"optional" json:"mailingAddress3" yaml:"mailingAddress3"`
	// The fourth line of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_address_4 WisdomMessageTemplate#mailing_address_4}
	MailingAddress4 *string `field:"optional" json:"mailingAddress4" yaml:"mailingAddress4"`
	// The city of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_city WisdomMessageTemplate#mailing_city}
	MailingCity *string `field:"optional" json:"mailingCity" yaml:"mailingCity"`
	// The country of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_country WisdomMessageTemplate#mailing_country}
	MailingCountry *string `field:"optional" json:"mailingCountry" yaml:"mailingCountry"`
	// The county of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_county WisdomMessageTemplate#mailing_county}
	MailingCounty *string `field:"optional" json:"mailingCounty" yaml:"mailingCounty"`
	// The postal code of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_postal_code WisdomMessageTemplate#mailing_postal_code}
	MailingPostalCode *string `field:"optional" json:"mailingPostalCode" yaml:"mailingPostalCode"`
	// The province of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_province WisdomMessageTemplate#mailing_province}
	MailingProvince *string `field:"optional" json:"mailingProvince" yaml:"mailingProvince"`
	// The state of a customer?s mailing address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mailing_state WisdomMessageTemplate#mailing_state}
	MailingState *string `field:"optional" json:"mailingState" yaml:"mailingState"`
	// The customer's middle name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#middle_name WisdomMessageTemplate#middle_name}
	MiddleName *string `field:"optional" json:"middleName" yaml:"middleName"`
	// The customer's mobile phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#mobile_phone_number WisdomMessageTemplate#mobile_phone_number}
	MobilePhoneNumber *string `field:"optional" json:"mobilePhoneNumber" yaml:"mobilePhoneNumber"`
	// The customer's party type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#party_type WisdomMessageTemplate#party_type}
	PartyType *string `field:"optional" json:"partyType" yaml:"partyType"`
	// The customer's phone number, which has not been specified as a mobile, home, or business number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#phone_number WisdomMessageTemplate#phone_number}
	PhoneNumber *string `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
	// The postal code of a customer address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#postal_code WisdomMessageTemplate#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The ARN of a customer profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#profile_arn WisdomMessageTemplate#profile_arn}
	ProfileArn *string `field:"optional" json:"profileArn" yaml:"profileArn"`
	// The unique identifier of a customer profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#profile_id WisdomMessageTemplate#profile_id}
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// The province in which a customer lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#province WisdomMessageTemplate#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The first line of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_address_1 WisdomMessageTemplate#shipping_address_1}
	ShippingAddress1 *string `field:"optional" json:"shippingAddress1" yaml:"shippingAddress1"`
	// The second line of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_address_2 WisdomMessageTemplate#shipping_address_2}
	ShippingAddress2 *string `field:"optional" json:"shippingAddress2" yaml:"shippingAddress2"`
	// The third line of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_address_3 WisdomMessageTemplate#shipping_address_3}
	ShippingAddress3 *string `field:"optional" json:"shippingAddress3" yaml:"shippingAddress3"`
	// The fourth line of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_address_4 WisdomMessageTemplate#shipping_address_4}
	ShippingAddress4 *string `field:"optional" json:"shippingAddress4" yaml:"shippingAddress4"`
	// The city of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_city WisdomMessageTemplate#shipping_city}
	ShippingCity *string `field:"optional" json:"shippingCity" yaml:"shippingCity"`
	// The country of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_country WisdomMessageTemplate#shipping_country}
	ShippingCountry *string `field:"optional" json:"shippingCountry" yaml:"shippingCountry"`
	// The county of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_county WisdomMessageTemplate#shipping_county}
	ShippingCounty *string `field:"optional" json:"shippingCounty" yaml:"shippingCounty"`
	// The postal code of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_postal_code WisdomMessageTemplate#shipping_postal_code}
	ShippingPostalCode *string `field:"optional" json:"shippingPostalCode" yaml:"shippingPostalCode"`
	// The province of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_province WisdomMessageTemplate#shipping_province}
	ShippingProvince *string `field:"optional" json:"shippingProvince" yaml:"shippingProvince"`
	// The state of a customer?s shipping address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#shipping_state WisdomMessageTemplate#shipping_state}
	ShippingState *string `field:"optional" json:"shippingState" yaml:"shippingState"`
	// The state in which a customer lives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_message_template#state WisdomMessageTemplate#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

