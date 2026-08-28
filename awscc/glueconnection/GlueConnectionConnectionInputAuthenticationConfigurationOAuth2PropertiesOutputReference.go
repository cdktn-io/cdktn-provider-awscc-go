// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthorizationCodeProperties() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
	AuthorizationCodePropertiesInput() interface{}
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
	OAuth2ClientApplication() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplicationOutputReference
	OAuth2ClientApplicationInput() interface{}
	OAuth2Credentials() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2CredentialsOutputReference
	OAuth2CredentialsInput() interface{}
	OAuth2GrantType() *string
	SetOAuth2GrantType(val *string)
	OAuth2GrantTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
	TokenUrlParametersMap() *string
	SetTokenUrlParametersMap(val *string)
	TokenUrlParametersMapInput() *string
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
	PutAuthorizationCodeProperties(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties)
	PutOAuth2ClientApplication(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplication)
	PutOAuth2Credentials(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2Credentials)
	ResetAuthorizationCodeProperties()
	ResetOAuth2ClientApplication()
	ResetOAuth2Credentials()
	ResetOAuth2GrantType()
	ResetTokenUrl()
	ResetTokenUrlParametersMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference
type jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) AuthorizationCodeProperties() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference {
	var returns GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) AuthorizationCodePropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationCodePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2ClientApplication() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplicationOutputReference {
	var returns GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplicationOutputReference
	_jsii_.Get(
		j,
		"oAuth2ClientApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2ClientApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuth2ClientApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2Credentials() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2CredentialsOutputReference {
	var returns GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2CredentialsOutputReference
	_jsii_.Get(
		j,
		"oAuth2Credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2CredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuth2CredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuth2GrantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TokenUrlParametersMap() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlParametersMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TokenUrlParametersMapInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlParametersMapInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnection.GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference_Override(g GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnection.GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetOAuth2GrantType(val *string) {
	if err := j.validateSetOAuth2GrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oAuth2GrantType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTokenUrlParametersMap(val *string) {
	if err := j.validateSetTokenUrlParametersMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrlParametersMap",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) PutAuthorizationCodeProperties(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties) {
	if err := g.validatePutAuthorizationCodePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizationCodeProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) PutOAuth2ClientApplication(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplication) {
	if err := g.validatePutOAuth2ClientApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOAuth2ClientApplication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) PutOAuth2Credentials(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2Credentials) {
	if err := g.validatePutOAuth2CredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOAuth2Credentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetAuthorizationCodeProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationCodeProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetOAuth2ClientApplication() {
	_jsii_.InvokeVoid(
		g,
		"resetOAuth2ClientApplication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetOAuth2Credentials() {
	_jsii_.InvokeVoid(
		g,
		"resetOAuth2Credentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetOAuth2GrantType() {
	_jsii_.InvokeVoid(
		g,
		"resetOAuth2GrantType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetTokenUrlParametersMap() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrlParametersMap",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

