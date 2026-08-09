// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stepfunctionsstatemachinealias

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/stepfunctionsstatemachinealias/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference interface {
	cdktn.ComplexObject
	Alarms() *[]*string
	SetAlarms(val *[]*string)
	AlarmsInput() *[]*string
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
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Percentage() *float64
	SetPercentage(val *float64)
	PercentageInput() *float64
	StateMachineVersionArn() *string
	SetStateMachineVersionArn(val *string)
	StateMachineVersionArnInput() *string
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
	ResetAlarms()
	ResetInterval()
	ResetPercentage()
	ResetStateMachineVersionArn()
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

// The jsii proxy struct for StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference
type jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) Alarms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) AlarmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) PercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) StateMachineVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) StateMachineVersionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineVersionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewStepfunctionsStateMachineAliasDeploymentPreferenceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference {
	_init_.Initialize()

	if err := validateNewStepfunctionsStateMachineAliasDeploymentPreferenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.stepfunctionsStateMachineAlias.StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStepfunctionsStateMachineAliasDeploymentPreferenceOutputReference_Override(s StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.stepfunctionsStateMachineAlias.StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetAlarms(val *[]*string) {
	if err := j.validateSetAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarms",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetPercentage(val *float64) {
	if err := j.validateSetPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentage",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetStateMachineVersionArn(val *string) {
	if err := j.validateSetStateMachineVersionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateMachineVersionArn",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		s,
		"resetAlarms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		s,
		"resetInterval",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetPercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ResetStateMachineVersionArn() {
	_jsii_.InvokeVoid(
		s,
		"resetStateMachineVersionArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepfunctionsStateMachineAliasDeploymentPreferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

