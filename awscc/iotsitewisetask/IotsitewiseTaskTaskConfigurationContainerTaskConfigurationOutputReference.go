// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotsitewisetask/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference interface {
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
	EcrUri() *string
	SetEcrUri(val *string)
	EcrUriInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	EphemeralStorageConfiguration() IotsitewiseTaskTaskConfigurationContainerTaskConfigurationEphemeralStorageConfigurationOutputReference
	EphemeralStorageConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mounts() IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsList
	MountsInput() interface{}
	ProcessingType() *string
	SetProcessingType(val *string)
	ProcessingTypeInput() *string
	ProcessingUnit() *string
	SetProcessingUnit(val *string)
	ProcessingUnitInput() *string
	TaskExecutionRole() *string
	SetTaskExecutionRole(val *string)
	TaskExecutionRoleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutEphemeralStorageConfiguration(value *IotsitewiseTaskTaskConfigurationContainerTaskConfigurationEphemeralStorageConfiguration)
	PutMounts(value interface{})
	ResetCommand()
	ResetEnvironmentVariables()
	ResetEphemeralStorageConfiguration()
	ResetMounts()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference
type jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) EcrUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) EcrUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) EphemeralStorageConfiguration() IotsitewiseTaskTaskConfigurationContainerTaskConfigurationEphemeralStorageConfigurationOutputReference {
	var returns IotsitewiseTaskTaskConfigurationContainerTaskConfigurationEphemeralStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) EphemeralStorageConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) Mounts() IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsList {
	var returns IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsList
	_jsii_.Get(
		j,
		"mounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) MountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ProcessingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ProcessingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ProcessingUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ProcessingUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) TaskExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) TaskExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskExecutionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewIotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewIotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseTask.IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference_Override(i IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseTask.IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetEcrUri(val *string) {
	if err := j.validateSetEcrUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecrUri",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetProcessingType(val *string) {
	if err := j.validateSetProcessingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processingType",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetProcessingUnit(val *string) {
	if err := j.validateSetProcessingUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processingUnit",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetTaskExecutionRole(val *string) {
	if err := j.validateSetTaskExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskExecutionRole",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) PutEphemeralStorageConfiguration(value *IotsitewiseTaskTaskConfigurationContainerTaskConfigurationEphemeralStorageConfiguration) {
	if err := i.validatePutEphemeralStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEphemeralStorageConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) PutMounts(value interface{}) {
	if err := i.validatePutMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMounts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		i,
		"resetCommand",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		i,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ResetEphemeralStorageConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetEphemeralStorageConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ResetMounts() {
	_jsii_.InvokeVoid(
		i,
		"resetMounts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseTaskTaskConfigurationContainerTaskConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

