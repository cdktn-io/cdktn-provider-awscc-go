// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecttasktemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connecttasktemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectTaskTemplateConstraintsOutputReference interface {
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
	InvisibleFields() ConnectTaskTemplateConstraintsInvisibleFieldsList
	InvisibleFieldsInput() interface{}
	ReadOnlyFields() ConnectTaskTemplateConstraintsReadOnlyFieldsList
	ReadOnlyFieldsInput() interface{}
	RequiredFields() ConnectTaskTemplateConstraintsRequiredFieldsList
	RequiredFieldsInput() interface{}
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
	PutInvisibleFields(value interface{})
	PutReadOnlyFields(value interface{})
	PutRequiredFields(value interface{})
	ResetInvisibleFields()
	ResetReadOnlyFields()
	ResetRequiredFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectTaskTemplateConstraintsOutputReference
type jsiiProxy_ConnectTaskTemplateConstraintsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) InvisibleFields() ConnectTaskTemplateConstraintsInvisibleFieldsList {
	var returns ConnectTaskTemplateConstraintsInvisibleFieldsList
	_jsii_.Get(
		j,
		"invisibleFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) InvisibleFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invisibleFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ReadOnlyFields() ConnectTaskTemplateConstraintsReadOnlyFieldsList {
	var returns ConnectTaskTemplateConstraintsReadOnlyFieldsList
	_jsii_.Get(
		j,
		"readOnlyFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ReadOnlyFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) RequiredFields() ConnectTaskTemplateConstraintsRequiredFieldsList {
	var returns ConnectTaskTemplateConstraintsRequiredFieldsList
	_jsii_.Get(
		j,
		"requiredFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) RequiredFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectTaskTemplateConstraintsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectTaskTemplateConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewConnectTaskTemplateConstraintsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectTaskTemplateConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectTaskTemplate.ConnectTaskTemplateConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectTaskTemplateConstraintsOutputReference_Override(c ConnectTaskTemplateConstraintsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectTaskTemplate.ConnectTaskTemplateConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) PutInvisibleFields(value interface{}) {
	if err := c.validatePutInvisibleFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInvisibleFields",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) PutReadOnlyFields(value interface{}) {
	if err := c.validatePutReadOnlyFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReadOnlyFields",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) PutRequiredFields(value interface{}) {
	if err := c.validatePutRequiredFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequiredFields",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ResetInvisibleFields() {
	_jsii_.InvokeVoid(
		c,
		"resetInvisibleFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ResetReadOnlyFields() {
	_jsii_.InvokeVoid(
		c,
		"resetReadOnlyFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ResetRequiredFields() {
	_jsii_.InvokeVoid(
		c,
		"resetRequiredFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectTaskTemplateConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

