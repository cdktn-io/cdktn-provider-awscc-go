// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdataautomationproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationProjectStandardOutputConfigurationOutputReference interface {
	cdktn.ComplexObject
	Audio() BedrockDataAutomationProjectStandardOutputConfigurationAudioOutputReference
	AudioInput() interface{}
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
	Document() BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputReference
	DocumentInput() interface{}
	// Experimental.
	Fqn() *string
	Image() BedrockDataAutomationProjectStandardOutputConfigurationImageOutputReference
	ImageInput() interface{}
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
	Video() BedrockDataAutomationProjectStandardOutputConfigurationVideoOutputReference
	VideoInput() interface{}
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
	PutAudio(value *BedrockDataAutomationProjectStandardOutputConfigurationAudio)
	PutDocument(value *BedrockDataAutomationProjectStandardOutputConfigurationDocument)
	PutImage(value *BedrockDataAutomationProjectStandardOutputConfigurationImage)
	PutVideo(value *BedrockDataAutomationProjectStandardOutputConfigurationVideo)
	ResetAudio()
	ResetDocument()
	ResetImage()
	ResetVideo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataAutomationProjectStandardOutputConfigurationOutputReference
type jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) Audio() BedrockDataAutomationProjectStandardOutputConfigurationAudioOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationAudioOutputReference
	_jsii_.Get(
		j,
		"audio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) AudioInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) Document() BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputReference
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) DocumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) Image() BedrockDataAutomationProjectStandardOutputConfigurationImageOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationImageOutputReference
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) Video() BedrockDataAutomationProjectStandardOutputConfigurationVideoOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationVideoOutputReference
	_jsii_.Get(
		j,
		"video",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) VideoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoInput",
		&returns,
	)
	return returns
}


func NewBedrockDataAutomationProjectStandardOutputConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataAutomationProjectStandardOutputConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataAutomationProjectStandardOutputConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectStandardOutputConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataAutomationProjectStandardOutputConfigurationOutputReference_Override(b BedrockDataAutomationProjectStandardOutputConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectStandardOutputConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) PutAudio(value *BedrockDataAutomationProjectStandardOutputConfigurationAudio) {
	if err := b.validatePutAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAudio",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) PutDocument(value *BedrockDataAutomationProjectStandardOutputConfigurationDocument) {
	if err := b.validatePutDocumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDocument",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) PutImage(value *BedrockDataAutomationProjectStandardOutputConfigurationImage) {
	if err := b.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putImage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) PutVideo(value *BedrockDataAutomationProjectStandardOutputConfigurationVideo) {
	if err := b.validatePutVideoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVideo",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ResetAudio() {
	_jsii_.InvokeVoid(
		b,
		"resetAudio",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ResetDocument() {
	_jsii_.InvokeVoid(
		b,
		"resetDocument",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		b,
		"resetImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ResetVideo() {
	_jsii_.InvokeVoid(
		b,
		"resetVideo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

