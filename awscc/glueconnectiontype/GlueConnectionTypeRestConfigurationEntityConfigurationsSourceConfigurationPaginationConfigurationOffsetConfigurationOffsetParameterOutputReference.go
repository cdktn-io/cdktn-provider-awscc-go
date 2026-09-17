// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference interface {
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
	DefaultValue() *string
	SetDefaultValue(val *string)
	DefaultValueInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	PropertyLocation() *string
	SetPropertyLocation(val *string)
	PropertyLocationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference
	ValueInput() interface{}
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
	PutValue(value *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValue)
	ResetDefaultValue()
	ResetKey()
	ResetPropertyLocation()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) PropertyLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) PropertyLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) Value() GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference {
	var returns GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference_Override(g GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetPropertyLocation(val *string) {
	if err := j.validateSetPropertyLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyLocation",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) PutValue(value *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValue) {
	if err := g.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ResetPropertyLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetPropertyLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		g,
		"resetValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

