// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccecsdaemontaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccecsdaemontaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference interface {
	cdktn.ComplexObject
	Command() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsDependsOnList
	EntryPoint() *[]*string
	Environment() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList
	EnvironmentFiles() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList
	Essential() cdktn.IResolvable
	FirelensConfiguration() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	// Experimental.
	Fqn() *string
	HealthCheck() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	Image() *string
	Interactive() cdktn.IResolvable
	InternalValue() *DataAwsccEcsDaemonTaskDefinitionContainerDefinitions
	SetInternalValue(val *DataAwsccEcsDaemonTaskDefinitionContainerDefinitions)
	LinuxParameters() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	LogConfiguration() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	Memory() *float64
	MemoryReservation() *float64
	MountPoints() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsMountPointsList
	Name() *string
	Privileged() cdktn.IResolvable
	PseudoTerminal() cdktn.IResolvable
	ReadonlyRootFilesystem() cdktn.IResolvable
	RepositoryCredentials() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	RestartPolicy() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference
	Secrets() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsSecretsList
	StartTimeout() *float64
	StopTimeout() *float64
	SystemControls() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ulimits() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsUlimitsList
	User() *string
	WorkingDirectory() *string
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

// The jsii proxy struct for DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference
type jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) DependsOn() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsDependsOnList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EntryPoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Environment() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) EnvironmentFiles() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList
	_jsii_.Get(
		j,
		"environmentFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Essential() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) FirelensConfiguration() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) HealthCheck() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Interactive() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"interactive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InternalValue() *DataAwsccEcsDaemonTaskDefinitionContainerDefinitions {
	var returns *DataAwsccEcsDaemonTaskDefinitionContainerDefinitions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) LinuxParameters() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) LogConfiguration() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MemoryReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) MountPoints() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsMountPointsList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Privileged() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) PseudoTerminal() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"pseudoTerminal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ReadonlyRootFilesystem() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) RepositoryCredentials() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) RestartPolicy() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Secrets() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsSecretsList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) SystemControls() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList
	_jsii_.Get(
		j,
		"systemControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Ulimits() DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsUlimitsList {
	var returns DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}


func NewDataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEcsDaemonTaskDefinition.DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference_Override(d DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEcsDaemonTaskDefinition.DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetInternalValue(val *DataAwsccEcsDaemonTaskDefinitionContainerDefinitions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEcsDaemonTaskDefinitionContainerDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

