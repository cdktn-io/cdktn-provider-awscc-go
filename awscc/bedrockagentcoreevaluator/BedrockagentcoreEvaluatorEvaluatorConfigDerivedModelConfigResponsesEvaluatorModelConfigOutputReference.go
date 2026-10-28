// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/bedrockagentcoreevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference interface {
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
	MaxOutputTokens() *float64
	SetMaxOutputTokens(val *float64)
	MaxOutputTokensInput() *float64
	ModelId() *string
	SetModelId(val *string)
	ModelIdInput() *string
	Reasoning() BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigReasoningOutputReference
	ReasoningInput() interface{}
	Temperature() *float64
	SetTemperature(val *float64)
	TemperatureInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TopP() *float64
	SetTopP(val *float64)
	TopPInput() *float64
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
	PutReasoning(value *BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigReasoning)
	ResetMaxOutputTokens()
	ResetModelId()
	ResetReasoning()
	ResetTemperature()
	ResetTopP()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference
type jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) MaxOutputTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOutputTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) MaxOutputTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOutputTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) Reasoning() BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigReasoningOutputReference {
	var returns BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigReasoningOutputReference
	_jsii_.Get(
		j,
		"reasoning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ReasoningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reasoningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) TemperatureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) TopPInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topPInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreEvaluator.BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference_Override(b BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreEvaluator.BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetMaxOutputTokens(val *float64) {
	if err := j.validateSetMaxOutputTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOutputTokens",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetTemperature(val *float64) {
	if err := j.validateSetTemperatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temperature",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference)SetTopP(val *float64) {
	if err := j.validateSetTopPParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topP",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) PutReasoning(value *BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigReasoning) {
	if err := b.validatePutReasoningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putReasoning",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ResetMaxOutputTokens() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxOutputTokens",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ResetModelId() {
	_jsii_.InvokeVoid(
		b,
		"resetModelId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ResetReasoning() {
	_jsii_.InvokeVoid(
		b,
		"resetReasoning",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ResetTemperature() {
	_jsii_.InvokeVoid(
		b,
		"resetTemperature",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ResetTopP() {
	_jsii_.InvokeVoid(
		b,
		"resetTopP",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigDerivedModelConfigResponsesEvaluatorModelConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

