// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerClusterRestrictedInstanceGroupsOutputReference interface {
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
	CurrentCount() *float64
	SetCurrentCount(val *float64)
	CurrentCountInput() *float64
	EnvironmentConfig() SagemakerClusterRestrictedInstanceGroupsEnvironmentConfigOutputReference
	EnvironmentConfigInput() interface{}
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceGroupName() *string
	SetInstanceGroupName(val *string)
	InstanceGroupNameInput() *string
	InstanceStorageConfigs() SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsList
	InstanceStorageConfigsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OnStartDeepHealthChecks() *[]*string
	SetOnStartDeepHealthChecks(val *[]*string)
	OnStartDeepHealthChecksInput() *[]*string
	OverrideVpcConfig() SagemakerClusterRestrictedInstanceGroupsOverrideVpcConfigOutputReference
	OverrideVpcConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThreadsPerCore() *float64
	SetThreadsPerCore(val *float64)
	ThreadsPerCoreInput() *float64
	TrainingPlanArn() *string
	SetTrainingPlanArn(val *string)
	TrainingPlanArnInput() *string
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
	PutEnvironmentConfig(value *SagemakerClusterRestrictedInstanceGroupsEnvironmentConfig)
	PutInstanceStorageConfigs(value interface{})
	PutOverrideVpcConfig(value *SagemakerClusterRestrictedInstanceGroupsOverrideVpcConfig)
	ResetCurrentCount()
	ResetEnvironmentConfig()
	ResetExecutionRole()
	ResetInstanceCount()
	ResetInstanceGroupName()
	ResetInstanceStorageConfigs()
	ResetInstanceType()
	ResetOnStartDeepHealthChecks()
	ResetOverrideVpcConfig()
	ResetThreadsPerCore()
	ResetTrainingPlanArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerClusterRestrictedInstanceGroupsOutputReference
type jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) CurrentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) CurrentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) EnvironmentConfig() SagemakerClusterRestrictedInstanceGroupsEnvironmentConfigOutputReference {
	var returns SagemakerClusterRestrictedInstanceGroupsEnvironmentConfigOutputReference
	_jsii_.Get(
		j,
		"environmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) EnvironmentConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceStorageConfigs() SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsList {
	var returns SagemakerClusterRestrictedInstanceGroupsInstanceStorageConfigsList
	_jsii_.Get(
		j,
		"instanceStorageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceStorageConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceStorageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) OnStartDeepHealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartDeepHealthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) OnStartDeepHealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartDeepHealthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) OverrideVpcConfig() SagemakerClusterRestrictedInstanceGroupsOverrideVpcConfigOutputReference {
	var returns SagemakerClusterRestrictedInstanceGroupsOverrideVpcConfigOutputReference
	_jsii_.Get(
		j,
		"overrideVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) OverrideVpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideVpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) TrainingPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) TrainingPlanArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingPlanArnInput",
		&returns,
	)
	return returns
}


func NewSagemakerClusterRestrictedInstanceGroupsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerClusterRestrictedInstanceGroupsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerClusterRestrictedInstanceGroupsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterRestrictedInstanceGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerClusterRestrictedInstanceGroupsOutputReference_Override(s SagemakerClusterRestrictedInstanceGroupsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterRestrictedInstanceGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetCurrentCount(val *float64) {
	if err := j.validateSetCurrentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetInstanceGroupName(val *string) {
	if err := j.validateSetInstanceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGroupName",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetOnStartDeepHealthChecks(val *[]*string) {
	if err := j.validateSetOnStartDeepHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStartDeepHealthChecks",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference)SetTrainingPlanArn(val *string) {
	if err := j.validateSetTrainingPlanArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingPlanArn",
		val,
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) PutEnvironmentConfig(value *SagemakerClusterRestrictedInstanceGroupsEnvironmentConfig) {
	if err := s.validatePutEnvironmentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEnvironmentConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) PutInstanceStorageConfigs(value interface{}) {
	if err := s.validatePutInstanceStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstanceStorageConfigs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) PutOverrideVpcConfig(value *SagemakerClusterRestrictedInstanceGroupsOverrideVpcConfig) {
	if err := s.validatePutOverrideVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOverrideVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetCurrentCount() {
	_jsii_.InvokeVoid(
		s,
		"resetCurrentCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetEnvironmentConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironmentConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		s,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetInstanceGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetInstanceStorageConfigs() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceStorageConfigs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetOnStartDeepHealthChecks() {
	_jsii_.InvokeVoid(
		s,
		"resetOnStartDeepHealthChecks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetOverrideVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		s,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ResetTrainingPlanArn() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingPlanArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

