// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediatailorprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediatailorProgramAdBreaksOutputReference interface {
	cdktn.ComplexObject
	AdBreakMetadata() MediatailorProgramAdBreaksAdBreakMetadataList
	AdBreakMetadataInput() interface{}
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
	MessageType() *string
	SetMessageType(val *string)
	MessageTypeInput() *string
	OffsetMillis() *float64
	SetOffsetMillis(val *float64)
	OffsetMillisInput() *float64
	Slate() MediatailorProgramAdBreaksSlateOutputReference
	SlateInput() interface{}
	SpliceInsertMessage() MediatailorProgramAdBreaksSpliceInsertMessageOutputReference
	SpliceInsertMessageInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeSignalMessage() MediatailorProgramAdBreaksTimeSignalMessageOutputReference
	TimeSignalMessageInput() interface{}
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
	PutAdBreakMetadata(value interface{})
	PutSlate(value *MediatailorProgramAdBreaksSlate)
	PutSpliceInsertMessage(value *MediatailorProgramAdBreaksSpliceInsertMessage)
	PutTimeSignalMessage(value *MediatailorProgramAdBreaksTimeSignalMessage)
	ResetAdBreakMetadata()
	ResetMessageType()
	ResetOffsetMillis()
	ResetSlate()
	ResetSpliceInsertMessage()
	ResetTimeSignalMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediatailorProgramAdBreaksOutputReference
type jsiiProxy_MediatailorProgramAdBreaksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) AdBreakMetadata() MediatailorProgramAdBreaksAdBreakMetadataList {
	var returns MediatailorProgramAdBreaksAdBreakMetadataList
	_jsii_.Get(
		j,
		"adBreakMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) AdBreakMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adBreakMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) MessageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) MessageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) OffsetMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) OffsetMillisInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMillisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) Slate() MediatailorProgramAdBreaksSlateOutputReference {
	var returns MediatailorProgramAdBreaksSlateOutputReference
	_jsii_.Get(
		j,
		"slate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) SlateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) SpliceInsertMessage() MediatailorProgramAdBreaksSpliceInsertMessageOutputReference {
	var returns MediatailorProgramAdBreaksSpliceInsertMessageOutputReference
	_jsii_.Get(
		j,
		"spliceInsertMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) SpliceInsertMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spliceInsertMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) TimeSignalMessage() MediatailorProgramAdBreaksTimeSignalMessageOutputReference {
	var returns MediatailorProgramAdBreaksTimeSignalMessageOutputReference
	_jsii_.Get(
		j,
		"timeSignalMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference) TimeSignalMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeSignalMessageInput",
		&returns,
	)
	return returns
}


func NewMediatailorProgramAdBreaksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediatailorProgramAdBreaksOutputReference {
	_init_.Initialize()

	if err := validateNewMediatailorProgramAdBreaksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediatailorProgramAdBreaksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorProgram.MediatailorProgramAdBreaksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediatailorProgramAdBreaksOutputReference_Override(m MediatailorProgramAdBreaksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediatailorProgram.MediatailorProgramAdBreaksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetMessageType(val *string) {
	if err := j.validateSetMessageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageType",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetOffsetMillis(val *float64) {
	if err := j.validateSetOffsetMillisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offsetMillis",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediatailorProgramAdBreaksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) PutAdBreakMetadata(value interface{}) {
	if err := m.validatePutAdBreakMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAdBreakMetadata",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) PutSlate(value *MediatailorProgramAdBreaksSlate) {
	if err := m.validatePutSlateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSlate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) PutSpliceInsertMessage(value *MediatailorProgramAdBreaksSpliceInsertMessage) {
	if err := m.validatePutSpliceInsertMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSpliceInsertMessage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) PutTimeSignalMessage(value *MediatailorProgramAdBreaksTimeSignalMessage) {
	if err := m.validatePutTimeSignalMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeSignalMessage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ResetAdBreakMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetAdBreakMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ResetMessageType() {
	_jsii_.InvokeVoid(
		m,
		"resetMessageType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ResetOffsetMillis() {
	_jsii_.InvokeVoid(
		m,
		"resetOffsetMillis",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ResetSlate() {
	_jsii_.InvokeVoid(
		m,
		"resetSlate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ResetSpliceInsertMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetSpliceInsertMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ResetTimeSignalMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeSignalMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediatailorProgramAdBreaksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

