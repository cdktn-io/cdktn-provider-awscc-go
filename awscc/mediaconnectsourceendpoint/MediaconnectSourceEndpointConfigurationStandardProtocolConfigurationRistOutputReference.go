// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectsourceendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference interface {
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
	DestinationIngestPort() *float64
	SetDestinationIngestPort(val *float64)
	DestinationIngestPortInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecoveryLatencyMilliseconds() *float64
	SetRecoveryLatencyMilliseconds(val *float64)
	RecoveryLatencyMillisecondsInput() *float64
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
	ResetDestinationIngestPort()
	ResetRecoveryLatencyMilliseconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference
type jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) DestinationIngestPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) DestinationIngestPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationIngestPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) RecoveryLatencyMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryLatencyMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) RecoveryLatencyMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryLatencyMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference_Override(m MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetDestinationIngestPort(val *float64) {
	if err := j.validateSetDestinationIngestPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIngestPort",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetRecoveryLatencyMilliseconds(val *float64) {
	if err := j.validateSetRecoveryLatencyMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryLatencyMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) ResetDestinationIngestPort() {
	_jsii_.InvokeVoid(
		m,
		"resetDestinationIngestPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) ResetRecoveryLatencyMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetRecoveryLatencyMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRistOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

