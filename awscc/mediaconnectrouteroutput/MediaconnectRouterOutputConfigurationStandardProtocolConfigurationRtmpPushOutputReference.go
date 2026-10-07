// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectrouteroutput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference interface {
	cdktn.ComplexObject
	ApplicationName() *string
	SetApplicationName(val *string)
	ApplicationNameInput() *string
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
	DestinationPort() *float64
	SetDestinationPort(val *float64)
	DestinationPortInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StreamName() *string
	SetStreamName(val *string)
	StreamNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlsEncryption() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryptionOutputReference
	TlsEncryptionInput() interface{}
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
	PutTlsEncryption(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryption)
	ResetApplicationName()
	ResetDestinationAddress()
	ResetDestinationPort()
	ResetStreamName()
	ResetTlsEncryption()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference
type jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ApplicationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) DestinationPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) DestinationPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) StreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) TlsEncryption() MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryptionOutputReference {
	var returns MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryptionOutputReference
	_jsii_.Get(
		j,
		"tlsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) TlsEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEncryptionInput",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference_Override(m MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetApplicationName(val *string) {
	if err := j.validateSetApplicationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetDestinationPort(val *float64) {
	if err := j.validateSetDestinationPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetStreamName(val *string) {
	if err := j.validateSetStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) PutTlsEncryption(value *MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushTlsEncryption) {
	if err := m.validatePutTlsEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTlsEncryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ResetApplicationName() {
	_jsii_.InvokeVoid(
		m,
		"resetApplicationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ResetDestinationPort() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ResetStreamName() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ResetTlsEncryption() {
	_jsii_.InvokeVoid(
		m,
		"resetTlsEncryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationStandardProtocolConfigurationRtmpPushOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

