// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsprivacybudgettemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsprivacybudgettemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsPrivacyBudgetTemplateParametersOutputReference interface {
	cdktn.ComplexObject
	BudgetParameters() CleanroomsPrivacyBudgetTemplateParametersBudgetParametersList
	BudgetParametersInput() interface{}
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
	Epsilon() *float64
	SetEpsilon(val *float64)
	EpsilonInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResourceArn() *string
	SetResourceArn(val *string)
	ResourceArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UsersNoisePerQuery() *float64
	SetUsersNoisePerQuery(val *float64)
	UsersNoisePerQueryInput() *float64
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
	PutBudgetParameters(value interface{})
	ResetBudgetParameters()
	ResetEpsilon()
	ResetResourceArn()
	ResetUsersNoisePerQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsPrivacyBudgetTemplateParametersOutputReference
type jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) BudgetParameters() CleanroomsPrivacyBudgetTemplateParametersBudgetParametersList {
	var returns CleanroomsPrivacyBudgetTemplateParametersBudgetParametersList
	_jsii_.Get(
		j,
		"budgetParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) BudgetParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"budgetParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) Epsilon() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epsilon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) EpsilonInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epsilonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) UsersNoisePerQuery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usersNoisePerQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) UsersNoisePerQueryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usersNoisePerQueryInput",
		&returns,
	)
	return returns
}


func NewCleanroomsPrivacyBudgetTemplateParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsPrivacyBudgetTemplateParametersOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsPrivacyBudgetTemplateParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsPrivacyBudgetTemplate.CleanroomsPrivacyBudgetTemplateParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsPrivacyBudgetTemplateParametersOutputReference_Override(c CleanroomsPrivacyBudgetTemplateParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsPrivacyBudgetTemplate.CleanroomsPrivacyBudgetTemplateParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetEpsilon(val *float64) {
	if err := j.validateSetEpsilonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"epsilon",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetResourceArn(val *string) {
	if err := j.validateSetResourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference)SetUsersNoisePerQuery(val *float64) {
	if err := j.validateSetUsersNoisePerQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usersNoisePerQuery",
		val,
	)
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) PutBudgetParameters(value interface{}) {
	if err := c.validatePutBudgetParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBudgetParameters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ResetBudgetParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetBudgetParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ResetEpsilon() {
	_jsii_.InvokeVoid(
		c,
		"resetEpsilon",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ResetResourceArn() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ResetUsersNoisePerQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetUsersNoisePerQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsPrivacyBudgetTemplateParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

