// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference interface {
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
	Devices() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList
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
	Tmpfs() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersTmpfsList
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
	PutDevices(value interface{})
	PutTmpfs(value interface{})
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

// The jsii proxy struct for BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference
type jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) Devices() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) DevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) InitProcessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) InitProcessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) MaxSwap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSwap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) MaxSwapInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSwapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) SharedMemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedMemorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) SharedMemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedMemorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) Swappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) SwappinessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swappinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) Tmpfs() BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersTmpfsList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersTmpfsList
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) TmpfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tmpfsInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference_Override(b BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetInitProcessEnabled(val interface{}) {
	if err := j.validateSetInitProcessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initProcessEnabled",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetMaxSwap(val *float64) {
	if err := j.validateSetMaxSwapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSwap",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetSharedMemorySize(val *float64) {
	if err := j.validateSetSharedMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedMemorySize",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetSwappiness(val *float64) {
	if err := j.validateSetSwappinessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swappiness",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) PutDevices(value interface{}) {
	if err := b.validatePutDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDevices",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) PutTmpfs(value interface{}) {
	if err := b.validatePutTmpfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTmpfs",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ResetDevices() {
	_jsii_.InvokeVoid(
		b,
		"resetDevices",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ResetInitProcessEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetInitProcessEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ResetMaxSwap() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxSwap",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ResetSharedMemorySize() {
	_jsii_.InvokeVoid(
		b,
		"resetSharedMemorySize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ResetSwappiness() {
	_jsii_.InvokeVoid(
		b,
		"resetSwappiness",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ResetTmpfs() {
	_jsii_.InvokeVoid(
		b,
		"resetTmpfs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

