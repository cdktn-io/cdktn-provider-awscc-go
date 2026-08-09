// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockprompt/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockPromptVariantsTemplateConfigurationChatOutputReference interface {
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
	InputVariables() BedrockPromptVariantsTemplateConfigurationChatInputVariablesList
	InputVariablesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Messages() BedrockPromptVariantsTemplateConfigurationChatMessagesList
	MessagesInput() interface{}
	SystemAttribute() BedrockPromptVariantsTemplateConfigurationChatSystemList
	SystemAttributeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToolConfiguration() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationOutputReference
	ToolConfigurationInput() interface{}
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
	PutInputVariables(value interface{})
	PutMessages(value interface{})
	PutSystemAttribute(value interface{})
	PutToolConfiguration(value *BedrockPromptVariantsTemplateConfigurationChatToolConfiguration)
	ResetInputVariables()
	ResetMessages()
	ResetSystemAttribute()
	ResetToolConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockPromptVariantsTemplateConfigurationChatOutputReference
type jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) InputVariables() BedrockPromptVariantsTemplateConfigurationChatInputVariablesList {
	var returns BedrockPromptVariantsTemplateConfigurationChatInputVariablesList
	_jsii_.Get(
		j,
		"inputVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) InputVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) Messages() BedrockPromptVariantsTemplateConfigurationChatMessagesList {
	var returns BedrockPromptVariantsTemplateConfigurationChatMessagesList
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) MessagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) SystemAttribute() BedrockPromptVariantsTemplateConfigurationChatSystemList {
	var returns BedrockPromptVariantsTemplateConfigurationChatSystemList
	_jsii_.Get(
		j,
		"systemAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) SystemAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ToolConfiguration() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationOutputReference {
	var returns BedrockPromptVariantsTemplateConfigurationChatToolConfigurationOutputReference
	_jsii_.Get(
		j,
		"toolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ToolConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolConfigurationInput",
		&returns,
	)
	return returns
}


func NewBedrockPromptVariantsTemplateConfigurationChatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockPromptVariantsTemplateConfigurationChatOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockPromptVariantsTemplateConfigurationChatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockPrompt.BedrockPromptVariantsTemplateConfigurationChatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockPromptVariantsTemplateConfigurationChatOutputReference_Override(b BedrockPromptVariantsTemplateConfigurationChatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockPrompt.BedrockPromptVariantsTemplateConfigurationChatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) PutInputVariables(value interface{}) {
	if err := b.validatePutInputVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInputVariables",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) PutMessages(value interface{}) {
	if err := b.validatePutMessagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMessages",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) PutSystemAttribute(value interface{}) {
	if err := b.validatePutSystemAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSystemAttribute",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) PutToolConfiguration(value *BedrockPromptVariantsTemplateConfigurationChatToolConfiguration) {
	if err := b.validatePutToolConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putToolConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ResetInputVariables() {
	_jsii_.InvokeVoid(
		b,
		"resetInputVariables",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ResetMessages() {
	_jsii_.InvokeVoid(
		b,
		"resetMessages",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ResetSystemAttribute() {
	_jsii_.InvokeVoid(
		b,
		"resetSystemAttribute",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ResetToolConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetToolConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

