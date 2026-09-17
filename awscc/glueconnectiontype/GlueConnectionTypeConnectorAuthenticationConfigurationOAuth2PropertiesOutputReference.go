// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthorizationCodeProperties() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
	AuthorizationCodePropertiesInput() interface{}
	ClientCredentialsProperties() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesOutputReference
	ClientCredentialsPropertiesInput() interface{}
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
	JwtBearerProperties() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesOutputReference
	JwtBearerPropertiesInput() interface{}
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
	PutAuthorizationCodeProperties(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties)
	PutClientCredentialsProperties(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsProperties)
	PutJwtBearerProperties(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerProperties)
	ResetAuthorizationCodeProperties()
	ResetClientCredentialsProperties()
	ResetJwtBearerProperties()
	ResetOAuth2GrantType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference
type jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) AuthorizationCodeProperties() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) AuthorizationCodePropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationCodePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ClientCredentialsProperties() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsPropertiesOutputReference
	_jsii_.Get(
		j,
		"clientCredentialsProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ClientCredentialsPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCredentialsPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) JwtBearerProperties() GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesOutputReference {
	var returns GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesOutputReference
	_jsii_.Get(
		j,
		"jwtBearerProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) JwtBearerPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtBearerPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuth2GrantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference_Override(g GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference)SetOAuth2GrantType(val *string) {
	if err := j.validateSetOAuth2GrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oAuth2GrantType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) PutAuthorizationCodeProperties(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties) {
	if err := g.validatePutAuthorizationCodePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizationCodeProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) PutClientCredentialsProperties(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesClientCredentialsProperties) {
	if err := g.validatePutClientCredentialsPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientCredentialsProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) PutJwtBearerProperties(value *GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerProperties) {
	if err := g.validatePutJwtBearerPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJwtBearerProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetAuthorizationCodeProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationCodeProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetClientCredentialsProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCredentialsProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetJwtBearerProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetJwtBearerProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ResetOAuth2GrantType() {
	_jsii_.InvokeVoid(
		g,
		"resetOAuth2GrantType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

