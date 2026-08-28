// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mskreplicator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference interface {
	cdktn.ComplexObject
	ClientCredentials() MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsOutputReference
	ClientCredentialsAssertion() MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertionOutputReference
	ClientCredentialsAssertionInput() interface{}
	ClientCredentialsInput() interface{}
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
	IamJwtBearer() MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearerOutputReference
	IamJwtBearerInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenEndpointAuthenticationMethod() *string
	SetTokenEndpointAuthenticationMethod(val *string)
	TokenEndpointAuthenticationMethodInput() *string
	TokenEndpointTlsCertificateArn() *string
	SetTokenEndpointTlsCertificateArn(val *string)
	TokenEndpointTlsCertificateArnInput() *string
	TokenEndpointUrl() *string
	SetTokenEndpointUrl(val *string)
	TokenEndpointUrlInput() *string
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
	PutClientCredentials(value *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentials)
	PutClientCredentialsAssertion(value *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertion)
	PutIamJwtBearer(value *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearer)
	ResetClientCredentials()
	ResetClientCredentialsAssertion()
	ResetIamJwtBearer()
	ResetScope()
	ResetTokenEndpointAuthenticationMethod()
	ResetTokenEndpointTlsCertificateArn()
	ResetTokenEndpointUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference
type jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ClientCredentials() MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsOutputReference {
	var returns MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsOutputReference
	_jsii_.Get(
		j,
		"clientCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ClientCredentialsAssertion() MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertionOutputReference {
	var returns MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertionOutputReference
	_jsii_.Get(
		j,
		"clientCredentialsAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ClientCredentialsAssertionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCredentialsAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ClientCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) IamJwtBearer() MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearerOutputReference {
	var returns MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearerOutputReference
	_jsii_.Get(
		j,
		"iamJwtBearer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) IamJwtBearerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamJwtBearerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TokenEndpointAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TokenEndpointAuthenticationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointAuthenticationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TokenEndpointTlsCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointTlsCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TokenEndpointTlsCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointTlsCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TokenEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) TokenEndpointUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointUrlInput",
		&returns,
	)
	return returns
}


func NewMskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference {
	_init_.Initialize()

	if err := validateNewMskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference_Override(m MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mskReplicator.MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetTokenEndpointAuthenticationMethod(val *string) {
	if err := j.validateSetTokenEndpointAuthenticationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpointAuthenticationMethod",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetTokenEndpointTlsCertificateArn(val *string) {
	if err := j.validateSetTokenEndpointTlsCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpointTlsCertificateArn",
		val,
	)
}

func (j *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference)SetTokenEndpointUrl(val *string) {
	if err := j.validateSetTokenEndpointUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpointUrl",
		val,
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) PutClientCredentials(value *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentials) {
	if err := m.validatePutClientCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClientCredentials",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) PutClientCredentialsAssertion(value *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerClientCredentialsAssertion) {
	if err := m.validatePutClientCredentialsAssertionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putClientCredentialsAssertion",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) PutIamJwtBearer(value *MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerIamJwtBearer) {
	if err := m.validatePutIamJwtBearerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIamJwtBearer",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetClientCredentials() {
	_jsii_.InvokeVoid(
		m,
		"resetClientCredentials",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetClientCredentialsAssertion() {
	_jsii_.InvokeVoid(
		m,
		"resetClientCredentialsAssertion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetIamJwtBearer() {
	_jsii_.InvokeVoid(
		m,
		"resetIamJwtBearer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		m,
		"resetScope",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetTokenEndpointAuthenticationMethod() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenEndpointAuthenticationMethod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetTokenEndpointTlsCertificateArn() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenEndpointTlsCertificateArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ResetTokenEndpointUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenEndpointUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MskReplicatorKafkaClustersClientAuthenticationSaslOAuthBearerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

