// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsdaemontaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference interface {
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
	ContainerPath() *string
	SetContainerPath(val *string)
	ContainerPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	SourceVolume() *string
	SetSourceVolume(val *string)
	SourceVolumeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetContainerPath()
	ResetReadOnly()
	ResetSourceVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference
type jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ContainerPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) SourceVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) SourceVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference {
	_init_.Initialize()

	if err := validateNewEcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference_Override(e EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetContainerPath(val *string) {
	if err := j.validateSetContainerPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerPath",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetSourceVolume(val *string) {
	if err := j.validateSetSourceVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVolume",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ResetContainerPath() {
	_jsii_.InvokeVoid(
		e,
		"resetContainerPath",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		e,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ResetSourceVolume() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceVolume",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

