// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerClusterInstanceGroupsOutputReference interface {
	cdktn.ComplexObject
	AutoPatchConfig() SagemakerClusterInstanceGroupsAutoPatchConfigOutputReference
	AutoPatchConfigInput() interface{}
	CapacityRequirements() SagemakerClusterInstanceGroupsCapacityRequirementsOutputReference
	CapacityRequirementsInput() interface{}
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
	ExecutionRole() *string
	SetExecutionRole(val *string)
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceGroupName() *string
	SetInstanceGroupName(val *string)
	InstanceGroupNameInput() *string
	InstanceRequirements() SagemakerClusterInstanceGroupsInstanceRequirementsOutputReference
	InstanceRequirementsInput() interface{}
	InstanceStorageConfigs() SagemakerClusterInstanceGroupsInstanceStorageConfigsList
	InstanceStorageConfigsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KubernetesConfig() SagemakerClusterInstanceGroupsKubernetesConfigOutputReference
	KubernetesConfigInput() interface{}
	LifeCycleConfig() SagemakerClusterInstanceGroupsLifeCycleConfigOutputReference
	LifeCycleConfigInput() interface{}
	MinInstanceCount() *float64
	SetMinInstanceCount(val *float64)
	MinInstanceCountInput() *float64
	NetworkInterface() SagemakerClusterInstanceGroupsNetworkInterfaceOutputReference
	NetworkInterfaceInput() interface{}
	OnStartDeepHealthChecks() *[]*string
	SetOnStartDeepHealthChecks(val *[]*string)
	OnStartDeepHealthChecksInput() *[]*string
	OverrideVpcConfig() SagemakerClusterInstanceGroupsOverrideVpcConfigOutputReference
	OverrideVpcConfigInput() interface{}
	ScheduledUpdateConfig() SagemakerClusterInstanceGroupsScheduledUpdateConfigOutputReference
	ScheduledUpdateConfigInput() interface{}
	SlurmConfig() SagemakerClusterInstanceGroupsSlurmConfigOutputReference
	SlurmConfigInput() interface{}
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
	PutAutoPatchConfig(value *SagemakerClusterInstanceGroupsAutoPatchConfig)
	PutCapacityRequirements(value *SagemakerClusterInstanceGroupsCapacityRequirements)
	PutInstanceRequirements(value *SagemakerClusterInstanceGroupsInstanceRequirements)
	PutInstanceStorageConfigs(value interface{})
	PutKubernetesConfig(value *SagemakerClusterInstanceGroupsKubernetesConfig)
	PutLifeCycleConfig(value *SagemakerClusterInstanceGroupsLifeCycleConfig)
	PutNetworkInterface(value *SagemakerClusterInstanceGroupsNetworkInterface)
	PutOverrideVpcConfig(value *SagemakerClusterInstanceGroupsOverrideVpcConfig)
	PutScheduledUpdateConfig(value *SagemakerClusterInstanceGroupsScheduledUpdateConfig)
	PutSlurmConfig(value *SagemakerClusterInstanceGroupsSlurmConfig)
	ResetAutoPatchConfig()
	ResetCapacityRequirements()
	ResetCurrentCount()
	ResetExecutionRole()
	ResetImageId()
	ResetInstanceCount()
	ResetInstanceGroupName()
	ResetInstanceRequirements()
	ResetInstanceStorageConfigs()
	ResetInstanceType()
	ResetKubernetesConfig()
	ResetLifeCycleConfig()
	ResetMinInstanceCount()
	ResetNetworkInterface()
	ResetOnStartDeepHealthChecks()
	ResetOverrideVpcConfig()
	ResetScheduledUpdateConfig()
	ResetSlurmConfig()
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

// The jsii proxy struct for SagemakerClusterInstanceGroupsOutputReference
type jsiiProxy_SagemakerClusterInstanceGroupsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) AutoPatchConfig() SagemakerClusterInstanceGroupsAutoPatchConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsAutoPatchConfigOutputReference
	_jsii_.Get(
		j,
		"autoPatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) AutoPatchConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CapacityRequirements() SagemakerClusterInstanceGroupsCapacityRequirementsOutputReference {
	var returns SagemakerClusterInstanceGroupsCapacityRequirementsOutputReference
	_jsii_.Get(
		j,
		"capacityRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CapacityRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CurrentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) CurrentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceRequirements() SagemakerClusterInstanceGroupsInstanceRequirementsOutputReference {
	var returns SagemakerClusterInstanceGroupsInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceStorageConfigs() SagemakerClusterInstanceGroupsInstanceStorageConfigsList {
	var returns SagemakerClusterInstanceGroupsInstanceStorageConfigsList
	_jsii_.Get(
		j,
		"instanceStorageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceStorageConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceStorageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) KubernetesConfig() SagemakerClusterInstanceGroupsKubernetesConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsKubernetesConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) KubernetesConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) LifeCycleConfig() SagemakerClusterInstanceGroupsLifeCycleConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsLifeCycleConfigOutputReference
	_jsii_.Get(
		j,
		"lifeCycleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) LifeCycleConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifeCycleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) MinInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) MinInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) NetworkInterface() SagemakerClusterInstanceGroupsNetworkInterfaceOutputReference {
	var returns SagemakerClusterInstanceGroupsNetworkInterfaceOutputReference
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OnStartDeepHealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartDeepHealthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OnStartDeepHealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"onStartDeepHealthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OverrideVpcConfig() SagemakerClusterInstanceGroupsOverrideVpcConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsOverrideVpcConfigOutputReference
	_jsii_.Get(
		j,
		"overrideVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) OverrideVpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideVpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ScheduledUpdateConfig() SagemakerClusterInstanceGroupsScheduledUpdateConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsScheduledUpdateConfigOutputReference
	_jsii_.Get(
		j,
		"scheduledUpdateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ScheduledUpdateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledUpdateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) SlurmConfig() SagemakerClusterInstanceGroupsSlurmConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsSlurmConfigOutputReference
	_jsii_.Get(
		j,
		"slurmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) SlurmConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slurmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) TrainingPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) TrainingPlanArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingPlanArnInput",
		&returns,
	)
	return returns
}


func NewSagemakerClusterInstanceGroupsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerClusterInstanceGroupsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerClusterInstanceGroupsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerClusterInstanceGroupsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterInstanceGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerClusterInstanceGroupsOutputReference_Override(s SagemakerClusterInstanceGroupsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterInstanceGroupsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetCurrentCount(val *float64) {
	if err := j.validateSetCurrentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"currentCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInstanceGroupName(val *string) {
	if err := j.validateSetInstanceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGroupName",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetMinInstanceCount(val *float64) {
	if err := j.validateSetMinInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstanceCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetOnStartDeepHealthChecks(val *[]*string) {
	if err := j.validateSetOnStartDeepHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onStartDeepHealthChecks",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference)SetTrainingPlanArn(val *string) {
	if err := j.validateSetTrainingPlanArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingPlanArn",
		val,
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutAutoPatchConfig(value *SagemakerClusterInstanceGroupsAutoPatchConfig) {
	if err := s.validatePutAutoPatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoPatchConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutCapacityRequirements(value *SagemakerClusterInstanceGroupsCapacityRequirements) {
	if err := s.validatePutCapacityRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCapacityRequirements",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutInstanceRequirements(value *SagemakerClusterInstanceGroupsInstanceRequirements) {
	if err := s.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutInstanceStorageConfigs(value interface{}) {
	if err := s.validatePutInstanceStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstanceStorageConfigs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutKubernetesConfig(value *SagemakerClusterInstanceGroupsKubernetesConfig) {
	if err := s.validatePutKubernetesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKubernetesConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutLifeCycleConfig(value *SagemakerClusterInstanceGroupsLifeCycleConfig) {
	if err := s.validatePutLifeCycleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLifeCycleConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutNetworkInterface(value *SagemakerClusterInstanceGroupsNetworkInterface) {
	if err := s.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutOverrideVpcConfig(value *SagemakerClusterInstanceGroupsOverrideVpcConfig) {
	if err := s.validatePutOverrideVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOverrideVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutScheduledUpdateConfig(value *SagemakerClusterInstanceGroupsScheduledUpdateConfig) {
	if err := s.validatePutScheduledUpdateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putScheduledUpdateConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) PutSlurmConfig(value *SagemakerClusterInstanceGroupsSlurmConfig) {
	if err := s.validatePutSlurmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSlurmConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetAutoPatchConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoPatchConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetCapacityRequirements() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacityRequirements",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetCurrentCount() {
	_jsii_.InvokeVoid(
		s,
		"resetCurrentCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		s,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetImageId() {
	_jsii_.InvokeVoid(
		s,
		"resetImageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetInstanceGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetInstanceStorageConfigs() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceStorageConfigs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetKubernetesConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetKubernetesConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetLifeCycleConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetLifeCycleConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetMinInstanceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMinInstanceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetOnStartDeepHealthChecks() {
	_jsii_.InvokeVoid(
		s,
		"resetOnStartDeepHealthChecks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetOverrideVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetScheduledUpdateConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduledUpdateConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetSlurmConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSlurmConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		s,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ResetTrainingPlanArn() {
	_jsii_.InvokeVoid(
		s,
		"resetTrainingPlanArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

