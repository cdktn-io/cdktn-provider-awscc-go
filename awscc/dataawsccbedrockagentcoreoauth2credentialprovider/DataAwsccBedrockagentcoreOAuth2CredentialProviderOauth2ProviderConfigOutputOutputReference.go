// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbedrockagentcoreoauth2credentialprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbedrockagentcoreoauth2credentialprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference interface {
	cdktn.ComplexObject
	ClientAuthenticationMethod() *string
	ClientId() *string
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
	InternalValue() *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutput
	SetInternalValue(val *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutput)
	OauthDiscovery() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOauthDiscoveryOutputReference
	OnBehalfOfTokenExchangeConfig() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOnBehalfOfTokenExchangeConfigOutputReference
	PrivateEndpoint() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOutputReference
	PrivateEndpointOverrides() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesList
	PrivateKeyJwtConfig() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateKeyJwtConfigOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference
type jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) ClientAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) InternalValue() *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutput {
	var returns *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) OauthDiscovery() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOauthDiscoveryOutputReference {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOauthDiscoveryOutputReference
	_jsii_.Get(
		j,
		"oauthDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) OnBehalfOfTokenExchangeConfig() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOnBehalfOfTokenExchangeConfigOutputReference {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOnBehalfOfTokenExchangeConfigOutputReference
	_jsii_.Get(
		j,
		"onBehalfOfTokenExchangeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) PrivateEndpoint() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOutputReference {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOutputReference
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) PrivateEndpointOverrides() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesList {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateEndpointOverridesList
	_jsii_.Get(
		j,
		"privateEndpointOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) PrivateKeyJwtConfig() DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateKeyJwtConfigOutputReference {
	var returns DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputPrivateKeyJwtConfigOutputReference
	_jsii_.Get(
		j,
		"privateKeyJwtConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockagentcoreOAuth2CredentialProvider.DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference_Override(d DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockagentcoreOAuth2CredentialProvider.DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference)SetInternalValue(val *DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

