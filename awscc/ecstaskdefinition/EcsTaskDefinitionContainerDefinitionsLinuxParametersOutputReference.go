// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecstaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference interface {
	cdktn.ComplexObject
	Capabilities() EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference
	CapabilitiesInput() interface{}
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
	Devices() EcsTaskDefinitionContainerDefinitionsLinuxParametersDevicesList
	DevicesInput() interface{}
	// Experimental.
	Fqn() *string
	InitProcessEnabled() interface{}
	SetInitProcessEnabled(val interface{})
	InitProcessEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxSwap() *float64
	SetMaxSwap(val *float64)
	MaxSwapInput() *float64
	SharedMemorySize() *float64
	SetSharedMemorySize(val *float64)
	SharedMemorySizeInput() *float64
	Swappiness() *float64
	SetSwappiness(val *float64)
	SwappinessInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tmpfs() EcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList
	TmpfsInput() interface{}
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
	PutCapabilities(value *EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilities)
	PutDevices(value interface{})
	PutTmpfs(value interface{})
	ResetCapabilities()
	ResetDevices()
	ResetInitProcessEnabled()
	ResetMaxSwap()
	ResetSharedMemorySize()
	ResetSwappiness()
	ResetTmpfs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
type jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Capabilities() EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference {
	var returns EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) CapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Devices() EcsTaskDefinitionContainerDefinitionsLinuxParametersDevicesList {
	var returns EcsTaskDefinitionContainerDefinitionsLinuxParametersDevicesList
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) DevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InitProcessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InitProcessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) MaxSwap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSwap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) MaxSwapInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSwapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) SharedMemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedMemorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) SharedMemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedMemorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Swappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) SwappinessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Tmpfs() EcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList {
	var returns EcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) TmpfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tmpfsInput",
		&returns,
	)
	return returns
}


func NewEcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsTaskDefinition.EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference_Override(e EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsTaskDefinition.EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetInitProcessEnabled(val interface{}) {
	if err := j.validateSetInitProcessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initProcessEnabled",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetMaxSwap(val *float64) {
	if err := j.validateSetMaxSwapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSwap",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetSharedMemorySize(val *float64) {
	if err := j.validateSetSharedMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedMemorySize",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetSwappiness(val *float64) {
	if err := j.validateSetSwappinessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swappiness",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) PutCapabilities(value *EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilities) {
	if err := e.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) PutDevices(value interface{}) {
	if err := e.validatePutDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDevices",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) PutTmpfs(value interface{}) {
	if err := e.validatePutTmpfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTmpfs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		e,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetDevices() {
	_jsii_.InvokeVoid(
		e,
		"resetDevices",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetInitProcessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetInitProcessEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetMaxSwap() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxSwap",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetSharedMemorySize() {
	_jsii_.InvokeVoid(
		e,
		"resetSharedMemorySize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetSwappiness() {
	_jsii_.InvokeVoid(
		e,
		"resetSwappiness",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetTmpfs() {
	_jsii_.InvokeVoid(
		e,
		"resetTmpfs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

