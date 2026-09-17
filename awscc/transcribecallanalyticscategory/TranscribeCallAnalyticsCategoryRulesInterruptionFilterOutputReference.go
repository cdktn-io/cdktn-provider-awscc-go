// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transcribecallanalyticscategory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/transcribecallanalyticscategory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference interface {
	cdktn.ComplexObject
	AbsoluteTimeRange() TranscribeCallAnalyticsCategoryRulesInterruptionFilterAbsoluteTimeRangeOutputReference
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
	RelativeTimeRange() TranscribeCallAnalyticsCategoryRulesInterruptionFilterRelativeTimeRangeOutputReference
	RelativeTimeRangeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
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
	PutAbsoluteTimeRange(value *TranscribeCallAnalyticsCategoryRulesInterruptionFilterAbsoluteTimeRange)
	PutRelativeTimeRange(value *TranscribeCallAnalyticsCategoryRulesInterruptionFilterRelativeTimeRange)
	ResetAbsoluteTimeRange()
	ResetNegate()
	ResetParticipantRole()
	ResetRelativeTimeRange()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference
type jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) AbsoluteTimeRange() TranscribeCallAnalyticsCategoryRulesInterruptionFilterAbsoluteTimeRangeOutputReference {
	var returns TranscribeCallAnalyticsCategoryRulesInterruptionFilterAbsoluteTimeRangeOutputReference
	_jsii_.Get(
		j,
		"absoluteTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) AbsoluteTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"absoluteTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) Negate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) NegateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ParticipantRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"participantRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ParticipantRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"participantRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) RelativeTimeRange() TranscribeCallAnalyticsCategoryRulesInterruptionFilterRelativeTimeRangeOutputReference {
	var returns TranscribeCallAnalyticsCategoryRulesInterruptionFilterRelativeTimeRangeOutputReference
	_jsii_.Get(
		j,
		"relativeTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) RelativeTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relativeTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewTranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference {
	_init_.Initialize()

	if err := validateNewTranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference_Override(t TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.transcribeCallAnalyticsCategory.TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetNegate(val interface{}) {
	if err := j.validateSetNegateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negate",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetParticipantRole(val *string) {
	if err := j.validateSetParticipantRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"participantRole",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) PutAbsoluteTimeRange(value *TranscribeCallAnalyticsCategoryRulesInterruptionFilterAbsoluteTimeRange) {
	if err := t.validatePutAbsoluteTimeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAbsoluteTimeRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) PutRelativeTimeRange(value *TranscribeCallAnalyticsCategoryRulesInterruptionFilterRelativeTimeRange) {
	if err := t.validatePutRelativeTimeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelativeTimeRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ResetAbsoluteTimeRange() {
	_jsii_.InvokeVoid(
		t,
		"resetAbsoluteTimeRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ResetNegate() {
	_jsii_.InvokeVoid(
		t,
		"resetNegate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ResetParticipantRole() {
	_jsii_.InvokeVoid(
		t,
		"resetParticipantRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ResetRelativeTimeRange() {
	_jsii_.InvokeVoid(
		t,
		"resetRelativeTimeRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		t,
		"resetThreshold",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscribeCallAnalyticsCategoryRulesInterruptionFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

