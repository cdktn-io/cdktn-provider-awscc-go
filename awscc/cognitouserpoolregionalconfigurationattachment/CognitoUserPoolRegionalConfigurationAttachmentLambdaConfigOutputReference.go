// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolregionalconfigurationattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cognitouserpoolregionalconfigurationattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference interface {
	cdktn.ComplexObject
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
	CreateAuthChallenge() *string
	SetCreateAuthChallenge(val *string)
	CreateAuthChallengeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomEmailSender() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomEmailSenderOutputReference
	CustomEmailSenderInput() interface{}
	CustomMessage() *string
	SetCustomMessage(val *string)
	CustomMessageInput() *string
	CustomSmsSender() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomSmsSenderOutputReference
	CustomSmsSenderInput() interface{}
	DefineAuthChallenge() *string
	SetDefineAuthChallenge(val *string)
	DefineAuthChallengeInput() *string
	// Experimental.
	Fqn() *string
	InboundFederation() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigInboundFederationOutputReference
	InboundFederationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	PostAuthentication() *string
	SetPostAuthentication(val *string)
	PostAuthenticationInput() *string
	PostConfirmation() *string
	SetPostConfirmation(val *string)
	PostConfirmationInput() *string
	PreAuthentication() *string
	SetPreAuthentication(val *string)
	PreAuthenticationInput() *string
	PreSignUp() *string
	SetPreSignUp(val *string)
	PreSignUpInput() *string
	PreTokenGeneration() *string
	SetPreTokenGeneration(val *string)
	PreTokenGenerationConfig() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigPreTokenGenerationConfigOutputReference
	PreTokenGenerationConfigInput() interface{}
	PreTokenGenerationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserMigration() *string
	SetUserMigration(val *string)
	UserMigrationInput() *string
	VerifyAuthChallengeResponse() *string
	SetVerifyAuthChallengeResponse(val *string)
	VerifyAuthChallengeResponseInput() *string
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
	PutCustomEmailSender(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomEmailSender)
	PutCustomSmsSender(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomSmsSender)
	PutInboundFederation(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigInboundFederation)
	PutPreTokenGenerationConfig(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigPreTokenGenerationConfig)
	ResetCreateAuthChallenge()
	ResetCustomEmailSender()
	ResetCustomMessage()
	ResetCustomSmsSender()
	ResetDefineAuthChallenge()
	ResetInboundFederation()
	ResetKmsKeyId()
	ResetPostAuthentication()
	ResetPostConfirmation()
	ResetPreAuthentication()
	ResetPreSignUp()
	ResetPreTokenGeneration()
	ResetPreTokenGenerationConfig()
	ResetUserMigration()
	ResetVerifyAuthChallengeResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference
type jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CreateAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CreateAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CustomEmailSender() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomEmailSenderOutputReference {
	var returns CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomEmailSenderOutputReference
	_jsii_.Get(
		j,
		"customEmailSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CustomEmailSenderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEmailSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CustomMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CustomMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CustomSmsSender() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomSmsSenderOutputReference {
	var returns CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomSmsSenderOutputReference
	_jsii_.Get(
		j,
		"customSmsSender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) CustomSmsSenderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSmsSenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) DefineAuthChallenge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) DefineAuthChallengeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defineAuthChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) InboundFederation() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigInboundFederationOutputReference {
	var returns CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigInboundFederationOutputReference
	_jsii_.Get(
		j,
		"inboundFederation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) InboundFederationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundFederationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PostAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PostAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PostConfirmation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PostConfirmationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postConfirmationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreAuthentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreAuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreSignUp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreSignUpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preSignUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreTokenGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreTokenGenerationConfig() CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigPreTokenGenerationConfigOutputReference {
	var returns CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigPreTokenGenerationConfigOutputReference
	_jsii_.Get(
		j,
		"preTokenGenerationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreTokenGenerationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preTokenGenerationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PreTokenGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preTokenGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) UserMigration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) UserMigrationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) VerifyAuthChallengeResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) VerifyAuthChallengeResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyAuthChallengeResponseInput",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference_Override(c CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetCreateAuthChallenge(val *string) {
	if err := j.validateSetCreateAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetCustomMessage(val *string) {
	if err := j.validateSetCustomMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMessage",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetDefineAuthChallenge(val *string) {
	if err := j.validateSetDefineAuthChallengeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defineAuthChallenge",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetPostAuthentication(val *string) {
	if err := j.validateSetPostAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAuthentication",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetPostConfirmation(val *string) {
	if err := j.validateSetPostConfirmationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postConfirmation",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetPreAuthentication(val *string) {
	if err := j.validateSetPreAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthentication",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetPreSignUp(val *string) {
	if err := j.validateSetPreSignUpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preSignUp",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetPreTokenGeneration(val *string) {
	if err := j.validateSetPreTokenGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preTokenGeneration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetUserMigration(val *string) {
	if err := j.validateSetUserMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userMigration",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference)SetVerifyAuthChallengeResponse(val *string) {
	if err := j.validateSetVerifyAuthChallengeResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyAuthChallengeResponse",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PutCustomEmailSender(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomEmailSender) {
	if err := c.validatePutCustomEmailSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomEmailSender",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PutCustomSmsSender(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigCustomSmsSender) {
	if err := c.validatePutCustomSmsSenderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomSmsSender",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PutInboundFederation(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigInboundFederation) {
	if err := c.validatePutInboundFederationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInboundFederation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) PutPreTokenGenerationConfig(value *CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigPreTokenGenerationConfig) {
	if err := c.validatePutPreTokenGenerationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreTokenGenerationConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetCreateAuthChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetCreateAuthChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetCustomEmailSender() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomEmailSender",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetCustomMessage() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomMessage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetCustomSmsSender() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomSmsSender",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetDefineAuthChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetDefineAuthChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetInboundFederation() {
	_jsii_.InvokeVoid(
		c,
		"resetInboundFederation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetPostAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetPostAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetPostConfirmation() {
	_jsii_.InvokeVoid(
		c,
		"resetPostConfirmation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetPreAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetPreAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetPreSignUp() {
	_jsii_.InvokeVoid(
		c,
		"resetPreSignUp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetPreTokenGeneration() {
	_jsii_.InvokeVoid(
		c,
		"resetPreTokenGeneration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetPreTokenGenerationConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPreTokenGenerationConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetUserMigration() {
	_jsii_.InvokeVoid(
		c,
		"resetUserMigration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ResetVerifyAuthChallengeResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetVerifyAuthChallengeResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentLambdaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

