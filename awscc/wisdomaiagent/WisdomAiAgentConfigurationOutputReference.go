// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationOutputReference interface {
	cdktn.ComplexObject
	AnswerRecommendationAiAgentConfiguration() WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference
	AnswerRecommendationAiAgentConfigurationInput() interface{}
	CaseSummarizationAiAgentConfiguration() WisdomAiAgentConfigurationCaseSummarizationAiAgentConfigurationOutputReference
	CaseSummarizationAiAgentConfigurationInput() interface{}
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
	EmailGenerativeAnswerAiAgentConfiguration() WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference
	EmailGenerativeAnswerAiAgentConfigurationInput() interface{}
	EmailOverviewAiAgentConfiguration() WisdomAiAgentConfigurationEmailOverviewAiAgentConfigurationOutputReference
	EmailOverviewAiAgentConfigurationInput() interface{}
	EmailResponseAiAgentConfiguration() WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationOutputReference
	EmailResponseAiAgentConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManualSearchAiAgentConfiguration() WisdomAiAgentConfigurationManualSearchAiAgentConfigurationOutputReference
	ManualSearchAiAgentConfigurationInput() interface{}
	NoteTakingAiAgentConfiguration() WisdomAiAgentConfigurationNoteTakingAiAgentConfigurationOutputReference
	NoteTakingAiAgentConfigurationInput() interface{}
	OrchestrationAiAgentConfiguration() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference
	OrchestrationAiAgentConfigurationInput() interface{}
	SelfServiceAiAgentConfiguration() WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationOutputReference
	SelfServiceAiAgentConfigurationInput() interface{}
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
	PutAnswerRecommendationAiAgentConfiguration(value *WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfiguration)
	PutCaseSummarizationAiAgentConfiguration(value *WisdomAiAgentConfigurationCaseSummarizationAiAgentConfiguration)
	PutEmailGenerativeAnswerAiAgentConfiguration(value *WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfiguration)
	PutEmailOverviewAiAgentConfiguration(value *WisdomAiAgentConfigurationEmailOverviewAiAgentConfiguration)
	PutEmailResponseAiAgentConfiguration(value *WisdomAiAgentConfigurationEmailResponseAiAgentConfiguration)
	PutManualSearchAiAgentConfiguration(value *WisdomAiAgentConfigurationManualSearchAiAgentConfiguration)
	PutNoteTakingAiAgentConfiguration(value *WisdomAiAgentConfigurationNoteTakingAiAgentConfiguration)
	PutOrchestrationAiAgentConfiguration(value *WisdomAiAgentConfigurationOrchestrationAiAgentConfiguration)
	PutSelfServiceAiAgentConfiguration(value *WisdomAiAgentConfigurationSelfServiceAiAgentConfiguration)
	ResetAnswerRecommendationAiAgentConfiguration()
	ResetCaseSummarizationAiAgentConfiguration()
	ResetEmailGenerativeAnswerAiAgentConfiguration()
	ResetEmailOverviewAiAgentConfiguration()
	ResetEmailResponseAiAgentConfiguration()
	ResetManualSearchAiAgentConfiguration()
	ResetNoteTakingAiAgentConfiguration()
	ResetOrchestrationAiAgentConfiguration()
	ResetSelfServiceAiAgentConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationOutputReference
type jsiiProxy_WisdomAiAgentConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) AnswerRecommendationAiAgentConfiguration() WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"answerRecommendationAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) AnswerRecommendationAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"answerRecommendationAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) CaseSummarizationAiAgentConfiguration() WisdomAiAgentConfigurationCaseSummarizationAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationCaseSummarizationAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"caseSummarizationAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) CaseSummarizationAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSummarizationAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) EmailGenerativeAnswerAiAgentConfiguration() WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"emailGenerativeAnswerAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) EmailGenerativeAnswerAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailGenerativeAnswerAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) EmailOverviewAiAgentConfiguration() WisdomAiAgentConfigurationEmailOverviewAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationEmailOverviewAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"emailOverviewAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) EmailOverviewAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailOverviewAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) EmailResponseAiAgentConfiguration() WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"emailResponseAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) EmailResponseAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailResponseAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ManualSearchAiAgentConfiguration() WisdomAiAgentConfigurationManualSearchAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationManualSearchAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"manualSearchAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ManualSearchAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualSearchAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) NoteTakingAiAgentConfiguration() WisdomAiAgentConfigurationNoteTakingAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationNoteTakingAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"noteTakingAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) NoteTakingAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTakingAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) OrchestrationAiAgentConfiguration() WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"orchestrationAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) OrchestrationAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orchestrationAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) SelfServiceAiAgentConfiguration() WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationOutputReference {
	var returns WisdomAiAgentConfigurationSelfServiceAiAgentConfigurationOutputReference
	_jsii_.Get(
		j,
		"selfServiceAiAgentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) SelfServiceAiAgentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfServiceAiAgentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WisdomAiAgentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationOutputReference_Override(w WisdomAiAgentConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutAnswerRecommendationAiAgentConfiguration(value *WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfiguration) {
	if err := w.validatePutAnswerRecommendationAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAnswerRecommendationAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutCaseSummarizationAiAgentConfiguration(value *WisdomAiAgentConfigurationCaseSummarizationAiAgentConfiguration) {
	if err := w.validatePutCaseSummarizationAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCaseSummarizationAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutEmailGenerativeAnswerAiAgentConfiguration(value *WisdomAiAgentConfigurationEmailGenerativeAnswerAiAgentConfiguration) {
	if err := w.validatePutEmailGenerativeAnswerAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEmailGenerativeAnswerAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutEmailOverviewAiAgentConfiguration(value *WisdomAiAgentConfigurationEmailOverviewAiAgentConfiguration) {
	if err := w.validatePutEmailOverviewAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEmailOverviewAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutEmailResponseAiAgentConfiguration(value *WisdomAiAgentConfigurationEmailResponseAiAgentConfiguration) {
	if err := w.validatePutEmailResponseAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEmailResponseAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutManualSearchAiAgentConfiguration(value *WisdomAiAgentConfigurationManualSearchAiAgentConfiguration) {
	if err := w.validatePutManualSearchAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putManualSearchAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutNoteTakingAiAgentConfiguration(value *WisdomAiAgentConfigurationNoteTakingAiAgentConfiguration) {
	if err := w.validatePutNoteTakingAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNoteTakingAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutOrchestrationAiAgentConfiguration(value *WisdomAiAgentConfigurationOrchestrationAiAgentConfiguration) {
	if err := w.validatePutOrchestrationAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOrchestrationAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) PutSelfServiceAiAgentConfiguration(value *WisdomAiAgentConfigurationSelfServiceAiAgentConfiguration) {
	if err := w.validatePutSelfServiceAiAgentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSelfServiceAiAgentConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetAnswerRecommendationAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetAnswerRecommendationAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetCaseSummarizationAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetCaseSummarizationAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetEmailGenerativeAnswerAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetEmailGenerativeAnswerAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetEmailOverviewAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetEmailOverviewAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetEmailResponseAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetEmailResponseAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetManualSearchAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetManualSearchAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetNoteTakingAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetNoteTakingAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetOrchestrationAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetOrchestrationAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ResetSelfServiceAiAgentConfiguration() {
	_jsii_.InvokeVoid(
		w,
		"resetSelfServiceAiAgentConfiguration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WisdomAiAgentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

