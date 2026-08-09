// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdommessagetemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference interface {
	cdktn.ComplexObject
	AccountNumber() *string
	SetAccountNumber(val *string)
	AccountNumberInput() *string
	AdditionalInformation() *string
	SetAdditionalInformation(val *string)
	AdditionalInformationInput() *string
	Address1() *string
	SetAddress1(val *string)
	Address1Input() *string
	Address2() *string
	SetAddress2(val *string)
	Address2Input() *string
	Address3() *string
	SetAddress3(val *string)
	Address3Input() *string
	Address4() *string
	SetAddress4(val *string)
	Address4Input() *string
	BillingAddress1() *string
	SetBillingAddress1(val *string)
	BillingAddress1Input() *string
	BillingAddress2() *string
	SetBillingAddress2(val *string)
	BillingAddress2Input() *string
	BillingAddress3() *string
	SetBillingAddress3(val *string)
	BillingAddress3Input() *string
	BillingAddress4() *string
	SetBillingAddress4(val *string)
	BillingAddress4Input() *string
	BillingCity() *string
	SetBillingCity(val *string)
	BillingCityInput() *string
	BillingCountry() *string
	SetBillingCountry(val *string)
	BillingCountryInput() *string
	BillingCounty() *string
	SetBillingCounty(val *string)
	BillingCountyInput() *string
	BillingPostalCode() *string
	SetBillingPostalCode(val *string)
	BillingPostalCodeInput() *string
	BillingProvince() *string
	SetBillingProvince(val *string)
	BillingProvinceInput() *string
	BillingState() *string
	SetBillingState(val *string)
	BillingStateInput() *string
	BirthDate() *string
	SetBirthDate(val *string)
	BirthDateInput() *string
	BusinessEmailAddress() *string
	SetBusinessEmailAddress(val *string)
	BusinessEmailAddressInput() *string
	BusinessName() *string
	SetBusinessName(val *string)
	BusinessNameInput() *string
	BusinessPhoneNumber() *string
	SetBusinessPhoneNumber(val *string)
	BusinessPhoneNumberInput() *string
	City() *string
	SetCity(val *string)
	CityInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	County() *string
	SetCounty(val *string)
	CountyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Custom() *map[string]*string
	SetCustom(val *map[string]*string)
	CustomInput() *map[string]*string
	EmailAddress() *string
	SetEmailAddress(val *string)
	EmailAddressInput() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	Fqn() *string
	Gender() *string
	SetGender(val *string)
	GenderInput() *string
	HomePhoneNumber() *string
	SetHomePhoneNumber(val *string)
	HomePhoneNumberInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	MailingAddress1() *string
	SetMailingAddress1(val *string)
	MailingAddress1Input() *string
	MailingAddress2() *string
	SetMailingAddress2(val *string)
	MailingAddress2Input() *string
	MailingAddress3() *string
	SetMailingAddress3(val *string)
	MailingAddress3Input() *string
	MailingAddress4() *string
	SetMailingAddress4(val *string)
	MailingAddress4Input() *string
	MailingCity() *string
	SetMailingCity(val *string)
	MailingCityInput() *string
	MailingCountry() *string
	SetMailingCountry(val *string)
	MailingCountryInput() *string
	MailingCounty() *string
	SetMailingCounty(val *string)
	MailingCountyInput() *string
	MailingPostalCode() *string
	SetMailingPostalCode(val *string)
	MailingPostalCodeInput() *string
	MailingProvince() *string
	SetMailingProvince(val *string)
	MailingProvinceInput() *string
	MailingState() *string
	SetMailingState(val *string)
	MailingStateInput() *string
	MiddleName() *string
	SetMiddleName(val *string)
	MiddleNameInput() *string
	MobilePhoneNumber() *string
	SetMobilePhoneNumber(val *string)
	MobilePhoneNumberInput() *string
	PartyType() *string
	SetPartyType(val *string)
	PartyTypeInput() *string
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	ProfileArn() *string
	SetProfileArn(val *string)
	ProfileArnInput() *string
	ProfileId() *string
	SetProfileId(val *string)
	ProfileIdInput() *string
	Province() *string
	SetProvince(val *string)
	ProvinceInput() *string
	ShippingAddress1() *string
	SetShippingAddress1(val *string)
	ShippingAddress1Input() *string
	ShippingAddress2() *string
	SetShippingAddress2(val *string)
	ShippingAddress2Input() *string
	ShippingAddress3() *string
	SetShippingAddress3(val *string)
	ShippingAddress3Input() *string
	ShippingAddress4() *string
	SetShippingAddress4(val *string)
	ShippingAddress4Input() *string
	ShippingCity() *string
	SetShippingCity(val *string)
	ShippingCityInput() *string
	ShippingCountry() *string
	SetShippingCountry(val *string)
	ShippingCountryInput() *string
	ShippingCounty() *string
	SetShippingCounty(val *string)
	ShippingCountyInput() *string
	ShippingPostalCode() *string
	SetShippingPostalCode(val *string)
	ShippingPostalCodeInput() *string
	ShippingProvince() *string
	SetShippingProvince(val *string)
	ShippingProvinceInput() *string
	ShippingState() *string
	SetShippingState(val *string)
	ShippingStateInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetAccountNumber()
	ResetAdditionalInformation()
	ResetAddress1()
	ResetAddress2()
	ResetAddress3()
	ResetAddress4()
	ResetBillingAddress1()
	ResetBillingAddress2()
	ResetBillingAddress3()
	ResetBillingAddress4()
	ResetBillingCity()
	ResetBillingCountry()
	ResetBillingCounty()
	ResetBillingPostalCode()
	ResetBillingProvince()
	ResetBillingState()
	ResetBirthDate()
	ResetBusinessEmailAddress()
	ResetBusinessName()
	ResetBusinessPhoneNumber()
	ResetCity()
	ResetCountry()
	ResetCounty()
	ResetCustom()
	ResetEmailAddress()
	ResetFirstName()
	ResetGender()
	ResetHomePhoneNumber()
	ResetLastName()
	ResetMailingAddress1()
	ResetMailingAddress2()
	ResetMailingAddress3()
	ResetMailingAddress4()
	ResetMailingCity()
	ResetMailingCountry()
	ResetMailingCounty()
	ResetMailingPostalCode()
	ResetMailingProvince()
	ResetMailingState()
	ResetMiddleName()
	ResetMobilePhoneNumber()
	ResetPartyType()
	ResetPhoneNumber()
	ResetPostalCode()
	ResetProfileArn()
	ResetProfileId()
	ResetProvince()
	ResetShippingAddress1()
	ResetShippingAddress2()
	ResetShippingAddress3()
	ResetShippingAddress4()
	ResetShippingCity()
	ResetShippingCountry()
	ResetShippingCounty()
	ResetShippingPostalCode()
	ResetShippingProvince()
	ResetShippingState()
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference
type jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) AccountNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) AccountNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) AdditionalInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) AdditionalInformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCountry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCountry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCountryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCountyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingPostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingPostalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingPostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingPostalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingProvince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProvince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProvinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BirthDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BirthDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) County() *string {
	var returns *string
	_jsii_.Get(
		j,
		"county",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) CountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Custom() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) CustomInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Gender() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GenderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) HomePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) HomePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCountry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCountry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCountryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCountyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingPostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingPostalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingPostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingPostalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingProvince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingProvince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingProvinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MobilePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MobilePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PartyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PartyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCountry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCountry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCountryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCountyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingPostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingPostalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingPostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingPostalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingProvince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingProvince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingProvinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomMessageTemplate.WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference_Override(w WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomMessageTemplate.WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetAccountNumber(val *string) {
	if err := j.validateSetAccountNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountNumber",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetAdditionalInformation(val *string) {
	if err := j.validateSetAdditionalInformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInformation",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetAddress1(val *string) {
	if err := j.validateSetAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address1",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetAddress2(val *string) {
	if err := j.validateSetAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address2",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetAddress3(val *string) {
	if err := j.validateSetAddress3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address3",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetAddress4(val *string) {
	if err := j.validateSetAddress4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address4",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingAddress1(val *string) {
	if err := j.validateSetBillingAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAddress1",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingAddress2(val *string) {
	if err := j.validateSetBillingAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAddress2",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingAddress3(val *string) {
	if err := j.validateSetBillingAddress3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAddress3",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingAddress4(val *string) {
	if err := j.validateSetBillingAddress4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingAddress4",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingCity(val *string) {
	if err := j.validateSetBillingCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingCity",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingCountry(val *string) {
	if err := j.validateSetBillingCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingCountry",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingCounty(val *string) {
	if err := j.validateSetBillingCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingCounty",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingPostalCode(val *string) {
	if err := j.validateSetBillingPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingPostalCode",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingProvince(val *string) {
	if err := j.validateSetBillingProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingProvince",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBillingState(val *string) {
	if err := j.validateSetBillingStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingState",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBirthDate(val *string) {
	if err := j.validateSetBirthDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"birthDate",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBusinessEmailAddress(val *string) {
	if err := j.validateSetBusinessEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessEmailAddress",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBusinessName(val *string) {
	if err := j.validateSetBusinessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessName",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetBusinessPhoneNumber(val *string) {
	if err := j.validateSetBusinessPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetCounty(val *string) {
	if err := j.validateSetCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"county",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetCustom(val *map[string]*string) {
	if err := j.validateSetCustomParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"custom",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetGender(val *string) {
	if err := j.validateSetGenderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gender",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetHomePhoneNumber(val *string) {
	if err := j.validateSetHomePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingAddress1(val *string) {
	if err := j.validateSetMailingAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingAddress1",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingAddress2(val *string) {
	if err := j.validateSetMailingAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingAddress2",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingAddress3(val *string) {
	if err := j.validateSetMailingAddress3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingAddress3",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingAddress4(val *string) {
	if err := j.validateSetMailingAddress4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingAddress4",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingCity(val *string) {
	if err := j.validateSetMailingCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingCity",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingCountry(val *string) {
	if err := j.validateSetMailingCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingCountry",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingCounty(val *string) {
	if err := j.validateSetMailingCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingCounty",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingPostalCode(val *string) {
	if err := j.validateSetMailingPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingPostalCode",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingProvince(val *string) {
	if err := j.validateSetMailingProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingProvince",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMailingState(val *string) {
	if err := j.validateSetMailingStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailingState",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetMobilePhoneNumber(val *string) {
	if err := j.validateSetMobilePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetPartyType(val *string) {
	if err := j.validateSetPartyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partyType",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetProfileArn(val *string) {
	if err := j.validateSetProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileArn",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetProfileId(val *string) {
	if err := j.validateSetProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileId",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetProvince(val *string) {
	if err := j.validateSetProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingAddress1(val *string) {
	if err := j.validateSetShippingAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingAddress1",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingAddress2(val *string) {
	if err := j.validateSetShippingAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingAddress2",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingAddress3(val *string) {
	if err := j.validateSetShippingAddress3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingAddress3",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingAddress4(val *string) {
	if err := j.validateSetShippingAddress4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingAddress4",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingCity(val *string) {
	if err := j.validateSetShippingCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingCity",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingCountry(val *string) {
	if err := j.validateSetShippingCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingCountry",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingCounty(val *string) {
	if err := j.validateSetShippingCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingCounty",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingPostalCode(val *string) {
	if err := j.validateSetShippingPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingPostalCode",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingProvince(val *string) {
	if err := j.validateSetShippingProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingProvince",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetShippingState(val *string) {
	if err := j.validateSetShippingStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shippingState",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetAccountNumber() {
	_jsii_.InvokeVoid(
		w,
		"resetAccountNumber",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetAdditionalInformation() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalInformation",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetAddress1() {
	_jsii_.InvokeVoid(
		w,
		"resetAddress1",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetAddress2() {
	_jsii_.InvokeVoid(
		w,
		"resetAddress2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetAddress3() {
	_jsii_.InvokeVoid(
		w,
		"resetAddress3",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetAddress4() {
	_jsii_.InvokeVoid(
		w,
		"resetAddress4",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingAddress1() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingAddress1",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingAddress2() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingAddress2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingAddress3() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingAddress3",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingAddress4() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingAddress4",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingCity() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingCity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingCountry() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingCountry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingCounty() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingCounty",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingPostalCode() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingPostalCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingProvince() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingProvince",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBillingState() {
	_jsii_.InvokeVoid(
		w,
		"resetBillingState",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBirthDate() {
	_jsii_.InvokeVoid(
		w,
		"resetBirthDate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBusinessEmailAddress() {
	_jsii_.InvokeVoid(
		w,
		"resetBusinessEmailAddress",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBusinessName() {
	_jsii_.InvokeVoid(
		w,
		"resetBusinessName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetBusinessPhoneNumber() {
	_jsii_.InvokeVoid(
		w,
		"resetBusinessPhoneNumber",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		w,
		"resetCity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		w,
		"resetCountry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetCounty() {
	_jsii_.InvokeVoid(
		w,
		"resetCounty",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		w,
		"resetCustom",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		w,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetFirstName() {
	_jsii_.InvokeVoid(
		w,
		"resetFirstName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetGender() {
	_jsii_.InvokeVoid(
		w,
		"resetGender",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetHomePhoneNumber() {
	_jsii_.InvokeVoid(
		w,
		"resetHomePhoneNumber",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetLastName() {
	_jsii_.InvokeVoid(
		w,
		"resetLastName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingAddress1() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingAddress1",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingAddress2() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingAddress2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingAddress3() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingAddress3",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingAddress4() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingAddress4",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingCity() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingCity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingCountry() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingCountry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingCounty() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingCounty",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingPostalCode() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingPostalCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingProvince() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingProvince",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMailingState() {
	_jsii_.InvokeVoid(
		w,
		"resetMailingState",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMiddleName() {
	_jsii_.InvokeVoid(
		w,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetMobilePhoneNumber() {
	_jsii_.InvokeVoid(
		w,
		"resetMobilePhoneNumber",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetPartyType() {
	_jsii_.InvokeVoid(
		w,
		"resetPartyType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		w,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		w,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetProfileArn() {
	_jsii_.InvokeVoid(
		w,
		"resetProfileArn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetProfileId() {
	_jsii_.InvokeVoid(
		w,
		"resetProfileId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetProvince() {
	_jsii_.InvokeVoid(
		w,
		"resetProvince",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingAddress1() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingAddress1",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingAddress2() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingAddress2",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingAddress3() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingAddress3",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingAddress4() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingAddress4",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingCity() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingCity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingCountry() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingCountry",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingCounty() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingCounty",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingPostalCode() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingPostalCode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingProvince() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingProvince",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetShippingState() {
	_jsii_.InvokeVoid(
		w,
		"resetShippingState",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		w,
		"resetState",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

