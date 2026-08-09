// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectrouteroutput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference interface {
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
	Rist() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRistOutputReference
	RistInput() interface{}
	Rtp() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtpOutputReference
	RtpInput() interface{}
	SrtCaller() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	SrtCallerInput() interface{}
	SrtListener() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListenerOutputReference
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
	PutRist(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRist)
	PutRtp(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtp)
	PutSrtCaller(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtCaller)
	PutSrtListener(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListener)
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

// The jsii proxy struct for MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference
type jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) Rist() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRistOutputReference {
	var returns MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRistOutputReference
	_jsii_.Get(
		j,
		"rist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) RistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ristInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) Rtp() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtpOutputReference {
	var returns MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtpOutputReference
	_jsii_.Get(
		j,
		"rtp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) RtpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rtpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) SrtCaller() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtCallerOutputReference {
	var returns MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtCallerOutputReference
	_jsii_.Get(
		j,
		"srtCaller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) SrtCallerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srtCallerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) SrtListener() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	var returns MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListenerOutputReference
	_jsii_.Get(
		j,
		"srtListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) SrtListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srtListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference_Override(m MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) PutRist(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRist) {
	if err := m.validatePutRistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRist",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) PutRtp(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtp) {
	if err := m.validatePutRtpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRtp",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) PutSrtCaller(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtCaller) {
	if err := m.validatePutSrtCallerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSrtCaller",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) PutSrtListener(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationSrtListener) {
	if err := m.validatePutSrtListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSrtListener",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ResetRist() {
	_jsii_.InvokeVoid(
		m,
		"resetRist",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ResetRtp() {
	_jsii_.InvokeVoid(
		m,
		"resetRtp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ResetSrtCaller() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtCaller",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ResetSrtListener() {
	_jsii_.InvokeVoid(
		m,
		"resetSrtListener",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

