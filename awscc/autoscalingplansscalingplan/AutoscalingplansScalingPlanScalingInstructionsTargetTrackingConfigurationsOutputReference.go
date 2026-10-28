// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/autoscalingplansscalingplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference interface {
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
	CustomizedScalingMetricSpecification() AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsCustomizedScalingMetricSpecificationOutputReference
	CustomizedScalingMetricSpecificationInput() interface{}
	DisableScaleIn() interface{}
	SetDisableScaleIn(val interface{})
	DisableScaleInInput() interface{}
	EstimatedInstanceWarmup() *float64
	SetEstimatedInstanceWarmup(val *float64)
	EstimatedInstanceWarmupInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PredefinedScalingMetricSpecification() AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference
	PredefinedScalingMetricSpecificationInput() interface{}
	ScaleInCooldown() *float64
	SetScaleInCooldown(val *float64)
	ScaleInCooldownInput() *float64
	ScaleOutCooldown() *float64
	SetScaleOutCooldown(val *float64)
	ScaleOutCooldownInput() *float64
	TargetValue() *float64
	SetTargetValue(val *float64)
	TargetValueInput() *float64
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
	PutCustomizedScalingMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsCustomizedScalingMetricSpecification)
	PutPredefinedScalingMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecification)
	ResetCustomizedScalingMetricSpecification()
	ResetDisableScaleIn()
	ResetEstimatedInstanceWarmup()
	ResetPredefinedScalingMetricSpecification()
	ResetScaleInCooldown()
	ResetScaleOutCooldown()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference
type jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) CustomizedScalingMetricSpecification() AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsCustomizedScalingMetricSpecificationOutputReference {
	var returns AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsCustomizedScalingMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) CustomizedScalingMetricSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) DisableScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) DisableScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) EstimatedInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) EstimatedInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) PredefinedScalingMetricSpecification() AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference {
	var returns AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) PredefinedScalingMetricSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ScaleInCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ScaleInCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ScaleOutCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ScaleOutCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingplansScalingPlan.AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference_Override(a AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingplansScalingPlan.AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetDisableScaleIn(val interface{}) {
	if err := j.validateSetDisableScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableScaleIn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetEstimatedInstanceWarmup(val *float64) {
	if err := j.validateSetEstimatedInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"estimatedInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetScaleInCooldown(val *float64) {
	if err := j.validateSetScaleInCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInCooldown",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetScaleOutCooldown(val *float64) {
	if err := j.validateSetScaleOutCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutCooldown",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) PutCustomizedScalingMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsCustomizedScalingMetricSpecification) {
	if err := a.validatePutCustomizedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) PutPredefinedScalingMetricSpecification(value *AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecification) {
	if err := a.validatePutPredefinedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ResetCustomizedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ResetDisableScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ResetEstimatedInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetEstimatedInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ResetPredefinedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ResetScaleInCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleInCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ResetScaleOutCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleOutCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

