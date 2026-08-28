// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprefetchschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference interface {
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
	DelayAfterAvailEndSeconds() *float64
	SetDelayAfterAvailEndSeconds(val *float64)
	DelayAfterAvailEndSecondsInput() *float64
	DynamicVariables() *map[string]*string
	SetDynamicVariables(val *map[string]*string)
	DynamicVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
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
	TrafficShapingRetrievalWindow() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference
	TrafficShapingRetrievalWindowInput() interface{}
	TrafficShapingTpsConfiguration() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfigurationOutputReference
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
	PutTrafficShapingRetrievalWindow(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow)
	PutTrafficShapingTpsConfiguration(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfiguration)
	ResetDelayAfterAvailEndSeconds()
	ResetDynamicVariables()
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

// The jsii proxy struct for MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference
type jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) DelayAfterAvailEndSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayAfterAvailEndSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) DelayAfterAvailEndSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayAfterAvailEndSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) DynamicVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) DynamicVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TrafficShapingRetrievalWindow() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference {
	var returns MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference
	_jsii_.Get(
		j,
		"trafficShapingRetrievalWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TrafficShapingRetrievalWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficShapingRetrievalWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TrafficShapingTpsConfiguration() MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfigurationOutputReference {
	var returns MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfigurationOutputReference
	_jsii_.Get(
		j,
		"trafficShapingTpsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TrafficShapingTpsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficShapingTpsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TrafficShapingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficShapingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) TrafficShapingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficShapingTypeInput",
		&returns,
	)
	return returns
}


func NewMediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference_Override(m MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetDelayAfterAvailEndSeconds(val *float64) {
	if err := j.validateSetDelayAfterAvailEndSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delayAfterAvailEndSeconds",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetDynamicVariables(val *map[string]*string) {
	if err := j.validateSetDynamicVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicVariables",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference)SetTrafficShapingType(val *string) {
	if err := j.validateSetTrafficShapingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficShapingType",
		val,
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) PutTrafficShapingRetrievalWindow(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow) {
	if err := m.validatePutTrafficShapingRetrievalWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTrafficShapingRetrievalWindow",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) PutTrafficShapingTpsConfiguration(value *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingTpsConfiguration) {
	if err := m.validatePutTrafficShapingTpsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTrafficShapingTpsConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ResetDelayAfterAvailEndSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetDelayAfterAvailEndSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ResetDynamicVariables() {
	_jsii_.InvokeVoid(
		m,
		"resetDynamicVariables",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ResetTrafficShapingRetrievalWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficShapingRetrievalWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ResetTrafficShapingTpsConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficShapingTpsConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ResetTrafficShapingType() {
	_jsii_.InvokeVoid(
		m,
		"resetTrafficShapingType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

