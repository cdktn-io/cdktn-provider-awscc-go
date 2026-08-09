// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cleanroomsconfiguredtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference interface {
	cdktn.ComplexObject
	AdditionalAnalyses() *string
	SetAdditionalAnalyses(val *string)
	AdditionalAnalysesInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JoinColumns() *[]*string
	SetJoinColumns(val *[]*string)
	JoinColumnsInput() *[]*string
	ListColumns() *[]*string
	SetListColumns(val *[]*string)
	ListColumnsInput() *[]*string
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
	ResetAdditionalAnalyses()
	ResetAllowedJoinOperators()
	ResetJoinColumns()
	ResetListColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference
type jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) AdditionalAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) AdditionalAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) AllowedJoinOperators() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedJoinOperators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) AllowedJoinOperatorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedJoinOperatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) JoinColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"joinColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) JoinColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"joinColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ListColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ListColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference {
	_init_.Initialize()

	if err := validateNewCleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference_Override(c CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cleanroomsConfiguredTable.CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetAdditionalAnalyses(val *string) {
	if err := j.validateSetAdditionalAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalAnalyses",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetAllowedJoinOperators(val *[]*string) {
	if err := j.validateSetAllowedJoinOperatorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedJoinOperators",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetJoinColumns(val *[]*string) {
	if err := j.validateSetJoinColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"joinColumns",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetListColumns(val *[]*string) {
	if err := j.validateSetListColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listColumns",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ResetAdditionalAnalyses() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalAnalyses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ResetAllowedJoinOperators() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedJoinOperators",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ResetJoinColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetJoinColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ResetListColumns() {
	_jsii_.InvokeVoid(
		c,
		"resetListColumns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

