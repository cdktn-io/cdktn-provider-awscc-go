// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutyfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/guarddutyfilter/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GuarddutyFilterFindingCriteriaCriterionOutputReference interface {
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
	Eq() *[]*string
	SetEq(val *[]*string)
	EqInput() *[]*string
	EqualTo() *[]*string
	SetEqualTo(val *[]*string)
	EqualToInput() *[]*string
	// Experimental.
	Fqn() *string
	GreaterThan() *float64
	SetGreaterThan(val *float64)
	GreaterThanInput() *float64
	GreaterThanOrEqual() *float64
	SetGreaterThanOrEqual(val *float64)
	GreaterThanOrEqualInput() *float64
	Gt() *float64
	SetGt(val *float64)
	Gte() *float64
	SetGte(val *float64)
	GteInput() *float64
	GtInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LessThan() *float64
	SetLessThan(val *float64)
	LessThanInput() *float64
	LessThanOrEqual() *float64
	SetLessThanOrEqual(val *float64)
	LessThanOrEqualInput() *float64
	Lt() *float64
	SetLt(val *float64)
	Lte() *float64
	SetLte(val *float64)
	LteInput() *float64
	LtInput() *float64
	Matches() *[]*string
	SetMatches(val *[]*string)
	MatchesInput() *[]*string
	Neq() *[]*string
	SetNeq(val *[]*string)
	NeqInput() *[]*string
	NotEquals() *[]*string
	SetNotEquals(val *[]*string)
	NotEqualsInput() *[]*string
	NotMatches() *[]*string
	SetNotMatches(val *[]*string)
	NotMatchesInput() *[]*string
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
	ResetEq()
	ResetEqualTo()
	ResetGreaterThan()
	ResetGreaterThanOrEqual()
	ResetGt()
	ResetGte()
	ResetLessThan()
	ResetLessThanOrEqual()
	ResetLt()
	ResetLte()
	ResetMatches()
	ResetNeq()
	ResetNotEquals()
	ResetNotMatches()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuarddutyFilterFindingCriteriaCriterionOutputReference
type jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Eq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) EqInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanOrEqual() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanOrEqualInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"greaterThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Gt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Gte() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanOrEqual() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanOrEqualInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lessThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Lt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Lte() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Matches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) MatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Neq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NeqInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NotMatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NotMatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGuarddutyFilterFindingCriteriaCriterionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) GuarddutyFilterFindingCriteriaCriterionOutputReference {
	_init_.Initialize()

	if err := validateNewGuarddutyFilterFindingCriteriaCriterionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.guarddutyFilter.GuarddutyFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewGuarddutyFilterFindingCriteriaCriterionOutputReference_Override(g GuarddutyFilterFindingCriteriaCriterionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.guarddutyFilter.GuarddutyFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		g,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetEq(val *[]*string) {
	if err := j.validateSetEqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eq",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetGreaterThan(val *float64) {
	if err := j.validateSetGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThan",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetGreaterThanOrEqual(val *float64) {
	if err := j.validateSetGreaterThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetGt(val *float64) {
	if err := j.validateSetGtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gt",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetGte(val *float64) {
	if err := j.validateSetGteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gte",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetLessThan(val *float64) {
	if err := j.validateSetLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThan",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetLessThanOrEqual(val *float64) {
	if err := j.validateSetLessThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetLt(val *float64) {
	if err := j.validateSetLtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lt",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetLte(val *float64) {
	if err := j.validateSetLteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lte",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetMatches(val *[]*string) {
	if err := j.validateSetMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matches",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetNeq(val *[]*string) {
	if err := j.validateSetNeqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"neq",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetNotMatches(val *[]*string) {
	if err := j.validateSetNotMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notMatches",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetEq() {
	_jsii_.InvokeVoid(
		g,
		"resetEq",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		g,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetGreaterThan() {
	_jsii_.InvokeVoid(
		g,
		"resetGreaterThan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetGreaterThanOrEqual() {
	_jsii_.InvokeVoid(
		g,
		"resetGreaterThanOrEqual",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetGt() {
	_jsii_.InvokeVoid(
		g,
		"resetGt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetGte() {
	_jsii_.InvokeVoid(
		g,
		"resetGte",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetLessThan() {
	_jsii_.InvokeVoid(
		g,
		"resetLessThan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetLessThanOrEqual() {
	_jsii_.InvokeVoid(
		g,
		"resetLessThanOrEqual",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetLt() {
	_jsii_.InvokeVoid(
		g,
		"resetLt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetLte() {
	_jsii_.InvokeVoid(
		g,
		"resetLte",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetMatches() {
	_jsii_.InvokeVoid(
		g,
		"resetMatches",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetNeq() {
	_jsii_.InvokeVoid(
		g,
		"resetNeq",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		g,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetNotMatches() {
	_jsii_.InvokeVoid(
		g,
		"resetNotMatches",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

