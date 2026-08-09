// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Email() Connectcampaignsv2CampaignChannelSubtypeConfigEmailOutputReference
	EmailInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Sms() Connectcampaignsv2CampaignChannelSubtypeConfigSmsOutputReference
	SmsInput() interface{}
	Telephony() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutputReference
	TelephonyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WhatsApp() Connectcampaignsv2CampaignChannelSubtypeConfigWhatsAppOutputReference
	WhatsAppInput() interface{}
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
	PutEmail(value *Connectcampaignsv2CampaignChannelSubtypeConfigEmail)
	PutSms(value *Connectcampaignsv2CampaignChannelSubtypeConfigSms)
	PutTelephony(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephony)
	PutWhatsApp(value *Connectcampaignsv2CampaignChannelSubtypeConfigWhatsApp)
	ResetEmail()
	ResetSms()
	ResetTelephony()
	ResetWhatsApp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference
type jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) Email() Connectcampaignsv2CampaignChannelSubtypeConfigEmailOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) Sms() Connectcampaignsv2CampaignChannelSubtypeConfigSmsOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigSmsOutputReference
	_jsii_.Get(
		j,
		"sms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) SmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) Telephony() Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutputReference
	_jsii_.Get(
		j,
		"telephony",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) TelephonyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"telephonyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) WhatsApp() Connectcampaignsv2CampaignChannelSubtypeConfigWhatsAppOutputReference {
	var returns Connectcampaignsv2CampaignChannelSubtypeConfigWhatsAppOutputReference
	_jsii_.Get(
		j,
		"whatsApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) WhatsAppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whatsAppInput",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignChannelSubtypeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignChannelSubtypeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignChannelSubtypeConfigOutputReference_Override(c Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) PutEmail(value *Connectcampaignsv2CampaignChannelSubtypeConfigEmail) {
	if err := c.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmail",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) PutSms(value *Connectcampaignsv2CampaignChannelSubtypeConfigSms) {
	if err := c.validatePutSmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSms",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) PutTelephony(value *Connectcampaignsv2CampaignChannelSubtypeConfigTelephony) {
	if err := c.validatePutTelephonyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTelephony",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) PutWhatsApp(value *Connectcampaignsv2CampaignChannelSubtypeConfigWhatsApp) {
	if err := c.validatePutWhatsAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWhatsApp",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ResetSms() {
	_jsii_.InvokeVoid(
		c,
		"resetSms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ResetTelephony() {
	_jsii_.InvokeVoid(
		c,
		"resetTelephony",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ResetWhatsApp() {
	_jsii_.InvokeVoid(
		c,
		"resetWhatsApp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignChannelSubtypeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

