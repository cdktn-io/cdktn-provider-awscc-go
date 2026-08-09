// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitouserpoolregionalconfigurationattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cognitouserpoolregionalconfigurationattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference interface {
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
	ConfigurationSet() *string
	SetConfigurationSet(val *string)
	ConfigurationSetInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EmailSendingAccount() *string
	SetEmailSendingAccount(val *string)
	EmailSendingAccountInput() *string
	// Experimental.
	Fqn() *string
	From() *string
	SetFrom(val *string)
	FromInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ReplyToEmailAddress() *string
	SetReplyToEmailAddress(val *string)
	ReplyToEmailAddressInput() *string
	SourceArn() *string
	SetSourceArn(val *string)
	SourceArnInput() *string
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
	ResetConfigurationSet()
	ResetEmailSendingAccount()
	ResetFrom()
	ResetReplyToEmailAddress()
	ResetSourceArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference
type jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ConfigurationSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ConfigurationSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) EmailSendingAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) EmailSendingAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSendingAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ReplyToEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ReplyToEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference_Override(c CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cognitoUserPoolRegionalConfigurationAttachment.CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetConfigurationSet(val *string) {
	if err := j.validateSetConfigurationSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetEmailSendingAccount(val *string) {
	if err := j.validateSetEmailSendingAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSendingAccount",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetReplyToEmailAddress(val *string) {
	if err := j.validateSetReplyToEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyToEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ResetConfigurationSet() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigurationSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ResetEmailSendingAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailSendingAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ResetFrom() {
	_jsii_.InvokeVoid(
		c,
		"resetFrom",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ResetReplyToEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetReplyToEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ResetSourceArn() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CognitoUserPoolRegionalConfigurationAttachmentEmailConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

