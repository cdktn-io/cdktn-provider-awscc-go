// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicDataSetsOutputReference interface {
	cdktn.ComplexObject
	CalculatedFields() QuicksightTopicDataSetsCalculatedFieldsList
	CalculatedFieldsInput() interface{}
	Columns() QuicksightTopicDataSetsColumnsList
	ColumnsInput() interface{}
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
	DataAggregation() QuicksightTopicDataSetsDataAggregationOutputReference
	DataAggregationInput() interface{}
	DatasetArn() *string
	SetDatasetArn(val *string)
	DatasetArnInput() *string
	DatasetDescription() *string
	SetDatasetDescription(val *string)
	DatasetDescriptionInput() *string
	DatasetName() *string
	SetDatasetName(val *string)
	DatasetNameInput() *string
	Filters() QuicksightTopicDataSetsFiltersList
	FiltersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NamedEntities() QuicksightTopicDataSetsNamedEntitiesList
	NamedEntitiesInput() interface{}
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
	PutCalculatedFields(value interface{})
	PutColumns(value interface{})
	PutDataAggregation(value *QuicksightTopicDataSetsDataAggregation)
	PutFilters(value interface{})
	PutNamedEntities(value interface{})
	ResetCalculatedFields()
	ResetColumns()
	ResetDataAggregation()
	ResetDatasetArn()
	ResetDatasetDescription()
	ResetDatasetName()
	ResetFilters()
	ResetNamedEntities()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTopicDataSetsOutputReference
type jsiiProxy_QuicksightTopicDataSetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) CalculatedFields() QuicksightTopicDataSetsCalculatedFieldsList {
	var returns QuicksightTopicDataSetsCalculatedFieldsList
	_jsii_.Get(
		j,
		"calculatedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) CalculatedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"calculatedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) Columns() QuicksightTopicDataSetsColumnsList {
	var returns QuicksightTopicDataSetsColumnsList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DataAggregation() QuicksightTopicDataSetsDataAggregationOutputReference {
	var returns QuicksightTopicDataSetsDataAggregationOutputReference
	_jsii_.Get(
		j,
		"dataAggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DataAggregationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataAggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DatasetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DatasetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DatasetDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DatasetDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) DatasetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) Filters() QuicksightTopicDataSetsFiltersList {
	var returns QuicksightTopicDataSetsFiltersList
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) FiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) NamedEntities() QuicksightTopicDataSetsNamedEntitiesList {
	var returns QuicksightTopicDataSetsNamedEntitiesList
	_jsii_.Get(
		j,
		"namedEntities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) NamedEntitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namedEntitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightTopicDataSetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightTopicDataSetsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightTopicDataSetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTopicDataSetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightTopicDataSetsOutputReference_Override(q QuicksightTopicDataSetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTopic.QuicksightTopicDataSetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetDatasetArn(val *string) {
	if err := j.validateSetDatasetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetDatasetDescription(val *string) {
	if err := j.validateSetDatasetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetDescription",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetDatasetName(val *string) {
	if err := j.validateSetDatasetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTopicDataSetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) PutCalculatedFields(value interface{}) {
	if err := q.validatePutCalculatedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCalculatedFields",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) PutColumns(value interface{}) {
	if err := q.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putColumns",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) PutDataAggregation(value *QuicksightTopicDataSetsDataAggregation) {
	if err := q.validatePutDataAggregationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataAggregation",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) PutFilters(value interface{}) {
	if err := q.validatePutFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFilters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) PutNamedEntities(value interface{}) {
	if err := q.validatePutNamedEntitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putNamedEntities",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetCalculatedFields() {
	_jsii_.InvokeVoid(
		q,
		"resetCalculatedFields",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		q,
		"resetColumns",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetDataAggregation() {
	_jsii_.InvokeVoid(
		q,
		"resetDataAggregation",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetDatasetArn() {
	_jsii_.InvokeVoid(
		q,
		"resetDatasetArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetDatasetDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetDatasetDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetDatasetName() {
	_jsii_.InvokeVoid(
		q,
		"resetDatasetName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		q,
		"resetFilters",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ResetNamedEntities() {
	_jsii_.InvokeVoid(
		q,
		"resetNamedEntities",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightTopicDataSetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

