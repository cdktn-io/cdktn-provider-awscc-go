// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplexprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/medialivemultiplexprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MedialiveMultiplexprogramPacketIdentifiersMapOutputReference interface {
	cdktn.ComplexObject
	AudioPids() *[]*float64
	SetAudioPids(val *[]*float64)
	AudioPidsInput() *[]*float64
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
	SetDvbSubPids(val *[]*float64)
	DvbSubPidsInput() *[]*float64
	DvbTeletextPid() *float64
	SetDvbTeletextPid(val *float64)
	DvbTeletextPidInput() *float64
	EtvPlatformPid() *float64
	SetEtvPlatformPid(val *float64)
	EtvPlatformPidInput() *float64
	EtvSignalPid() *float64
	SetEtvSignalPid(val *float64)
	EtvSignalPidInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KlvDataPids() *[]*float64
	SetKlvDataPids(val *[]*float64)
	KlvDataPidsInput() *[]*float64
	PcrPid() *float64
	SetPcrPid(val *float64)
	PcrPidInput() *float64
	PmtPid() *float64
	SetPmtPid(val *float64)
	PmtPidInput() *float64
	PrivateMetadataPid() *float64
	SetPrivateMetadataPid(val *float64)
	PrivateMetadataPidInput() *float64
	Scte27Pids() *[]*float64
	SetScte27Pids(val *[]*float64)
	Scte27PidsInput() *[]*float64
	Scte35Pid() *float64
	SetScte35Pid(val *float64)
	Scte35PidInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimedMetadataPid() *float64
	SetTimedMetadataPid(val *float64)
	TimedMetadataPidInput() *float64
	VideoPid() *float64
	SetVideoPid(val *float64)
	VideoPidInput() *float64
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
	ResetAudioPids()
	ResetDvbSubPids()
	ResetDvbTeletextPid()
	ResetEtvPlatformPid()
	ResetEtvSignalPid()
	ResetKlvDataPids()
	ResetPcrPid()
	ResetPmtPid()
	ResetPrivateMetadataPid()
	ResetScte27Pids()
	ResetScte35Pid()
	ResetTimedMetadataPid()
	ResetVideoPid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveMultiplexprogramPacketIdentifiersMapOutputReference
type jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) AudioPids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"audioPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) AudioPidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"audioPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) DvbSubPids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dvbSubPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) DvbSubPidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dvbSubPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) DvbTeletextPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dvbTeletextPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) DvbTeletextPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dvbTeletextPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) EtvPlatformPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"etvPlatformPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) EtvPlatformPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"etvPlatformPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) EtvSignalPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"etvSignalPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) EtvSignalPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"etvSignalPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) KlvDataPids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"klvDataPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) KlvDataPidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"klvDataPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PcrPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PcrPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PmtPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PmtPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PrivateMetadataPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMetadataPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) PrivateMetadataPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"privateMetadataPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Scte27Pids() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"scte27Pids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Scte27PidsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"scte27PidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Scte35Pid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Scte35PidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scte35PidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TimedMetadataPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) TimedMetadataPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) VideoPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"videoPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) VideoPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"videoPidInput",
		&returns,
	)
	return returns
}


func NewMedialiveMultiplexprogramPacketIdentifiersMapOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MedialiveMultiplexprogramPacketIdentifiersMapOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexprogramPacketIdentifiersMapOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPacketIdentifiersMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveMultiplexprogramPacketIdentifiersMapOutputReference_Override(m MedialiveMultiplexprogramPacketIdentifiersMapOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPacketIdentifiersMapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetAudioPids(val *[]*float64) {
	if err := j.validateSetAudioPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioPids",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetDvbSubPids(val *[]*float64) {
	if err := j.validateSetDvbSubPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dvbSubPids",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetDvbTeletextPid(val *float64) {
	if err := j.validateSetDvbTeletextPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dvbTeletextPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetEtvPlatformPid(val *float64) {
	if err := j.validateSetEtvPlatformPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etvPlatformPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetEtvSignalPid(val *float64) {
	if err := j.validateSetEtvSignalPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etvSignalPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetKlvDataPids(val *[]*float64) {
	if err := j.validateSetKlvDataPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"klvDataPids",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetPcrPid(val *float64) {
	if err := j.validateSetPcrPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetPmtPid(val *float64) {
	if err := j.validateSetPmtPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pmtPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetPrivateMetadataPid(val *float64) {
	if err := j.validateSetPrivateMetadataPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateMetadataPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetScte27Pids(val *[]*float64) {
	if err := j.validateSetScte27PidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte27Pids",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetScte35Pid(val *float64) {
	if err := j.validateSetScte35PidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Pid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetTimedMetadataPid(val *float64) {
	if err := j.validateSetTimedMetadataPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataPid",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)SetVideoPid(val *float64) {
	if err := j.validateSetVideoPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoPid",
		val,
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetAudioPids() {
	_jsii_.InvokeVoid(
		m,
		"resetAudioPids",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetDvbSubPids() {
	_jsii_.InvokeVoid(
		m,
		"resetDvbSubPids",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetDvbTeletextPid() {
	_jsii_.InvokeVoid(
		m,
		"resetDvbTeletextPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetEtvPlatformPid() {
	_jsii_.InvokeVoid(
		m,
		"resetEtvPlatformPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetEtvSignalPid() {
	_jsii_.InvokeVoid(
		m,
		"resetEtvSignalPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetKlvDataPids() {
	_jsii_.InvokeVoid(
		m,
		"resetKlvDataPids",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetPcrPid() {
	_jsii_.InvokeVoid(
		m,
		"resetPcrPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetPmtPid() {
	_jsii_.InvokeVoid(
		m,
		"resetPmtPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetPrivateMetadataPid() {
	_jsii_.InvokeVoid(
		m,
		"resetPrivateMetadataPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetScte27Pids() {
	_jsii_.InvokeVoid(
		m,
		"resetScte27Pids",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetScte35Pid() {
	_jsii_.InvokeVoid(
		m,
		"resetScte35Pid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetTimedMetadataPid() {
	_jsii_.InvokeVoid(
		m,
		"resetTimedMetadataPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ResetVideoPid() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoPid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

