// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference interface {
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
	IdentityProviderResourceUri() *string
	SetIdentityProviderResourceUri(val *string)
	IdentityProviderResourceUriInput() *string
	IdentityProviderVpcConnectionProperties() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersIdentityProviderVpcConnectionPropertiesOutputReference
	IdentityProviderVpcConnectionPropertiesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OAuthScope() *string
	SetOAuthScope(val *string)
	OAuthScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenProviderUrl() *string
	SetTokenProviderUrl(val *string)
	TokenProviderUrlInput() *string
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
	PutIdentityProviderVpcConnectionProperties(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersIdentityProviderVpcConnectionProperties)
	ResetIdentityProviderResourceUri()
	ResetIdentityProviderVpcConnectionProperties()
	ResetOAuthScope()
	ResetTokenProviderUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference
type jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) IdentityProviderResourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderResourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) IdentityProviderResourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderResourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) IdentityProviderVpcConnectionProperties() QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersIdentityProviderVpcConnectionPropertiesOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersIdentityProviderVpcConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"identityProviderVpcConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) IdentityProviderVpcConnectionPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderVpcConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) OAuthScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuthScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) OAuthScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuthScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) TokenProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenProviderUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) TokenProviderUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenProviderUrlInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference_Override(q QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSource.QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetIdentityProviderResourceUri(val *string) {
	if err := j.validateSetIdentityProviderResourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderResourceUri",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetOAuthScope(val *string) {
	if err := j.validateSetOAuthScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oAuthScope",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference)SetTokenProviderUrl(val *string) {
	if err := j.validateSetTokenProviderUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenProviderUrl",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) PutIdentityProviderVpcConnectionProperties(value *QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersIdentityProviderVpcConnectionProperties) {
	if err := q.validatePutIdentityProviderVpcConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIdentityProviderVpcConnectionProperties",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ResetIdentityProviderResourceUri() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityProviderResourceUri",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ResetIdentityProviderVpcConnectionProperties() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityProviderVpcConnectionProperties",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ResetOAuthScope() {
	_jsii_.InvokeVoid(
		q,
		"resetOAuthScope",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ResetTokenProviderUrl() {
	_jsii_.InvokeVoid(
		q,
		"resetTokenProviderUrl",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairAlternateDataSourceParametersStarburstParametersOAuthParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

