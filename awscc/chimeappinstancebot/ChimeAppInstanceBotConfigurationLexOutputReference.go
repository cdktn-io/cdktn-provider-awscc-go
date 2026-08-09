// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstancebot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/chimeappinstancebot/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeAppInstanceBotConfigurationLexOutputReference interface {
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
	InvokedBy() ChimeAppInstanceBotConfigurationLexInvokedByOutputReference
	InvokedByInput() interface{}
	LexBotAliasArn() *string
	SetLexBotAliasArn(val *string)
	LexBotAliasArnInput() *string
	LocaleId() *string
	SetLocaleId(val *string)
	LocaleIdInput() *string
	RespondsTo() *string
	SetRespondsTo(val *string)
	RespondsToInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WelcomeIntent() *string
	SetWelcomeIntent(val *string)
	WelcomeIntentInput() *string
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
	PutInvokedBy(value *ChimeAppInstanceBotConfigurationLexInvokedBy)
	ResetInvokedBy()
	ResetRespondsTo()
	ResetWelcomeIntent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimeAppInstanceBotConfigurationLexOutputReference
type jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) InvokedBy() ChimeAppInstanceBotConfigurationLexInvokedByOutputReference {
	var returns ChimeAppInstanceBotConfigurationLexInvokedByOutputReference
	_jsii_.Get(
		j,
		"invokedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) InvokedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invokedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) LexBotAliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lexBotAliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) LexBotAliasArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lexBotAliasArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) LocaleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) LocaleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) RespondsTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) RespondsToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) WelcomeIntent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"welcomeIntent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) WelcomeIntentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"welcomeIntentInput",
		&returns,
	)
	return returns
}


func NewChimeAppInstanceBotConfigurationLexOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ChimeAppInstanceBotConfigurationLexOutputReference {
	_init_.Initialize()

	if err := validateNewChimeAppInstanceBotConfigurationLexOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeAppInstanceBot.ChimeAppInstanceBotConfigurationLexOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChimeAppInstanceBotConfigurationLexOutputReference_Override(c ChimeAppInstanceBotConfigurationLexOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.chimeAppInstanceBot.ChimeAppInstanceBotConfigurationLexOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetLexBotAliasArn(val *string) {
	if err := j.validateSetLexBotAliasArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lexBotAliasArn",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetLocaleId(val *string) {
	if err := j.validateSetLocaleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localeId",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetRespondsTo(val *string) {
	if err := j.validateSetRespondsToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respondsTo",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference)SetWelcomeIntent(val *string) {
	if err := j.validateSetWelcomeIntentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"welcomeIntent",
		val,
	)
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) PutInvokedBy(value *ChimeAppInstanceBotConfigurationLexInvokedBy) {
	if err := c.validatePutInvokedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInvokedBy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ResetInvokedBy() {
	_jsii_.InvokeVoid(
		c,
		"resetInvokedBy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ResetRespondsTo() {
	_jsii_.InvokeVoid(
		c,
		"resetRespondsTo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ResetWelcomeIntent() {
	_jsii_.InvokeVoid(
		c,
		"resetWelcomeIntent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChimeAppInstanceBotConfigurationLexOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

