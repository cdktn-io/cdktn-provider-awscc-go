// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridgesource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectbridgesource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectBridgeSourceNetworkSourceOutputReference interface {
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
	MulticastIp() *string
	SetMulticastIp(val *string)
	MulticastIpInput() *string
	MulticastSourceSettings() MediaconnectBridgeSourceNetworkSourceMulticastSourceSettingsOutputReference
	MulticastSourceSettingsInput() interface{}
	NetworkName() *string
	SetNetworkName(val *string)
	NetworkNameInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	PutMulticastSourceSettings(value *MediaconnectBridgeSourceNetworkSourceMulticastSourceSettings)
	ResetMulticastIp()
	ResetMulticastSourceSettings()
	ResetNetworkName()
	ResetPort()
	ResetProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectBridgeSourceNetworkSourceOutputReference
type jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) MulticastIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multicastIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) MulticastIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multicastIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) MulticastSourceSettings() MediaconnectBridgeSourceNetworkSourceMulticastSourceSettingsOutputReference {
	var returns MediaconnectBridgeSourceNetworkSourceMulticastSourceSettingsOutputReference
	_jsii_.Get(
		j,
		"multicastSourceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) MulticastSourceSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multicastSourceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) NetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) NetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectBridgeSourceNetworkSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectBridgeSourceNetworkSourceOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectBridgeSourceNetworkSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectBridgeSource.MediaconnectBridgeSourceNetworkSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectBridgeSourceNetworkSourceOutputReference_Override(m MediaconnectBridgeSourceNetworkSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectBridgeSource.MediaconnectBridgeSourceNetworkSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetMulticastIp(val *string) {
	if err := j.validateSetMulticastIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multicastIp",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetNetworkName(val *string) {
	if err := j.validateSetNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkName",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) PutMulticastSourceSettings(value *MediaconnectBridgeSourceNetworkSourceMulticastSourceSettings) {
	if err := m.validatePutMulticastSourceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMulticastSourceSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ResetMulticastIp() {
	_jsii_.InvokeVoid(
		m,
		"resetMulticastIp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ResetMulticastSourceSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetMulticastSourceSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ResetNetworkName() {
	_jsii_.InvokeVoid(
		m,
		"resetNetworkName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		m,
		"resetPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectBridgeSourceNetworkSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

