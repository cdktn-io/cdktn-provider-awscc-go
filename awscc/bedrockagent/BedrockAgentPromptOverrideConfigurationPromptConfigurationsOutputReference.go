// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference interface {
	cdktn.ComplexObject
	AdditionalModelRequestFields() *string
	SetAdditionalModelRequestFields(val *string)
	AdditionalModelRequestFieldsInput() *string
	BasePromptTemplate() *string
	SetBasePromptTemplate(val *string)
	BasePromptTemplateInput() *string
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
	FoundationModel() *string
	SetFoundationModel(val *string)
	FoundationModelInput() *string
	// Experimental.
	Fqn() *string
	InferenceConfiguration() BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfigurationOutputReference
	InferenceConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ParserMode() *string
	SetParserMode(val *string)
	ParserModeInput() *string
	PromptCreationMode() *string
	SetPromptCreationMode(val *string)
	PromptCreationModeInput() *string
	PromptState() *string
	SetPromptState(val *string)
	PromptStateInput() *string
	PromptType() *string
	SetPromptType(val *string)
	PromptTypeInput() *string
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
	PutInferenceConfiguration(value *BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfiguration)
	ResetAdditionalModelRequestFields()
	ResetBasePromptTemplate()
	ResetFoundationModel()
	ResetInferenceConfiguration()
	ResetParserMode()
	ResetPromptCreationMode()
	ResetPromptState()
	ResetPromptType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference
type jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) AdditionalModelRequestFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) AdditionalModelRequestFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalModelRequestFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) BasePromptTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePromptTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) BasePromptTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePromptTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) FoundationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) FoundationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InferenceConfiguration() BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfigurationOutputReference {
	var returns BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfigurationOutputReference
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InferenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ParserMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ParserModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptCreationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptCreationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptCreationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptCreationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PromptTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference_Override(b BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockAgent.BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetAdditionalModelRequestFields(val *string) {
	if err := j.validateSetAdditionalModelRequestFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalModelRequestFields",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetBasePromptTemplate(val *string) {
	if err := j.validateSetBasePromptTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basePromptTemplate",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetFoundationModel(val *string) {
	if err := j.validateSetFoundationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationModel",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetParserMode(val *string) {
	if err := j.validateSetParserModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parserMode",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetPromptCreationMode(val *string) {
	if err := j.validateSetPromptCreationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptCreationMode",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetPromptState(val *string) {
	if err := j.validateSetPromptStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptState",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetPromptType(val *string) {
	if err := j.validateSetPromptTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptType",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) PutInferenceConfiguration(value *BedrockAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfiguration) {
	if err := b.validatePutInferenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInferenceConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetAdditionalModelRequestFields() {
	_jsii_.InvokeVoid(
		b,
		"resetAdditionalModelRequestFields",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetBasePromptTemplate() {
	_jsii_.InvokeVoid(
		b,
		"resetBasePromptTemplate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetFoundationModel() {
	_jsii_.InvokeVoid(
		b,
		"resetFoundationModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetInferenceConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetInferenceConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetParserMode() {
	_jsii_.InvokeVoid(
		b,
		"resetParserMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetPromptCreationMode() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptCreationMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetPromptState() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptState",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ResetPromptType() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockAgentPromptOverrideConfigurationPromptConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

