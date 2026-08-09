// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference interface {
	cdktn.ComplexObject
	AnswerGenerationAiGuardrailId() *string
	SetAnswerGenerationAiGuardrailId(val *string)
	AnswerGenerationAiGuardrailIdInput() *string
	AnswerGenerationAiPromptId() *string
	SetAnswerGenerationAiPromptId(val *string)
	AnswerGenerationAiPromptIdInput() *string
	AssociationConfigurations() WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationAssociationConfigurationsList
	AssociationConfigurationsInput() interface{}
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
	IntentLabelingGenerationAiPromptId() *string
	SetIntentLabelingGenerationAiPromptId(val *string)
	IntentLabelingGenerationAiPromptIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	QueryReformulationAiPromptId() *string
	SetQueryReformulationAiPromptId(val *string)
	QueryReformulationAiPromptIdInput() *string
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
	PutAssociationConfigurations(value interface{})
	ResetAnswerGenerationAiGuardrailId()
	ResetAnswerGenerationAiPromptId()
	ResetAssociationConfigurations()
	ResetIntentLabelingGenerationAiPromptId()
	ResetLocale()
	ResetQueryReformulationAiPromptId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference
type jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) AnswerGenerationAiGuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerGenerationAiGuardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) AnswerGenerationAiGuardrailIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerGenerationAiGuardrailIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) AnswerGenerationAiPromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerGenerationAiPromptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) AnswerGenerationAiPromptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"answerGenerationAiPromptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) AssociationConfigurations() WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationAssociationConfigurationsList {
	var returns WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationAssociationConfigurationsList
	_jsii_.Get(
		j,
		"associationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) AssociationConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associationConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) IntentLabelingGenerationAiPromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentLabelingGenerationAiPromptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) IntentLabelingGenerationAiPromptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentLabelingGenerationAiPromptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) QueryReformulationAiPromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryReformulationAiPromptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) QueryReformulationAiPromptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryReformulationAiPromptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference_Override(w WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetAnswerGenerationAiGuardrailId(val *string) {
	if err := j.validateSetAnswerGenerationAiGuardrailIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"answerGenerationAiGuardrailId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetAnswerGenerationAiPromptId(val *string) {
	if err := j.validateSetAnswerGenerationAiPromptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"answerGenerationAiPromptId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetIntentLabelingGenerationAiPromptId(val *string) {
	if err := j.validateSetIntentLabelingGenerationAiPromptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intentLabelingGenerationAiPromptId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetQueryReformulationAiPromptId(val *string) {
	if err := j.validateSetQueryReformulationAiPromptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryReformulationAiPromptId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) PutAssociationConfigurations(value interface{}) {
	if err := w.validatePutAssociationConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAssociationConfigurations",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ResetAnswerGenerationAiGuardrailId() {
	_jsii_.InvokeVoid(
		w,
		"resetAnswerGenerationAiGuardrailId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ResetAnswerGenerationAiPromptId() {
	_jsii_.InvokeVoid(
		w,
		"resetAnswerGenerationAiPromptId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ResetAssociationConfigurations() {
	_jsii_.InvokeVoid(
		w,
		"resetAssociationConfigurations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ResetIntentLabelingGenerationAiPromptId() {
	_jsii_.InvokeVoid(
		w,
		"resetIntentLabelingGenerationAiPromptId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ResetLocale() {
	_jsii_.InvokeVoid(
		w,
		"resetLocale",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ResetQueryReformulationAiPromptId() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryReformulationAiPromptId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

