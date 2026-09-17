// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/transcribecallanalyticscategory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference interface {
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
	EndTime() *float64
	SetEndTime(val *float64)
	EndTimeInput() *float64
	First() *float64
	SetFirst(val *float64)
	FirstInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Last() *float64
	SetLast(val *float64)
	LastInput() *float64
	StartTime() *float64
	SetStartTime(val *float64)
	StartTimeInput() *float64
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
	ResetEndTime()
	ResetFirst()
	ResetLast()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference
type jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) EndTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) EndTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) First() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) FirstInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) Last() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"last",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) LastInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) StartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) StartTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference {
	_init_.Initialize()

	if err := validateNewTranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference_Override(t TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetEndTime(val *float64) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetFirst(val *float64) {
	if err := j.validateSetFirstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"first",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetLast(val *float64) {
	if err := j.validateSetLastParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"last",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetStartTime(val *float64) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		t,
		"resetEndTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ResetFirst() {
	_jsii_.InvokeVoid(
		t,
		"resetFirst",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ResetLast() {
	_jsii_.InvokeVoid(
		t,
		"resetLast",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterAbsoluteTimeRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

