// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccwisdommessagetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccwisdommessagetemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference interface {
	cdktn.ComplexObject
	AccountNumber() *string
	AdditionalInformation() *string
	Address1() *string
	Address2() *string
	Address3() *string
	Address4() *string
	BillingAddress1() *string
	BillingAddress2() *string
	BillingAddress3() *string
	BillingAddress4() *string
	BillingCity() *string
	BillingCountry() *string
	BillingCounty() *string
	BillingPostalCode() *string
	BillingProvince() *string
	BillingState() *string
	BirthDate() *string
	BusinessEmailAddress() *string
	BusinessName() *string
	BusinessPhoneNumber() *string
	City() *string
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
	County() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Custom() cdktn.StringMap
	EmailAddress() *string
	FirstName() *string
	// Experimental.
	Fqn() *string
	Gender() *string
	HomePhoneNumber() *string
	InternalValue() *DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributes
	SetInternalValue(val *DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributes)
	LastName() *string
	MailingAddress1() *string
	MailingAddress2() *string
	MailingAddress3() *string
	MailingAddress4() *string
	MailingCity() *string
	MailingCountry() *string
	MailingCounty() *string
	MailingPostalCode() *string
	MailingProvince() *string
	MailingState() *string
	MiddleName() *string
	MobilePhoneNumber() *string
	PartyType() *string
	PhoneNumber() *string
	PostalCode() *string
	ProfileArn() *string
	ProfileId() *string
	Province() *string
	ShippingAddress1() *string
	ShippingAddress2() *string
	ShippingAddress3() *string
	ShippingAddress4() *string
	ShippingCity() *string
	ShippingCountry() *string
	ShippingCounty() *string
	ShippingPostalCode() *string
	ShippingProvince() *string
	ShippingState() *string
	State() *string
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

// The jsii proxy struct for DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference
type jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) AccountNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) AdditionalInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Address4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingAddress4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingAddress4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCountry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCountry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingPostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingPostalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingProvince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProvince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BillingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BirthDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) BusinessPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) County() *string {
	var returns *string
	_jsii_.Get(
		j,
		"county",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Custom() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Gender() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) HomePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) InternalValue() *DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributes {
	var returns *DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingAddress4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingAddress4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCountry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCountry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingPostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingPostalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingProvince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingProvince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MailingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) MobilePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PartyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingAddress4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingAddress4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCountry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCountry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingPostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingPostalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingProvince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingProvince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ShippingState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shippingState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWisdomMessageTemplate.DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference_Override(d DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWisdomMessageTemplate.DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetInternalValue(val *DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccWisdomMessageTemplateDefaultAttributesCustomerProfileAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

