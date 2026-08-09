// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotsecurityprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotSecurityProfileBehaviorsCriteriaValueOutputReference interface {
	cdktn.ComplexObject
	Cidrs() *[]*string
	SetCidrs(val *[]*string)
	CidrsInput() *[]*string
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
	Count() *string
	SetCount(val *string)
	CountInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Number() *float64
	SetNumber(val *float64)
	NumberInput() *float64
	Numbers() *[]*float64
	SetNumbers(val *[]*float64)
	NumbersInput() *[]*float64
	Ports() *[]*float64
	SetPorts(val *[]*float64)
	PortsInput() *[]*float64
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
	ResetCidrs()
	ResetCount()
	ResetNumber()
	ResetNumbers()
	ResetPorts()
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

// The jsii proxy struct for IotSecurityProfileBehaviorsCriteriaValueOutputReference
type jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) CidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Count() *string {
	var returns *string
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) CountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Number() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"number",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) NumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Numbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"numbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) NumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"numbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Ports() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) PortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Strings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"strings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) StringsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotSecurityProfileBehaviorsCriteriaValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotSecurityProfileBehaviorsCriteriaValueOutputReference {
	_init_.Initialize()

	if err := validateNewIotSecurityProfileBehaviorsCriteriaValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotSecurityProfile.IotSecurityProfileBehaviorsCriteriaValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecurityProfileBehaviorsCriteriaValueOutputReference_Override(i IotSecurityProfileBehaviorsCriteriaValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotSecurityProfile.IotSecurityProfileBehaviorsCriteriaValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetCidrs(val *[]*string) {
	if err := j.validateSetCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrs",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetCount(val *string) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetNumber(val *float64) {
	if err := j.validateSetNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"number",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetNumbers(val *[]*float64) {
	if err := j.validateSetNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numbers",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetPorts(val *[]*float64) {
	if err := j.validateSetPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ports",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetStrings(val *[]*string) {
	if err := j.validateSetStringsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strings",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ResetCidrs() {
	_jsii_.InvokeVoid(
		i,
		"resetCidrs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		i,
		"resetCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ResetNumber() {
	_jsii_.InvokeVoid(
		i,
		"resetNumber",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ResetNumbers() {
	_jsii_.InvokeVoid(
		i,
		"resetNumbers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		i,
		"resetPorts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ResetStrings() {
	_jsii_.InvokeVoid(
		i,
		"resetStrings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotSecurityProfileBehaviorsCriteriaValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

