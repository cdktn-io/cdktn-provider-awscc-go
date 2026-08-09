// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowSourceOutputReference interface {
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
	Decryption() MediaconnectFlowSourceDecryptionOutputReference
	DecryptionInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EntitlementArn() *string
	SetEntitlementArn(val *string)
	EntitlementArnInput() *string
	// Experimental.
	Fqn() *string
	GatewayBridgeSource() MediaconnectFlowSourceGatewayBridgeSourceOutputReference
	GatewayBridgeSourceInput() interface{}
	IngestIp() *string
	IngestPort() *float64
	SetIngestPort(val *float64)
	IngestPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxBitrate() *float64
	SetMaxBitrate(val *float64)
	MaxBitrateInput() *float64
	MaxLatency() *float64
	SetMaxLatency(val *float64)
	MaxLatencyInput() *float64
	MaxSyncBuffer() *float64
	SetMaxSyncBuffer(val *float64)
	MaxSyncBufferInput() *float64
	MediaStreamSourceConfigurations() MediaconnectFlowSourceMediaStreamSourceConfigurationsList
	MediaStreamSourceConfigurationsInput() interface{}
	MinLatency() *float64
	SetMinLatency(val *float64)
	MinLatencyInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SenderControlPort() *float64
	SetSenderControlPort(val *float64)
	SenderControlPortInput() *float64
	SenderIpAddress() *string
	SetSenderIpAddress(val *string)
	SenderIpAddressInput() *string
	SourceArn() *string
	SourceIngestPort() *string
	SourceListenerAddress() *string
	SetSourceListenerAddress(val *string)
	SourceListenerAddressInput() *string
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
	VpcInterfaceName() *string
	SetVpcInterfaceName(val *string)
	VpcInterfaceNameInput() *string
	WhitelistCidr() *string
	SetWhitelistCidr(val *string)
	WhitelistCidrInput() *string
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
	PutDecryption(value *MediaconnectFlowSourceDecryption)
	PutGatewayBridgeSource(value *MediaconnectFlowSourceGatewayBridgeSource)
	PutMediaStreamSourceConfigurations(value interface{})
	ResetDecryption()
	ResetDescription()
	ResetEntitlementArn()
	ResetGatewayBridgeSource()
	ResetIngestPort()
	ResetMaxBitrate()
	ResetMaxLatency()
	ResetMaxSyncBuffer()
	ResetMediaStreamSourceConfigurations()
	ResetMinLatency()
	ResetName()
	ResetProtocol()
	ResetSenderControlPort()
	ResetSenderIpAddress()
	ResetSourceListenerAddress()
	ResetSourceListenerPort()
	ResetStreamId()
	ResetVpcInterfaceName()
	ResetWhitelistCidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectFlowSourceOutputReference
type jsiiProxy_MediaconnectFlowSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) Decryption() MediaconnectFlowSourceDecryptionOutputReference {
	var returns MediaconnectFlowSourceDecryptionOutputReference
	_jsii_.Get(
		j,
		"decryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) DecryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) EntitlementArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) EntitlementArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) GatewayBridgeSource() MediaconnectFlowSourceGatewayBridgeSourceOutputReference {
	var returns MediaconnectFlowSourceGatewayBridgeSourceOutputReference
	_jsii_.Get(
		j,
		"gatewayBridgeSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) GatewayBridgeSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayBridgeSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) IngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) IngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) IngestPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingestPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MaxLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MaxLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MaxSyncBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSyncBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MaxSyncBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSyncBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MediaStreamSourceConfigurations() MediaconnectFlowSourceMediaStreamSourceConfigurationsList {
	var returns MediaconnectFlowSourceMediaStreamSourceConfigurationsList
	_jsii_.Get(
		j,
		"mediaStreamSourceConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MediaStreamSourceConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaStreamSourceConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MinLatency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) MinLatencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SenderControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"senderControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SenderControlPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"senderControlPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SenderIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SenderIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"senderIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SourceListenerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceListenerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SourceListenerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceListenerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SourceListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) SourceListenerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceListenerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) VpcInterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) VpcInterfaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) WhitelistCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whitelistCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference) WhitelistCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whitelistCidrInput",
		&returns,
	)
	return returns
}


func NewMediaconnectFlowSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectFlowSourceOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectFlowSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectFlowSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectFlowSourceOutputReference_Override(m MediaconnectFlowSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectFlow.MediaconnectFlowSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetEntitlementArn(val *string) {
	if err := j.validateSetEntitlementArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entitlementArn",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetIngestPort(val *float64) {
	if err := j.validateSetIngestPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetMaxLatency(val *float64) {
	if err := j.validateSetMaxLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetMaxSyncBuffer(val *float64) {
	if err := j.validateSetMaxSyncBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSyncBuffer",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetMinLatency(val *float64) {
	if err := j.validateSetMinLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLatency",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetSenderControlPort(val *float64) {
	if err := j.validateSetSenderControlPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"senderControlPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetSenderIpAddress(val *string) {
	if err := j.validateSetSenderIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"senderIpAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetSourceListenerAddress(val *string) {
	if err := j.validateSetSourceListenerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceListenerAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetSourceListenerPort(val *float64) {
	if err := j.validateSetSourceListenerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceListenerPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetVpcInterfaceName(val *string) {
	if err := j.validateSetVpcInterfaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcInterfaceName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectFlowSourceOutputReference)SetWhitelistCidr(val *string) {
	if err := j.validateSetWhitelistCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelistCidr",
		val,
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) PutDecryption(value *MediaconnectFlowSourceDecryption) {
	if err := m.validatePutDecryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDecryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) PutGatewayBridgeSource(value *MediaconnectFlowSourceGatewayBridgeSource) {
	if err := m.validatePutGatewayBridgeSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGatewayBridgeSource",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) PutMediaStreamSourceConfigurations(value interface{}) {
	if err := m.validatePutMediaStreamSourceConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMediaStreamSourceConfigurations",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetDecryption() {
	_jsii_.InvokeVoid(
		m,
		"resetDecryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetEntitlementArn() {
	_jsii_.InvokeVoid(
		m,
		"resetEntitlementArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetGatewayBridgeSource() {
	_jsii_.InvokeVoid(
		m,
		"resetGatewayBridgeSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetIngestPort() {
	_jsii_.InvokeVoid(
		m,
		"resetIngestPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetMaxLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetMaxSyncBuffer() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxSyncBuffer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetMediaStreamSourceConfigurations() {
	_jsii_.InvokeVoid(
		m,
		"resetMediaStreamSourceConfigurations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetMinLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetMinLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetSenderControlPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSenderControlPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetSenderIpAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSenderIpAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetSourceListenerAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceListenerAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetSourceListenerPort() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceListenerPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetVpcInterfaceName() {
	_jsii_.InvokeVoid(
		m,
		"resetVpcInterfaceName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ResetWhitelistCidr() {
	_jsii_.InvokeVoid(
		m,
		"resetWhitelistCidr",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectFlowSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

