// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprefetchschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprefetchschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PeakConcurrentUsers() *float64
	SetPeakConcurrentUsers(val *float64)
	PeakConcurrentUsersInput() *float64
	PeakTps() *float64
	SetPeakTps(val *float64)
	PeakTpsInput() *float64
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
	ResetPeakConcurrentUsers()
	ResetPeakTps()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference
type jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) PeakConcurrentUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peakConcurrentUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) PeakConcurrentUsersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peakConcurrentUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) PeakTps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peakTps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) PeakTpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peakTpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference_Override(m MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorPrefetchSchedule.MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetPeakConcurrentUsers(val *float64) {
	if err := j.validateSetPeakConcurrentUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peakConcurrentUsers",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetPeakTps(val *float64) {
	if err := j.validateSetPeakTpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peakTps",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) ResetPeakConcurrentUsers() {
	_jsii_.InvokeVoid(
		m,
		"resetPeakConcurrentUsers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) ResetPeakTps() {
	_jsii_.InvokeVoid(
		m,
		"resetPeakTps",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorPrefetchScheduleRetrievalTrafficShapingTpsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

