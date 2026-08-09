// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcoreoauth2credentialprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference interface {
	cdktn.ComplexObject
	AtlassianOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputAtlassianOauth2ProviderConfigOutputReference
	AtlassianOauth2ProviderConfigInput() interface{}
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
	CustomOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOutputReference
	CustomOauth2ProviderConfigInput() interface{}
	// Experimental.
	Fqn() *string
	GithubOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGithubOauth2ProviderConfigOutputReference
	GithubOauth2ProviderConfigInput() interface{}
	GoogleOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGoogleOauth2ProviderConfigOutputReference
	GoogleOauth2ProviderConfigInput() interface{}
	IncludedOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputIncludedOauth2ProviderConfigOutputReference
	IncludedOauth2ProviderConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinkedinOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfigOutputReference
	LinkedinOauth2ProviderConfigInput() interface{}
	MicrosoftOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputMicrosoftOauth2ProviderConfigOutputReference
	MicrosoftOauth2ProviderConfigInput() interface{}
	SalesforceOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSalesforceOauth2ProviderConfigOutputReference
	SalesforceOauth2ProviderConfigInput() interface{}
	SlackOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSlackOauth2ProviderConfigOutputReference
	SlackOauth2ProviderConfigInput() interface{}
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
	PutAtlassianOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputAtlassianOauth2ProviderConfig)
	PutCustomOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfig)
	PutGithubOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGithubOauth2ProviderConfig)
	PutGoogleOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGoogleOauth2ProviderConfig)
	PutIncludedOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputIncludedOauth2ProviderConfig)
	PutLinkedinOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfig)
	PutMicrosoftOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputMicrosoftOauth2ProviderConfig)
	PutSalesforceOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSalesforceOauth2ProviderConfig)
	PutSlackOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSlackOauth2ProviderConfig)
	ResetAtlassianOauth2ProviderConfig()
	ResetCustomOauth2ProviderConfig()
	ResetGithubOauth2ProviderConfig()
	ResetGoogleOauth2ProviderConfig()
	ResetIncludedOauth2ProviderConfig()
	ResetLinkedinOauth2ProviderConfig()
	ResetMicrosoftOauth2ProviderConfig()
	ResetSalesforceOauth2ProviderConfig()
	ResetSlackOauth2ProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference
type jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) AtlassianOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputAtlassianOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputAtlassianOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"atlassianOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) AtlassianOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atlassianOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) CustomOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"customOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) CustomOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GithubOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGithubOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGithubOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"githubOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GithubOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GoogleOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGoogleOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGoogleOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"googleOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GoogleOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) IncludedOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputIncludedOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputIncludedOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"includedOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) IncludedOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) LinkedinOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"linkedinOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) LinkedinOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkedinOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) MicrosoftOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputMicrosoftOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputMicrosoftOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"microsoftOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) MicrosoftOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) SalesforceOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSalesforceOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSalesforceOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"salesforceOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) SalesforceOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) SlackOauth2ProviderConfig() BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSlackOauth2ProviderConfigOutputReference {
	var returns BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSlackOauth2ProviderConfigOutputReference
	_jsii_.Get(
		j,
		"slackOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) SlackOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreOAuth2CredentialProvider.BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference_Override(b BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreOAuth2CredentialProvider.BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutAtlassianOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputAtlassianOauth2ProviderConfig) {
	if err := b.validatePutAtlassianOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAtlassianOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutCustomOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfig) {
	if err := b.validatePutCustomOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutGithubOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGithubOauth2ProviderConfig) {
	if err := b.validatePutGithubOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGithubOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutGoogleOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGoogleOauth2ProviderConfig) {
	if err := b.validatePutGoogleOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGoogleOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutIncludedOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputIncludedOauth2ProviderConfig) {
	if err := b.validatePutIncludedOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putIncludedOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutLinkedinOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfig) {
	if err := b.validatePutLinkedinOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLinkedinOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutMicrosoftOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputMicrosoftOauth2ProviderConfig) {
	if err := b.validatePutMicrosoftOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMicrosoftOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutSalesforceOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSalesforceOauth2ProviderConfig) {
	if err := b.validatePutSalesforceOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSalesforceOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) PutSlackOauth2ProviderConfig(value *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSlackOauth2ProviderConfig) {
	if err := b.validatePutSlackOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSlackOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetAtlassianOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetAtlassianOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetCustomOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetGithubOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetGithubOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetGoogleOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetGoogleOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetIncludedOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludedOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetLinkedinOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetLinkedinOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetMicrosoftOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetSalesforceOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetSalesforceOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ResetSlackOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetSlackOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

