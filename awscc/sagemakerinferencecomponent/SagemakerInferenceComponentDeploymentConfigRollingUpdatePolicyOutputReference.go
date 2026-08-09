// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerinferencecomponent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference interface {
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
	MaximumBatchSize() SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference
	MaximumBatchSizeInput() interface{}
	MaximumExecutionTimeoutInSeconds() *float64
	SetMaximumExecutionTimeoutInSeconds(val *float64)
	MaximumExecutionTimeoutInSecondsInput() *float64
	RollbackMaximumBatchSize() SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSizeOutputReference
	RollbackMaximumBatchSizeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WaitIntervalInSeconds() *float64
	SetWaitIntervalInSeconds(val *float64)
	WaitIntervalInSecondsInput() *float64
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
	PutMaximumBatchSize(value *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSize)
	PutRollbackMaximumBatchSize(value *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSize)
	ResetMaximumBatchSize()
	ResetMaximumExecutionTimeoutInSeconds()
	ResetRollbackMaximumBatchSize()
	ResetWaitIntervalInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference
type jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) MaximumBatchSize() SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference {
	var returns SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSizeOutputReference
	_jsii_.Get(
		j,
		"maximumBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) MaximumBatchSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) MaximumExecutionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) MaximumExecutionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) RollbackMaximumBatchSize() SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSizeOutputReference {
	var returns SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSizeOutputReference
	_jsii_.Get(
		j,
		"rollbackMaximumBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) RollbackMaximumBatchSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackMaximumBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) WaitIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) WaitIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSecondsInput",
		&returns,
	)
	return returns
}


func NewSagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerInferenceComponent.SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference_Override(s SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerInferenceComponent.SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetMaximumExecutionTimeoutInSeconds(val *float64) {
	if err := j.validateSetMaximumExecutionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference)SetWaitIntervalInSeconds(val *float64) {
	if err := j.validateSetWaitIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitIntervalInSeconds",
		val,
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) PutMaximumBatchSize(value *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSize) {
	if err := s.validatePutMaximumBatchSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMaximumBatchSize",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) PutRollbackMaximumBatchSize(value *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSize) {
	if err := s.validatePutRollbackMaximumBatchSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRollbackMaximumBatchSize",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ResetMaximumBatchSize() {
	_jsii_.InvokeVoid(
		s,
		"resetMaximumBatchSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ResetMaximumExecutionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMaximumExecutionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ResetRollbackMaximumBatchSize() {
	_jsii_.InvokeVoid(
		s,
		"resetRollbackMaximumBatchSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ResetWaitIntervalInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitIntervalInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

