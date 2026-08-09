// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftcontainergroupdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gameliftcontainergroupdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference interface {
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
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() GameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOnList
	DependsOnInput() interface{}
	EnvironmentOverride() GameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverrideList
	EnvironmentOverrideInput() interface{}
	Essential() interface{}
	SetEssential(val interface{})
	EssentialInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheck() GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheckOutputReference
	HealthCheckInput() interface{}
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinuxCapabilities() GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilitiesOutputReference
	LinuxCapabilitiesInput() interface{}
	MemoryHardLimitMebibytes() *float64
	SetMemoryHardLimitMebibytes(val *float64)
	MemoryHardLimitMebibytesInput() *float64
	MountPoints() GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList
	MountPointsInput() interface{}
	PortConfiguration() GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationOutputReference
	PortConfigurationInput() interface{}
	ResolvedImageDigest() *string
	SetResolvedImageDigest(val *string)
	ResolvedImageDigestInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vcpu() *float64
	SetVcpu(val *float64)
	VcpuInput() *float64
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
	PutEnvironmentOverride(value interface{})
	PutHealthCheck(value *GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheck)
	PutLinuxCapabilities(value *GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilities)
	PutMountPoints(value interface{})
	PutPortConfiguration(value *GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfiguration)
	ResetContainerName()
	ResetDependsOn()
	ResetEnvironmentOverride()
	ResetEssential()
	ResetHealthCheck()
	ResetImageUri()
	ResetLinuxCapabilities()
	ResetMemoryHardLimitMebibytes()
	ResetMountPoints()
	ResetPortConfiguration()
	ResetResolvedImageDigest()
	ResetVcpu()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference
type jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) DependsOn() GameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOnList {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) DependsOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependsOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) EnvironmentOverride() GameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverrideList {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverrideList
	_jsii_.Get(
		j,
		"environmentOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) EnvironmentOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Essential() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) EssentialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"essentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) HealthCheck() GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheckOutputReference {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) LinuxCapabilities() GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilitiesOutputReference {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"linuxCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) LinuxCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) MemoryHardLimitMebibytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryHardLimitMebibytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) MemoryHardLimitMebibytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryHardLimitMebibytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) MountPoints() GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) MountPointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PortConfiguration() GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationOutputReference {
	var returns GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationOutputReference
	_jsii_.Get(
		j,
		"portConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PortConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResolvedImageDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolvedImageDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResolvedImageDigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolvedImageDigestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Vcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) VcpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpuInput",
		&returns,
	)
	return returns
}


func NewGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference_Override(g GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gameliftContainerGroupDefinition.GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetEssential(val interface{}) {
	if err := j.validateSetEssentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"essential",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetMemoryHardLimitMebibytes(val *float64) {
	if err := j.validateSetMemoryHardLimitMebibytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryHardLimitMebibytes",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetResolvedImageDigest(val *string) {
	if err := j.validateSetResolvedImageDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolvedImageDigest",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetVcpu(val *float64) {
	if err := j.validateSetVcpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcpu",
		val,
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PutDependsOn(value interface{}) {
	if err := g.validatePutDependsOnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDependsOn",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PutEnvironmentOverride(value interface{}) {
	if err := g.validatePutEnvironmentOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnvironmentOverride",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PutHealthCheck(value *GameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheck) {
	if err := g.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PutLinuxCapabilities(value *GameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilities) {
	if err := g.validatePutLinuxCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxCapabilities",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PutMountPoints(value interface{}) {
	if err := g.validatePutMountPointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMountPoints",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PutPortConfiguration(value *GameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfiguration) {
	if err := g.validatePutPortConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPortConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetContainerName() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetDependsOn() {
	_jsii_.InvokeVoid(
		g,
		"resetDependsOn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetEnvironmentOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetEssential() {
	_jsii_.InvokeVoid(
		g,
		"resetEssential",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetImageUri() {
	_jsii_.InvokeVoid(
		g,
		"resetImageUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetLinuxCapabilities() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxCapabilities",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetMemoryHardLimitMebibytes() {
	_jsii_.InvokeVoid(
		g,
		"resetMemoryHardLimitMebibytes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetMountPoints() {
	_jsii_.InvokeVoid(
		g,
		"resetMountPoints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetPortConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetPortConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetResolvedImageDigest() {
	_jsii_.InvokeVoid(
		g,
		"resetResolvedImageDigest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResetVcpu() {
	_jsii_.InvokeVoid(
		g,
		"resetVcpu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

