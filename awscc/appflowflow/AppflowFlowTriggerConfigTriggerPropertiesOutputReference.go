// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appflowflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppflowFlowTriggerConfigTriggerPropertiesOutputReference interface {
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
	DataPullMode() *string
	SetDataPullMode(val *string)
	DataPullModeInput() *string
	FirstExecutionFrom() *float64
	SetFirstExecutionFrom(val *float64)
	FirstExecutionFromInput() *float64
	FlowErrorDeactivationThreshold() *float64
	SetFlowErrorDeactivationThreshold(val *float64)
	FlowErrorDeactivationThresholdInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ScheduleEndTime() *float64
	SetScheduleEndTime(val *float64)
	ScheduleEndTimeInput() *float64
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	ScheduleExpressionInput() *string
	ScheduleOffset() *float64
	SetScheduleOffset(val *float64)
	ScheduleOffsetInput() *float64
	ScheduleStartTime() *float64
	SetScheduleStartTime(val *float64)
	ScheduleStartTimeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetDataPullMode()
	ResetFirstExecutionFrom()
	ResetFlowErrorDeactivationThreshold()
	ResetScheduleEndTime()
	ResetScheduleExpression()
	ResetScheduleOffset()
	ResetScheduleStartTime()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowTriggerConfigTriggerPropertiesOutputReference
type jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) DataPullMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPullMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) DataPullModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPullModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) FirstExecutionFrom() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstExecutionFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) FirstExecutionFromInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstExecutionFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) FlowErrorDeactivationThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowErrorDeactivationThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) FlowErrorDeactivationThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowErrorDeactivationThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleEndTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleEndTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleEndTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleOffsetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleStartTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ScheduleStartTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduleStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewAppflowFlowTriggerConfigTriggerPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppflowFlowTriggerConfigTriggerPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowTriggerConfigTriggerPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowTriggerConfigTriggerPropertiesOutputReference_Override(a AppflowFlowTriggerConfigTriggerPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appflowFlow.AppflowFlowTriggerConfigTriggerPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetDataPullMode(val *string) {
	if err := j.validateSetDataPullModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPullMode",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetFirstExecutionFrom(val *float64) {
	if err := j.validateSetFirstExecutionFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstExecutionFrom",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetFlowErrorDeactivationThreshold(val *float64) {
	if err := j.validateSetFlowErrorDeactivationThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowErrorDeactivationThreshold",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetScheduleEndTime(val *float64) {
	if err := j.validateSetScheduleEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleEndTime",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetScheduleExpression(val *string) {
	if err := j.validateSetScheduleExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetScheduleOffset(val *float64) {
	if err := j.validateSetScheduleOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleOffset",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetScheduleStartTime(val *float64) {
	if err := j.validateSetScheduleStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleStartTime",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetDataPullMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPullMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetFirstExecutionFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstExecutionFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetFlowErrorDeactivationThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetFlowErrorDeactivationThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetScheduleEndTime() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleEndTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetScheduleExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetScheduleOffset() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleOffset",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetScheduleStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowTriggerConfigTriggerPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

