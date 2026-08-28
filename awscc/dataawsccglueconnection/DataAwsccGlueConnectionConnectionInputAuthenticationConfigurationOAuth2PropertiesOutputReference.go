// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccglueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccglueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthorizationCodeProperties() DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
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
	InternalValue() *DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties
	SetInternalValue(val *DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties)
	OAuth2ClientApplication() DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplicationOutputReference
	OAuth2Credentials() DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2CredentialsOutputReference
	OAuth2GrantType() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() *string
	TokenUrlParametersMap() *string
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

// The jsii proxy struct for DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference
type jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) AuthorizationCodeProperties() DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference {
	var returns DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesAuthorizationCodePropertiesOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) InternalValue() *DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties {
	var returns *DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2ClientApplication() DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplicationOutputReference {
	var returns DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2ClientApplicationOutputReference
	_jsii_.Get(
		j,
		"oAuth2ClientApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2Credentials() DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2CredentialsOutputReference {
	var returns DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOAuth2CredentialsOutputReference
	_jsii_.Get(
		j,
		"oAuth2Credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) OAuth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) TokenUrlParametersMap() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlParametersMap",
		&returns,
	)
	return returns
}


func NewDataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGlueConnection.DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference_Override(d DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGlueConnection.DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetInternalValue(val *DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2Properties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccGlueConnectionConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

