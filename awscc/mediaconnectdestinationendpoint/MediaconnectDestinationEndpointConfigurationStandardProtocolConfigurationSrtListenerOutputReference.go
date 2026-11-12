// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectdestinationendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference interface {
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
	DestinationAddress() *string
	SetDestinationAddress(val *string)
	DestinationAddressInput() *string
	DestinationIngestPort() *float64
	SetDestinationIngestPort(val *float64)
	DestinationIngestPortInput() *float64
	EncryptionPassphrase() MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference
	EncryptionPassphraseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumLatencyMilliseconds() *float64
	SetMinimumLatencyMilliseconds(val *float64)
	MinimumLatencyMillisecondsInput() *float64
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
	PutEncryptionPassphrase(value *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphrase)
	ResetDestinationAddress()
	ResetDestinationIngestPort()
	ResetEncryptionPassphrase()
	ResetMinimumLatencyMilliseconds()
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

// The jsii proxy struct for MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference
type jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) DestinationIngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) DestinationIngestPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) EncryptionPassphrase() MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference {
	var returns MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphraseOutputReference
	_jsii_.Get(
		j,
		"encryptionPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) EncryptionPassphraseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) MinimumLatencyMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) MinimumLatencyMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectDestinationEndpoint.MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference_Override(m MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectDestinationEndpoint.MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetDestinationIngestPort(val *float64) {
	if err := j.validateSetDestinationIngestPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIngestPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetMinimumLatencyMilliseconds(val *float64) {
	if err := j.validateSetMinimumLatencyMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLatencyMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) PutEncryptionPassphrase(value *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphrase) {
	if err := m.validatePutEncryptionPassphraseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionPassphrase",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetDestinationIngestPort() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationIngestPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetEncryptionPassphrase() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionPassphrase",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetMinimumLatencyMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumLatencyMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

