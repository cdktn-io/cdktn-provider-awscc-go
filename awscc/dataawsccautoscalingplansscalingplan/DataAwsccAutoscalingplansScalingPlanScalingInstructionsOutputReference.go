// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccautoscalingplansscalingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccautoscalingplansscalingplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference interface {
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
	CustomizedLoadMetricSpecification() DataAwsccAutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference
	DisableDynamicScaling() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccAutoscalingplansScalingPlanScalingInstructions
	SetInternalValue(val *DataAwsccAutoscalingplansScalingPlanScalingInstructions)
	MaxCapacity() *float64
	MinCapacity() *float64
	PredefinedLoadMetricSpecification() DataAwsccAutoscalingplansScalingPlanScalingInstructionsPredefinedLoadMetricSpecificationOutputReference
	PredictiveScalingMaxCapacityBehavior() *string
	PredictiveScalingMaxCapacityBuffer() *float64
	PredictiveScalingMode() *string
	ResourceId() *string
	ScalableDimension() *string
	ScalingPolicyUpdateBehavior() *string
	ScheduledActionBufferTime() *float64
	ServiceNamespace() *string
	TargetTrackingConfigurations() DataAwsccAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference
type jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) CustomizedLoadMetricSpecification() DataAwsccAutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference {
	var returns DataAwsccAutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) DisableDynamicScaling() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableDynamicScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) InternalValue() *DataAwsccAutoscalingplansScalingPlanScalingInstructions {
	var returns *DataAwsccAutoscalingplansScalingPlanScalingInstructions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) PredefinedLoadMetricSpecification() DataAwsccAutoscalingplansScalingPlanScalingInstructionsPredefinedLoadMetricSpecificationOutputReference {
	var returns DataAwsccAutoscalingplansScalingPlanScalingInstructionsPredefinedLoadMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) PredictiveScalingMaxCapacityBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) PredictiveScalingMaxCapacityBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) PredictiveScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ScalingPolicyUpdateBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ScheduledActionBufferTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) TargetTrackingConfigurations() DataAwsccAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsList {
	var returns DataAwsccAutoscalingplansScalingPlanScalingInstructionsTargetTrackingConfigurationsList
	_jsii_.Get(
		j,
		"targetTrackingConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAutoscalingplansScalingPlan.DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference_Override(d DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccAutoscalingplansScalingPlan.DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference)SetInternalValue(val *DataAwsccAutoscalingplansScalingPlanScalingInstructions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingplansScalingPlanScalingInstructionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

