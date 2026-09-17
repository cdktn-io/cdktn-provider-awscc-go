// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference interface {
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
	Value() GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValueOutputReference
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
	PutValue(value *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValue)
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

// The jsii proxy struct for GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) PropertyLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) PropertyLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Value() GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValueOutputReference {
	var returns GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference_Override(g GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetPropertyLocation(val *string) {
	if err := j.validateSetPropertyLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyLocation",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) PutValue(value *GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageValue) {
	if err := g.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ResetPropertyLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetPropertyLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		g,
		"resetValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSourceConfigurationPaginationConfigurationCursorConfigurationNextPageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

