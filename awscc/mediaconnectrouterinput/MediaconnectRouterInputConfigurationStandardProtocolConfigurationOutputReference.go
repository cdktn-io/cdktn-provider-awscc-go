// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectrouterinput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference interface {
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
	Rist() MediaconnectRouterInputConfigurationStandardProtocolConfigurationRistOutputReference
	RistInput() interface{}
	Rtp() MediaconnectRouterInputConfigurationStandardProtocolConfigurationRtpOutputReference
	RtpInput() interface{}
	SrtCaller() MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	SrtCallerInput() interface{}
	SrtListener() MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtListenerOutputReference
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
	PutRist(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationRist)
	PutRtp(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationRtp)
	PutSrtCaller(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCaller)
	PutSrtListener(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtListener)
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

// The jsii proxy struct for MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference
type jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) Rist() MediaconnectRouterInputConfigurationStandardProtocolConfigurationRistOutputReference {
	var returns MediaconnectRouterInputConfigurationStandardProtocolConfigurationRistOutputReference
	_jsii_.Get(
		j,
		"rist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) RistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ristInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) Rtp() MediaconnectRouterInputConfigurationStandardProtocolConfigurationRtpOutputReference {
	var returns MediaconnectRouterInputConfigurationStandardProtocolConfigurationRtpOutputReference
	_jsii_.Get(
		j,
		"rtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) RtpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) SrtCaller() MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference {
	var returns MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	_jsii_.Get(
		j,
		"srtCaller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) SrtCallerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srtCallerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) SrtListener() MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	var returns MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtListenerOutputReference
	_jsii_.Get(
		j,
		"srtListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) SrtListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srtListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterInput.MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference_Override(m MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterInput.MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) PutRist(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationRist) {
	if err := m.validatePutRistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRist",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) PutRtp(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationRtp) {
	if err := m.validatePutRtpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRtp",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) PutSrtCaller(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCaller) {
	if err := m.validatePutSrtCallerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSrtCaller",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) PutSrtListener(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtListener) {
	if err := m.validatePutSrtListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSrtListener",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ResetRist() {
	_jsii_.InvokeVoid(
		m,
		"resetRist",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ResetRtp() {
	_jsii_.InvokeVoid(
		m,
		"resetRtp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ResetSrtCaller() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtCaller",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ResetSrtListener() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtListener",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

