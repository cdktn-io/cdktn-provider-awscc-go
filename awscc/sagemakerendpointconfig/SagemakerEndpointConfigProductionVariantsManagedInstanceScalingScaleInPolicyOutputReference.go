// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerendpointconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference interface {
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
	CooldownInMinutes() *float64
	SetCooldownInMinutes(val *float64)
	CooldownInMinutesInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumStepSize() *float64
	SetMaximumStepSize(val *float64)
	MaximumStepSizeInput() *float64
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
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
	ResetCooldownInMinutes()
	ResetMaximumStepSize()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference
type jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) CooldownInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) CooldownInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) MaximumStepSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumStepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) MaximumStepSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumStepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference_Override(s SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetCooldownInMinutes(val *float64) {
	if err := j.validateSetCooldownInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownInMinutes",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetMaximumStepSize(val *float64) {
	if err := j.validateSetMaximumStepSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumStepSize",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ResetCooldownInMinutes() {
	_jsii_.InvokeVoid(
		s,
		"resetCooldownInMinutes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ResetMaximumStepSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMaximumStepSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigProductionVariantsManagedInstanceScalingScaleInPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

