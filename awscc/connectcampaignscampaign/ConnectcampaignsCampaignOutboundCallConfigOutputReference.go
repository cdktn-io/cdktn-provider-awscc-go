// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignscampaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectcampaignsCampaignOutboundCallConfigOutputReference interface {
	cdktn.ComplexObject
	AnswerMachineDetectionConfig() ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference
	AnswerMachineDetectionConfigInput() interface{}
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
	ConnectContactFlowArn() *string
	SetConnectContactFlowArn(val *string)
	ConnectContactFlowArnInput() *string
	ConnectQueueArn() *string
	SetConnectQueueArn(val *string)
	ConnectQueueArnInput() *string
	ConnectSourcePhoneNumber() *string
	SetConnectSourcePhoneNumber(val *string)
	ConnectSourcePhoneNumberInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutAnswerMachineDetectionConfig(value *ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfig)
	ResetAnswerMachineDetectionConfig()
	ResetConnectQueueArn()
	ResetConnectSourcePhoneNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectcampaignsCampaignOutboundCallConfigOutputReference
type jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) AnswerMachineDetectionConfig() ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference {
	var returns ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference
	_jsii_.Get(
		j,
		"answerMachineDetectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) AnswerMachineDetectionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"answerMachineDetectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ConnectContactFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectContactFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ConnectContactFlowArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectContactFlowArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ConnectQueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectQueueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ConnectQueueArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectQueueArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ConnectSourcePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSourcePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ConnectSourcePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectSourcePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsCampaignOutboundCallConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectcampaignsCampaignOutboundCallConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsCampaignOutboundCallConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsCampaign.ConnectcampaignsCampaignOutboundCallConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsCampaignOutboundCallConfigOutputReference_Override(c ConnectcampaignsCampaignOutboundCallConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsCampaign.ConnectcampaignsCampaignOutboundCallConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetConnectContactFlowArn(val *string) {
	if err := j.validateSetConnectContactFlowArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectContactFlowArn",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetConnectQueueArn(val *string) {
	if err := j.validateSetConnectQueueArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectQueueArn",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetConnectSourcePhoneNumber(val *string) {
	if err := j.validateSetConnectSourcePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectSourcePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) PutAnswerMachineDetectionConfig(value *ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfig) {
	if err := c.validatePutAnswerMachineDetectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnswerMachineDetectionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ResetAnswerMachineDetectionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAnswerMachineDetectionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ResetConnectQueueArn() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectQueueArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ResetConnectSourcePhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectSourcePhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

