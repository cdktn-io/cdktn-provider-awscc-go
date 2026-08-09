// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessdeviceprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotwirelessdeviceprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessDeviceProfileLoRaWanOutputReference interface {
	cdktn.ComplexObject
	ClassBTimeout() *float64
	SetClassBTimeout(val *float64)
	ClassBTimeoutInput() *float64
	ClassCTimeout() *float64
	SetClassCTimeout(val *float64)
	ClassCTimeoutInput() *float64
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
	SetFactoryPresetFreqsList(val *[]*float64)
	FactoryPresetFreqsListInput() *[]*float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MacVersion() *string
	SetMacVersion(val *string)
	MacVersionInput() *string
	MaxDutyCycle() *float64
	SetMaxDutyCycle(val *float64)
	MaxDutyCycleInput() *float64
	MaxEirp() *float64
	SetMaxEirp(val *float64)
	MaxEirpInput() *float64
	PingSlotDr() *float64
	SetPingSlotDr(val *float64)
	PingSlotDrInput() *float64
	PingSlotFreq() *float64
	SetPingSlotFreq(val *float64)
	PingSlotFreqInput() *float64
	PingSlotPeriod() *float64
	SetPingSlotPeriod(val *float64)
	PingSlotPeriodInput() *float64
	RegParamsRevision() *string
	SetRegParamsRevision(val *string)
	RegParamsRevisionInput() *string
	RfRegion() *string
	SetRfRegion(val *string)
	RfRegionInput() *string
	RxDataRate2() *float64
	SetRxDataRate2(val *float64)
	RxDataRate2Input() *float64
	RxDelay1() *float64
	SetRxDelay1(val *float64)
	RxDelay1Input() *float64
	RxDrOffset1() *float64
	SetRxDrOffset1(val *float64)
	RxDrOffset1Input() *float64
	RxFreq2() *float64
	SetRxFreq2(val *float64)
	RxFreq2Input() *float64
	Supports32BitFCnt() interface{}
	SetSupports32BitFCnt(val interface{})
	Supports32BitFCntInput() interface{}
	SupportsClassB() interface{}
	SetSupportsClassB(val interface{})
	SupportsClassBInput() interface{}
	SupportsClassC() interface{}
	SetSupportsClassC(val interface{})
	SupportsClassCInput() interface{}
	SupportsJoin() interface{}
	SetSupportsJoin(val interface{})
	SupportsJoinInput() interface{}
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
	ResetClassBTimeout()
	ResetClassCTimeout()
	ResetFactoryPresetFreqsList()
	ResetMacVersion()
	ResetMaxDutyCycle()
	ResetMaxEirp()
	ResetPingSlotDr()
	ResetPingSlotFreq()
	ResetPingSlotPeriod()
	ResetRegParamsRevision()
	ResetRfRegion()
	ResetRxDataRate2()
	ResetRxDelay1()
	ResetRxDrOffset1()
	ResetRxFreq2()
	ResetSupports32BitFCnt()
	ResetSupportsClassB()
	ResetSupportsClassC()
	ResetSupportsJoin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessDeviceProfileLoRaWanOutputReference
type jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ClassBTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classBTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ClassBTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classBTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ClassCTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classCTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ClassCTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classCTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) FactoryPresetFreqsList() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"factoryPresetFreqsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) FactoryPresetFreqsListInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"factoryPresetFreqsListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) MacVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) MacVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) MaxDutyCycle() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDutyCycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) MaxDutyCycleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDutyCycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) MaxEirp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEirp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) MaxEirpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEirpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) PingSlotDr() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotDr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) PingSlotDrInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotDrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) PingSlotFreq() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotFreq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) PingSlotFreqInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotFreqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) PingSlotPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) PingSlotPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pingSlotPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RegParamsRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regParamsRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RegParamsRevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regParamsRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RfRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RfRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rfRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxDataRate2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDataRate2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxDataRate2Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDataRate2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxDelay1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDelay1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxDelay1Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDelay1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxDrOffset1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDrOffset1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxDrOffset1Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxDrOffset1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxFreq2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxFreq2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) RxFreq2Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rxFreq2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) Supports32BitFCnt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supports32BitFCnt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) Supports32BitFCntInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supports32BitFCntInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) SupportsClassB() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsClassB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) SupportsClassBInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsClassBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) SupportsClassC() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsClassC",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) SupportsClassCInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsClassCInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) SupportsJoin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) SupportsJoinInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsJoinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotwirelessDeviceProfileLoRaWanOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotwirelessDeviceProfileLoRaWanOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessDeviceProfileLoRaWanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessDeviceProfile.IotwirelessDeviceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessDeviceProfileLoRaWanOutputReference_Override(i IotwirelessDeviceProfileLoRaWanOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessDeviceProfile.IotwirelessDeviceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetClassBTimeout(val *float64) {
	if err := j.validateSetClassBTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classBTimeout",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetClassCTimeout(val *float64) {
	if err := j.validateSetClassCTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classCTimeout",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetFactoryPresetFreqsList(val *[]*float64) {
	if err := j.validateSetFactoryPresetFreqsListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"factoryPresetFreqsList",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetMacVersion(val *string) {
	if err := j.validateSetMacVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macVersion",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetMaxDutyCycle(val *float64) {
	if err := j.validateSetMaxDutyCycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDutyCycle",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetMaxEirp(val *float64) {
	if err := j.validateSetMaxEirpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEirp",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetPingSlotDr(val *float64) {
	if err := j.validateSetPingSlotDrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pingSlotDr",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetPingSlotFreq(val *float64) {
	if err := j.validateSetPingSlotFreqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pingSlotFreq",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetPingSlotPeriod(val *float64) {
	if err := j.validateSetPingSlotPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pingSlotPeriod",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetRegParamsRevision(val *string) {
	if err := j.validateSetRegParamsRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regParamsRevision",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetRfRegion(val *string) {
	if err := j.validateSetRfRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfRegion",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetRxDataRate2(val *float64) {
	if err := j.validateSetRxDataRate2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rxDataRate2",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetRxDelay1(val *float64) {
	if err := j.validateSetRxDelay1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rxDelay1",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetRxDrOffset1(val *float64) {
	if err := j.validateSetRxDrOffset1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rxDrOffset1",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetRxFreq2(val *float64) {
	if err := j.validateSetRxFreq2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rxFreq2",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetSupports32BitFCnt(val interface{}) {
	if err := j.validateSetSupports32BitFCntParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supports32BitFCnt",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetSupportsClassB(val interface{}) {
	if err := j.validateSetSupportsClassBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsClassB",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetSupportsClassC(val interface{}) {
	if err := j.validateSetSupportsClassCParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsClassC",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetSupportsJoin(val interface{}) {
	if err := j.validateSetSupportsJoinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsJoin",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetClassBTimeout() {
	_jsii_.InvokeVoid(
		i,
		"resetClassBTimeout",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetClassCTimeout() {
	_jsii_.InvokeVoid(
		i,
		"resetClassCTimeout",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetFactoryPresetFreqsList() {
	_jsii_.InvokeVoid(
		i,
		"resetFactoryPresetFreqsList",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetMacVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetMacVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetMaxDutyCycle() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxDutyCycle",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetMaxEirp() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxEirp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetPingSlotDr() {
	_jsii_.InvokeVoid(
		i,
		"resetPingSlotDr",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetPingSlotFreq() {
	_jsii_.InvokeVoid(
		i,
		"resetPingSlotFreq",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetPingSlotPeriod() {
	_jsii_.InvokeVoid(
		i,
		"resetPingSlotPeriod",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetRegParamsRevision() {
	_jsii_.InvokeVoid(
		i,
		"resetRegParamsRevision",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetRfRegion() {
	_jsii_.InvokeVoid(
		i,
		"resetRfRegion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetRxDataRate2() {
	_jsii_.InvokeVoid(
		i,
		"resetRxDataRate2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetRxDelay1() {
	_jsii_.InvokeVoid(
		i,
		"resetRxDelay1",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetRxDrOffset1() {
	_jsii_.InvokeVoid(
		i,
		"resetRxDrOffset1",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetRxFreq2() {
	_jsii_.InvokeVoid(
		i,
		"resetRxFreq2",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetSupports32BitFCnt() {
	_jsii_.InvokeVoid(
		i,
		"resetSupports32BitFCnt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetSupportsClassB() {
	_jsii_.InvokeVoid(
		i,
		"resetSupportsClassB",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetSupportsClassC() {
	_jsii_.InvokeVoid(
		i,
		"resetSupportsClassC",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ResetSupportsJoin() {
	_jsii_.InvokeVoid(
		i,
		"resetSupportsJoin",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessDeviceProfileLoRaWanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

