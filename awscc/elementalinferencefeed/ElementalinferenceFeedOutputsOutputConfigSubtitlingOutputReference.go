// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elementalinferencefeed/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference interface {
	cdktn.ComplexObject
	AspectRatio() ElementalinferenceFeedOutputsOutputConfigSubtitlingAspectRatioOutputReference
	AspectRatioInput() interface{}
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
	Dictionary() *string
	SetDictionary(val *string)
	DictionaryInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	ProfanityFilter() *string
	SetProfanityFilter(val *string)
	ProfanityFilterInput() *string
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
	PutAspectRatio(value *ElementalinferenceFeedOutputsOutputConfigSubtitlingAspectRatio)
	ResetAspectRatio()
	ResetDictionary()
	ResetLanguage()
	ResetProfanityFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference
type jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) AspectRatio() ElementalinferenceFeedOutputsOutputConfigSubtitlingAspectRatioOutputReference {
	var returns ElementalinferenceFeedOutputsOutputConfigSubtitlingAspectRatioOutputReference
	_jsii_.Get(
		j,
		"aspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) AspectRatioInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) Dictionary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) DictionaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ProfanityFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profanityFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ProfanityFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profanityFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference {
	_init_.Initialize()

	if err := validateNewElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elementalinferenceFeed.ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference_Override(e ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elementalinferenceFeed.ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetDictionary(val *string) {
	if err := j.validateSetDictionaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dictionary",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetProfanityFilter(val *string) {
	if err := j.validateSetProfanityFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profanityFilter",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) PutAspectRatio(value *ElementalinferenceFeedOutputsOutputConfigSubtitlingAspectRatio) {
	if err := e.validatePutAspectRatioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAspectRatio",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ResetAspectRatio() {
	_jsii_.InvokeVoid(
		e,
		"resetAspectRatio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		e,
		"resetDictionary",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ResetLanguage() {
	_jsii_.InvokeVoid(
		e,
		"resetLanguage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ResetProfanityFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetProfanityFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElementalinferenceFeedOutputsOutputConfigSubtitlingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

