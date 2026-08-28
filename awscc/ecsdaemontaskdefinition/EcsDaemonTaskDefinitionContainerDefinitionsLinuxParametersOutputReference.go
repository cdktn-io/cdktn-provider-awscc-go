// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsdaemontaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference interface {
	cdktn.ComplexObject
	Capabilities() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference
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
	Devices() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesList
	DevicesInput() interface{}
	// Experimental.
	Fqn() *string
	InitProcessEnabled() interface{}
	SetInitProcessEnabled(val interface{})
	InitProcessEnabledInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tmpfs() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList
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
	PutCapabilities(value *EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilities)
	PutDevices(value interface{})
	PutTmpfs(value interface{})
	ResetCapabilities()
	ResetDevices()
	ResetInitProcessEnabled()
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

// The jsii proxy struct for EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference
type jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Capabilities() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) CapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Devices() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesList
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) DevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InitProcessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InitProcessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Tmpfs() EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList {
	var returns EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) TmpfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tmpfsInput",
		&returns,
	)
	return returns
}


func NewEcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference_Override(e EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetInitProcessEnabled(val interface{}) {
	if err := j.validateSetInitProcessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initProcessEnabled",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) PutCapabilities(value *EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilities) {
	if err := e.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) PutDevices(value interface{}) {
	if err := e.validatePutDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDevices",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) PutTmpfs(value interface{}) {
	if err := e.validatePutTmpfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTmpfs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		e,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetDevices() {
	_jsii_.InvokeVoid(
		e,
		"resetDevices",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetInitProcessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetInitProcessEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ResetTmpfs() {
	_jsii_.InvokeVoid(
		e,
		"resetTmpfs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

