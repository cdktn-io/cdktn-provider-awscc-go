// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainerfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gameliftcontainerfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftContainerFleetScalingPoliciesOutputReference interface {
	cdktn.ComplexObject
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	ComparisonOperatorInput() *string
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
	EvaluationPeriods() *float64
	SetEvaluationPeriods(val *float64)
	EvaluationPeriodsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
	ScalingAdjustment() *float64
	SetScalingAdjustment(val *float64)
	ScalingAdjustmentInput() *float64
	ScalingAdjustmentType() *string
	SetScalingAdjustmentType(val *string)
	ScalingAdjustmentTypeInput() *string
	TargetConfiguration() GameliftContainerFleetScalingPoliciesTargetConfigurationOutputReference
	TargetConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
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
	PutTargetConfiguration(value *GameliftContainerFleetScalingPoliciesTargetConfiguration)
	ResetComparisonOperator()
	ResetEvaluationPeriods()
	ResetMetricName()
	ResetName()
	ResetPolicyType()
	ResetScalingAdjustment()
	ResetScalingAdjustmentType()
	ResetTargetConfiguration()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameliftContainerFleetScalingPoliciesOutputReference
type jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ScalingAdjustment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scalingAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ScalingAdjustmentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scalingAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ScalingAdjustmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingAdjustmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ScalingAdjustmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingAdjustmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) TargetConfiguration() GameliftContainerFleetScalingPoliciesTargetConfigurationOutputReference {
	var returns GameliftContainerFleetScalingPoliciesTargetConfigurationOutputReference
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) TargetConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewGameliftContainerFleetScalingPoliciesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GameliftContainerFleetScalingPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewGameliftContainerFleetScalingPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleetScalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGameliftContainerFleetScalingPoliciesOutputReference_Override(g GameliftContainerFleetScalingPoliciesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerFleet.GameliftContainerFleetScalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetScalingAdjustment(val *float64) {
	if err := j.validateSetScalingAdjustmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingAdjustment",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetScalingAdjustmentType(val *string) {
	if err := j.validateSetScalingAdjustmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingAdjustmentType",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) PutTargetConfiguration(value *GameliftContainerFleetScalingPoliciesTargetConfiguration) {
	if err := g.validatePutTargetConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTargetConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetComparisonOperator() {
	_jsii_.InvokeVoid(
		g,
		"resetComparisonOperator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetEvaluationPeriods() {
	_jsii_.InvokeVoid(
		g,
		"resetEvaluationPeriods",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetMetricName() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetScalingAdjustment() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingAdjustment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetScalingAdjustmentType() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingAdjustmentType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetTargetConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerFleetScalingPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

