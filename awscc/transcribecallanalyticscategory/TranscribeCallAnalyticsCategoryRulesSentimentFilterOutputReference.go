// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/transcribecallanalyticscategory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference interface {
	cdktn.ComplexObject
	AbsoluteTimeRange() TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference
	AbsoluteTimeRangeInput() interface{}
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
	Negate() interface{}
	SetNegate(val interface{})
	NegateInput() interface{}
	ParticipantRole() *string
	SetParticipantRole(val *string)
	ParticipantRoleInput() *string
	RelativeTimeRange() TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference
	RelativeTimeRangeInput() interface{}
	Sentiments() *[]*string
	SetSentiments(val *[]*string)
	SentimentsInput() *[]*string
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
	PutAbsoluteTimeRange(value *TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRange)
	PutRelativeTimeRange(value *TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRange)
	ResetAbsoluteTimeRange()
	ResetNegate()
	ResetParticipantRole()
	ResetRelativeTimeRange()
	ResetSentiments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference
type jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) AbsoluteTimeRange() TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference {
	var returns TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference
	_jsii_.Get(
		j,
		"absoluteTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) AbsoluteTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"absoluteTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) Negate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) NegateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ParticipantRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"participantRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ParticipantRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"participantRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) RelativeTimeRange() TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference {
	var returns TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference
	_jsii_.Get(
		j,
		"relativeTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) RelativeTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relativeTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) Sentiments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sentiments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) SentimentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sentimentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference {
	_init_.Initialize()

	if err := validateNewTranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference_Override(t TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetNegate(val interface{}) {
	if err := j.validateSetNegateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negate",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetParticipantRole(val *string) {
	if err := j.validateSetParticipantRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"participantRole",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetSentiments(val *[]*string) {
	if err := j.validateSetSentimentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sentiments",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) PutAbsoluteTimeRange(value *TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRange) {
	if err := t.validatePutAbsoluteTimeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAbsoluteTimeRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) PutRelativeTimeRange(value *TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRange) {
	if err := t.validatePutRelativeTimeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelativeTimeRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ResetAbsoluteTimeRange() {
	_jsii_.InvokeVoid(
		t,
		"resetAbsoluteTimeRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ResetNegate() {
	_jsii_.InvokeVoid(
		t,
		"resetNegate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ResetParticipantRole() {
	_jsii_.InvokeVoid(
		t,
		"resetParticipantRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ResetRelativeTimeRange() {
	_jsii_.InvokeVoid(
		t,
		"resetRelativeTimeRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ResetSentiments() {
	_jsii_.InvokeVoid(
		t,
		"resetSentiments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

