// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobDefinitionContainerPropertiesOutputReference interface {
	cdktn.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	EnableExecuteCommandInput() interface{}
	Environment() BatchJobDefinitionContainerPropertiesEnvironmentList
	EnvironmentInput() interface{}
	EphemeralStorage() BatchJobDefinitionContainerPropertiesEphemeralStorageOutputReference
	EphemeralStorageInput() interface{}
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	FargatePlatformConfiguration() BatchJobDefinitionContainerPropertiesFargatePlatformConfigurationOutputReference
	FargatePlatformConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobRoleArn() *string
	SetJobRoleArn(val *string)
	JobRoleArnInput() *string
	LinuxParameters() BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference
	LinuxParametersInput() interface{}
	LogConfiguration() BatchJobDefinitionContainerPropertiesLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MountPoints() BatchJobDefinitionContainerPropertiesMountPointsList
	MountPointsInput() interface{}
	NetworkConfiguration() BatchJobDefinitionContainerPropertiesNetworkConfigurationOutputReference
	NetworkConfigurationInput() interface{}
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	ReadonlyRootFilesystem() interface{}
	SetReadonlyRootFilesystem(val interface{})
	ReadonlyRootFilesystemInput() interface{}
	RepositoryCredentials() BatchJobDefinitionContainerPropertiesRepositoryCredentialsOutputReference
	RepositoryCredentialsInput() interface{}
	ResourceRequirements() BatchJobDefinitionContainerPropertiesResourceRequirementsList
	ResourceRequirementsInput() interface{}
	RuntimePlatform() BatchJobDefinitionContainerPropertiesRuntimePlatformOutputReference
	RuntimePlatformInput() interface{}
	Secrets() BatchJobDefinitionContainerPropertiesSecretsList
	SecretsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ulimits() BatchJobDefinitionContainerPropertiesUlimitsList
	UlimitsInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
	Vcpus() *float64
	SetVcpus(val *float64)
	VcpusInput() *float64
	Volumes() BatchJobDefinitionContainerPropertiesVolumesList
	VolumesInput() interface{}
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
	PutEnvironment(value interface{})
	PutEphemeralStorage(value *BatchJobDefinitionContainerPropertiesEphemeralStorage)
	PutFargatePlatformConfiguration(value *BatchJobDefinitionContainerPropertiesFargatePlatformConfiguration)
	PutLinuxParameters(value *BatchJobDefinitionContainerPropertiesLinuxParameters)
	PutLogConfiguration(value *BatchJobDefinitionContainerPropertiesLogConfiguration)
	PutMountPoints(value interface{})
	PutNetworkConfiguration(value *BatchJobDefinitionContainerPropertiesNetworkConfiguration)
	PutRepositoryCredentials(value *BatchJobDefinitionContainerPropertiesRepositoryCredentials)
	PutResourceRequirements(value interface{})
	PutRuntimePlatform(value *BatchJobDefinitionContainerPropertiesRuntimePlatform)
	PutSecrets(value interface{})
	PutUlimits(value interface{})
	PutVolumes(value interface{})
	ResetCommand()
	ResetEnableExecuteCommand()
	ResetEnvironment()
	ResetEphemeralStorage()
	ResetExecutionRoleArn()
	ResetFargatePlatformConfiguration()
	ResetImage()
	ResetJobRoleArn()
	ResetLinuxParameters()
	ResetLogConfiguration()
	ResetMemory()
	ResetMountPoints()
	ResetNetworkConfiguration()
	ResetPrivileged()
	ResetReadonlyRootFilesystem()
	ResetRepositoryCredentials()
	ResetResourceRequirements()
	ResetRuntimePlatform()
	ResetSecrets()
	ResetUlimits()
	ResetUser()
	ResetVcpus()
	ResetVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionContainerPropertiesOutputReference
type jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Environment() BatchJobDefinitionContainerPropertiesEnvironmentList {
	var returns BatchJobDefinitionContainerPropertiesEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) EphemeralStorage() BatchJobDefinitionContainerPropertiesEphemeralStorageOutputReference {
	var returns BatchJobDefinitionContainerPropertiesEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) EphemeralStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) FargatePlatformConfiguration() BatchJobDefinitionContainerPropertiesFargatePlatformConfigurationOutputReference {
	var returns BatchJobDefinitionContainerPropertiesFargatePlatformConfigurationOutputReference
	_jsii_.Get(
		j,
		"fargatePlatformConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) FargatePlatformConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fargatePlatformConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) JobRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) JobRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) LinuxParameters() BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference {
	var returns BatchJobDefinitionContainerPropertiesLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) LogConfiguration() BatchJobDefinitionContainerPropertiesLogConfigurationOutputReference {
	var returns BatchJobDefinitionContainerPropertiesLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) MountPoints() BatchJobDefinitionContainerPropertiesMountPointsList {
	var returns BatchJobDefinitionContainerPropertiesMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) MountPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) NetworkConfiguration() BatchJobDefinitionContainerPropertiesNetworkConfigurationOutputReference {
	var returns BatchJobDefinitionContainerPropertiesNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) RepositoryCredentials() BatchJobDefinitionContainerPropertiesRepositoryCredentialsOutputReference {
	var returns BatchJobDefinitionContainerPropertiesRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResourceRequirements() BatchJobDefinitionContainerPropertiesResourceRequirementsList {
	var returns BatchJobDefinitionContainerPropertiesResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResourceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) RuntimePlatform() BatchJobDefinitionContainerPropertiesRuntimePlatformOutputReference {
	var returns BatchJobDefinitionContainerPropertiesRuntimePlatformOutputReference
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) RuntimePlatformInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Secrets() BatchJobDefinitionContainerPropertiesSecretsList {
	var returns BatchJobDefinitionContainerPropertiesSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Ulimits() BatchJobDefinitionContainerPropertiesUlimitsList {
	var returns BatchJobDefinitionContainerPropertiesUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) UlimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Vcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) VcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Volumes() BatchJobDefinitionContainerPropertiesVolumesList {
	var returns BatchJobDefinitionContainerPropertiesVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionContainerPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionContainerPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionContainerPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionContainerPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionContainerPropertiesOutputReference_Override(b BatchJobDefinitionContainerPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionContainerPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetJobRoleArn(val *string) {
	if err := j.validateSetJobRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobRoleArn",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference)SetVcpus(val *float64) {
	if err := j.validateSetVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcpus",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutEnvironment(value interface{}) {
	if err := b.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutEphemeralStorage(value *BatchJobDefinitionContainerPropertiesEphemeralStorage) {
	if err := b.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutFargatePlatformConfiguration(value *BatchJobDefinitionContainerPropertiesFargatePlatformConfiguration) {
	if err := b.validatePutFargatePlatformConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFargatePlatformConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutLinuxParameters(value *BatchJobDefinitionContainerPropertiesLinuxParameters) {
	if err := b.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutLogConfiguration(value *BatchJobDefinitionContainerPropertiesLogConfiguration) {
	if err := b.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutMountPoints(value interface{}) {
	if err := b.validatePutMountPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMountPoints",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutNetworkConfiguration(value *BatchJobDefinitionContainerPropertiesNetworkConfiguration) {
	if err := b.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutRepositoryCredentials(value *BatchJobDefinitionContainerPropertiesRepositoryCredentials) {
	if err := b.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutResourceRequirements(value interface{}) {
	if err := b.validatePutResourceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResourceRequirements",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutRuntimePlatform(value *BatchJobDefinitionContainerPropertiesRuntimePlatform) {
	if err := b.validatePutRuntimePlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRuntimePlatform",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutSecrets(value interface{}) {
	if err := b.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSecrets",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutUlimits(value interface{}) {
	if err := b.validatePutUlimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUlimits",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) PutVolumes(value interface{}) {
	if err := b.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVolumes",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		b,
		"resetCommand",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		b,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		b,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		b,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetFargatePlatformConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetFargatePlatformConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		b,
		"resetImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetJobRoleArn() {
	_jsii_.InvokeVoid(
		b,
		"resetJobRoleArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		b,
		"resetMemory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetMountPoints() {
	_jsii_.InvokeVoid(
		b,
		"resetMountPoints",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		b,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		b,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetResourceRequirements() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceRequirements",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetRuntimePlatform() {
	_jsii_.InvokeVoid(
		b,
		"resetRuntimePlatform",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetSecrets() {
	_jsii_.InvokeVoid(
		b,
		"resetSecrets",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetUlimits() {
	_jsii_.InvokeVoid(
		b,
		"resetUlimits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		b,
		"resetUser",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetVcpus() {
	_jsii_.InvokeVoid(
		b,
		"resetVcpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionContainerPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

