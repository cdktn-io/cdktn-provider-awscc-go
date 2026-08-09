// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casescaserule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/casescaserule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference interface {
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
	OperandOne() CasesCaseRuleRuleHiddenConditionsNotEqualToOperandOneOutputReference
	OperandOneInput() interface{}
	OperandTwo() CasesCaseRuleRuleHiddenConditionsNotEqualToOperandTwoOutputReference
	OperandTwoInput() interface{}
	Result() interface{}
	SetResult(val interface{})
	ResultInput() interface{}
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
	PutOperandOne(value *CasesCaseRuleRuleHiddenConditionsNotEqualToOperandOne)
	PutOperandTwo(value *CasesCaseRuleRuleHiddenConditionsNotEqualToOperandTwo)
	ResetOperandOne()
	ResetOperandTwo()
	ResetResult()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference
type jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) OperandOne() CasesCaseRuleRuleHiddenConditionsNotEqualToOperandOneOutputReference {
	var returns CasesCaseRuleRuleHiddenConditionsNotEqualToOperandOneOutputReference
	_jsii_.Get(
		j,
		"operandOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) OperandOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operandOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) OperandTwo() CasesCaseRuleRuleHiddenConditionsNotEqualToOperandTwoOutputReference {
	var returns CasesCaseRuleRuleHiddenConditionsNotEqualToOperandTwoOutputReference
	_jsii_.Get(
		j,
		"operandTwo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) OperandTwoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operandTwoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) Result() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ResultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference {
	_init_.Initialize()

	if err := validateNewCasesCaseRuleRuleHiddenConditionsNotEqualToOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.casesCaseRule.CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference_Override(c CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.casesCaseRule.CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference)SetResult(val interface{}) {
	if err := j.validateSetResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) PutOperandOne(value *CasesCaseRuleRuleHiddenConditionsNotEqualToOperandOne) {
	if err := c.validatePutOperandOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOperandOne",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) PutOperandTwo(value *CasesCaseRuleRuleHiddenConditionsNotEqualToOperandTwo) {
	if err := c.validatePutOperandTwoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOperandTwo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ResetOperandOne() {
	_jsii_.InvokeVoid(
		c,
		"resetOperandOne",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ResetOperandTwo() {
	_jsii_.InvokeVoid(
		c,
		"resetOperandTwo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ResetResult() {
	_jsii_.InvokeVoid(
		c,
		"resetResult",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CasesCaseRuleRuleHiddenConditionsNotEqualToOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

