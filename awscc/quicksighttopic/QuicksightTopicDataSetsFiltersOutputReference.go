// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicDataSetsFiltersOutputReference interface {
	cdktn.ComplexObject
	CategoryFilter() QuicksightTopicDataSetsFiltersCategoryFilterOutputReference
	CategoryFilterInput() interface{}
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
	DateRangeFilter() QuicksightTopicDataSetsFiltersDateRangeFilterOutputReference
	DateRangeFilterInput() interface{}
	FilterClass() *string
	SetFilterClass(val *string)
	FilterClassInput() *string
	FilterDescription() *string
	SetFilterDescription(val *string)
	FilterDescriptionInput() *string
	FilterName() *string
	SetFilterName(val *string)
	FilterNameInput() *string
	FilterSynonyms() *[]*string
	SetFilterSynonyms(val *[]*string)
	FilterSynonymsInput() *[]*string
	FilterType() *string
	SetFilterType(val *string)
	FilterTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NumericEqualityFilter() QuicksightTopicDataSetsFiltersNumericEqualityFilterOutputReference
	NumericEqualityFilterInput() interface{}
	NumericRangeFilter() QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference
	NumericRangeFilterInput() interface{}
	OperandFieldName() *string
	SetOperandFieldName(val *string)
	OperandFieldNameInput() *string
	RelativeDateFilter() QuicksightTopicDataSetsFiltersRelativeDateFilterOutputReference
	RelativeDateFilterInput() interface{}
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
	PutCategoryFilter(value *QuicksightTopicDataSetsFiltersCategoryFilter)
	PutDateRangeFilter(value *QuicksightTopicDataSetsFiltersDateRangeFilter)
	PutNumericEqualityFilter(value *QuicksightTopicDataSetsFiltersNumericEqualityFilter)
	PutNumericRangeFilter(value *QuicksightTopicDataSetsFiltersNumericRangeFilter)
	PutRelativeDateFilter(value *QuicksightTopicDataSetsFiltersRelativeDateFilter)
	ResetCategoryFilter()
	ResetDateRangeFilter()
	ResetFilterClass()
	ResetFilterDescription()
	ResetFilterName()
	ResetFilterSynonyms()
	ResetFilterType()
	ResetNumericEqualityFilter()
	ResetNumericRangeFilter()
	ResetOperandFieldName()
	ResetRelativeDateFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsFiltersOutputReference
type jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) CategoryFilter() QuicksightTopicDataSetsFiltersCategoryFilterOutputReference {
	var returns QuicksightTopicDataSetsFiltersCategoryFilterOutputReference
	_jsii_.Get(
		j,
		"categoryFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) CategoryFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"categoryFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) DateRangeFilter() QuicksightTopicDataSetsFiltersDateRangeFilterOutputReference {
	var returns QuicksightTopicDataSetsFiltersDateRangeFilterOutputReference
	_jsii_.Get(
		j,
		"dateRangeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) DateRangeFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateRangeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterSynonymsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterSynonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) FilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) NumericEqualityFilter() QuicksightTopicDataSetsFiltersNumericEqualityFilterOutputReference {
	var returns QuicksightTopicDataSetsFiltersNumericEqualityFilterOutputReference
	_jsii_.Get(
		j,
		"numericEqualityFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) NumericEqualityFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numericEqualityFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) NumericRangeFilter() QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference {
	var returns QuicksightTopicDataSetsFiltersNumericRangeFilterOutputReference
	_jsii_.Get(
		j,
		"numericRangeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) NumericRangeFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numericRangeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) OperandFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operandFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) OperandFieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operandFieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) RelativeDateFilter() QuicksightTopicDataSetsFiltersRelativeDateFilterOutputReference {
	var returns QuicksightTopicDataSetsFiltersRelativeDateFilterOutputReference
	_jsii_.Get(
		j,
		"relativeDateFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) RelativeDateFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relativeDateFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightTopicDataSetsFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsFiltersOutputReference_Override(q QuicksightTopicDataSetsFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetFilterClass(val *string) {
	if err := j.validateSetFilterClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterClass",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetFilterDescription(val *string) {
	if err := j.validateSetFilterDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterDescription",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetFilterName(val *string) {
	if err := j.validateSetFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetFilterSynonyms(val *[]*string) {
	if err := j.validateSetFilterSynonymsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterSynonyms",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetFilterType(val *string) {
	if err := j.validateSetFilterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterType",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetOperandFieldName(val *string) {
	if err := j.validateSetOperandFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operandFieldName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) PutCategoryFilter(value *QuicksightTopicDataSetsFiltersCategoryFilter) {
	if err := q.validatePutCategoryFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCategoryFilter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) PutDateRangeFilter(value *QuicksightTopicDataSetsFiltersDateRangeFilter) {
	if err := q.validatePutDateRangeFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDateRangeFilter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) PutNumericEqualityFilter(value *QuicksightTopicDataSetsFiltersNumericEqualityFilter) {
	if err := q.validatePutNumericEqualityFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNumericEqualityFilter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) PutNumericRangeFilter(value *QuicksightTopicDataSetsFiltersNumericRangeFilter) {
	if err := q.validatePutNumericRangeFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNumericRangeFilter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) PutRelativeDateFilter(value *QuicksightTopicDataSetsFiltersRelativeDateFilter) {
	if err := q.validatePutRelativeDateFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRelativeDateFilter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetCategoryFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetCategoryFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetDateRangeFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetDateRangeFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetFilterClass() {
	_jsii_.InvokeVoid(
		q,
		"resetFilterClass",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetFilterDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetFilterDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetFilterName() {
	_jsii_.InvokeVoid(
		q,
		"resetFilterName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetFilterSynonyms() {
	_jsii_.InvokeVoid(
		q,
		"resetFilterSynonyms",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetFilterType() {
	_jsii_.InvokeVoid(
		q,
		"resetFilterType",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetNumericEqualityFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetNumericEqualityFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetNumericRangeFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetNumericRangeFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetOperandFieldName() {
	_jsii_.InvokeVoid(
		q,
		"resetOperandFieldName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ResetRelativeDateFilter() {
	_jsii_.InvokeVoid(
		q,
		"resetRelativeDateFilter",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

