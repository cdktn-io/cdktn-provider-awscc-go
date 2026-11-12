// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectsourceendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectSourceEndpointConfigurationStandardOutputReference interface {
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolConfiguration() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference
	ProtocolConfigurationInput() interface{}
	ProtocolInput() *string
	SourceNetworkingConfiguration() MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationOutputReference
	SourceNetworkingConfigurationInput() interface{}
	SourceNetworkingType() *string
	SetSourceNetworkingType(val *string)
	SourceNetworkingTypeInput() *string
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
	PutProtocolConfiguration(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfiguration)
	PutSourceNetworkingConfiguration(value *MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfiguration)
	ResetName()
	ResetProtocol()
	ResetProtocolConfiguration()
	ResetSourceNetworkingConfiguration()
	ResetSourceNetworkingType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectSourceEndpointConfigurationStandardOutputReference
type jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ProtocolConfiguration() MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationOutputReference
	_jsii_.Get(
		j,
		"protocolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ProtocolConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) SourceNetworkingConfiguration() MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationOutputReference {
	var returns MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationOutputReference
	_jsii_.Get(
		j,
		"sourceNetworkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) SourceNetworkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceNetworkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) SourceNetworkingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceNetworkingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) SourceNetworkingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceNetworkingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectSourceEndpointConfigurationStandardOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectSourceEndpointConfigurationStandardOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectSourceEndpointConfigurationStandardOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectSourceEndpointConfigurationStandardOutputReference_Override(m MediaconnectSourceEndpointConfigurationStandardOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectSourceEndpoint.MediaconnectSourceEndpointConfigurationStandardOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetSourceNetworkingType(val *string) {
	if err := j.validateSetSourceNetworkingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceNetworkingType",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) PutProtocolConfiguration(value *MediaconnectSourceEndpointConfigurationStandardProtocolConfiguration) {
	if err := m.validatePutProtocolConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProtocolConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) PutSourceNetworkingConfiguration(value *MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfiguration) {
	if err := m.validatePutSourceNetworkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSourceNetworkingConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ResetProtocolConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetProtocolConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ResetSourceNetworkingConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceNetworkingConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ResetSourceNetworkingType() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceNetworkingType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

