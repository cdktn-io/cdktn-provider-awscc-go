// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcorememory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference interface {
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
	EpisodicOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideOutputReference
	EpisodicOverrideInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SelfManagedConfiguration() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationOutputReference
	SelfManagedConfigurationInput() interface{}
	SemanticOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSemanticOverrideOutputReference
	SemanticOverrideInput() interface{}
	SummaryOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSummaryOverrideOutputReference
	SummaryOverrideInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserPreferenceOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationUserPreferenceOverrideOutputReference
	UserPreferenceOverrideInput() interface{}
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
	PutEpisodicOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverride)
	PutSelfManagedConfiguration(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfiguration)
	PutSemanticOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSemanticOverride)
	PutSummaryOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSummaryOverride)
	PutUserPreferenceOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationUserPreferenceOverride)
	ResetEpisodicOverride()
	ResetSelfManagedConfiguration()
	ResetSemanticOverride()
	ResetSummaryOverride()
	ResetUserPreferenceOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference
type jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) EpisodicOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverrideOutputReference
	_jsii_.Get(
		j,
		"episodicOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) EpisodicOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"episodicOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) SelfManagedConfiguration() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfigurationOutputReference
	_jsii_.Get(
		j,
		"selfManagedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) SelfManagedConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) SemanticOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSemanticOverrideOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSemanticOverrideOutputReference
	_jsii_.Get(
		j,
		"semanticOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) SemanticOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) SummaryOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSummaryOverrideOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSummaryOverrideOutputReference
	_jsii_.Get(
		j,
		"summaryOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) SummaryOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summaryOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) UserPreferenceOverride() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationUserPreferenceOverrideOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationUserPreferenceOverrideOutputReference
	_jsii_.Get(
		j,
		"userPreferenceOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) UserPreferenceOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPreferenceOverrideInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference_Override(b BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) PutEpisodicOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationEpisodicOverride) {
	if err := b.validatePutEpisodicOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEpisodicOverride",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) PutSelfManagedConfiguration(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSelfManagedConfiguration) {
	if err := b.validatePutSelfManagedConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSelfManagedConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) PutSemanticOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSemanticOverride) {
	if err := b.validatePutSemanticOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSemanticOverride",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) PutSummaryOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationSummaryOverride) {
	if err := b.validatePutSummaryOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSummaryOverride",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) PutUserPreferenceOverride(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationUserPreferenceOverride) {
	if err := b.validatePutUserPreferenceOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserPreferenceOverride",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ResetEpisodicOverride() {
	_jsii_.InvokeVoid(
		b,
		"resetEpisodicOverride",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ResetSelfManagedConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetSelfManagedConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ResetSemanticOverride() {
	_jsii_.InvokeVoid(
		b,
		"resetSemanticOverride",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ResetSummaryOverride() {
	_jsii_.InvokeVoid(
		b,
		"resetSummaryOverride",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ResetUserPreferenceOverride() {
	_jsii_.InvokeVoid(
		b,
		"resetUserPreferenceOverride",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

