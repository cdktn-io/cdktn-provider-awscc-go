// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectuser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference interface {
	cdktn.ComplexObject
	AfterContactWorkMode() *string
	SetAfterContactWorkMode(val *string)
	AfterContactWorkModeInput() *string
	AfterContactWorkTimeLimit() *float64
	SetAfterContactWorkTimeLimit(val *float64)
	AfterContactWorkTimeLimitInput() *float64
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
	ResetAfterContactWorkMode()
	ResetAfterContactWorkTimeLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference
type jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) AfterContactWorkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterContactWorkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) AfterContactWorkModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterContactWorkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) AfterContactWorkTimeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"afterContactWorkTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) AfterContactWorkTimeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"afterContactWorkTimeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectUser.ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference_Override(c ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectUser.ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetAfterContactWorkMode(val *string) {
	if err := j.validateSetAfterContactWorkModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterContactWorkMode",
		val,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetAfterContactWorkTimeLimit(val *float64) {
	if err := j.validateSetAfterContactWorkTimeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterContactWorkTimeLimit",
		val,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) ResetAfterContactWorkMode() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterContactWorkMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) ResetAfterContactWorkTimeLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetAfterContactWorkTimeLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectUserAfterContactWorkConfigsAgentFirstCallbackAfterContactWorkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

