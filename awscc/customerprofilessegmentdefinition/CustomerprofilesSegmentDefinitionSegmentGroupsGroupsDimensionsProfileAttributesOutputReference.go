// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilessegmentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/customerprofilessegmentdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference interface {
	cdktn.ComplexObject
	AccountNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumberOutputReference
	AccountNumberInput() interface{}
	AdditionalInformation() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformationOutputReference
	AdditionalInformationInput() interface{}
	Address() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressOutputReference
	AddressInput() interface{}
	Attributes() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap
	AttributesInput() interface{}
	BillingAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference
	BillingAddressInput() interface{}
	BirthDate() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDateOutputReference
	BirthDateInput() interface{}
	BusinessEmailAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddressOutputReference
	BusinessEmailAddressInput() interface{}
	BusinessName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessNameOutputReference
	BusinessNameInput() interface{}
	BusinessPhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumberOutputReference
	BusinessPhoneNumberInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EmailAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddressOutputReference
	EmailAddressInput() interface{}
	FirstName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstNameOutputReference
	FirstNameInput() interface{}
	// Experimental.
	Fqn() *string
	GenderString() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderStringOutputReference
	GenderStringInput() interface{}
	HomePhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumberOutputReference
	HomePhoneNumberInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastNameOutputReference
	LastNameInput() interface{}
	MailingAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressOutputReference
	MailingAddressInput() interface{}
	MiddleName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleNameOutputReference
	MiddleNameInput() interface{}
	MobilePhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumberOutputReference
	MobilePhoneNumberInput() interface{}
	PartyTypeString() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeStringOutputReference
	PartyTypeStringInput() interface{}
	PersonalEmailAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddressOutputReference
	PersonalEmailAddressInput() interface{}
	PhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumberOutputReference
	PhoneNumberInput() interface{}
	ProfileType() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileTypeOutputReference
	ProfileTypeInput() interface{}
	ShippingAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference
	ShippingAddressInput() interface{}
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
	PutAccountNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumber)
	PutAdditionalInformation(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformation)
	PutAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddress)
	PutAttributes(value interface{})
	PutBillingAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress)
	PutBirthDate(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDate)
	PutBusinessEmailAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddress)
	PutBusinessName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessName)
	PutBusinessPhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumber)
	PutEmailAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddress)
	PutFirstName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstName)
	PutGenderString(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderString)
	PutHomePhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumber)
	PutLastName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastName)
	PutMailingAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddress)
	PutMiddleName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleName)
	PutMobilePhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumber)
	PutPartyTypeString(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeString)
	PutPersonalEmailAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddress)
	PutPhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumber)
	PutProfileType(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileType)
	PutShippingAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddress)
	ResetAccountNumber()
	ResetAdditionalInformation()
	ResetAddress()
	ResetAttributes()
	ResetBillingAddress()
	ResetBirthDate()
	ResetBusinessEmailAddress()
	ResetBusinessName()
	ResetBusinessPhoneNumber()
	ResetEmailAddress()
	ResetFirstName()
	ResetGenderString()
	ResetHomePhoneNumber()
	ResetLastName()
	ResetMailingAddress()
	ResetMiddleName()
	ResetMobilePhoneNumber()
	ResetPartyTypeString()
	ResetPersonalEmailAddress()
	ResetPhoneNumber()
	ResetProfileType()
	ResetShippingAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference
type jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AccountNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumberOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumberOutputReference
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AccountNumberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AdditionalInformation() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformationOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformationOutputReference
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AdditionalInformationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Address() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressOutputReference
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Attributes() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BillingAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference
	_jsii_.Get(
		j,
		"billingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BillingAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BirthDate() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDateOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDateOutputReference
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BirthDateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"birthDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessEmailAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddressOutputReference
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessEmailAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"businessEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessNameOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessNameOutputReference
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"businessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessPhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumberOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumberOutputReference
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessPhoneNumberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"businessPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) EmailAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddressOutputReference
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) EmailAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) FirstName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstNameOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstNameOutputReference
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) FirstNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GenderString() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderStringOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderStringOutputReference
	_jsii_.Get(
		j,
		"genderString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GenderStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"genderStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) HomePhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumberOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumberOutputReference
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) HomePhoneNumberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"homePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) LastName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastNameOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastNameOutputReference
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) LastNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MailingAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressOutputReference
	_jsii_.Get(
		j,
		"mailingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MailingAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MiddleName() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleNameOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleNameOutputReference
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MiddleNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MobilePhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumberOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumberOutputReference
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MobilePhoneNumberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mobilePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PartyTypeString() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeStringOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeStringOutputReference
	_jsii_.Get(
		j,
		"partyTypeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PartyTypeStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partyTypeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PersonalEmailAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddressOutputReference
	_jsii_.Get(
		j,
		"personalEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PersonalEmailAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"personalEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PhoneNumber() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumberOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumberOutputReference
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PhoneNumberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ProfileType() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileTypeOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileTypeOutputReference
	_jsii_.Get(
		j,
		"profileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ProfileTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ShippingAddress() CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference {
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference
	_jsii_.Get(
		j,
		"shippingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ShippingAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shippingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesSegmentDefinition.CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference_Override(c CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesSegmentDefinition.CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutAccountNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumber) {
	if err := c.validatePutAccountNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccountNumber",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutAdditionalInformation(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformation) {
	if err := c.validatePutAdditionalInformationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdditionalInformation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddress) {
	if err := c.validatePutAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutAttributes(value interface{}) {
	if err := c.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAttributes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutBillingAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddress) {
	if err := c.validatePutBillingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBillingAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutBirthDate(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDate) {
	if err := c.validatePutBirthDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBirthDate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutBusinessEmailAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddress) {
	if err := c.validatePutBusinessEmailAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBusinessEmailAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutBusinessName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessName) {
	if err := c.validatePutBusinessNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBusinessName",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutBusinessPhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumber) {
	if err := c.validatePutBusinessPhoneNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBusinessPhoneNumber",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutEmailAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddress) {
	if err := c.validatePutEmailAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmailAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutFirstName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstName) {
	if err := c.validatePutFirstNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirstName",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutGenderString(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderString) {
	if err := c.validatePutGenderStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGenderString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutHomePhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumber) {
	if err := c.validatePutHomePhoneNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHomePhoneNumber",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutLastName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastName) {
	if err := c.validatePutLastNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLastName",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutMailingAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddress) {
	if err := c.validatePutMailingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMailingAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutMiddleName(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleName) {
	if err := c.validatePutMiddleNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMiddleName",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutMobilePhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumber) {
	if err := c.validatePutMobilePhoneNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMobilePhoneNumber",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutPartyTypeString(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeString) {
	if err := c.validatePutPartyTypeStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPartyTypeString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutPersonalEmailAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddress) {
	if err := c.validatePutPersonalEmailAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPersonalEmailAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutPhoneNumber(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumber) {
	if err := c.validatePutPhoneNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhoneNumber",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutProfileType(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileType) {
	if err := c.validatePutProfileTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProfileType",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PutShippingAddress(value *CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddress) {
	if err := c.validatePutShippingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShippingAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetAccountNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetAdditionalInformation() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalInformation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetBillingAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetBillingAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetBirthDate() {
	_jsii_.InvokeVoid(
		c,
		"resetBirthDate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetBusinessEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetBusinessEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetBusinessName() {
	_jsii_.InvokeVoid(
		c,
		"resetBusinessName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetBusinessPhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetBusinessPhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetFirstName() {
	_jsii_.InvokeVoid(
		c,
		"resetFirstName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetGenderString() {
	_jsii_.InvokeVoid(
		c,
		"resetGenderString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetHomePhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetHomePhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetLastName() {
	_jsii_.InvokeVoid(
		c,
		"resetLastName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetMailingAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetMailingAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetMiddleName() {
	_jsii_.InvokeVoid(
		c,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetMobilePhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetMobilePhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetPartyTypeString() {
	_jsii_.InvokeVoid(
		c,
		"resetPartyTypeString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetPersonalEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetPersonalEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetProfileType() {
	_jsii_.InvokeVoid(
		c,
		"resetProfileType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ResetShippingAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetShippingAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

