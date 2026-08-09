// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdataautomationproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference interface {
	cdktn.ComplexObject
	ChannelLabeling() BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptChannelLabelingOutputReference
	ChannelLabelingInput() interface{}
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
	SpeakerLabeling() BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptSpeakerLabelingOutputReference
	SpeakerLabelingInput() interface{}
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
	PutChannelLabeling(value *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptChannelLabeling)
	PutSpeakerLabeling(value *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptSpeakerLabeling)
	ResetChannelLabeling()
	ResetSpeakerLabeling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference
type jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ChannelLabeling() BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptChannelLabelingOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptChannelLabelingOutputReference
	_jsii_.Get(
		j,
		"channelLabeling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ChannelLabelingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelLabelingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) SpeakerLabeling() BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptSpeakerLabelingOutputReference {
	var returns BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptSpeakerLabelingOutputReference
	_jsii_.Get(
		j,
		"speakerLabeling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) SpeakerLabelingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"speakerLabelingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference_Override(b BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) PutChannelLabeling(value *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptChannelLabeling) {
	if err := b.validatePutChannelLabelingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putChannelLabeling",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) PutSpeakerLabeling(value *BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptSpeakerLabeling) {
	if err := b.validatePutSpeakerLabelingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSpeakerLabeling",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ResetChannelLabeling() {
	_jsii_.InvokeVoid(
		b,
		"resetChannelLabeling",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ResetSpeakerLabeling() {
	_jsii_.InvokeVoid(
		b,
		"resetSpeakerLabeling",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockDataAutomationProjectStandardOutputConfigurationAudioExtractionCategoryTypeConfigurationTranscriptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

