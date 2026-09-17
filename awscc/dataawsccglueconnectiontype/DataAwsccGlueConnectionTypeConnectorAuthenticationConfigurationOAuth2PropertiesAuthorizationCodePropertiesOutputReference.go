// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccglueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccglueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthorizationCode() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeOutputReference
	AuthorizationCodeUrl() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrlOutputReference
	ClientId() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientIdOutputReference
	ClientSecret() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecretOutputReference
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties
	SetInternalValue(val *DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties)
	Prompt() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPromptOutputReference
	RedirectUri() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUriOutputReference
	RequestMethod() *string
	Scope() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScopeOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlOutputReference
	TokenUrlParameters() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlParametersList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
type jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) AuthorizationCode() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeOutputReference
	_jsii_.Get(
		j,
		"authorizationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) AuthorizationCodeUrl() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrlOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesAuthorizationCodeUrlOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ClientId() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientIdOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientIdOutputReference
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ClientSecret() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecretOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) InternalValue() *DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties {
	var returns *DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Prompt() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPromptOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesPromptOutputReference
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) RedirectUri() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUriOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesRedirectUriOutputReference
	_jsii_.Get(
		j,
		"redirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Scope() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScopeOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesScopeOutputReference
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TokenUrl() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlOutputReference {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlOutputReference
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) TokenUrlParameters() DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlParametersList {
	var returns DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesTokenUrlParametersList
	_jsii_.Get(
		j,
		"tokenUrlParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGlueConnectionType.DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference_Override(d DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGlueConnectionType.DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetInternalValue(val *DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodeProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

