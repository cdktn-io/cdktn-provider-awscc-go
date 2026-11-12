// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/cloudwatchalert/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchAlertRuleTelemetryRuleConditionOutputReference interface {
	cdktn.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
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
	CriticalThreshold() *float64
	SetCriticalThreshold(val *float64)
	CriticalThresholdInput() *float64
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
	ThresholdField() *string
	SetThresholdField(val *string)
	ThresholdFieldInput() *string
	ThresholdMode() *string
	SetThresholdMode(val *string)
	ThresholdModeInput() *string
	WarningThreshold() *float64
	SetWarningThreshold(val *float64)
	WarningThresholdInput() *float64
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
	ResetComparator()
	ResetCriticalThreshold()
	ResetThresholdField()
	ResetThresholdMode()
	ResetWarningThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchAlertRuleTelemetryRuleConditionOutputReference
type jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) CriticalThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticalThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) CriticalThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticalThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ThresholdField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ThresholdFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ThresholdMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ThresholdModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) WarningThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warningThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) WarningThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warningThresholdInput",
		&returns,
	)
	return returns
}


func NewCloudwatchAlertRuleTelemetryRuleConditionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudwatchAlertRuleTelemetryRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchAlertRuleTelemetryRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchAlert.CloudwatchAlertRuleTelemetryRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudwatchAlertRuleTelemetryRuleConditionOutputReference_Override(c CloudwatchAlertRuleTelemetryRuleConditionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchAlert.CloudwatchAlertRuleTelemetryRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetComparator(val *string) {
	if err := j.validateSetComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetCriticalThreshold(val *float64) {
	if err := j.validateSetCriticalThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"criticalThreshold",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetThresholdField(val *string) {
	if err := j.validateSetThresholdFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdField",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetThresholdMode(val *string) {
	if err := j.validateSetThresholdModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdMode",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference)SetWarningThreshold(val *float64) {
	if err := j.validateSetWarningThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warningThreshold",
		val,
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		c,
		"resetComparator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ResetCriticalThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetCriticalThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ResetThresholdField() {
	_jsii_.InvokeVoid(
		c,
		"resetThresholdField",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ResetThresholdMode() {
	_jsii_.InvokeVoid(
		c,
		"resetThresholdMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ResetWarningThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetWarningThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

