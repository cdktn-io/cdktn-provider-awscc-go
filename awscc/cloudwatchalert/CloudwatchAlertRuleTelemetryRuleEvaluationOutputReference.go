// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/cloudwatchalert/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference interface {
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
	IntervalSeconds() *float64
	SetIntervalSeconds(val *float64)
	IntervalSecondsInput() *float64
	PendingDurationSeconds() *float64
	SetPendingDurationSeconds(val *float64)
	PendingDurationSecondsInput() *float64
	RecoveryDurationSeconds() *float64
	SetRecoveryDurationSeconds(val *float64)
	RecoveryDurationSecondsInput() *float64
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
	ResetPendingDurationSeconds()
	ResetRecoveryDurationSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference
type jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) PendingDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) PendingDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pendingDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) RecoveryDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) RecoveryDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudwatchAlertRuleTelemetryRuleEvaluationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchAlertRuleTelemetryRuleEvaluationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchAlert.CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudwatchAlertRuleTelemetryRuleEvaluationOutputReference_Override(c CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchAlert.CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetPendingDurationSeconds(val *float64) {
	if err := j.validateSetPendingDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pendingDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetRecoveryDurationSeconds(val *float64) {
	if err := j.validateSetRecoveryDurationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) ResetPendingDurationSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetPendingDurationSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) ResetRecoveryDurationSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetRecoveryDurationSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

