// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicDataSetsCalculatedFieldsOutputReference interface {
	cdktn.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
	AllowedAggregations() *[]*string
	SetAllowedAggregations(val *[]*string)
	AllowedAggregationsInput() *[]*string
	CalculatedFieldDescription() *string
	SetCalculatedFieldDescription(val *string)
	CalculatedFieldDescriptionInput() *string
	CalculatedFieldName() *string
	SetCalculatedFieldName(val *string)
	CalculatedFieldNameInput() *string
	CalculatedFieldSynonyms() *[]*string
	SetCalculatedFieldSynonyms(val *[]*string)
	CalculatedFieldSynonymsInput() *[]*string
	CellValueSynonyms() QuicksightTopicDataSetsCalculatedFieldsCellValueSynonymsList
	CellValueSynonymsInput() interface{}
	ColumnDataRole() *string
	SetColumnDataRole(val *string)
	ColumnDataRoleInput() *string
	ComparativeOrder() QuicksightTopicDataSetsCalculatedFieldsComparativeOrderOutputReference
	ComparativeOrderInput() interface{}
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
	DefaultFormatting() QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingOutputReference
	DefaultFormattingInput() interface{}
	DisableIndexing() interface{}
	SetDisableIndexing(val interface{})
	DisableIndexingInput() interface{}
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsIncludedInTopic() interface{}
	SetIsIncludedInTopic(val interface{})
	IsIncludedInTopicInput() interface{}
	NeverAggregateInFilter() interface{}
	SetNeverAggregateInFilter(val interface{})
	NeverAggregateInFilterInput() interface{}
	NonAdditive() interface{}
	SetNonAdditive(val interface{})
	NonAdditiveInput() interface{}
	NotAllowedAggregations() *[]*string
	SetNotAllowedAggregations(val *[]*string)
	NotAllowedAggregationsInput() *[]*string
	SemanticType() QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference
	SemanticTypeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeGranularity() *string
	SetTimeGranularity(val *string)
	TimeGranularityInput() *string
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
	PutCellValueSynonyms(value interface{})
	PutComparativeOrder(value *QuicksightTopicDataSetsCalculatedFieldsComparativeOrder)
	PutDefaultFormatting(value *QuicksightTopicDataSetsCalculatedFieldsDefaultFormatting)
	PutSemanticType(value *QuicksightTopicDataSetsCalculatedFieldsSemanticType)
	ResetAggregation()
	ResetAllowedAggregations()
	ResetCalculatedFieldDescription()
	ResetCalculatedFieldName()
	ResetCalculatedFieldSynonyms()
	ResetCellValueSynonyms()
	ResetColumnDataRole()
	ResetComparativeOrder()
	ResetDefaultFormatting()
	ResetDisableIndexing()
	ResetExpression()
	ResetIsIncludedInTopic()
	ResetNeverAggregateInFilter()
	ResetNonAdditive()
	ResetNotAllowedAggregations()
	ResetSemanticType()
	ResetTimeGranularity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsCalculatedFieldsOutputReference
type jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) AllowedAggregations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) AllowedAggregationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CalculatedFieldDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calculatedFieldDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CalculatedFieldDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calculatedFieldDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CalculatedFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calculatedFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CalculatedFieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calculatedFieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CalculatedFieldSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"calculatedFieldSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CalculatedFieldSynonymsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"calculatedFieldSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CellValueSynonyms() QuicksightTopicDataSetsCalculatedFieldsCellValueSynonymsList {
	var returns QuicksightTopicDataSetsCalculatedFieldsCellValueSynonymsList
	_jsii_.Get(
		j,
		"cellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CellValueSynonymsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cellValueSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ColumnDataRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDataRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ColumnDataRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDataRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ComparativeOrder() QuicksightTopicDataSetsCalculatedFieldsComparativeOrderOutputReference {
	var returns QuicksightTopicDataSetsCalculatedFieldsComparativeOrderOutputReference
	_jsii_.Get(
		j,
		"comparativeOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ComparativeOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"comparativeOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) DefaultFormatting() QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingOutputReference {
	var returns QuicksightTopicDataSetsCalculatedFieldsDefaultFormattingOutputReference
	_jsii_.Get(
		j,
		"defaultFormatting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) DefaultFormattingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultFormattingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) DisableIndexing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIndexing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) DisableIndexingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIndexingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) IsIncludedInTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncludedInTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) IsIncludedInTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncludedInTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) NeverAggregateInFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neverAggregateInFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) NeverAggregateInFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neverAggregateInFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) NonAdditive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonAdditive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) NonAdditiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonAdditiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) NotAllowedAggregations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notAllowedAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) NotAllowedAggregationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notAllowedAggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) SemanticType() QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference {
	var returns QuicksightTopicDataSetsCalculatedFieldsSemanticTypeOutputReference
	_jsii_.Get(
		j,
		"semanticType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) SemanticTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) TimeGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) TimeGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGranularityInput",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsCalculatedFieldsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightTopicDataSetsCalculatedFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsCalculatedFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsCalculatedFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsCalculatedFieldsOutputReference_Override(q QuicksightTopicDataSetsCalculatedFieldsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsCalculatedFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetAllowedAggregations(val *[]*string) {
	if err := j.validateSetAllowedAggregationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAggregations",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetCalculatedFieldDescription(val *string) {
	if err := j.validateSetCalculatedFieldDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calculatedFieldDescription",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetCalculatedFieldName(val *string) {
	if err := j.validateSetCalculatedFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calculatedFieldName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetCalculatedFieldSynonyms(val *[]*string) {
	if err := j.validateSetCalculatedFieldSynonymsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calculatedFieldSynonyms",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetColumnDataRole(val *string) {
	if err := j.validateSetColumnDataRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnDataRole",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetDisableIndexing(val interface{}) {
	if err := j.validateSetDisableIndexingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableIndexing",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetIsIncludedInTopic(val interface{}) {
	if err := j.validateSetIsIncludedInTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIncludedInTopic",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetNeverAggregateInFilter(val interface{}) {
	if err := j.validateSetNeverAggregateInFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"neverAggregateInFilter",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetNonAdditive(val interface{}) {
	if err := j.validateSetNonAdditiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonAdditive",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetNotAllowedAggregations(val *[]*string) {
	if err := j.validateSetNotAllowedAggregationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAllowedAggregations",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference)SetTimeGranularity(val *string) {
	if err := j.validateSetTimeGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeGranularity",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) PutCellValueSynonyms(value interface{}) {
	if err := q.validatePutCellValueSynonymsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCellValueSynonyms",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) PutComparativeOrder(value *QuicksightTopicDataSetsCalculatedFieldsComparativeOrder) {
	if err := q.validatePutComparativeOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putComparativeOrder",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) PutDefaultFormatting(value *QuicksightTopicDataSetsCalculatedFieldsDefaultFormatting) {
	if err := q.validatePutDefaultFormattingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDefaultFormatting",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) PutSemanticType(value *QuicksightTopicDataSetsCalculatedFieldsSemanticType) {
	if err := q.validatePutSemanticTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSemanticType",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetAggregation() {
	_jsii_.InvokeVoid(
		q,
		"resetAggregation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetAllowedAggregations() {
	_jsii_.InvokeVoid(
		q,
		"resetAllowedAggregations",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetCalculatedFieldDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetCalculatedFieldDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetCalculatedFieldName() {
	_jsii_.InvokeVoid(
		q,
		"resetCalculatedFieldName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetCalculatedFieldSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetCalculatedFieldSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetCellValueSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetCellValueSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetColumnDataRole() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnDataRole",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetComparativeOrder() {
	_jsii_.InvokeVoid(
		q,
		"resetComparativeOrder",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetDefaultFormatting() {
	_jsii_.InvokeVoid(
		q,
		"resetDefaultFormatting",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetDisableIndexing() {
	_jsii_.InvokeVoid(
		q,
		"resetDisableIndexing",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetExpression() {
	_jsii_.InvokeVoid(
		q,
		"resetExpression",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetIsIncludedInTopic() {
	_jsii_.InvokeVoid(
		q,
		"resetIsIncludedInTopic",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetNeverAggregateInFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetNeverAggregateInFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetNonAdditive() {
	_jsii_.InvokeVoid(
		q,
		"resetNonAdditive",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetNotAllowedAggregations() {
	_jsii_.InvokeVoid(
		q,
		"resetNotAllowedAggregations",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetSemanticType() {
	_jsii_.InvokeVoid(
		q,
		"resetSemanticType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ResetTimeGranularity() {
	_jsii_.InvokeVoid(
		q,
		"resetTimeGranularity",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsCalculatedFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

