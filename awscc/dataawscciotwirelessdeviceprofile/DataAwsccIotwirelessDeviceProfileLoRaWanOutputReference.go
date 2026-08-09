// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscciotwirelessdeviceprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscciotwirelessdeviceprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference interface {
	cdktn.ComplexObject
	ClassBTimeout() *float64
	ClassCTimeout() *float64
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
	FactoryPresetFreqsList() *[]*float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccIotwirelessDeviceProfileLoRaWan
	SetInternalValue(val *DataAwsccIotwirelessDeviceProfileLoRaWan)
	MacVersion() *string
	MaxDutyCycle() *float64
	MaxEirp() *float64
	PingSlotDr() *float64
	PingSlotFreq() *float64
	PingSlotPeriod() *float64
	RegParamsRevision() *string
	RfRegion() *string
	RxDataRate2() *float64
	RxDelay1() *float64
	RxDrOffset1() *float64
	RxFreq2() *float64
	Supports32BitFCnt() cdktn.IResolvable
	SupportsClassB() cdktn.IResolvable
	SupportsClassC() cdktn.IResolvable
	SupportsJoin() cdktn.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference
type jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) ClassBTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classBTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) ClassCTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classCTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) FactoryPresetFreqsList() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"factoryPresetFreqsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) InternalValue() *DataAwsccIotwirelessDeviceProfileLoRaWan {
	var returns *DataAwsccIotwirelessDeviceProfileLoRaWan
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) MacVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) MaxDutyCycle() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDutyCycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) MaxEirp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEirp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) PingSlotDr() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotDr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) PingSlotFreq() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotFreq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) PingSlotPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) RegParamsRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regParamsRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) RfRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) RxDataRate2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDataRate2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) RxDelay1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDelay1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) RxDrOffset1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDrOffset1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) RxFreq2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxFreq2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) Supports32BitFCnt() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supports32BitFCnt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) SupportsClassB() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsClassB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) SupportsClassC() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsClassC",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) SupportsJoin() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccIotwirelessDeviceProfileLoRaWanOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotwirelessDeviceProfileLoRaWanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotwirelessDeviceProfile.DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccIotwirelessDeviceProfileLoRaWanOutputReference_Override(d DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotwirelessDeviceProfile.DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference)SetInternalValue(val *DataAwsccIotwirelessDeviceProfileLoRaWan) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotwirelessDeviceProfileLoRaWanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

