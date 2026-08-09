// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lightsaildatabase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailDatabaseRelationalDatabaseParametersOutputReference interface {
	cdktn.ComplexObject
	AllowedValues() *string
	SetAllowedValues(val *string)
	AllowedValuesInput() *string
	ApplyMethod() *string
	SetApplyMethod(val *string)
	ApplyMethodInput() *string
	ApplyType() *string
	SetApplyType(val *string)
	ApplyTypeInput() *string
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
	DataType() *string
	SetDataType(val *string)
	DataTypeInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsModifiable() interface{}
	SetIsModifiable(val interface{})
	IsModifiableInput() interface{}
	ParameterName() *string
	SetParameterName(val *string)
	ParameterNameInput() *string
	ParameterValue() *string
	SetParameterValue(val *string)
	ParameterValueInput() *string
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
	ResetAllowedValues()
	ResetApplyMethod()
	ResetApplyType()
	ResetDataType()
	ResetDescription()
	ResetIsModifiable()
	ResetParameterName()
	ResetParameterValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LightsailDatabaseRelationalDatabaseParametersOutputReference
type jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) AllowedValues() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) AllowedValuesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ApplyMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ApplyMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ApplyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ApplyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) IsModifiable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isModifiable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) IsModifiableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isModifiableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ParameterValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ParameterValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLightsailDatabaseRelationalDatabaseParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LightsailDatabaseRelationalDatabaseParametersOutputReference {
	_init_.Initialize()

	if err := validateNewLightsailDatabaseRelationalDatabaseParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lightsailDatabase.LightsailDatabaseRelationalDatabaseParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLightsailDatabaseRelationalDatabaseParametersOutputReference_Override(l LightsailDatabaseRelationalDatabaseParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lightsailDatabase.LightsailDatabaseRelationalDatabaseParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetAllowedValues(val *string) {
	if err := j.validateSetAllowedValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedValues",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetApplyMethod(val *string) {
	if err := j.validateSetApplyMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyMethod",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetApplyType(val *string) {
	if err := j.validateSetApplyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyType",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetDataType(val *string) {
	if err := j.validateSetDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetIsModifiable(val interface{}) {
	if err := j.validateSetIsModifiableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isModifiable",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetParameterName(val *string) {
	if err := j.validateSetParameterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterName",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetParameterValue(val *string) {
	if err := j.validateSetParameterValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterValue",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetAllowedValues() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedValues",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetApplyMethod() {
	_jsii_.InvokeVoid(
		l,
		"resetApplyMethod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetApplyType() {
	_jsii_.InvokeVoid(
		l,
		"resetApplyType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetDataType() {
	_jsii_.InvokeVoid(
		l,
		"resetDataType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetIsModifiable() {
	_jsii_.InvokeVoid(
		l,
		"resetIsModifiable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetParameterName() {
	_jsii_.InvokeVoid(
		l,
		"resetParameterName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ResetParameterValue() {
	_jsii_.InvokeVoid(
		l,
		"resetParameterValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LightsailDatabaseRelationalDatabaseParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

