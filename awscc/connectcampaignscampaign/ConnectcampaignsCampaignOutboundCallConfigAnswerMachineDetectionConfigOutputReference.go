// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectcampaignscampaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference interface {
	cdktn.ComplexObject
	AwaitAnswerMachinePrompt() interface{}
	SetAwaitAnswerMachinePrompt(val interface{})
	AwaitAnswerMachinePromptInput() interface{}
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
	EnableAnswerMachineDetection() interface{}
	SetEnableAnswerMachineDetection(val interface{})
	EnableAnswerMachineDetectionInput() interface{}
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
	ResetAwaitAnswerMachinePrompt()
	ResetEnableAnswerMachineDetection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference
type jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) AwaitAnswerMachinePrompt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awaitAnswerMachinePrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) AwaitAnswerMachinePromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awaitAnswerMachinePromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) EnableAnswerMachineDetection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAnswerMachineDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) EnableAnswerMachineDetectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAnswerMachineDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsCampaign.ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference_Override(c ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectcampaignsCampaign.ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetAwaitAnswerMachinePrompt(val interface{}) {
	if err := j.validateSetAwaitAnswerMachinePromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awaitAnswerMachinePrompt",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetEnableAnswerMachineDetection(val interface{}) {
	if err := j.validateSetEnableAnswerMachineDetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAnswerMachineDetection",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) ResetAwaitAnswerMachinePrompt() {
	_jsii_.InvokeVoid(
		c,
		"resetAwaitAnswerMachinePrompt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) ResetEnableAnswerMachineDetection() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAnswerMachineDetection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

