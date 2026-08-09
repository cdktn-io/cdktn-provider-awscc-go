// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicDataSetsColumnsOutputReference interface {
	cdktn.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
	AllowedAggregations() *[]*string
	SetAllowedAggregations(val *[]*string)
	AllowedAggregationsInput() *[]*string
	CellValueSynonyms() QuicksightTopicDataSetsColumnsCellValueSynonymsList
	CellValueSynonymsInput() interface{}
	ColumnDataRole() *string
	SetColumnDataRole(val *string)
	ColumnDataRoleInput() *string
	ColumnDescription() *string
	SetColumnDescription(val *string)
	ColumnDescriptionInput() *string
	ColumnFriendlyName() *string
	SetColumnFriendlyName(val *string)
	ColumnFriendlyNameInput() *string
	ColumnName() *string
	SetColumnName(val *string)
	ColumnNameInput() *string
	ColumnSynonyms() *[]*string
	SetColumnSynonyms(val *[]*string)
	ColumnSynonymsInput() *[]*string
	ComparativeOrder() QuicksightTopicDataSetsColumnsComparativeOrderOutputReference
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
	DefaultFormatting() QuicksightTopicDataSetsColumnsDefaultFormattingOutputReference
	DefaultFormattingInput() interface{}
	DisableIndexing() interface{}
	SetDisableIndexing(val interface{})
	DisableIndexingInput() interface{}
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
	SemanticType() QuicksightTopicDataSetsColumnsSemanticTypeOutputReference
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
	PutComparativeOrder(value *QuicksightTopicDataSetsColumnsComparativeOrder)
	PutDefaultFormatting(value *QuicksightTopicDataSetsColumnsDefaultFormatting)
	PutSemanticType(value *QuicksightTopicDataSetsColumnsSemanticType)
	ResetAggregation()
	ResetAllowedAggregations()
	ResetCellValueSynonyms()
	ResetColumnDataRole()
	ResetColumnDescription()
	ResetColumnFriendlyName()
	ResetColumnName()
	ResetColumnSynonyms()
	ResetComparativeOrder()
	ResetDefaultFormatting()
	ResetDisableIndexing()
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

// The jsii proxy struct for QuicksightTopicDataSetsColumnsOutputReference
type jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) AllowedAggregations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) AllowedAggregationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) CellValueSynonyms() QuicksightTopicDataSetsColumnsCellValueSynonymsList {
	var returns QuicksightTopicDataSetsColumnsCellValueSynonymsList
	_jsii_.Get(
		j,
		"cellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) CellValueSynonymsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cellValueSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnDataRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDataRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnDataRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDataRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnFriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnFriendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnFriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnFriendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ColumnSynonymsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ComparativeOrder() QuicksightTopicDataSetsColumnsComparativeOrderOutputReference {
	var returns QuicksightTopicDataSetsColumnsComparativeOrderOutputReference
	_jsii_.Get(
		j,
		"comparativeOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ComparativeOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"comparativeOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) DefaultFormatting() QuicksightTopicDataSetsColumnsDefaultFormattingOutputReference {
	var returns QuicksightTopicDataSetsColumnsDefaultFormattingOutputReference
	_jsii_.Get(
		j,
		"defaultFormatting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) DefaultFormattingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultFormattingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) DisableIndexing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIndexing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) DisableIndexingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIndexingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) IsIncludedInTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncludedInTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) IsIncludedInTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncludedInTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) NeverAggregateInFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neverAggregateInFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) NeverAggregateInFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neverAggregateInFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) NonAdditive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonAdditive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) NonAdditiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonAdditiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) NotAllowedAggregations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notAllowedAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) NotAllowedAggregationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notAllowedAggregationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) SemanticType() QuicksightTopicDataSetsColumnsSemanticTypeOutputReference {
	var returns QuicksightTopicDataSetsColumnsSemanticTypeOutputReference
	_jsii_.Get(
		j,
		"semanticType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) SemanticTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) TimeGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) TimeGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGranularityInput",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsColumnsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightTopicDataSetsColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsColumnsOutputReference_Override(q QuicksightTopicDataSetsColumnsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetAllowedAggregations(val *[]*string) {
	if err := j.validateSetAllowedAggregationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAggregations",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetColumnDataRole(val *string) {
	if err := j.validateSetColumnDataRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnDataRole",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetColumnDescription(val *string) {
	if err := j.validateSetColumnDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnDescription",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetColumnFriendlyName(val *string) {
	if err := j.validateSetColumnFriendlyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnFriendlyName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetColumnName(val *string) {
	if err := j.validateSetColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetColumnSynonyms(val *[]*string) {
	if err := j.validateSetColumnSynonymsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnSynonyms",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetDisableIndexing(val interface{}) {
	if err := j.validateSetDisableIndexingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableIndexing",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetIsIncludedInTopic(val interface{}) {
	if err := j.validateSetIsIncludedInTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIncludedInTopic",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetNeverAggregateInFilter(val interface{}) {
	if err := j.validateSetNeverAggregateInFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"neverAggregateInFilter",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetNonAdditive(val interface{}) {
	if err := j.validateSetNonAdditiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonAdditive",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetNotAllowedAggregations(val *[]*string) {
	if err := j.validateSetNotAllowedAggregationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAllowedAggregations",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference)SetTimeGranularity(val *string) {
	if err := j.validateSetTimeGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeGranularity",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) PutCellValueSynonyms(value interface{}) {
	if err := q.validatePutCellValueSynonymsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCellValueSynonyms",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) PutComparativeOrder(value *QuicksightTopicDataSetsColumnsComparativeOrder) {
	if err := q.validatePutComparativeOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putComparativeOrder",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) PutDefaultFormatting(value *QuicksightTopicDataSetsColumnsDefaultFormatting) {
	if err := q.validatePutDefaultFormattingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDefaultFormatting",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) PutSemanticType(value *QuicksightTopicDataSetsColumnsSemanticType) {
	if err := q.validatePutSemanticTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSemanticType",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetAggregation() {
	_jsii_.InvokeVoid(
		q,
		"resetAggregation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetAllowedAggregations() {
	_jsii_.InvokeVoid(
		q,
		"resetAllowedAggregations",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetCellValueSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetCellValueSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetColumnDataRole() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnDataRole",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetColumnDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetColumnFriendlyName() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnFriendlyName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetColumnName() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetColumnSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetComparativeOrder() {
	_jsii_.InvokeVoid(
		q,
		"resetComparativeOrder",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetDefaultFormatting() {
	_jsii_.InvokeVoid(
		q,
		"resetDefaultFormatting",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetDisableIndexing() {
	_jsii_.InvokeVoid(
		q,
		"resetDisableIndexing",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetIsIncludedInTopic() {
	_jsii_.InvokeVoid(
		q,
		"resetIsIncludedInTopic",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetNeverAggregateInFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetNeverAggregateInFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetNonAdditive() {
	_jsii_.InvokeVoid(
		q,
		"resetNonAdditive",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetNotAllowedAggregations() {
	_jsii_.InvokeVoid(
		q,
		"resetNotAllowedAggregations",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetSemanticType() {
	_jsii_.InvokeVoid(
		q,
		"resetSemanticType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ResetTimeGranularity() {
	_jsii_.InvokeVoid(
		q,
		"resetTimeGranularity",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

