// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockprompt/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference interface {
	cdktn.ComplexObject
	Any() *string
	SetAny(val *string)
	AnyInput() *string
	Auto() *string
	SetAuto(val *string)
	AutoInput() *string
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tool() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceToolOutputReference
	ToolInput() interface{}
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
	PutTool(value *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceTool)
	ResetAny()
	ResetAuto()
	ResetTool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference
type jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) Any() *string {
	var returns *string
	_jsii_.Get(
		j,
		"any",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) AnyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) Auto() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) AutoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) Tool() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceToolOutputReference {
	var returns BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceToolOutputReference
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ToolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}


func NewBedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockPrompt.BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference_Override(b BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockPrompt.BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetAny(val *string) {
	if err := j.validateSetAnyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"any",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetAuto(val *string) {
	if err := j.validateSetAutoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auto",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) PutTool(value *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceTool) {
	if err := b.validatePutToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTool",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ResetAny() {
	_jsii_.InvokeVoid(
		b,
		"resetAny",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ResetAuto() {
	_jsii_.InvokeVoid(
		b,
		"resetAuto",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ResetTool() {
	_jsii_.InvokeVoid(
		b,
		"resetTool",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolChoiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

