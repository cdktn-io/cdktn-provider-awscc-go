// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccmedialivemultiplexprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccmedialivemultiplexprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference interface {
	cdktn.ComplexObject
	AudioPids() *[]*float64
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
	DvbSubPids() *[]*float64
	DvbTeletextPid() *float64
	EtvPlatformPid() *float64
	EtvSignalPid() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccMedialiveMultiplexprogramPacketIdentifiersMap
	SetInternalValue(val *DataAwsccMedialiveMultiplexprogramPacketIdentifiersMap)
	KlvDataPids() *[]*float64
	PcrPid() *float64
	PmtPid() *float64
	PrivateMetadataPid() *float64
	Scte27Pids() *[]*float64
	Scte35Pid() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimedMetadataPid() *float64
	VideoPid() *float64
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

// The jsii proxy struct for DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference
type jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) AudioPids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"audioPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) DvbSubPids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dvbSubPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) DvbTeletextPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dvbTeletextPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) EtvPlatformPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"etvPlatformPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) EtvSignalPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"etvSignalPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) InternalValue() *DataAwsccMedialiveMultiplexprogramPacketIdentifiersMap {
	var returns *DataAwsccMedialiveMultiplexprogramPacketIdentifiersMap
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) KlvDataPids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"klvDataPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PcrPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PmtPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PrivateMetadataPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMetadataPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Scte27Pids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"scte27Pids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Scte35Pid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TimedMetadataPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) VideoPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"videoPid",
		&returns,
	)
	return returns
}


func NewDataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMedialiveMultiplexprogram.DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference_Override(d DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMedialiveMultiplexprogram.DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetInternalValue(val *DataAwsccMedialiveMultiplexprogramPacketIdentifiersMap) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

