// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsintermediatetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsintermediatetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference interface {
	cdktn.ComplexObject
	AllowedAggregateExpressionType() *string
	SetAllowedAggregateExpressionType(val *string)
	AllowedAggregateExpressionTypeInput() *string
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
	IdentityColumns() *[]*string
	SetIdentityColumns(val *[]*string)
	IdentityColumnsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumIdentityCount() *float64
	SetMinimumIdentityCount(val *float64)
	MinimumIdentityCountInput() *float64
	OutputColumnThresholds() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholdsList
	OutputColumnThresholdsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutOutputColumnThresholds(value interface{})
	ResetAllowedAggregateExpressionType()
	ResetIdentityColumns()
	ResetMinimumIdentityCount()
	ResetOutputColumnThresholds()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference
type jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) AllowedAggregateExpressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedAggregateExpressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) AllowedAggregateExpressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedAggregateExpressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) IdentityColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) IdentityColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) MinimumIdentityCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumIdentityCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) MinimumIdentityCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumIdentityCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) OutputColumnThresholds() CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholdsList {
	var returns CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholdsList
	_jsii_.Get(
		j,
		"outputColumnThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) OutputColumnThresholdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputColumnThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsIntermediateTable.CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference_Override(c CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsIntermediateTable.CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetAllowedAggregateExpressionType(val *string) {
	if err := j.validateSetAllowedAggregateExpressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAggregateExpressionType",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetIdentityColumns(val *[]*string) {
	if err := j.validateSetIdentityColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityColumns",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetMinimumIdentityCount(val *float64) {
	if err := j.validateSetMinimumIdentityCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumIdentityCount",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) PutOutputColumnThresholds(value interface{}) {
	if err := c.validatePutOutputColumnThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOutputColumnThresholds",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ResetAllowedAggregateExpressionType() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedAggregateExpressionType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ResetIdentityColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ResetMinimumIdentityCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumIdentityCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ResetOutputColumnThresholds() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputColumnThresholds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsIntermediateTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

