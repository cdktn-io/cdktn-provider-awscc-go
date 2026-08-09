// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilescalculatedattributedefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/customerprofilescalculatedattributedefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference interface {
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
	ObjectCount() *float64
	SetObjectCount(val *float64)
	ObjectCountInput() *float64
	Range() CustomerprofilesCalculatedAttributeDefinitionConditionsRangeOutputReference
	RangeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Threshold() CustomerprofilesCalculatedAttributeDefinitionConditionsThresholdOutputReference
	ThresholdInput() interface{}
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
	PutRange(value *CustomerprofilesCalculatedAttributeDefinitionConditionsRange)
	PutThreshold(value *CustomerprofilesCalculatedAttributeDefinitionConditionsThreshold)
	ResetObjectCount()
	ResetRange()
	ResetThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference
type jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ObjectCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"objectCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ObjectCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"objectCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) Range() CustomerprofilesCalculatedAttributeDefinitionConditionsRangeOutputReference {
	var returns CustomerprofilesCalculatedAttributeDefinitionConditionsRangeOutputReference
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) RangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) Threshold() CustomerprofilesCalculatedAttributeDefinitionConditionsThresholdOutputReference {
	var returns CustomerprofilesCalculatedAttributeDefinitionConditionsThresholdOutputReference
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ThresholdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}


func NewCustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesCalculatedAttributeDefinitionConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesCalculatedAttributeDefinition.CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference_Override(c CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.customerprofilesCalculatedAttributeDefinition.CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference)SetObjectCount(val *float64) {
	if err := j.validateSetObjectCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectCount",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) PutRange(value *CustomerprofilesCalculatedAttributeDefinitionConditionsRange) {
	if err := c.validatePutRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRange",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) PutThreshold(value *CustomerprofilesCalculatedAttributeDefinitionConditionsThreshold) {
	if err := c.validatePutThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putThreshold",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ResetObjectCount() {
	_jsii_.InvokeVoid(
		c,
		"resetObjectCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		c,
		"resetRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesCalculatedAttributeDefinitionConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

