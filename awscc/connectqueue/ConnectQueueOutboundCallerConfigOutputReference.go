// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectqueue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectQueueOutboundCallerConfigOutputReference interface {
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
	OutboundCallerIdName() *string
	SetOutboundCallerIdName(val *string)
	OutboundCallerIdNameInput() *string
	OutboundCallerIdNumberArn() *string
	SetOutboundCallerIdNumberArn(val *string)
	OutboundCallerIdNumberArnInput() *string
	OutboundFlowArn() *string
	SetOutboundFlowArn(val *string)
	OutboundFlowArnInput() *string
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
	ResetOutboundCallerIdName()
	ResetOutboundCallerIdNumberArn()
	ResetOutboundFlowArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectQueueOutboundCallerConfigOutputReference
type jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdNumberArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNumberArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundCallerIdNumberArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCallerIdNumberArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) OutboundFlowArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundFlowArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectQueueOutboundCallerConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectQueueOutboundCallerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectQueueOutboundCallerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectQueue.ConnectQueueOutboundCallerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectQueueOutboundCallerConfigOutputReference_Override(c ConnectQueueOutboundCallerConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectQueue.ConnectQueueOutboundCallerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetOutboundCallerIdName(val *string) {
	if err := j.validateSetOutboundCallerIdNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallerIdName",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetOutboundCallerIdNumberArn(val *string) {
	if err := j.validateSetOutboundCallerIdNumberArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallerIdNumberArn",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetOutboundFlowArn(val *string) {
	if err := j.validateSetOutboundFlowArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundFlowArn",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ResetOutboundCallerIdName() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundCallerIdName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ResetOutboundCallerIdNumberArn() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundCallerIdNumberArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ResetOutboundFlowArn() {
	_jsii_.InvokeVoid(
		c,
		"resetOutboundFlowArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectQueueOutboundCallerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

