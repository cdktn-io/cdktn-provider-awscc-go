// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SegmentationEventId() *float64
	SetSegmentationEventId(val *float64)
	SegmentationEventIdInput() *float64
	SegmentationTypeId() *float64
	SetSegmentationTypeId(val *float64)
	SegmentationTypeIdInput() *float64
	SegmentationUpid() *string
	SetSegmentationUpid(val *string)
	SegmentationUpidInput() *string
	SegmentationUpidType() *float64
	SetSegmentationUpidType(val *float64)
	SegmentationUpidTypeInput() *float64
	SegmentNum() *float64
	SetSegmentNum(val *float64)
	SegmentNumInput() *float64
	SegmentsExpected() *float64
	SetSegmentsExpected(val *float64)
	SegmentsExpectedInput() *float64
	SubSegmentNum() *float64
	SetSubSegmentNum(val *float64)
	SubSegmentNumInput() *float64
	SubSegmentsExpected() *float64
	SetSubSegmentsExpected(val *float64)
	SubSegmentsExpectedInput() *float64
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
	ResetSegmentationEventId()
	ResetSegmentationTypeId()
	ResetSegmentationUpid()
	ResetSegmentationUpidType()
	ResetSegmentNum()
	ResetSegmentsExpected()
	ResetSubSegmentNum()
	ResetSubSegmentsExpected()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference
type jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationEventId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationEventId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationEventIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationEventIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationTypeId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationTypeIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationUpid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationUpid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationUpidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationUpidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationUpidType() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationUpidType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentationUpidTypeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationUpidTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentsExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SegmentsExpectedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsExpectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SubSegmentNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subSegmentNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SubSegmentNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subSegmentNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SubSegmentsExpected() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subSegmentsExpected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) SubSegmentsExpectedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subSegmentsExpectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorProgram.MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference_Override(m MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorProgram.MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSegmentationEventId(val *float64) {
	if err := j.validateSetSegmentationEventIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationEventId",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSegmentationTypeId(val *float64) {
	if err := j.validateSetSegmentationTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationTypeId",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSegmentationUpid(val *string) {
	if err := j.validateSetSegmentationUpidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationUpid",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSegmentationUpidType(val *float64) {
	if err := j.validateSetSegmentationUpidTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationUpidType",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSegmentNum(val *float64) {
	if err := j.validateSetSegmentNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentNum",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSegmentsExpected(val *float64) {
	if err := j.validateSetSegmentsExpectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentsExpected",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSubSegmentNum(val *float64) {
	if err := j.validateSetSubSegmentNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subSegmentNum",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetSubSegmentsExpected(val *float64) {
	if err := j.validateSetSubSegmentsExpectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subSegmentsExpected",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSegmentationEventId() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentationEventId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSegmentationTypeId() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentationTypeId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSegmentationUpid() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentationUpid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSegmentationUpidType() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentationUpidType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSegmentNum() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentNum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSegmentsExpected() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentsExpected",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSubSegmentNum() {
	_jsii_.InvokeVoid(
		m,
		"resetSubSegmentNum",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ResetSubSegmentsExpected() {
	_jsii_.InvokeVoid(
		m,
		"resetSubSegmentsExpected",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

