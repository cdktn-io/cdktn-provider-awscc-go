// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessserviceprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotwirelessserviceprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotwirelessServiceProfileLoRaWanOutputReference interface {
	cdktn.ComplexObject
	AddGwMetadata() interface{}
	SetAddGwMetadata(val interface{})
	AddGwMetadataInput() interface{}
	ChannelMask() *string
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
	DevStatusReqFreq() *float64
	DlBucketSize() *float64
	DlRate() *float64
	DlRatePolicy() *string
	DrMax() *float64
	DrMin() *float64
	// Experimental.
	Fqn() *string
	HrAllowed() cdktn.IResolvable
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinGwDiversity() *float64
	NwkGeoLoc() cdktn.IResolvable
	PrAllowed() interface{}
	SetPrAllowed(val interface{})
	PrAllowedInput() interface{}
	RaAllowed() interface{}
	SetRaAllowed(val interface{})
	RaAllowedInput() interface{}
	ReportDevStatusBattery() cdktn.IResolvable
	ReportDevStatusMargin() cdktn.IResolvable
	TargetPer() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UlBucketSize() *float64
	UlRate() *float64
	UlRatePolicy() *string
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
	ResetAddGwMetadata()
	ResetPrAllowed()
	ResetRaAllowed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotwirelessServiceProfileLoRaWanOutputReference
type jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) AddGwMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addGwMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) AddGwMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addGwMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ChannelMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) DevStatusReqFreq() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"devStatusReqFreq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) DlBucketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dlBucketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) DlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dlRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) DlRatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dlRatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) DrMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) DrMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) HrAllowed() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"hrAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) MinGwDiversity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minGwDiversity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) NwkGeoLoc() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"nwkGeoLoc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) PrAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) PrAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) RaAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"raAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) RaAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"raAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ReportDevStatusBattery() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reportDevStatusBattery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ReportDevStatusMargin() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reportDevStatusMargin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) TargetPer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) UlBucketSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ulBucketSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) UlRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ulRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) UlRatePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ulRatePolicy",
		&returns,
	)
	return returns
}


func NewIotwirelessServiceProfileLoRaWanOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotwirelessServiceProfileLoRaWanOutputReference {
	_init_.Initialize()

	if err := validateNewIotwirelessServiceProfileLoRaWanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessServiceProfile.IotwirelessServiceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotwirelessServiceProfileLoRaWanOutputReference_Override(i IotwirelessServiceProfileLoRaWanOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotwirelessServiceProfile.IotwirelessServiceProfileLoRaWanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetAddGwMetadata(val interface{}) {
	if err := j.validateSetAddGwMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addGwMetadata",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetPrAllowed(val interface{}) {
	if err := j.validateSetPrAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prAllowed",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetRaAllowed(val interface{}) {
	if err := j.validateSetRaAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"raAllowed",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ResetAddGwMetadata() {
	_jsii_.InvokeVoid(
		i,
		"resetAddGwMetadata",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ResetPrAllowed() {
	_jsii_.InvokeVoid(
		i,
		"resetPrAllowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ResetRaAllowed() {
	_jsii_.InvokeVoid(
		i,
		"resetRaAllowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotwirelessServiceProfileLoRaWanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

