// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/autoscalingplansscalingplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference interface {
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
	PredefinedScalingMetricType() *string
	SetPredefinedScalingMetricType(val *string)
	PredefinedScalingMetricTypeInput() *string
	ResourceLabel() *string
	SetResourceLabel(val *string)
	ResourceLabelInput() *string
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
	ResetPredefinedScalingMetricType()
	ResetResourceLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference
type jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) PredefinedScalingMetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedScalingMetricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) PredefinedScalingMetricTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedScalingMetricTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ResourceLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ResourceLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingplansScalingPlan.AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference_Override(a AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingplansScalingPlan.AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetPredefinedScalingMetricType(val *string) {
	if err := j.validateSetPredefinedScalingMetricTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedScalingMetricType",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetResourceLabel(val *string) {
	if err := j.validateSetResourceLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabel",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ResetPredefinedScalingMetricType() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedScalingMetricType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ResetResourceLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsPredefinedScalingMetricSpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

