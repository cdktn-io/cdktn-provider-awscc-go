// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package outpostssite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/outpostssite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OutpostsSiteOperatingAddressOutputReference interface {
	cdktn.ComplexObject
	AddressLine1() *string
	SetAddressLine1(val *string)
	AddressLine1Input() *string
	AddressLine2() *string
	SetAddressLine2(val *string)
	AddressLine2Input() *string
	AddressLine3() *string
	SetAddressLine3(val *string)
	AddressLine3Input() *string
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
	ContactName() *string
	SetContactName(val *string)
	ContactNameInput() *string
	ContactPhoneNumber() *string
	SetContactPhoneNumber(val *string)
	ContactPhoneNumberInput() *string
	CountryCode() *string
	SetCountryCode(val *string)
	CountryCodeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DistrictOrCounty() *string
	SetDistrictOrCounty(val *string)
	DistrictOrCountyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Municipality() *string
	SetMunicipality(val *string)
	MunicipalityInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	StateOrRegion() *string
	SetStateOrRegion(val *string)
	StateOrRegionInput() *string
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
	ResetAddressLine1()
	ResetAddressLine2()
	ResetAddressLine3()
	ResetCity()
	ResetContactName()
	ResetContactPhoneNumber()
	ResetCountryCode()
	ResetDistrictOrCounty()
	ResetMunicipality()
	ResetPostalCode()
	ResetStateOrRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OutpostsSiteOperatingAddressOutputReference
type jsiiProxy_OutpostsSiteOperatingAddressOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) AddressLine1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) AddressLine1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) AddressLine2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) AddressLine2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) AddressLine3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) AddressLine3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ContactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ContactNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ContactPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ContactPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) DistrictOrCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"districtOrCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) DistrictOrCountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"districtOrCountyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) Municipality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"municipality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) MunicipalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"municipalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) StateOrRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateOrRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) StateOrRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateOrRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOutpostsSiteOperatingAddressOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OutpostsSiteOperatingAddressOutputReference {
	_init_.Initialize()

	if err := validateNewOutpostsSiteOperatingAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OutpostsSiteOperatingAddressOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.outpostsSite.OutpostsSiteOperatingAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOutpostsSiteOperatingAddressOutputReference_Override(o OutpostsSiteOperatingAddressOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.outpostsSite.OutpostsSiteOperatingAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetAddressLine1(val *string) {
	if err := j.validateSetAddressLine1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine1",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetAddressLine2(val *string) {
	if err := j.validateSetAddressLine2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine2",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetAddressLine3(val *string) {
	if err := j.validateSetAddressLine3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine3",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetContactName(val *string) {
	if err := j.validateSetContactNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactName",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetContactPhoneNumber(val *string) {
	if err := j.validateSetContactPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetDistrictOrCounty(val *string) {
	if err := j.validateSetDistrictOrCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"districtOrCounty",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetMunicipality(val *string) {
	if err := j.validateSetMunicipalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"municipality",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetStateOrRegion(val *string) {
	if err := j.validateSetStateOrRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateOrRegion",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OutpostsSiteOperatingAddressOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetAddressLine1() {
	_jsii_.InvokeVoid(
		o,
		"resetAddressLine1",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetAddressLine2() {
	_jsii_.InvokeVoid(
		o,
		"resetAddressLine2",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetAddressLine3() {
	_jsii_.InvokeVoid(
		o,
		"resetAddressLine3",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		o,
		"resetCity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetContactName() {
	_jsii_.InvokeVoid(
		o,
		"resetContactName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetContactPhoneNumber() {
	_jsii_.InvokeVoid(
		o,
		"resetContactPhoneNumber",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		o,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetDistrictOrCounty() {
	_jsii_.InvokeVoid(
		o,
		"resetDistrictOrCounty",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetMunicipality() {
	_jsii_.InvokeVoid(
		o,
		"resetMunicipality",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		o,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ResetStateOrRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetStateOrRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OutpostsSiteOperatingAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

