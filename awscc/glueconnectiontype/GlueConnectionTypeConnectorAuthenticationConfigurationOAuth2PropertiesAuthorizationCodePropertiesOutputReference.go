// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthorizationCode() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeOutputReference
	AuthorizationCodeInput() interface{}
	AuthorizationCodeUrl() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrlOutputReference
	AuthorizationCodeUrlInput() interface{}
	ClientId() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientIdOutputReference
	ClientIdInput() interface{}
	ClientSecret() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecretOutputReference
	ClientSecretInput() interface{}
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
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Prompt() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPromptOutputReference
	PromptInput() interface{}
	RedirectUri() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUriOutputReference
	RedirectUriInput() interface{}
	RequestMethod() *string
	SetRequestMethod(val *string)
	RequestMethodInput() *string
	Scope() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScopeOutputReference
	ScopeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlOutputReference
	TokenUrlInput() interface{}
	TokenUrlParameters() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlParametersList
	TokenUrlParametersInput() interface{}
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
	PutAuthorizationCode(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCode)
	PutAuthorizationCodeUrl(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrl)
	PutClientId(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientId)
	PutClientSecret(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecret)
	PutPrompt(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPrompt)
	PutRedirectUri(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUri)
	PutScope(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScope)
	PutTokenUrl(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrl)
	PutTokenUrlParameters(value interface{})
	ResetAuthorizationCode()
	ResetAuthorizationCodeUrl()
	ResetClientId()
	ResetClientSecret()
	ResetContentType()
	ResetPrompt()
	ResetRedirectUri()
	ResetRequestMethod()
	ResetScope()
	ResetTokenUrl()
	ResetTokenUrlParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
type jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) AuthorizationCode() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeOutputReference
	_jsii_.Get(
		j,
		"authorizationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) AuthorizationCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) AuthorizationCodeUrl() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrlOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrlOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) AuthorizationCodeUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationCodeUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ClientId() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientIdOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientIdOutputReference
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ClientIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ClientSecret() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecretOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Prompt() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPromptOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPromptOutputReference
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) RedirectUri() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUriOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUriOutputReference
	_jsii_.Get(
		j,
		"redirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) RedirectUriInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) RequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Scope() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScopeOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScopeOutputReference
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TokenUrl() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlOutputReference
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TokenUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TokenUrlParameters() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlParametersList {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlParametersList
	_jsii_.Get(
		j,
		"tokenUrlParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TokenUrlParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenUrlParametersInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference_Override(g GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetRequestMethod(val *string) {
	if err := j.validateSetRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutAuthorizationCode(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCode) {
	if err := g.validatePutAuthorizationCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizationCode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutAuthorizationCodeUrl(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrl) {
	if err := g.validatePutAuthorizationCodeUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizationCodeUrl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutClientId(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientId) {
	if err := g.validatePutClientIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientId",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutClientSecret(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecret) {
	if err := g.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutPrompt(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPrompt) {
	if err := g.validatePutPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrompt",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutRedirectUri(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUri) {
	if err := g.validatePutRedirectUriParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRedirectUri",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutScope(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScope) {
	if err := g.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScope",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutTokenUrl(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrl) {
	if err := g.validatePutTokenUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTokenUrl",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) PutTokenUrlParameters(value interface{}) {
	if err := g.validatePutTokenUrlParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTokenUrlParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetAuthorizationCode() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetAuthorizationCodeUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationCodeUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		g,
		"resetClientId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		g,
		"resetContentType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		g,
		"resetPrompt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetRedirectUri() {
	_jsii_.InvokeVoid(
		g,
		"resetRedirectUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetRequestMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		g,
		"resetScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ResetTokenUrlParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrlParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

