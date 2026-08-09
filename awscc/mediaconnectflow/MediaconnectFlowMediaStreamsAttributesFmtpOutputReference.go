// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowMediaStreamsAttributesFmtpOutputReference interface {
	cdktn.ComplexObject
	ChannelOrder() *string
	SetChannelOrder(val *string)
	ChannelOrderInput() *string
	Colorimetry() *string
	SetColorimetry(val *string)
	ColorimetryInput() *string
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
	ExactFramerate() *string
	SetExactFramerate(val *string)
	ExactFramerateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Par() *string
	SetPar(val *string)
	ParInput() *string
	Range() *string
	SetRange(val *string)
	RangeInput() *string
	ScanMode() *string
	SetScanMode(val *string)
	ScanModeInput() *string
	Tcs() *string
	SetTcs(val *string)
	TcsInput() *string
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
	ResetChannelOrder()
	ResetColorimetry()
	ResetExactFramerate()
	ResetPar()
	ResetRange()
	ResetScanMode()
	ResetTcs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectFlowMediaStreamsAttributesFmtpOutputReference
type jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ChannelOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ChannelOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) Colorimetry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorimetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ColorimetryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorimetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ExactFramerate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactFramerate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ExactFramerateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactFramerateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) Par() *string {
	var returns *string
	_jsii_.Get(
		j,
		"par",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ParInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ScanMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ScanModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) Tcs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) TcsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectFlowMediaStreamsAttributesFmtpOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectFlowMediaStreamsAttributesFmtpOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowMediaStreamsAttributesFmtpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowMediaStreamsAttributesFmtpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectFlowMediaStreamsAttributesFmtpOutputReference_Override(m MediaconnectFlowMediaStreamsAttributesFmtpOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowMediaStreamsAttributesFmtpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetChannelOrder(val *string) {
	if err := j.validateSetChannelOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelOrder",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetColorimetry(val *string) {
	if err := j.validateSetColorimetryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorimetry",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetExactFramerate(val *string) {
	if err := j.validateSetExactFramerateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactFramerate",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetPar(val *string) {
	if err := j.validateSetParParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"par",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetRange(val *string) {
	if err := j.validateSetRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"range",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetScanMode(val *string) {
	if err := j.validateSetScanModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanMode",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetTcs(val *string) {
	if err := j.validateSetTcsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcs",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetChannelOrder() {
	_jsii_.InvokeVoid(
		m,
		"resetChannelOrder",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetColorimetry() {
	_jsii_.InvokeVoid(
		m,
		"resetColorimetry",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetExactFramerate() {
	_jsii_.InvokeVoid(
		m,
		"resetExactFramerate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetPar() {
	_jsii_.InvokeVoid(
		m,
		"resetPar",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		m,
		"resetRange",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetScanMode() {
	_jsii_.InvokeVoid(
		m,
		"resetScanMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ResetTcs() {
	_jsii_.InvokeVoid(
		m,
		"resetTcs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectFlowMediaStreamsAttributesFmtpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

