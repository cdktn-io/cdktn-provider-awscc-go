// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecstaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsTaskDefinitionContainerDefinitionsOutputReference interface {
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
	CredentialSpecs() *[]*string
	SetCredentialSpecs(val *[]*string)
	CredentialSpecsInput() *[]*string
	DependsOn() EcsTaskDefinitionContainerDefinitionsDependsOnList
	DependsOnInput() interface{}
	DisableNetworking() interface{}
	SetDisableNetworking(val interface{})
	DisableNetworkingInput() interface{}
	DnsSearchDomains() *[]*string
	SetDnsSearchDomains(val *[]*string)
	DnsSearchDomainsInput() *[]*string
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
	DockerLabels() *map[string]*string
	SetDockerLabels(val *map[string]*string)
	DockerLabelsInput() *map[string]*string
	DockerSecurityOptions() *[]*string
	SetDockerSecurityOptions(val *[]*string)
	DockerSecurityOptionsInput() *[]*string
	EntryPoint() *[]*string
	SetEntryPoint(val *[]*string)
	EntryPointInput() *[]*string
	Environment() EcsTaskDefinitionContainerDefinitionsEnvironmentList
	EnvironmentFiles() EcsTaskDefinitionContainerDefinitionsEnvironmentFilesList
	EnvironmentFilesInput() interface{}
	EnvironmentInput() interface{}
	Essential() interface{}
	SetEssential(val interface{})
	EssentialInput() interface{}
	ExtraHosts() EcsTaskDefinitionContainerDefinitionsExtraHostsList
	ExtraHostsInput() interface{}
	FirelensConfiguration() EcsTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	FirelensConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheck() EcsTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	HealthCheckInput() interface{}
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	Interactive() interface{}
	SetInteractive(val interface{})
	InteractiveInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Links() *[]*string
	SetLinks(val *[]*string)
	LinksInput() *[]*string
	LinuxParameters() EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	LinuxParametersInput() interface{}
	LogConfiguration() EcsTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	LogConfigurationInput() interface{}
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MemoryReservation() *float64
	SetMemoryReservation(val *float64)
	MemoryReservationInput() *float64
	MountPoints() EcsTaskDefinitionContainerDefinitionsMountPointsList
	MountPointsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PortMappings() EcsTaskDefinitionContainerDefinitionsPortMappingsList
	PortMappingsInput() interface{}
	Privileged() interface{}
	SetPrivileged(val interface{})
	PrivilegedInput() interface{}
	PseudoTerminal() interface{}
	SetPseudoTerminal(val interface{})
	PseudoTerminalInput() interface{}
	ReadonlyRootFilesystem() interface{}
	SetReadonlyRootFilesystem(val interface{})
	ReadonlyRootFilesystemInput() interface{}
	RepositoryCredentials() EcsTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	RepositoryCredentialsInput() interface{}
	ResourceRequirements() EcsTaskDefinitionContainerDefinitionsResourceRequirementsList
	ResourceRequirementsInput() interface{}
	RestartPolicy() EcsTaskDefinitionContainerDefinitionsRestartPolicyOutputReference
	RestartPolicyInput() interface{}
	Secrets() EcsTaskDefinitionContainerDefinitionsSecretsList
	SecretsInput() interface{}
	StartTimeout() *float64
	SetStartTimeout(val *float64)
	StartTimeoutInput() *float64
	StopTimeout() *float64
	SetStopTimeout(val *float64)
	StopTimeoutInput() *float64
	SystemControls() EcsTaskDefinitionContainerDefinitionsSystemControlsList
	SystemControlsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ulimits() EcsTaskDefinitionContainerDefinitionsUlimitsList
	UlimitsInput() interface{}
	User() *string
	SetUser(val *string)
	UserInput() *string
	VersionConsistency() *string
	SetVersionConsistency(val *string)
	VersionConsistencyInput() *string
	VolumesFrom() EcsTaskDefinitionContainerDefinitionsVolumesFromList
	VolumesFromInput() interface{}
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
	PutExtraHosts(value interface{})
	PutFirelensConfiguration(value *EcsTaskDefinitionContainerDefinitionsFirelensConfiguration)
	PutHealthCheck(value *EcsTaskDefinitionContainerDefinitionsHealthCheck)
	PutLinuxParameters(value *EcsTaskDefinitionContainerDefinitionsLinuxParameters)
	PutLogConfiguration(value *EcsTaskDefinitionContainerDefinitionsLogConfiguration)
	PutMountPoints(value interface{})
	PutPortMappings(value interface{})
	PutRepositoryCredentials(value *EcsTaskDefinitionContainerDefinitionsRepositoryCredentials)
	PutResourceRequirements(value interface{})
	PutRestartPolicy(value *EcsTaskDefinitionContainerDefinitionsRestartPolicy)
	PutSecrets(value interface{})
	PutSystemControls(value interface{})
	PutUlimits(value interface{})
	PutVolumesFrom(value interface{})
	ResetCommand()
	ResetCpu()
	ResetCredentialSpecs()
	ResetDependsOn()
	ResetDisableNetworking()
	ResetDnsSearchDomains()
	ResetDnsServers()
	ResetDockerLabels()
	ResetDockerSecurityOptions()
	ResetEntryPoint()
	ResetEnvironment()
	ResetEnvironmentFiles()
	ResetEssential()
	ResetExtraHosts()
	ResetFirelensConfiguration()
	ResetHealthCheck()
	ResetHostname()
	ResetImage()
	ResetInteractive()
	ResetLinks()
	ResetLinuxParameters()
	ResetLogConfiguration()
	ResetMemory()
	ResetMemoryReservation()
	ResetMountPoints()
	ResetName()
	ResetPortMappings()
	ResetPrivileged()
	ResetPseudoTerminal()
	ResetReadonlyRootFilesystem()
	ResetRepositoryCredentials()
	ResetResourceRequirements()
	ResetRestartPolicy()
	ResetSecrets()
	ResetStartTimeout()
	ResetStopTimeout()
	ResetSystemControls()
	ResetUlimits()
	ResetUser()
	ResetVersionConsistency()
	ResetVolumesFrom()
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

// The jsii proxy struct for EcsTaskDefinitionContainerDefinitionsOutputReference
type jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) CredentialSpecs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"credentialSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) CredentialSpecsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"credentialSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DependsOn() EcsTaskDefinitionContainerDefinitionsDependsOnList {
	var returns EcsTaskDefinitionContainerDefinitionsDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DisableNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DisableNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DnsSearchDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSearchDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DnsSearchDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsSearchDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DockerLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dockerLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DockerLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dockerLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DockerSecurityOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dockerSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) DockerSecurityOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dockerSecurityOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) EntryPointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Environment() EcsTaskDefinitionContainerDefinitionsEnvironmentList {
	var returns EcsTaskDefinitionContainerDefinitionsEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) EnvironmentFiles() EcsTaskDefinitionContainerDefinitionsEnvironmentFilesList {
	var returns EcsTaskDefinitionContainerDefinitionsEnvironmentFilesList
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) EnvironmentFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ExtraHosts() EcsTaskDefinitionContainerDefinitionsExtraHostsList {
	var returns EcsTaskDefinitionContainerDefinitionsExtraHostsList
	_jsii_.Get(
		j,
		"extraHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ExtraHostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extraHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) FirelensConfiguration() EcsTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) FirelensConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firelensConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) HealthCheck() EcsTaskDefinitionContainerDefinitionsHealthCheckOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Interactive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) InteractiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Links() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) LinksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) LinuxParameters() EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) LinuxParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) LogConfiguration() EcsTaskDefinitionContainerDefinitionsLogConfigurationOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) LogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) MemoryReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) MountPoints() EcsTaskDefinitionContainerDefinitionsMountPointsList {
	var returns EcsTaskDefinitionContainerDefinitionsMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) MountPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PortMappings() EcsTaskDefinitionContainerDefinitionsPortMappingsList {
	var returns EcsTaskDefinitionContainerDefinitionsPortMappingsList
	_jsii_.Get(
		j,
		"portMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PortMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PseudoTerminal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PseudoTerminalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pseudoTerminalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ReadonlyRootFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ReadonlyRootFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readonlyRootFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) RepositoryCredentials() EcsTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) RepositoryCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResourceRequirements() EcsTaskDefinitionContainerDefinitionsResourceRequirementsList {
	var returns EcsTaskDefinitionContainerDefinitionsResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResourceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) RestartPolicy() EcsTaskDefinitionContainerDefinitionsRestartPolicyOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsRestartPolicyOutputReference
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) RestartPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Secrets() EcsTaskDefinitionContainerDefinitionsSecretsList {
	var returns EcsTaskDefinitionContainerDefinitionsSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) SecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) StartTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) StopTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) SystemControls() EcsTaskDefinitionContainerDefinitionsSystemControlsList {
	var returns EcsTaskDefinitionContainerDefinitionsSystemControlsList
	_jsii_.Get(
		j,
		"systemControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) SystemControlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Ulimits() EcsTaskDefinitionContainerDefinitionsUlimitsList {
	var returns EcsTaskDefinitionContainerDefinitionsUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) UlimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) VersionConsistency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionConsistency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) VersionConsistencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionConsistencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) VolumesFrom() EcsTaskDefinitionContainerDefinitionsVolumesFromList {
	var returns EcsTaskDefinitionContainerDefinitionsVolumesFromList
	_jsii_.Get(
		j,
		"volumesFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) VolumesFromInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) WorkingDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectoryInput",
		&returns,
	)
	return returns
}


func NewEcsTaskDefinitionContainerDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsTaskDefinitionContainerDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionContainerDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsTaskDefinition.EcsTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsTaskDefinitionContainerDefinitionsOutputReference_Override(e EcsTaskDefinitionContainerDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsTaskDefinition.EcsTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetCredentialSpecs(val *[]*string) {
	if err := j.validateSetCredentialSpecsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialSpecs",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetDisableNetworking(val interface{}) {
	if err := j.validateSetDisableNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableNetworking",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetDnsSearchDomains(val *[]*string) {
	if err := j.validateSetDnsSearchDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSearchDomains",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetDockerLabels(val *map[string]*string) {
	if err := j.validateSetDockerLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerLabels",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetDockerSecurityOptions(val *[]*string) {
	if err := j.validateSetDockerSecurityOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerSecurityOptions",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetEntryPoint(val *[]*string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetInteractive(val interface{}) {
	if err := j.validateSetInteractiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactive",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetLinks(val *[]*string) {
	if err := j.validateSetLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"links",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetMemoryReservation(val *float64) {
	if err := j.validateSetMemoryReservationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryReservation",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetPseudoTerminal(val interface{}) {
	if err := j.validateSetPseudoTerminalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pseudoTerminal",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetReadonlyRootFilesystem(val interface{}) {
	if err := j.validateSetReadonlyRootFilesystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonlyRootFilesystem",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetStartTimeout(val *float64) {
	if err := j.validateSetStartTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeout",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetStopTimeout(val *float64) {
	if err := j.validateSetStopTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopTimeout",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetVersionConsistency(val *string) {
	if err := j.validateSetVersionConsistencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionConsistency",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference)SetWorkingDirectory(val *string) {
	if err := j.validateSetWorkingDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutDependsOn(value interface{}) {
	if err := e.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutEnvironment(value interface{}) {
	if err := e.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutEnvironmentFiles(value interface{}) {
	if err := e.validatePutEnvironmentFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnvironmentFiles",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutExtraHosts(value interface{}) {
	if err := e.validatePutExtraHostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putExtraHosts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutFirelensConfiguration(value *EcsTaskDefinitionContainerDefinitionsFirelensConfiguration) {
	if err := e.validatePutFirelensConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFirelensConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutHealthCheck(value *EcsTaskDefinitionContainerDefinitionsHealthCheck) {
	if err := e.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutLinuxParameters(value *EcsTaskDefinitionContainerDefinitionsLinuxParameters) {
	if err := e.validatePutLinuxParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLinuxParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutLogConfiguration(value *EcsTaskDefinitionContainerDefinitionsLogConfiguration) {
	if err := e.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutMountPoints(value interface{}) {
	if err := e.validatePutMountPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMountPoints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutPortMappings(value interface{}) {
	if err := e.validatePutPortMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPortMappings",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutRepositoryCredentials(value *EcsTaskDefinitionContainerDefinitionsRepositoryCredentials) {
	if err := e.validatePutRepositoryCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRepositoryCredentials",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutResourceRequirements(value interface{}) {
	if err := e.validatePutResourceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourceRequirements",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutRestartPolicy(value *EcsTaskDefinitionContainerDefinitionsRestartPolicy) {
	if err := e.validatePutRestartPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRestartPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutSecrets(value interface{}) {
	if err := e.validatePutSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSecrets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutSystemControls(value interface{}) {
	if err := e.validatePutSystemControlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSystemControls",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutUlimits(value interface{}) {
	if err := e.validatePutUlimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUlimits",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) PutVolumesFrom(value interface{}) {
	if err := e.validatePutVolumesFromParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVolumesFrom",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetCredentialSpecs() {
	_jsii_.InvokeVoid(
		e,
		"resetCredentialSpecs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		e,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetDisableNetworking() {
	_jsii_.InvokeVoid(
		e,
		"resetDisableNetworking",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetDnsSearchDomains() {
	_jsii_.InvokeVoid(
		e,
		"resetDnsSearchDomains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetDnsServers() {
	_jsii_.InvokeVoid(
		e,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetDockerLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetDockerLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetDockerSecurityOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetDockerSecurityOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		e,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetEnvironmentFiles() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironmentFiles",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		e,
		"resetEssential",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetExtraHosts() {
	_jsii_.InvokeVoid(
		e,
		"resetExtraHosts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetFirelensConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetFirelensConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		e,
		"resetHostname",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		e,
		"resetImage",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetInteractive() {
	_jsii_.InvokeVoid(
		e,
		"resetInteractive",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetLinks() {
	_jsii_.InvokeVoid(
		e,
		"resetLinks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetLinuxParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetLinuxParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetMemoryReservation() {
	_jsii_.InvokeVoid(
		e,
		"resetMemoryReservation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetMountPoints() {
	_jsii_.InvokeVoid(
		e,
		"resetMountPoints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetPortMappings() {
	_jsii_.InvokeVoid(
		e,
		"resetPortMappings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetPseudoTerminal() {
	_jsii_.InvokeVoid(
		e,
		"resetPseudoTerminal",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetReadonlyRootFilesystem() {
	_jsii_.InvokeVoid(
		e,
		"resetReadonlyRootFilesystem",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetRepositoryCredentials() {
	_jsii_.InvokeVoid(
		e,
		"resetRepositoryCredentials",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetResourceRequirements() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceRequirements",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetSecrets() {
	_jsii_.InvokeVoid(
		e,
		"resetSecrets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetStartTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetStartTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetStopTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetStopTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetSystemControls() {
	_jsii_.InvokeVoid(
		e,
		"resetSystemControls",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetUlimits() {
	_jsii_.InvokeVoid(
		e,
		"resetUlimits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		e,
		"resetUser",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetVersionConsistency() {
	_jsii_.InvokeVoid(
		e,
		"resetVersionConsistency",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetVolumesFrom() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumesFrom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ResetWorkingDirectory() {
	_jsii_.InvokeVoid(
		e,
		"resetWorkingDirectory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

