// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectdatatableattribute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectdatatableattribute/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectDataTableAttributeValidationOutputReference interface {
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
	Enum() ConnectDataTableAttributeValidationEnumOutputReference
	EnumInput() interface{}
	ExclusiveMaximum() *float64
	SetExclusiveMaximum(val *float64)
	ExclusiveMaximumInput() *float64
	ExclusiveMinimum() *float64
	SetExclusiveMinimum(val *float64)
	ExclusiveMinimumInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Maximum() *float64
	SetMaximum(val *float64)
	MaximumInput() *float64
	MaxLength() *float64
	SetMaxLength(val *float64)
	MaxLengthInput() *float64
	MaxValues() *float64
	SetMaxValues(val *float64)
	MaxValuesInput() *float64
	Minimum() *float64
	SetMinimum(val *float64)
	MinimumInput() *float64
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	MinValues() *float64
	SetMinValues(val *float64)
	MinValuesInput() *float64
	MultipleOf() *float64
	SetMultipleOf(val *float64)
	MultipleOfInput() *float64
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
	PutEnum(value *ConnectDataTableAttributeValidationEnum)
	ResetEnum()
	ResetExclusiveMaximum()
	ResetExclusiveMinimum()
	ResetMaximum()
	ResetMaxLength()
	ResetMaxValues()
	ResetMinimum()
	ResetMinLength()
	ResetMinValues()
	ResetMultipleOf()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectDataTableAttributeValidationOutputReference
type jsiiProxy_ConnectDataTableAttributeValidationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) Enum() ConnectDataTableAttributeValidationEnumOutputReference {
	var returns ConnectDataTableAttributeValidationEnumOutputReference
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) EnumInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ExclusiveMaximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exclusiveMaximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ExclusiveMaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exclusiveMaximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ExclusiveMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exclusiveMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ExclusiveMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exclusiveMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) Maximum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MaximumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MaxLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MaxValues() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MaxValuesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) Minimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MinValues() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MinValuesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MultipleOf() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multipleOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) MultipleOfInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multipleOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectDataTableAttributeValidationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectDataTableAttributeValidationOutputReference {
	_init_.Initialize()

	if err := validateNewConnectDataTableAttributeValidationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectDataTableAttributeValidationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectDataTableAttribute.ConnectDataTableAttributeValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectDataTableAttributeValidationOutputReference_Override(c ConnectDataTableAttributeValidationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectDataTableAttribute.ConnectDataTableAttributeValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetExclusiveMaximum(val *float64) {
	if err := j.validateSetExclusiveMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusiveMaximum",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetExclusiveMinimum(val *float64) {
	if err := j.validateSetExclusiveMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusiveMinimum",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMaximum(val *float64) {
	if err := j.validateSetMaximumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximum",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMaxLength(val *float64) {
	if err := j.validateSetMaxLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMaxValues(val *float64) {
	if err := j.validateSetMaxValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxValues",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMinimum(val *float64) {
	if err := j.validateSetMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimum",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMinLength(val *float64) {
	if err := j.validateSetMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMinValues(val *float64) {
	if err := j.validateSetMinValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minValues",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetMultipleOf(val *float64) {
	if err := j.validateSetMultipleOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multipleOf",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectDataTableAttributeValidationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) PutEnum(value *ConnectDataTableAttributeValidationEnum) {
	if err := c.validatePutEnumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnum",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		c,
		"resetEnum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetExclusiveMaximum() {
	_jsii_.InvokeVoid(
		c,
		"resetExclusiveMaximum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetExclusiveMinimum() {
	_jsii_.InvokeVoid(
		c,
		"resetExclusiveMinimum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMaximum() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMaxLength() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMaxValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMinimum() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimum",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMinLength() {
	_jsii_.InvokeVoid(
		c,
		"resetMinLength",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMinValues() {
	_jsii_.InvokeVoid(
		c,
		"resetMinValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ResetMultipleOf() {
	_jsii_.InvokeVoid(
		c,
		"resetMultipleOf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectDataTableAttributeValidationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

