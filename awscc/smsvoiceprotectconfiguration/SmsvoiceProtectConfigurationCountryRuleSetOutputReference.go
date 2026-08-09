// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceprotectconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/smsvoiceprotectconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SmsvoiceProtectConfigurationCountryRuleSetOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mms() SmsvoiceProtectConfigurationCountryRuleSetMmsList
	MmsInput() interface{}
	Sms() SmsvoiceProtectConfigurationCountryRuleSetSmsList
	SmsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Voice() SmsvoiceProtectConfigurationCountryRuleSetVoiceList
	VoiceInput() interface{}
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
	PutMms(value interface{})
	PutSms(value interface{})
	PutVoice(value interface{})
	ResetMms()
	ResetSms()
	ResetVoice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SmsvoiceProtectConfigurationCountryRuleSetOutputReference
type jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) Mms() SmsvoiceProtectConfigurationCountryRuleSetMmsList {
	var returns SmsvoiceProtectConfigurationCountryRuleSetMmsList
	_jsii_.Get(
		j,
		"mms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) MmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) Sms() SmsvoiceProtectConfigurationCountryRuleSetSmsList {
	var returns SmsvoiceProtectConfigurationCountryRuleSetSmsList
	_jsii_.Get(
		j,
		"sms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) SmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) Voice() SmsvoiceProtectConfigurationCountryRuleSetVoiceList {
	var returns SmsvoiceProtectConfigurationCountryRuleSetVoiceList
	_jsii_.Get(
		j,
		"voice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) VoiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"voiceInput",
		&returns,
	)
	return returns
}


func NewSmsvoiceProtectConfigurationCountryRuleSetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SmsvoiceProtectConfigurationCountryRuleSetOutputReference {
	_init_.Initialize()

	if err := validateNewSmsvoiceProtectConfigurationCountryRuleSetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.smsvoiceProtectConfiguration.SmsvoiceProtectConfigurationCountryRuleSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSmsvoiceProtectConfigurationCountryRuleSetOutputReference_Override(s SmsvoiceProtectConfigurationCountryRuleSetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.smsvoiceProtectConfiguration.SmsvoiceProtectConfigurationCountryRuleSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) PutMms(value interface{}) {
	if err := s.validatePutMmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMms",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) PutSms(value interface{}) {
	if err := s.validatePutSmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSms",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) PutVoice(value interface{}) {
	if err := s.validatePutVoiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVoice",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ResetMms() {
	_jsii_.InvokeVoid(
		s,
		"resetMms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ResetSms() {
	_jsii_.InvokeVoid(
		s,
		"resetSms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ResetVoice() {
	_jsii_.InvokeVoid(
		s,
		"resetVoice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceProtectConfigurationCountryRuleSetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

