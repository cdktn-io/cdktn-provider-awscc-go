// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectsourceendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference interface {
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
	Rist() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference
	RistInput() interface{}
	Rtp() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtpOutputReference
	RtpInput() interface{}
	SrtCaller() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	SrtCallerInput() interface{}
	SrtListener() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
	SrtListenerInput() interface{}
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
	PutRist(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRist)
	PutRtp(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtp)
	PutSrtCaller(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCaller)
	PutSrtListener(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListener)
	ResetRist()
	ResetRtp()
	ResetSrtCaller()
	ResetSrtListener()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference
type jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) Rist() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference
	_jsii_.Get(
		j,
		"rist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) RistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ristInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) Rtp() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtpOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtpOutputReference
	_jsii_.Get(
		j,
		"rtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) RtpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) SrtCaller() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	_jsii_.Get(
		j,
		"srtCaller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) SrtCallerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srtCallerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) SrtListener() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
	_jsii_.Get(
		j,
		"srtListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) SrtListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srtListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference_Override(m MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) PutRist(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRist) {
	if err := m.validatePutRistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRist",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) PutRtp(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtp) {
	if err := m.validatePutRtpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRtp",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) PutSrtCaller(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCaller) {
	if err := m.validatePutSrtCallerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSrtCaller",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) PutSrtListener(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListener) {
	if err := m.validatePutSrtListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSrtListener",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ResetRist() {
	_jsii_.InvokeVoid(
		m,
		"resetRist",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ResetRtp() {
	_jsii_.InvokeVoid(
		m,
		"resetRtp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ResetSrtCaller() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtCaller",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ResetSrtListener() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtListener",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

