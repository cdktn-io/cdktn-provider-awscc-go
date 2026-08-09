// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apsworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApsWorkspaceWorkspaceConfigurationOutputReference interface {
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
	LimitsPerLabelSets() ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsList
	LimitsPerLabelSetsInput() interface{}
	OutOfOrderTimeWindowInSeconds() *float64
	SetOutOfOrderTimeWindowInSeconds(val *float64)
	OutOfOrderTimeWindowInSecondsInput() *float64
	RetentionPeriodInDays() *float64
	SetRetentionPeriodInDays(val *float64)
	RetentionPeriodInDaysInput() *float64
	RuleQueryOffsetInSeconds() *float64
	SetRuleQueryOffsetInSeconds(val *float64)
	RuleQueryOffsetInSecondsInput() *float64
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
	PutLimitsPerLabelSets(value interface{})
	ResetLimitsPerLabelSets()
	ResetOutOfOrderTimeWindowInSeconds()
	ResetRetentionPeriodInDays()
	ResetRuleQueryOffsetInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApsWorkspaceWorkspaceConfigurationOutputReference
type jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) LimitsPerLabelSets() ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsList {
	var returns ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsList
	_jsii_.Get(
		j,
		"limitsPerLabelSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) LimitsPerLabelSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitsPerLabelSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) OutOfOrderTimeWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outOfOrderTimeWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) OutOfOrderTimeWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outOfOrderTimeWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) RetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) RetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) RuleQueryOffsetInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleQueryOffsetInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) RuleQueryOffsetInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleQueryOffsetInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApsWorkspaceWorkspaceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApsWorkspaceWorkspaceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewApsWorkspaceWorkspaceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apsWorkspace.ApsWorkspaceWorkspaceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApsWorkspaceWorkspaceConfigurationOutputReference_Override(a ApsWorkspaceWorkspaceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apsWorkspace.ApsWorkspaceWorkspaceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetOutOfOrderTimeWindowInSeconds(val *float64) {
	if err := j.validateSetOutOfOrderTimeWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outOfOrderTimeWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetRetentionPeriodInDays(val *float64) {
	if err := j.validateSetRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetRuleQueryOffsetInSeconds(val *float64) {
	if err := j.validateSetRuleQueryOffsetInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleQueryOffsetInSeconds",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) PutLimitsPerLabelSets(value interface{}) {
	if err := a.validatePutLimitsPerLabelSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLimitsPerLabelSets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ResetLimitsPerLabelSets() {
	_jsii_.InvokeVoid(
		a,
		"resetLimitsPerLabelSets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ResetOutOfOrderTimeWindowInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetOutOfOrderTimeWindowInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ResetRetentionPeriodInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionPeriodInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ResetRuleQueryOffsetInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleQueryOffsetInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApsWorkspaceWorkspaceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

