// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectdestinationendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference interface {
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
	ForwardErrorCorrectionState() *string
	SetForwardErrorCorrectionState(val *string)
	ForwardErrorCorrectionStateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetDestinationAddress()
	ResetDestinationIngestPort()
	ResetForwardErrorCorrectionState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference
type jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) DestinationIngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) DestinationIngestPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ForwardErrorCorrectionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardErrorCorrectionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ForwardErrorCorrectionStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardErrorCorrectionStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectDestinationEndpoint.MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference_Override(m MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectDestinationEndpoint.MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetDestinationIngestPort(val *float64) {
	if err := j.validateSetDestinationIngestPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIngestPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetForwardErrorCorrectionState(val *string) {
	if err := j.validateSetForwardErrorCorrectionStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardErrorCorrectionState",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ResetDestinationIngestPort() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationIngestPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ResetForwardErrorCorrectionState() {
	_jsii_.InvokeVoid(
		m,
		"resetForwardErrorCorrectionState",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

