// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazoneconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	BasicAuthenticationCredentials() DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference
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
	CustomAuthenticationCredentials() *map[string]*string
	SetCustomAuthenticationCredentials(val *map[string]*string)
	CustomAuthenticationCredentialsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	OAuth2Properties() DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference
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
	PutBasicAuthenticationCredentials(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationBasicAuthenticationCredentials)
	PutOAuth2Properties(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2Properties)
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

// The jsii proxy struct for DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference
type jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) BasicAuthenticationCredentials() DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference {
	var returns DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationBasicAuthenticationCredentialsOutputReference
	_jsii_.Get(
		j,
		"basicAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) BasicAuthenticationCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) CustomAuthenticationCredentials() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) CustomAuthenticationCredentialsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) OAuth2Properties() DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference {
	var returns DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2PropertiesOutputReference
	_jsii_.Get(
		j,
		"oAuth2Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) OAuth2PropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oAuth2PropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference_Override(d DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetCustomAuthenticationCredentials(val *map[string]*string) {
	if err := j.validateSetCustomAuthenticationCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAuthenticationCredentials",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) PutBasicAuthenticationCredentials(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationBasicAuthenticationCredentials) {
	if err := d.validatePutBasicAuthenticationCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBasicAuthenticationCredentials",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) PutOAuth2Properties(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOAuth2Properties) {
	if err := d.validatePutOAuth2PropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOAuth2Properties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ResetBasicAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetBasicAuthenticationCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ResetCustomAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomAuthenticationCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ResetOAuth2Properties() {
	_jsii_.InvokeVoid(
		d,
		"resetOAuth2Properties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

