// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsconfiguredtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference interface {
	cdktn.ComplexObject
	AdditionalAnalyses() *string
	SetAdditionalAnalyses(val *string)
	AdditionalAnalysesInput() *string
	AggregateColumns() CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList
	AggregateColumnsInput() interface{}
	AllowedJoinOperators() *[]*string
	SetAllowedJoinOperators(val *[]*string)
	AllowedJoinOperatorsInput() *[]*string
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
	DimensionColumns() *[]*string
	SetDimensionColumns(val *[]*string)
	DimensionColumnsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JoinColumns() *[]*string
	SetJoinColumns(val *[]*string)
	JoinColumnsInput() *[]*string
	JoinRequired() *string
	SetJoinRequired(val *string)
	JoinRequiredInput() *string
	OutputConstraints() CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputConstraintsList
	OutputConstraintsInput() interface{}
	ScalarFunctions() *[]*string
	SetScalarFunctions(val *[]*string)
	ScalarFunctionsInput() *[]*string
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
	PutAggregateColumns(value interface{})
	PutOutputConstraints(value interface{})
	ResetAdditionalAnalyses()
	ResetAggregateColumns()
	ResetAllowedJoinOperators()
	ResetDimensionColumns()
	ResetJoinColumns()
	ResetJoinRequired()
	ResetOutputConstraints()
	ResetScalarFunctions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference
type jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) AdditionalAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) AdditionalAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) AggregateColumns() CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList {
	var returns CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumnsList
	_jsii_.Get(
		j,
		"aggregateColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) AggregateColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregateColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) AllowedJoinOperators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedJoinOperators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) AllowedJoinOperatorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedJoinOperatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) DimensionColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dimensionColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) DimensionColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dimensionColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) JoinColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"joinColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) JoinColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"joinColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) JoinRequired() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) JoinRequiredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) OutputConstraints() CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputConstraintsList {
	var returns CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputConstraintsList
	_jsii_.Get(
		j,
		"outputConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) OutputConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ScalarFunctions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scalarFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ScalarFunctionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scalarFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference_Override(c CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetAdditionalAnalyses(val *string) {
	if err := j.validateSetAdditionalAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalAnalyses",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetAllowedJoinOperators(val *[]*string) {
	if err := j.validateSetAllowedJoinOperatorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedJoinOperators",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetDimensionColumns(val *[]*string) {
	if err := j.validateSetDimensionColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionColumns",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetJoinColumns(val *[]*string) {
	if err := j.validateSetJoinColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinColumns",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetJoinRequired(val *string) {
	if err := j.validateSetJoinRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinRequired",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetScalarFunctions(val *[]*string) {
	if err := j.validateSetScalarFunctionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalarFunctions",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) PutAggregateColumns(value interface{}) {
	if err := c.validatePutAggregateColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAggregateColumns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) PutOutputConstraints(value interface{}) {
	if err := c.validatePutOutputConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOutputConstraints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetAdditionalAnalyses() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalAnalyses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetAggregateColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetAggregateColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetAllowedJoinOperators() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedJoinOperators",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetDimensionColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetDimensionColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetJoinColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetJoinColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetJoinRequired() {
	_jsii_.InvokeVoid(
		c,
		"resetJoinRequired",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetOutputConstraints() {
	_jsii_.InvokeVoid(
		c,
		"resetOutputConstraints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ResetScalarFunctions() {
	_jsii_.InvokeVoid(
		c,
		"resetScalarFunctions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

