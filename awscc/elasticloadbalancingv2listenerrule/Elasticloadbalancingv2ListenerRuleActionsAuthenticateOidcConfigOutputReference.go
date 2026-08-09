// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingv2listenerrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference interface {
	cdktn.ComplexObject
	AuthenticationRequestExtraParams() *map[string]*string
	SetAuthenticationRequestExtraParams(val *map[string]*string)
	AuthenticationRequestExtraParamsInput() *map[string]*string
	AuthorizationEndpoint() *string
	SetAuthorizationEndpoint(val *string)
	AuthorizationEndpointInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	OnUnauthenticatedRequest() *string
	SetOnUnauthenticatedRequest(val *string)
	OnUnauthenticatedRequestInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	SessionCookieName() *string
	SetSessionCookieName(val *string)
	SessionCookieNameInput() *string
	SessionTimeout() *float64
	SetSessionTimeout(val *float64)
	SessionTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenEndpoint() *string
	SetTokenEndpoint(val *string)
	TokenEndpointInput() *string
	UseExistingClientSecret() interface{}
	SetUseExistingClientSecret(val interface{})
	UseExistingClientSecretInput() interface{}
	UserInfoEndpoint() *string
	SetUserInfoEndpoint(val *string)
	UserInfoEndpointInput() *string
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
	ResetAuthenticationRequestExtraParams()
	ResetAuthorizationEndpoint()
	ResetClientId()
	ResetClientSecret()
	ResetIssuer()
	ResetOnUnauthenticatedRequest()
	ResetScope()
	ResetSessionCookieName()
	ResetSessionTimeout()
	ResetTokenEndpoint()
	ResetUseExistingClientSecret()
	ResetUserInfoEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference
type jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) AuthenticationRequestExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) AuthenticationRequestExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) OnUnauthenticatedRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) OnUnauthenticatedRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) SessionCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) SessionCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) SessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) SessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) UseExistingClientSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExistingClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) UseExistingClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useExistingClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) UserInfoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) UserInfoEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpointInput",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2ListenerRule.Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference_Override(e Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2ListenerRule.Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetAuthenticationRequestExtraParams(val *map[string]*string) {
	if err := j.validateSetAuthenticationRequestExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationRequestExtraParams",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetOnUnauthenticatedRequest(val *string) {
	if err := j.validateSetOnUnauthenticatedRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUnauthenticatedRequest",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetSessionCookieName(val *string) {
	if err := j.validateSetSessionCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionCookieName",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetSessionTimeout(val *float64) {
	if err := j.validateSetSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeout",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetUseExistingClientSecret(val interface{}) {
	if err := j.validateSetUseExistingClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExistingClientSecret",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference)SetUserInfoEndpoint(val *string) {
	if err := j.validateSetUserInfoEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInfoEndpoint",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetAuthenticationRequestExtraParams() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticationRequestExtraParams",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		e,
		"resetClientId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		e,
		"resetIssuer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetOnUnauthenticatedRequest() {
	_jsii_.InvokeVoid(
		e,
		"resetOnUnauthenticatedRequest",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		e,
		"resetScope",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetSessionCookieName() {
	_jsii_.InvokeVoid(
		e,
		"resetSessionCookieName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetSessionTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetSessionTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetTokenEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetTokenEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetUseExistingClientSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetUseExistingClientSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ResetUserInfoEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetUserInfoEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleActionsAuthenticateOidcConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

