// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotcommand/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotCommandMandatoryParametersValueConditionsOperandOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Number() *string
	SetNumber(val *string)
	NumberInput() *string
	NumberRange() IotCommandMandatoryParametersValueConditionsOperandNumberRangeOutputReference
	NumberRangeInput() interface{}
	Numbers() *[]*string
	SetNumbers(val *[]*string)
	NumbersInput() *[]*string
	String() *string
	SetString(val *string)
	StringInput() *string
	Strings() *[]*string
	SetStrings(val *[]*string)
	StringsInput() *[]*string
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
	PutNumberRange(value *IotCommandMandatoryParametersValueConditionsOperandNumberRange)
	ResetNumber()
	ResetNumberRange()
	ResetNumbers()
	ResetString()
	ResetStrings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotCommandMandatoryParametersValueConditionsOperandOutputReference
type jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) Number() *string {
	var returns *string
	_jsii_.Get(
		j,
		"number",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) NumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) NumberRange() IotCommandMandatoryParametersValueConditionsOperandNumberRangeOutputReference {
	var returns IotCommandMandatoryParametersValueConditionsOperandNumberRangeOutputReference
	_jsii_.Get(
		j,
		"numberRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) NumberRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) Numbers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"numbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) NumbersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"numbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) String() *string {
	var returns *string
	_jsii_.Get(
		j,
		"string",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) StringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) Strings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"strings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) StringsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotCommandMandatoryParametersValueConditionsOperandOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotCommandMandatoryParametersValueConditionsOperandOutputReference {
	_init_.Initialize()

	if err := validateNewIotCommandMandatoryParametersValueConditionsOperandOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotCommand.IotCommandMandatoryParametersValueConditionsOperandOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotCommandMandatoryParametersValueConditionsOperandOutputReference_Override(i IotCommandMandatoryParametersValueConditionsOperandOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotCommand.IotCommandMandatoryParametersValueConditionsOperandOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetNumber(val *string) {
	if err := j.validateSetNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"number",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetNumbers(val *[]*string) {
	if err := j.validateSetNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numbers",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetString(val *string) {
	if err := j.validateSetStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"string",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetStrings(val *[]*string) {
	if err := j.validateSetStringsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strings",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) PutNumberRange(value *IotCommandMandatoryParametersValueConditionsOperandNumberRange) {
	if err := i.validatePutNumberRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNumberRange",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ResetNumber() {
	_jsii_.InvokeVoid(
		i,
		"resetNumber",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ResetNumberRange() {
	_jsii_.InvokeVoid(
		i,
		"resetNumberRange",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ResetNumbers() {
	_jsii_.InvokeVoid(
		i,
		"resetNumbers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ResetString() {
	_jsii_.InvokeVoid(
		i,
		"resetString",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ResetStrings() {
	_jsii_.InvokeVoid(
		i,
		"resetStrings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueConditionsOperandOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

