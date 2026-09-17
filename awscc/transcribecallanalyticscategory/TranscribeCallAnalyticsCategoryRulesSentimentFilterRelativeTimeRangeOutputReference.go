// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/transcribecallanalyticscategory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference interface {
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
	EndPercentage() *float64
	SetEndPercentage(val *float64)
	EndPercentageInput() *float64
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
	StartPercentage() *float64
	SetStartPercentage(val *float64)
	StartPercentageInput() *float64
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
	ResetEndPercentage()
	ResetFirst()
	ResetLast()
	ResetStartPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference
type jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) EndPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) EndPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) First() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"first",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) FirstInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) Last() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"last",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) LastInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) StartPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) StartPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference {
	_init_.Initialize()

	if err := validateNewTranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference_Override(t TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetEndPercentage(val *float64) {
	if err := j.validateSetEndPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endPercentage",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetFirst(val *float64) {
	if err := j.validateSetFirstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"first",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetLast(val *float64) {
	if err := j.validateSetLastParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"last",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetStartPercentage(val *float64) {
	if err := j.validateSetStartPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startPercentage",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ResetEndPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetEndPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ResetFirst() {
	_jsii_.InvokeVoid(
		t,
		"resetFirst",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ResetLast() {
	_jsii_.InvokeVoid(
		t,
		"resetLast",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ResetStartPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetStartPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesSentimentFilterRelativeTimeRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

