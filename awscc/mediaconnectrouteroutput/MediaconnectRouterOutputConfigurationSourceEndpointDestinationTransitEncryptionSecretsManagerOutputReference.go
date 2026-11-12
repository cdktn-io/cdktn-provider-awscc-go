// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectrouteroutput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference interface {
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
	EncryptionSecret() MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerEncryptionSecretOutputReference
	EncryptionSecretInput() interface{}
	EncryptionSecretType() *string
	SetEncryptionSecretType(val *string)
	EncryptionSecretTypeInput() *string
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
	PutEncryptionSecret(value *MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerEncryptionSecret)
	ResetEncryptionSecret()
	ResetEncryptionSecretType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference
type jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) EncryptionSecret() MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerEncryptionSecretOutputReference {
	var returns MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerEncryptionSecretOutputReference
	_jsii_.Get(
		j,
		"encryptionSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) EncryptionSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) EncryptionSecretType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionSecretType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) EncryptionSecretTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionSecretTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference_Override(m MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterOutput.MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference)SetEncryptionSecretType(val *string) {
	if err := j.validateSetEncryptionSecretTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionSecretType",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) PutEncryptionSecret(value *MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerEncryptionSecret) {
	if err := m.validatePutEncryptionSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEncryptionSecret",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) ResetEncryptionSecret() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionSecret",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) ResetEncryptionSecretType() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionSecretType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterOutputConfigurationSourceEndpointDestinationTransitEncryptionSecretsManagerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

