// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelesswirelessdevice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotwirelesswirelessdevice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference interface {
	cdktn.ComplexObject
	AppKey() *string
	SetAppKey(val *string)
	AppKeyInput() *string
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
	JoinEui() *string
	SetJoinEui(val *string)
	JoinEuiInput() *string
	NwkKey() *string
	SetNwkKey(val *string)
	NwkKeyInput() *string
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
	ResetAppKey()
	ResetJoinEui()
	ResetNwkKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference
type jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) AppKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) AppKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) JoinEui() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinEui",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) JoinEuiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinEuiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) NwkKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nwkKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) NwkKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nwkKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessWirelessDeviceLoRaWanOtaaV11OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference_Override(i IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessWirelessDevice.IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetAppKey(val *string) {
	if err := j.validateSetAppKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetJoinEui(val *string) {
	if err := j.validateSetJoinEuiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinEui",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetNwkKey(val *string) {
	if err := j.validateSetNwkKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nwkKey",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ResetAppKey() {
	_jsii_.InvokeVoid(
		i,
		"resetAppKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ResetJoinEui() {
	_jsii_.InvokeVoid(
		i,
		"resetJoinEui",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ResetNwkKey() {
	_jsii_.InvokeVoid(
		i,
		"resetNwkKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessWirelessDeviceLoRaWanOtaaV11OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

