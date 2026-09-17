// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference interface {
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
	FieldDataType() *string
	SetFieldDataType(val *string)
	FieldDataTypeInput() *string
	FilterOverrides() GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverridesOutputReference
	FilterOverridesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsNullable() interface{}
	SetIsNullable(val interface{})
	IsNullableInput() interface{}
	IsOrderable() interface{}
	SetIsOrderable(val interface{})
	IsOrderableInput() interface{}
	IsPartitionable() interface{}
	SetIsPartitionable(val interface{})
	IsPartitionableInput() interface{}
	IsQueryable() interface{}
	SetIsQueryable(val interface{})
	IsQueryableInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	ResponseDateFormat() *string
	SetResponseDateFormat(val *string)
	ResponseDateFormatInput() *string
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
	PutFilterOverrides(value *GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverrides)
	ResetFieldDataType()
	ResetFilterOverrides()
	ResetIsNullable()
	ResetIsOrderable()
	ResetIsPartitionable()
	ResetIsQueryable()
	ResetName()
	ResetResponseDateFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) FieldDataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) FieldDataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) FilterOverrides() GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverridesOutputReference {
	var returns GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverridesOutputReference
	_jsii_.Get(
		j,
		"filterOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) FilterOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsNullable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsNullableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsOrderable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrderable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsOrderableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrderableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsPartitionable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPartitionable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsPartitionableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPartitionableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsQueryable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isQueryable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) IsQueryableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isQueryableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResponseDateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseDateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResponseDateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseDateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference_Override(g GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetFieldDataType(val *string) {
	if err := j.validateSetFieldDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDataType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetIsNullable(val interface{}) {
	if err := j.validateSetIsNullableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNullable",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetIsOrderable(val interface{}) {
	if err := j.validateSetIsOrderableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOrderable",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetIsPartitionable(val interface{}) {
	if err := j.validateSetIsPartitionableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPartitionable",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetIsQueryable(val interface{}) {
	if err := j.validateSetIsQueryableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isQueryable",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetResponseDateFormat(val *string) {
	if err := j.validateSetResponseDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseDateFormat",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) PutFilterOverrides(value *GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaFilterOverrides) {
	if err := g.validatePutFilterOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilterOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetFieldDataType() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldDataType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetFilterOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetIsNullable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsNullable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetIsOrderable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsOrderable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetIsPartitionable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsPartitionable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetIsQueryable() {
	_jsii_.InvokeVoid(
		g,
		"resetIsQueryable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ResetResponseDateFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetResponseDateFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationEntityConfigurationsSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

