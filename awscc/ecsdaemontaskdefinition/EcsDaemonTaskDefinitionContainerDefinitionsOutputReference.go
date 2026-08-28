// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsdaemontaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsDaemonTaskDefinitionContainerDefinitionsOutputReference interface {
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
	Cpu() *float64
	SetCpu(val *float64)
	CpuInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() EcsDaemonTaskDefinitionContainerDefinitionsDependsOnList
	DependsOnInput() interface{}
	EntryPoint() *[]*string
	SetEntryPoint(val *[]*string)
	EntryPointInput() *[]*string
	Environment() EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList
	EnvironmentFiles() EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList
	EnvironmentFilesInput() interface{}
	EnvironmentInput() interface{}
	Essential() interface{}
	SetEssential(val interface{})
	EssentialInput() interface{}
	FirelensConfiguration() EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	FirelensConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheck() EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	HealthCheckInput() interface{}
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	Interactive() interface{}
	SetInteractive(val interface{})
	InteractiveInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxParameters() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	LinuxParametersInput() interface{}
	LogConfiguration() EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	MountPoints() EcsDaemonTaskDefinitionContainerDefinitionsMountPointsList
	MountPointsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	PseudoTerminal() interface{}
	SetPseudoTerminal(val interface{})
	PseudoTerminalInput() interface{}
	ReadonlyRootFilesystem() interface{}
	SetReadonlyRootFilesystem(val interface{})
	ReadonlyRootFilesystemInput() interface{}
	RepositoryCredentials() EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	RepositoryCredentialsInput() interface{}
	RestartPolicy() EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference
	RestartPolicyInput() interface{}
	Secrets() EcsDaemonTaskDefinitionContainerDefinitionsSecretsList
	SecretsInput() interface{}
	StartTimeout() *float64
	SetStartTimeout(val *float64)
	StartTimeoutInput() *float64
	StopTimeout() *float64
	SetStopTimeout(val *float64)
	StopTimeoutInput() *float64
	SystemControls() EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList
	SystemControlsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ulimits() EcsDaemonTaskDefinitionContainerDefinitionsUlimitsList
	UlimitsInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
	WorkingDirectoryInput() *string
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
	PutDependsOn(value interface{})
	PutEnvironment(value interface{})
	PutEnvironmentFiles(value interface{})
	PutFirelensConfiguration(value *EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfiguration)
	PutHealthCheck(value *EcsDaemonTaskDefinitionContainerDefinitionsHealthCheck)
	PutLinuxParameters(value *EcsDaemonTaskDefinitionContainerDefinitionsLinuxParameters)
	PutLogConfiguration(value *EcsDaemonTaskDefinitionContainerDefinitionsLogConfiguration)
	PutMountPoints(value interface{})
	PutRepositoryCredentials(value *EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentials)
	PutRestartPolicy(value *EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicy)
	PutSecrets(value interface{})
	PutSystemControls(value interface{})
	PutUlimits(value interface{})
	ResetCommand()
	ResetCpu()
	ResetDependsOn()
	ResetEntryPoint()
	ResetEnvironment()
	ResetEnvironmentFiles()
	ResetEssential()
	ResetFirelensConfiguration()
	ResetHealthCheck()
	ResetImage()
	ResetInteractive()
	ResetLinuxParameters()
	ResetLogConfiguration()
	ResetMemory()
	ResetMemoryReservation()
	ResetMountPoints()
	ResetName()
	ResetPrivileged()
	ResetPseudoTerminal()
	ResetReadonlyRootFilesystem()
	ResetRepositoryCredentials()
	ResetRestartPolicy()
	ResetSecrets()
	ResetStartTimeout()
	ResetStopTimeout()
	ResetSystemControls()
	ResetUlimits()
	ResetUser()
	ResetWorkingDirectory()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsDaemonTaskDefinitionContainerDefinitionsOutputReference
type jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) DependsOn() EcsDaemonTaskDefinitionContainerDefinitionsDependsOnList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EntryPointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Environment() EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EnvironmentFiles() EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EnvironmentFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) FirelensConfiguration() EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) FirelensConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firelensConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) HealthCheck() EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Interactive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InteractiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) LinuxParameters() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) LogConfiguration() EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MountPoints() EcsDaemonTaskDefinitionContainerDefinitionsMountPointsList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MountPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PseudoTerminal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PseudoTerminalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) RepositoryCredentials() EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) RestartPolicy() EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) RestartPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Secrets() EcsDaemonTaskDefinitionContainerDefinitionsSecretsList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) StartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) StopTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) SystemControls() EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList
	_jsii_.Get(
		j,
		"systemControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) SystemControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Ulimits() EcsDaemonTaskDefinitionContainerDefinitionsUlimitsList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) UlimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


func NewEcsDaemonTaskDefinitionContainerDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsDaemonTaskDefinitionContainerDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewEcsDaemonTaskDefinitionContainerDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsDaemonTaskDefinitionContainerDefinitionsOutputReference_Override(e EcsDaemonTaskDefinitionContainerDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetEntryPoint(val *[]*string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetInteractive(val interface{}) {
	if err := j.validateSetInteractiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactive",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetPseudoTerminal(val interface{}) {
	if err := j.validateSetPseudoTerminalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudoTerminal",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetStartTimeout(val *float64) {
	if err := j.validateSetStartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeout",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetStopTimeout(val *float64) {
	if err := j.validateSetStopTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopTimeout",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutDependsOn(value interface{}) {
	if err := e.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutEnvironment(value interface{}) {
	if err := e.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutEnvironmentFiles(value interface{}) {
	if err := e.validatePutEnvironmentFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironmentFiles",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutFirelensConfiguration(value *EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfiguration) {
	if err := e.validatePutFirelensConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFirelensConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutHealthCheck(value *EcsDaemonTaskDefinitionContainerDefinitionsHealthCheck) {
	if err := e.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutLinuxParameters(value *EcsDaemonTaskDefinitionContainerDefinitionsLinuxParameters) {
	if err := e.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutLogConfiguration(value *EcsDaemonTaskDefinitionContainerDefinitionsLogConfiguration) {
	if err := e.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutMountPoints(value interface{}) {
	if err := e.validatePutMountPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMountPoints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutRepositoryCredentials(value *EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentials) {
	if err := e.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutRestartPolicy(value *EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicy) {
	if err := e.validatePutRestartPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRestartPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutSecrets(value interface{}) {
	if err := e.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecrets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutSystemControls(value interface{}) {
	if err := e.validatePutSystemControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSystemControls",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PutUlimits(value interface{}) {
	if err := e.validatePutUlimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUlimits",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		e,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		e,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetEnvironmentFiles() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironmentFiles",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		e,
		"resetEssential",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetFirelensConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetFirelensConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		e,
		"resetImage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetInteractive() {
	_jsii_.InvokeVoid(
		e,
		"resetInteractive",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetMountPoints() {
	_jsii_.InvokeVoid(
		e,
		"resetMountPoints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetPseudoTerminal() {
	_jsii_.InvokeVoid(
		e,
		"resetPseudoTerminal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		e,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		e,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetSecrets() {
	_jsii_.InvokeVoid(
		e,
		"resetSecrets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetStartTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetStartTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetStopTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetStopTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetSystemControls() {
	_jsii_.InvokeVoid(
		e,
		"resetSystemControls",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetUlimits() {
	_jsii_.InvokeVoid(
		e,
		"resetUlimits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		e,
		"resetUser",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		e,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

