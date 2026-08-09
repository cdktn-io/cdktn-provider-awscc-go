// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscccustomerprofilessegmentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscccustomerprofilessegmentdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference interface {
	cdktn.ComplexObject
	AccountNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumberOutputReference
	AdditionalInformation() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformationOutputReference
	Address() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressOutputReference
	Attributes() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap
	BillingAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference
	BirthDate() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDateOutputReference
	BusinessEmailAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddressOutputReference
	BusinessName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessNameOutputReference
	BusinessPhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumberOutputReference
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
	EmailAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddressOutputReference
	FirstName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstNameOutputReference
	// Experimental.
	Fqn() *string
	GenderString() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderStringOutputReference
	HomePhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumberOutputReference
	InternalValue() *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes
	SetInternalValue(val *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes)
	LastName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastNameOutputReference
	MailingAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressOutputReference
	MiddleName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleNameOutputReference
	MobilePhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumberOutputReference
	PartyTypeString() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeStringOutputReference
	PersonalEmailAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddressOutputReference
	PhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumberOutputReference
	ProfileType() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileTypeOutputReference
	ShippingAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference
type jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AccountNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumberOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAccountNumberOutputReference
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) AdditionalInformation() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformationOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAdditionalInformationOutputReference
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Address() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAddressOutputReference
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Attributes() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BillingAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBillingAddressOutputReference
	_jsii_.Get(
		j,
		"billingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BirthDate() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDateOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBirthDateOutputReference
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessEmailAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessEmailAddressOutputReference
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessNameOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessNameOutputReference
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) BusinessPhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumberOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesBusinessPhoneNumberOutputReference
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) EmailAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesEmailAddressOutputReference
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) FirstName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstNameOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesFirstNameOutputReference
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GenderString() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderStringOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesGenderStringOutputReference
	_jsii_.Get(
		j,
		"genderString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) HomePhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumberOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesHomePhoneNumberOutputReference
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) InternalValue() *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes {
	var returns *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) LastName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastNameOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesLastNameOutputReference
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MailingAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMailingAddressOutputReference
	_jsii_.Get(
		j,
		"mailingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MiddleName() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleNameOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMiddleNameOutputReference
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) MobilePhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumberOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesMobilePhoneNumberOutputReference
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PartyTypeString() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeStringOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPartyTypeStringOutputReference
	_jsii_.Get(
		j,
		"partyTypeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PersonalEmailAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPersonalEmailAddressOutputReference
	_jsii_.Get(
		j,
		"personalEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) PhoneNumber() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumberOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesPhoneNumberOutputReference
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ProfileType() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileTypeOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesProfileTypeOutputReference
	_jsii_.Get(
		j,
		"profileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ShippingAddress() DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference {
	var returns DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesShippingAddressOutputReference
	_jsii_.Get(
		j,
		"shippingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCustomerprofilesSegmentDefinition.DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference_Override(d DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCustomerprofilesSegmentDefinition.DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetInternalValue(val *DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

