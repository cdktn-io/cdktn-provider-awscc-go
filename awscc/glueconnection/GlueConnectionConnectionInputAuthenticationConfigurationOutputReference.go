// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionConnectionInputAuthenticationConfigurationOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	BasicAuthenticationCredentials() GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference
	BasicAuthenticationCredentialsInput() interface{}
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
	CustomAuthenticationCredentials() *string
	SetCustomAuthenticationCredentials(val *string)
	CustomAuthenticationCredentialsInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	OAuth2Properties() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference
	OAuth2PropertiesInput() interface{}
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
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
	PutBasicAuthenticationCredentials(value *GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentials)
	PutOAuth2Properties(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties)
	ResetAuthenticationType()
	ResetBasicAuthenticationCredentials()
	ResetCustomAuthenticationCredentials()
	ResetKmsKeyArn()
	ResetOAuth2Properties()
	ResetSecretArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionConnectionInputAuthenticationConfigurationOutputReference
type jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) BasicAuthenticationCredentials() GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference {
	var returns GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference
	_jsii_.Get(
		j,
		"basicAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) BasicAuthenticationCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) CustomAuthenticationCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) CustomAuthenticationCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) OAuth2Properties() GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference {
	var returns GlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference
	_jsii_.Get(
		j,
		"oAuth2Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) OAuth2PropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuth2PropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionConnectionInputAuthenticationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionConnectionInputAuthenticationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionConnectionInputAuthenticationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnection.GlueConnectionConnectionInputAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionConnectionInputAuthenticationConfigurationOutputReference_Override(g GlueConnectionConnectionInputAuthenticationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnection.GlueConnectionConnectionInputAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetCustomAuthenticationCredentials(val *string) {
	if err := j.validateSetCustomAuthenticationCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAuthenticationCredentials",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) PutBasicAuthenticationCredentials(value *GlueConnectionConnectionInputAuthenticationConfigurationBasicAuthenticationCredentials) {
	if err := g.validatePutBasicAuthenticationCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasicAuthenticationCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) PutOAuth2Properties(value *GlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties) {
	if err := g.validatePutOAuth2PropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOAuth2Properties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ResetBasicAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetBasicAuthenticationCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ResetCustomAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomAuthenticationCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ResetOAuth2Properties() {
	_jsii_.InvokeVoid(
		g,
		"resetOAuth2Properties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionConnectionInputAuthenticationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

