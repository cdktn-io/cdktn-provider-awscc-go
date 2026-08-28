// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcoreruntime/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountPath() *string
	SetMountPath(val *string)
	MountPathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VolumeName() *string
	SetVolumeName(val *string)
	VolumeNameInput() *string
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
	ResetMountPath()
	ResetVolumeName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference
type jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) MountPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) MountPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) VolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) VolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeNameInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreRuntime.BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference_Override(b BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreRuntime.BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetMountPath(val *string) {
	if err := j.validateSetMountPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountPath",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference)SetVolumeName(val *string) {
	if err := j.validateSetVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeName",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) ResetMountPath() {
	_jsii_.InvokeVoid(
		b,
		"resetMountPath",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) ResetVolumeName() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumeName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

