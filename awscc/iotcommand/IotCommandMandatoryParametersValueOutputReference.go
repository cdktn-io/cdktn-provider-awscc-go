// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcommand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotcommand/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotCommandMandatoryParametersValueOutputReference interface {
	cdktn.ComplexObject
	B() interface{}
	SetB(val interface{})
	Bin() *string
	SetBin(val *string)
	BinInput() *string
	BInput() interface{}
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
	D() *float64
	SetD(val *float64)
	DInput() *float64
	// Experimental.
	Fqn() *string
	I() *float64
	SetI(val *float64)
	IInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	L() *string
	SetL(val *string)
	LInput() *string
	S() *string
	SetS(val *string)
	SInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ul() *string
	SetUl(val *string)
	UlInput() *string
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
	ResetB()
	ResetBin()
	ResetD()
	ResetI()
	ResetL()
	ResetS()
	ResetUl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotCommandMandatoryParametersValueOutputReference
type jsiiProxy_IotCommandMandatoryParametersValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) B() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"b",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) Bin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) BinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) BInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) D() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"d",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) DInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) I() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"i",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) IInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) L() *string {
	var returns *string
	_jsii_.Get(
		j,
		"l",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) LInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) S() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) SInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) Ul() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ul",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) UlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ulInput",
		&returns,
	)
	return returns
}


func NewIotCommandMandatoryParametersValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotCommandMandatoryParametersValueOutputReference {
	_init_.Initialize()

	if err := validateNewIotCommandMandatoryParametersValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotCommandMandatoryParametersValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotCommand.IotCommandMandatoryParametersValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotCommandMandatoryParametersValueOutputReference_Override(i IotCommandMandatoryParametersValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotCommand.IotCommandMandatoryParametersValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetB(val interface{}) {
	if err := j.validateSetBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"b",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetBin(val *string) {
	if err := j.validateSetBinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bin",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetD(val *float64) {
	if err := j.validateSetDParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"d",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetI(val *float64) {
	if err := j.validateSetIParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"i",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetL(val *string) {
	if err := j.validateSetLParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"l",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetS(val *string) {
	if err := j.validateSetSParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotCommandMandatoryParametersValueOutputReference)SetUl(val *string) {
	if err := j.validateSetUlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ul",
		val,
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetB() {
	_jsii_.InvokeVoid(
		i,
		"resetB",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetBin() {
	_jsii_.InvokeVoid(
		i,
		"resetBin",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetD() {
	_jsii_.InvokeVoid(
		i,
		"resetD",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetI() {
	_jsii_.InvokeVoid(
		i,
		"resetI",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetL() {
	_jsii_.InvokeVoid(
		i,
		"resetL",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetS() {
	_jsii_.InvokeVoid(
		i,
		"resetS",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ResetUl() {
	_jsii_.InvokeVoid(
		i,
		"resetUl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotCommandMandatoryParametersValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

