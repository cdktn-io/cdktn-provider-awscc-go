// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorProgramAudienceMediaAlternateMediaOutputReference interface {
	cdktn.ComplexObject
	AdBreaks() MediatailorProgramAudienceMediaAlternateMediaAdBreaksList
	AdBreaksInput() interface{}
	ClipRange() MediatailorProgramAudienceMediaAlternateMediaClipRangeOutputReference
	ClipRangeInput() interface{}
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
	DurationMillis() *float64
	SetDurationMillis(val *float64)
	DurationMillisInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LiveSourceName() *string
	SetLiveSourceName(val *string)
	LiveSourceNameInput() *string
	ScheduledStartTimeMillis() *float64
	SetScheduledStartTimeMillis(val *float64)
	ScheduledStartTimeMillisInput() *float64
	SourceLocationName() *string
	SetSourceLocationName(val *string)
	SourceLocationNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VodSourceName() *string
	SetVodSourceName(val *string)
	VodSourceNameInput() *string
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
	PutAdBreaks(value interface{})
	PutClipRange(value *MediatailorProgramAudienceMediaAlternateMediaClipRange)
	ResetAdBreaks()
	ResetClipRange()
	ResetDurationMillis()
	ResetLiveSourceName()
	ResetScheduledStartTimeMillis()
	ResetSourceLocationName()
	ResetVodSourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorProgramAudienceMediaAlternateMediaOutputReference
type jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) AdBreaks() MediatailorProgramAudienceMediaAlternateMediaAdBreaksList {
	var returns MediatailorProgramAudienceMediaAlternateMediaAdBreaksList
	_jsii_.Get(
		j,
		"adBreaks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) AdBreaksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adBreaksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ClipRange() MediatailorProgramAudienceMediaAlternateMediaClipRangeOutputReference {
	var returns MediatailorProgramAudienceMediaAlternateMediaClipRangeOutputReference
	_jsii_.Get(
		j,
		"clipRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ClipRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clipRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) DurationMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) DurationMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) LiveSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) LiveSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ScheduledStartTimeMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledStartTimeMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ScheduledStartTimeMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledStartTimeMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) SourceLocationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) SourceLocationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) VodSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vodSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) VodSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vodSourceNameInput",
		&returns,
	)
	return returns
}


func NewMediatailorProgramAudienceMediaAlternateMediaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediatailorProgramAudienceMediaAlternateMediaOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorProgramAudienceMediaAlternateMediaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorProgram.MediatailorProgramAudienceMediaAlternateMediaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediatailorProgramAudienceMediaAlternateMediaOutputReference_Override(m MediatailorProgramAudienceMediaAlternateMediaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorProgram.MediatailorProgramAudienceMediaAlternateMediaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetDurationMillis(val *float64) {
	if err := j.validateSetDurationMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationMillis",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetLiveSourceName(val *string) {
	if err := j.validateSetLiveSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSourceName",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetScheduledStartTimeMillis(val *float64) {
	if err := j.validateSetScheduledStartTimeMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledStartTimeMillis",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetSourceLocationName(val *string) {
	if err := j.validateSetSourceLocationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceLocationName",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference)SetVodSourceName(val *string) {
	if err := j.validateSetVodSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vodSourceName",
		val,
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) PutAdBreaks(value interface{}) {
	if err := m.validatePutAdBreaksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAdBreaks",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) PutClipRange(value *MediatailorProgramAudienceMediaAlternateMediaClipRange) {
	if err := m.validatePutClipRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClipRange",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetAdBreaks() {
	_jsii_.InvokeVoid(
		m,
		"resetAdBreaks",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetClipRange() {
	_jsii_.InvokeVoid(
		m,
		"resetClipRange",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetDurationMillis() {
	_jsii_.InvokeVoid(
		m,
		"resetDurationMillis",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetLiveSourceName() {
	_jsii_.InvokeVoid(
		m,
		"resetLiveSourceName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetScheduledStartTimeMillis() {
	_jsii_.InvokeVoid(
		m,
		"resetScheduledStartTimeMillis",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetSourceLocationName() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceLocationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ResetVodSourceName() {
	_jsii_.InvokeVoid(
		m,
		"resetVodSourceName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorProgramAudienceMediaAlternateMediaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

