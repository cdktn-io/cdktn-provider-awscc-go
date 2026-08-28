// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprefetchschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorPrefetchScheduleRetrievalOutputReference interface {
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
	DynamicVariables() *map[string]*string
	SetDynamicVariables(val *map[string]*string)
	DynamicVariablesInput() *map[string]*string
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	TrafficShapingRetrievalWindow() MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindowOutputReference
	TrafficShapingRetrievalWindowInput() interface{}
	TrafficShapingTpsConfiguration() MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference
	TrafficShapingTpsConfigurationInput() interface{}
	TrafficShapingType() *string
	SetTrafficShapingType(val *string)
	TrafficShapingTypeInput() *string
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
	PutTrafficShapingRetrievalWindow(value *MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindow)
	PutTrafficShapingTpsConfiguration(value *MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfiguration)
	ResetDynamicVariables()
	ResetEndTime()
	ResetStartTime()
	ResetTrafficShapingRetrievalWindow()
	ResetTrafficShapingTpsConfiguration()
	ResetTrafficShapingType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorPrefetchScheduleRetrievalOutputReference
type jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) DynamicVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) DynamicVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TrafficShapingRetrievalWindow() MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindowOutputReference {
	var returns MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindowOutputReference
	_jsii_.Get(
		j,
		"trafficShapingRetrievalWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TrafficShapingRetrievalWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficShapingRetrievalWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TrafficShapingTpsConfiguration() MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference {
	var returns MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference
	_jsii_.Get(
		j,
		"trafficShapingTpsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TrafficShapingTpsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficShapingTpsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TrafficShapingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficShapingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) TrafficShapingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficShapingTypeInput",
		&returns,
	)
	return returns
}


func NewMediatailorPrefetchScheduleRetrievalOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediatailorPrefetchScheduleRetrievalOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorPrefetchScheduleRetrievalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediatailorPrefetchScheduleRetrievalOutputReference_Override(m MediatailorPrefetchScheduleRetrievalOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetDynamicVariables(val *map[string]*string) {
	if err := j.validateSetDynamicVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicVariables",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference)SetTrafficShapingType(val *string) {
	if err := j.validateSetTrafficShapingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficShapingType",
		val,
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) PutTrafficShapingRetrievalWindow(value *MediatailorPrefetchScheduleRetrievalTrafficShapingRetrievalWindow) {
	if err := m.validatePutTrafficShapingRetrievalWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTrafficShapingRetrievalWindow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) PutTrafficShapingTpsConfiguration(value *MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfiguration) {
	if err := m.validatePutTrafficShapingTpsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTrafficShapingTpsConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ResetDynamicVariables() {
	_jsii_.InvokeVoid(
		m,
		"resetDynamicVariables",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		m,
		"resetEndTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ResetTrafficShapingRetrievalWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficShapingRetrievalWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ResetTrafficShapingTpsConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficShapingTpsConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ResetTrafficShapingType() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficShapingType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

