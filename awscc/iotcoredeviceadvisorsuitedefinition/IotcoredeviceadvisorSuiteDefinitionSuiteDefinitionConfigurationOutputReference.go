// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotcoredeviceadvisorsuitedefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotcoredeviceadvisorsuitedefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference interface {
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
	DevicePermissionRoleArn() *string
	SetDevicePermissionRoleArn(val *string)
	DevicePermissionRoleArnInput() *string
	Devices() IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationDevicesList
	DevicesInput() interface{}
	// Experimental.
	Fqn() *string
	IntendedForQualification() interface{}
	SetIntendedForQualification(val interface{})
	IntendedForQualificationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RootGroup() *string
	SetRootGroup(val *string)
	RootGroupInput() *string
	SuiteDefinitionName() *string
	SetSuiteDefinitionName(val *string)
	SuiteDefinitionNameInput() *string
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
	PutDevices(value interface{})
	ResetDevices()
	ResetIntendedForQualification()
	ResetSuiteDefinitionName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference
type jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) DevicePermissionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devicePermissionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) DevicePermissionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devicePermissionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) Devices() IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationDevicesList {
	var returns IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationDevicesList
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) DevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) IntendedForQualification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intendedForQualification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) IntendedForQualificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intendedForQualificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) RootGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) RootGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) SuiteDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suiteDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) SuiteDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suiteDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewIotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotcoredeviceadvisorSuiteDefinition.IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference_Override(i IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotcoredeviceadvisorSuiteDefinition.IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetDevicePermissionRoleArn(val *string) {
	if err := j.validateSetDevicePermissionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devicePermissionRoleArn",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetIntendedForQualification(val interface{}) {
	if err := j.validateSetIntendedForQualificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intendedForQualification",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetRootGroup(val *string) {
	if err := j.validateSetRootGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootGroup",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetSuiteDefinitionName(val *string) {
	if err := j.validateSetSuiteDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suiteDefinitionName",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) PutDevices(value interface{}) {
	if err := i.validatePutDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDevices",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ResetDevices() {
	_jsii_.InvokeVoid(
		i,
		"resetDevices",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ResetIntendedForQualification() {
	_jsii_.InvokeVoid(
		i,
		"resetIntendedForQualification",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ResetSuiteDefinitionName() {
	_jsii_.InvokeVoid(
		i,
		"resetSuiteDefinitionName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

