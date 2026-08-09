// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesthinclientenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspacesthinclientenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference interface {
	cdktn.ComplexObject
	ApplyTimeOf() *string
	SetApplyTimeOf(val *string)
	ApplyTimeOfInput() *string
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
	DaysOfTheWeek() *[]*string
	SetDaysOfTheWeek(val *[]*string)
	DaysOfTheWeekInput() *[]*string
	EndTimeHour() *float64
	SetEndTimeHour(val *float64)
	EndTimeHourInput() *float64
	EndTimeMinute() *float64
	SetEndTimeMinute(val *float64)
	EndTimeMinuteInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartTimeHour() *float64
	SetStartTimeHour(val *float64)
	StartTimeHourInput() *float64
	StartTimeMinute() *float64
	SetStartTimeMinute(val *float64)
	StartTimeMinuteInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetApplyTimeOf()
	ResetDaysOfTheWeek()
	ResetEndTimeHour()
	ResetEndTimeMinute()
	ResetStartTimeHour()
	ResetStartTimeMinute()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference
type jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ApplyTimeOf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyTimeOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ApplyTimeOfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applyTimeOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) DaysOfTheWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfTheWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) DaysOfTheWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfTheWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) EndTimeHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTimeHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) EndTimeHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTimeHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) EndTimeMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTimeMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) EndTimeMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endTimeMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) StartTimeHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) StartTimeHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) StartTimeMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) StartTimeMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewWorkspacesthinclientEnvironmentMaintenanceWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesthinclientEnvironmentMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesthinclientEnvironmentMaintenanceWindowOutputReference_Override(w WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetApplyTimeOf(val *string) {
	if err := j.validateSetApplyTimeOfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyTimeOf",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetDaysOfTheWeek(val *[]*string) {
	if err := j.validateSetDaysOfTheWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfTheWeek",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetEndTimeHour(val *float64) {
	if err := j.validateSetEndTimeHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeHour",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetEndTimeMinute(val *float64) {
	if err := j.validateSetEndTimeMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeMinute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetStartTimeHour(val *float64) {
	if err := j.validateSetStartTimeHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeHour",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetStartTimeMinute(val *float64) {
	if err := j.validateSetStartTimeMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeMinute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetApplyTimeOf() {
	_jsii_.InvokeVoid(
		w,
		"resetApplyTimeOf",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetDaysOfTheWeek() {
	_jsii_.InvokeVoid(
		w,
		"resetDaysOfTheWeek",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetEndTimeHour() {
	_jsii_.InvokeVoid(
		w,
		"resetEndTimeHour",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetEndTimeMinute() {
	_jsii_.InvokeVoid(
		w,
		"resetEndTimeMinute",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetStartTimeHour() {
	_jsii_.InvokeVoid(
		w,
		"resetStartTimeHour",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetStartTimeMinute() {
	_jsii_.InvokeVoid(
		w,
		"resetStartTimeMinute",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		w,
		"resetType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

