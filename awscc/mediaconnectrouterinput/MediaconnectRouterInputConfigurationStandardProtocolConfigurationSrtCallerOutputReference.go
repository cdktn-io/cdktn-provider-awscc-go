// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mediaconnectrouterinput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference interface {
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
	DecryptionConfiguration() MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerDecryptionConfigurationOutputReference
	DecryptionConfigurationInput() interface{}
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
	SourcePort() *float64
	SetSourcePort(val *float64)
	SourcePortInput() *float64
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
	PutDecryptionConfiguration(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerDecryptionConfiguration)
	ResetDecryptionConfiguration()
	ResetMinimumLatencyMilliseconds()
	ResetSourceAddress()
	ResetSourcePort()
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

// The jsii proxy struct for MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference
type jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) DecryptionConfiguration() MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerDecryptionConfigurationOutputReference {
	var returns MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerDecryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"decryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) DecryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) MinimumLatencyMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) MinimumLatencyMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumLatencyMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) SourceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) SourceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) SourcePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) SourcePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) StreamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) StreamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterInput.MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference_Override(m MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterInput.MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetMinimumLatencyMilliseconds(val *float64) {
	if err := j.validateSetMinimumLatencyMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumLatencyMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetSourceAddress(val *string) {
	if err := j.validateSetSourceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetSourcePort(val *float64) {
	if err := j.validateSetSourcePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetStreamId(val *string) {
	if err := j.validateSetStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamId",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) PutDecryptionConfiguration(value *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerDecryptionConfiguration) {
	if err := m.validatePutDecryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDecryptionConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ResetDecryptionConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetDecryptionConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ResetMinimumLatencyMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMinimumLatencyMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ResetSourceAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ResetSourcePort() {
	_jsii_.InvokeVoid(
		m,
		"resetSourcePort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ResetStreamId() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

