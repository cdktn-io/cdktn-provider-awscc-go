// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotwirelesswirelessdevice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessWirelessDeviceLoRaWanOutputReference interface {
	cdktn.ComplexObject
	AbpV10X() IotwirelessWirelessDeviceLoRaWanAbpV10XOutputReference
	AbpV10XInput() interface{}
	AbpV11() IotwirelessWirelessDeviceLoRaWanAbpV11OutputReference
	AbpV11Input() interface{}
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
	DevEui() *string
	SetDevEui(val *string)
	DevEuiInput() *string
	DeviceProfileId() *string
	SetDeviceProfileId(val *string)
	DeviceProfileIdInput() *string
	FPorts() IotwirelessWirelessDeviceLoRaWanFPortsOutputReference
	FPortsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OtaaV10X() IotwirelessWirelessDeviceLoRaWanOtaaV10XOutputReference
	OtaaV10XInput() interface{}
	OtaaV11() IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference
	OtaaV11Input() interface{}
	ServiceProfileId() *string
	SetServiceProfileId(val *string)
	ServiceProfileIdInput() *string
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
	PutAbpV10X(value *IotwirelessWirelessDeviceLoRaWanAbpV10X)
	PutAbpV11(value *IotwirelessWirelessDeviceLoRaWanAbpV11)
	PutFPorts(value *IotwirelessWirelessDeviceLoRaWanFPorts)
	PutOtaaV10X(value *IotwirelessWirelessDeviceLoRaWanOtaaV10X)
	PutOtaaV11(value *IotwirelessWirelessDeviceLoRaWanOtaaV11)
	ResetAbpV10X()
	ResetAbpV11()
	ResetDevEui()
	ResetDeviceProfileId()
	ResetFPorts()
	ResetOtaaV10X()
	ResetOtaaV11()
	ResetServiceProfileId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessWirelessDeviceLoRaWanOutputReference
type jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) AbpV10X() IotwirelessWirelessDeviceLoRaWanAbpV10XOutputReference {
	var returns IotwirelessWirelessDeviceLoRaWanAbpV10XOutputReference
	_jsii_.Get(
		j,
		"abpV10X",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) AbpV10XInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abpV10XInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) AbpV11() IotwirelessWirelessDeviceLoRaWanAbpV11OutputReference {
	var returns IotwirelessWirelessDeviceLoRaWanAbpV11OutputReference
	_jsii_.Get(
		j,
		"abpV11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) AbpV11Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abpV11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) DevEui() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devEui",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) DevEuiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devEuiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) DeviceProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) DeviceProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) FPorts() IotwirelessWirelessDeviceLoRaWanFPortsOutputReference {
	var returns IotwirelessWirelessDeviceLoRaWanFPortsOutputReference
	_jsii_.Get(
		j,
		"fPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) FPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) OtaaV10X() IotwirelessWirelessDeviceLoRaWanOtaaV10XOutputReference {
	var returns IotwirelessWirelessDeviceLoRaWanOtaaV10XOutputReference
	_jsii_.Get(
		j,
		"otaaV10X",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) OtaaV10XInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"otaaV10XInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) OtaaV11() IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference {
	var returns IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference
	_jsii_.Get(
		j,
		"otaaV11",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) OtaaV11Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"otaaV11Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ServiceProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ServiceProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotwirelessWirelessDeviceLoRaWanOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotwirelessWirelessDeviceLoRaWanOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessWirelessDeviceLoRaWanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessWirelessDeviceLoRaWanOutputReference_Override(i IotwirelessWirelessDeviceLoRaWanOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetDevEui(val *string) {
	if err := j.validateSetDevEuiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devEui",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetDeviceProfileId(val *string) {
	if err := j.validateSetDeviceProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceProfileId",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetServiceProfileId(val *string) {
	if err := j.validateSetServiceProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceProfileId",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) PutAbpV10X(value *IotwirelessWirelessDeviceLoRaWanAbpV10X) {
	if err := i.validatePutAbpV10XParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAbpV10X",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) PutAbpV11(value *IotwirelessWirelessDeviceLoRaWanAbpV11) {
	if err := i.validatePutAbpV11Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAbpV11",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) PutFPorts(value *IotwirelessWirelessDeviceLoRaWanFPorts) {
	if err := i.validatePutFPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFPorts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) PutOtaaV10X(value *IotwirelessWirelessDeviceLoRaWanOtaaV10X) {
	if err := i.validatePutOtaaV10XParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOtaaV10X",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) PutOtaaV11(value *IotwirelessWirelessDeviceLoRaWanOtaaV11) {
	if err := i.validatePutOtaaV11Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOtaaV11",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetAbpV10X() {
	_jsii_.InvokeVoid(
		i,
		"resetAbpV10X",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetAbpV11() {
	_jsii_.InvokeVoid(
		i,
		"resetAbpV11",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetDevEui() {
	_jsii_.InvokeVoid(
		i,
		"resetDevEui",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetDeviceProfileId() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceProfileId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetFPorts() {
	_jsii_.InvokeVoid(
		i,
		"resetFPorts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetOtaaV10X() {
	_jsii_.InvokeVoid(
		i,
		"resetOtaaV10X",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetOtaaV11() {
	_jsii_.InvokeVoid(
		i,
		"resetOtaaV11",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ResetServiceProfileId() {
	_jsii_.InvokeVoid(
		i,
		"resetServiceProfileId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

