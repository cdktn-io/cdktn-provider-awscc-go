// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectsourceendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference interface {
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
	EncryptionPassphrase() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference
	EncryptionPassphraseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumLatencyMilliseconds() *float64
	SetMinimumLatencyMilliseconds(val *float64)
	MinimumLatencyMillisecondsInput() *float64
	SourceAddress() *string
	SetSourceAddress(val *string)
	SourceAddressInput() *string
	SourceListenerPort() *float64
	SetSourceListenerPort(val *float64)
	SourceListenerPortInput() *float64
	StreamId() *string
	SetStreamId(val *string)
	StreamIdInput() *string
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
	PutEncryptionPassphrase(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphrase)
	ResetEncryptionPassphrase()
	ResetMinimumLatencyMilliseconds()
	ResetSourceAddress()
	ResetSourceListenerPort()
	ResetStreamId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
type jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) EncryptionPassphrase() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference
	_jsii_.Get(
		j,
		"encryptionPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) EncryptionPassphraseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) MinimumLatencyMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) MinimumLatencyMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) SourceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) SourceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) SourceListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) SourceListenerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference_Override(m MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetMinimumLatencyMilliseconds(val *float64) {
	if err := j.validateSetMinimumLatencyMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLatencyMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetSourceAddress(val *string) {
	if err := j.validateSetSourceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetSourceListenerPort(val *float64) {
	if err := j.validateSetSourceListenerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceListenerPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) PutEncryptionPassphrase(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphrase) {
	if err := m.validatePutEncryptionPassphraseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionPassphrase",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetEncryptionPassphrase() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionPassphrase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetMinimumLatencyMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumLatencyMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetSourceAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetSourceListenerPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceListenerPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

