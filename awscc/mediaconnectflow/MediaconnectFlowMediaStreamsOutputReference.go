// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowMediaStreamsOutputReference interface {
	cdktn.ComplexObject
	Attributes() MediaconnectFlowMediaStreamsAttributesOutputReference
	AttributesInput() interface{}
	ClockRate() *float64
	SetClockRate(val *float64)
	ClockRateInput() *float64
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fmt() *float64
	SetFmt(val *float64)
	FmtInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MediaStreamId() *float64
	SetMediaStreamId(val *float64)
	MediaStreamIdInput() *float64
	MediaStreamName() *string
	SetMediaStreamName(val *string)
	MediaStreamNameInput() *string
	MediaStreamType() *string
	SetMediaStreamType(val *string)
	MediaStreamTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VideoFormat() *string
	SetVideoFormat(val *string)
	VideoFormatInput() *string
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
	PutAttributes(value *MediaconnectFlowMediaStreamsAttributes)
	ResetAttributes()
	ResetClockRate()
	ResetDescription()
	ResetFmt()
	ResetMediaStreamId()
	ResetMediaStreamName()
	ResetMediaStreamType()
	ResetVideoFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectFlowMediaStreamsOutputReference
type jsiiProxy_MediaconnectFlowMediaStreamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) Attributes() MediaconnectFlowMediaStreamsAttributesOutputReference {
	var returns MediaconnectFlowMediaStreamsAttributesOutputReference
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ClockRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clockRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ClockRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clockRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) Fmt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fmt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) FmtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fmtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) MediaStreamId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mediaStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) MediaStreamIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mediaStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) MediaStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) MediaStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) MediaStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) MediaStreamTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStreamTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) VideoFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) VideoFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoFormatInput",
		&returns,
	)
	return returns
}


func NewMediaconnectFlowMediaStreamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaconnectFlowMediaStreamsOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowMediaStreamsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowMediaStreamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowMediaStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaconnectFlowMediaStreamsOutputReference_Override(m MediaconnectFlowMediaStreamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowMediaStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetClockRate(val *float64) {
	if err := j.validateSetClockRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clockRate",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetFmt(val *float64) {
	if err := j.validateSetFmtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fmt",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetMediaStreamId(val *float64) {
	if err := j.validateSetMediaStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaStreamId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetMediaStreamName(val *string) {
	if err := j.validateSetMediaStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaStreamName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetMediaStreamType(val *string) {
	if err := j.validateSetMediaStreamTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaStreamType",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference)SetVideoFormat(val *string) {
	if err := j.validateSetVideoFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoFormat",
		val,
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) PutAttributes(value *MediaconnectFlowMediaStreamsAttributes) {
	if err := m.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAttributes",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		m,
		"resetAttributes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetClockRate() {
	_jsii_.InvokeVoid(
		m,
		"resetClockRate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetFmt() {
	_jsii_.InvokeVoid(
		m,
		"resetFmt",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetMediaStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetMediaStreamName() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaStreamName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetMediaStreamType() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaStreamType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ResetVideoFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetVideoFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

