// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pipespipe/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipesPipeTargetParametersTimestreamParametersOutputReference interface {
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
	DimensionMappings() PipesPipeTargetParametersTimestreamParametersDimensionMappingsList
	DimensionMappingsInput() interface{}
	EpochTimeUnit() *string
	SetEpochTimeUnit(val *string)
	EpochTimeUnitInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiMeasureMappings() PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsList
	MultiMeasureMappingsInput() interface{}
	SingleMeasureMappings() PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsList
	SingleMeasureMappingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeFieldType() *string
	SetTimeFieldType(val *string)
	TimeFieldTypeInput() *string
	TimestampFormat() *string
	SetTimestampFormat(val *string)
	TimestampFormatInput() *string
	TimeValue() *string
	SetTimeValue(val *string)
	TimeValueInput() *string
	VersionValue() *string
	SetVersionValue(val *string)
	VersionValueInput() *string
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
	PutDimensionMappings(value interface{})
	PutMultiMeasureMappings(value interface{})
	PutSingleMeasureMappings(value interface{})
	ResetDimensionMappings()
	ResetEpochTimeUnit()
	ResetMultiMeasureMappings()
	ResetSingleMeasureMappings()
	ResetTimeFieldType()
	ResetTimestampFormat()
	ResetTimeValue()
	ResetVersionValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipesPipeTargetParametersTimestreamParametersOutputReference
type jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) DimensionMappings() PipesPipeTargetParametersTimestreamParametersDimensionMappingsList {
	var returns PipesPipeTargetParametersTimestreamParametersDimensionMappingsList
	_jsii_.Get(
		j,
		"dimensionMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) DimensionMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) EpochTimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"epochTimeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) EpochTimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"epochTimeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) MultiMeasureMappings() PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsList {
	var returns PipesPipeTargetParametersTimestreamParametersMultiMeasureMappingsList
	_jsii_.Get(
		j,
		"multiMeasureMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) MultiMeasureMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiMeasureMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) SingleMeasureMappings() PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsList {
	var returns PipesPipeTargetParametersTimestreamParametersSingleMeasureMappingsList
	_jsii_.Get(
		j,
		"singleMeasureMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) SingleMeasureMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleMeasureMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TimeFieldType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFieldType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TimeFieldTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFieldTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TimeValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) TimeValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) VersionValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) VersionValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionValueInput",
		&returns,
	)
	return returns
}


func NewPipesPipeTargetParametersTimestreamParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PipesPipeTargetParametersTimestreamParametersOutputReference {
	_init_.Initialize()

	if err := validateNewPipesPipeTargetParametersTimestreamParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pipesPipe.PipesPipeTargetParametersTimestreamParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipesPipeTargetParametersTimestreamParametersOutputReference_Override(p PipesPipeTargetParametersTimestreamParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pipesPipe.PipesPipeTargetParametersTimestreamParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetEpochTimeUnit(val *string) {
	if err := j.validateSetEpochTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epochTimeUnit",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetTimeFieldType(val *string) {
	if err := j.validateSetTimeFieldTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFieldType",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetTimeValue(val *string) {
	if err := j.validateSetTimeValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeValue",
		val,
	)
}

func (j *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference)SetVersionValue(val *string) {
	if err := j.validateSetVersionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionValue",
		val,
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) PutDimensionMappings(value interface{}) {
	if err := p.validatePutDimensionMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDimensionMappings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) PutMultiMeasureMappings(value interface{}) {
	if err := p.validatePutMultiMeasureMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMultiMeasureMappings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) PutSingleMeasureMappings(value interface{}) {
	if err := p.validatePutSingleMeasureMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSingleMeasureMappings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetDimensionMappings() {
	_jsii_.InvokeVoid(
		p,
		"resetDimensionMappings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetEpochTimeUnit() {
	_jsii_.InvokeVoid(
		p,
		"resetEpochTimeUnit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetMultiMeasureMappings() {
	_jsii_.InvokeVoid(
		p,
		"resetMultiMeasureMappings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetSingleMeasureMappings() {
	_jsii_.InvokeVoid(
		p,
		"resetSingleMeasureMappings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetTimeFieldType() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeFieldType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetTimeValue() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ResetVersionValue() {
	_jsii_.InvokeVoid(
		p,
		"resetVersionValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipesPipeTargetParametersTimestreamParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

