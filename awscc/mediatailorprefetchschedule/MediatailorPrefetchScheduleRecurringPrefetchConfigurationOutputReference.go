// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprefetchschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference interface {
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
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecurringConsumption() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionOutputReference
	RecurringConsumptionInput() interface{}
	RecurringRetrieval() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference
	RecurringRetrievalInput() interface{}
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
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
	PutRecurringConsumption(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumption)
	PutRecurringRetrieval(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrieval)
	ResetEndTime()
	ResetRecurringConsumption()
	ResetRecurringRetrieval()
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

// The jsii proxy struct for MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference
type jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) RecurringConsumption() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionOutputReference {
	var returns MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumptionOutputReference
	_jsii_.Get(
		j,
		"recurringConsumption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) RecurringConsumptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recurringConsumptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) RecurringRetrieval() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference {
	var returns MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference
	_jsii_.Get(
		j,
		"recurringRetrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) RecurringRetrievalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recurringRetrievalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference_Override(m MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) PutRecurringConsumption(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringConsumption) {
	if err := m.validatePutRecurringConsumptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRecurringConsumption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) PutRecurringRetrieval(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrieval) {
	if err := m.validatePutRecurringRetrievalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRecurringRetrieval",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		m,
		"resetEndTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ResetRecurringConsumption() {
	_jsii_.InvokeVoid(
		m,
		"resetRecurringConsumption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ResetRecurringRetrieval() {
	_jsii_.InvokeVoid(
		m,
		"resetRecurringRetrieval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

