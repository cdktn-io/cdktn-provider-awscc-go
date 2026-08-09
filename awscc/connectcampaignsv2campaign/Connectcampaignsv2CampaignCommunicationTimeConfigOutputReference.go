// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignsv2campaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference interface {
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
	Email() Connectcampaignsv2CampaignCommunicationTimeConfigEmailOutputReference
	EmailInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalTimeZoneConfig() Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference
	LocalTimeZoneConfigInput() interface{}
	Sms() Connectcampaignsv2CampaignCommunicationTimeConfigSmsOutputReference
	SmsInput() interface{}
	Telephony() Connectcampaignsv2CampaignCommunicationTimeConfigTelephonyOutputReference
	TelephonyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WhatsApp() Connectcampaignsv2CampaignCommunicationTimeConfigWhatsAppOutputReference
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
	PutEmail(value *Connectcampaignsv2CampaignCommunicationTimeConfigEmail)
	PutLocalTimeZoneConfig(value *Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfig)
	PutSms(value *Connectcampaignsv2CampaignCommunicationTimeConfigSms)
	PutTelephony(value *Connectcampaignsv2CampaignCommunicationTimeConfigTelephony)
	PutWhatsApp(value *Connectcampaignsv2CampaignCommunicationTimeConfigWhatsApp)
	ResetEmail()
	ResetLocalTimeZoneConfig()
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

// The jsii proxy struct for Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference
type jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) Email() Connectcampaignsv2CampaignCommunicationTimeConfigEmailOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationTimeConfigEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) LocalTimeZoneConfig() Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfigOutputReference
	_jsii_.Get(
		j,
		"localTimeZoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) LocalTimeZoneConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localTimeZoneConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) Sms() Connectcampaignsv2CampaignCommunicationTimeConfigSmsOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationTimeConfigSmsOutputReference
	_jsii_.Get(
		j,
		"sms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) SmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) Telephony() Connectcampaignsv2CampaignCommunicationTimeConfigTelephonyOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationTimeConfigTelephonyOutputReference
	_jsii_.Get(
		j,
		"telephony",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) TelephonyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"telephonyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) WhatsApp() Connectcampaignsv2CampaignCommunicationTimeConfigWhatsAppOutputReference {
	var returns Connectcampaignsv2CampaignCommunicationTimeConfigWhatsAppOutputReference
	_jsii_.Get(
		j,
		"whatsApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) WhatsAppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"whatsAppInput",
		&returns,
	)
	return returns
}


func NewConnectcampaignsv2CampaignCommunicationTimeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsv2CampaignCommunicationTimeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsv2CampaignCommunicationTimeConfigOutputReference_Override(c Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsv2Campaign.Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) PutEmail(value *Connectcampaignsv2CampaignCommunicationTimeConfigEmail) {
	if err := c.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmail",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) PutLocalTimeZoneConfig(value *Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfig) {
	if err := c.validatePutLocalTimeZoneConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocalTimeZoneConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) PutSms(value *Connectcampaignsv2CampaignCommunicationTimeConfigSms) {
	if err := c.validatePutSmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSms",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) PutTelephony(value *Connectcampaignsv2CampaignCommunicationTimeConfigTelephony) {
	if err := c.validatePutTelephonyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTelephony",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) PutWhatsApp(value *Connectcampaignsv2CampaignCommunicationTimeConfigWhatsApp) {
	if err := c.validatePutWhatsAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWhatsApp",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ResetLocalTimeZoneConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalTimeZoneConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ResetSms() {
	_jsii_.InvokeVoid(
		c,
		"resetSms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ResetTelephony() {
	_jsii_.InvokeVoid(
		c,
		"resetTelephony",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ResetWhatsApp() {
	_jsii_.InvokeVoid(
		c,
		"resetWhatsApp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

