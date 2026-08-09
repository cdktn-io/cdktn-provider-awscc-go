// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockprompt

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockprompt/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference interface {
	cdktn.ComplexObject
	CachePoint() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsCachePointOutputReference
	CachePointInput() interface{}
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
	ToolSpec() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpecOutputReference
	ToolSpecInput() interface{}
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
	PutCachePoint(value *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsCachePoint)
	PutToolSpec(value *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpec)
	ResetCachePoint()
	ResetToolSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference
type jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) CachePoint() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsCachePointOutputReference {
	var returns BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsCachePointOutputReference
	_jsii_.Get(
		j,
		"cachePoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) CachePointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachePointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ToolSpec() BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpecOutputReference {
	var returns BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpecOutputReference
	_jsii_.Get(
		j,
		"toolSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ToolSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolSpecInput",
		&returns,
	)
	return returns
}


func NewBedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockPrompt.BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference_Override(b BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockPrompt.BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) PutCachePoint(value *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsCachePoint) {
	if err := b.validatePutCachePointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCachePoint",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) PutToolSpec(value *BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsToolSpec) {
	if err := b.validatePutToolSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putToolSpec",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ResetCachePoint() {
	_jsii_.InvokeVoid(
		b,
		"resetCachePoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ResetToolSpec() {
	_jsii_.InvokeVoid(
		b,
		"resetToolSpec",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockPromptVariantsTemplateConfigurationChatToolConfigurationToolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

