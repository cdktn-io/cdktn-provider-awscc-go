// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/autoscalingautoscalinggroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference interface {
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
	OnDemandAllocationStrategy() *string
	SetOnDemandAllocationStrategy(val *string)
	OnDemandAllocationStrategyInput() *string
	OnDemandBaseCapacity() *float64
	SetOnDemandBaseCapacity(val *float64)
	OnDemandBaseCapacityInput() *float64
	OnDemandPercentageAboveBaseCapacity() *float64
	SetOnDemandPercentageAboveBaseCapacity(val *float64)
	OnDemandPercentageAboveBaseCapacityInput() *float64
	SpotAllocationStrategy() *string
	SetSpotAllocationStrategy(val *string)
	SpotAllocationStrategyInput() *string
	SpotInstancePools() *float64
	SetSpotInstancePools(val *float64)
	SpotInstancePoolsInput() *float64
	SpotMaxPrice() *string
	SetSpotMaxPrice(val *string)
	SpotMaxPriceInput() *string
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
	ResetOnDemandAllocationStrategy()
	ResetOnDemandBaseCapacity()
	ResetOnDemandPercentageAboveBaseCapacity()
	ResetSpotAllocationStrategy()
	ResetSpotInstancePools()
	ResetSpotMaxPrice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference
type jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandPercentageAboveBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandPercentageAboveBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotInstancePools() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotInstancePoolsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotMaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotMaxPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference_Override(a AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetOnDemandAllocationStrategy(val *string) {
	if err := j.validateSetOnDemandAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetOnDemandBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetOnDemandPercentageAboveBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandPercentageAboveBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandPercentageAboveBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetSpotAllocationStrategy(val *string) {
	if err := j.validateSetSpotAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetSpotInstancePools(val *float64) {
	if err := j.validateSetSpotInstancePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstancePools",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetSpotMaxPrice(val *string) {
	if err := j.validateSetSpotMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPrice",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetOnDemandAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetOnDemandBaseCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandBaseCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetOnDemandPercentageAboveBaseCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandPercentageAboveBaseCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetSpotAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetSpotInstancePools() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotInstancePools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetSpotMaxPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotMaxPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutoscalingAutoScalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

