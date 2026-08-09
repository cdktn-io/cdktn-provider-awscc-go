// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccquicksighttopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccquicksighttopic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccQuicksightTopicDataSetsColumnsOutputReference interface {
	cdktn.ComplexObject
	Aggregation() *string
	AllowedAggregations() *[]*string
	CellValueSynonyms() DataAwsccQuicksightTopicDataSetsColumnsCellValueSynonymsList
	ColumnDataRole() *string
	ColumnDescription() *string
	ColumnFriendlyName() *string
	ColumnName() *string
	ColumnSynonyms() *[]*string
	ComparativeOrder() DataAwsccQuicksightTopicDataSetsColumnsComparativeOrderOutputReference
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
	DefaultFormatting() DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingOutputReference
	DisableIndexing() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccQuicksightTopicDataSetsColumns
	SetInternalValue(val *DataAwsccQuicksightTopicDataSetsColumns)
	IsIncludedInTopic() cdktn.IResolvable
	NeverAggregateInFilter() cdktn.IResolvable
	NonAdditive() cdktn.IResolvable
	NotAllowedAggregations() *[]*string
	SemanticType() DataAwsccQuicksightTopicDataSetsColumnsSemanticTypeOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeGranularity() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccQuicksightTopicDataSetsColumnsOutputReference
type jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) AllowedAggregations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) CellValueSynonyms() DataAwsccQuicksightTopicDataSetsColumnsCellValueSynonymsList {
	var returns DataAwsccQuicksightTopicDataSetsColumnsCellValueSynonymsList
	_jsii_.Get(
		j,
		"cellValueSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ColumnDataRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDataRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ColumnDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ColumnFriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnFriendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ColumnSynonyms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnSynonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ComparativeOrder() DataAwsccQuicksightTopicDataSetsColumnsComparativeOrderOutputReference {
	var returns DataAwsccQuicksightTopicDataSetsColumnsComparativeOrderOutputReference
	_jsii_.Get(
		j,
		"comparativeOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) DefaultFormatting() DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingOutputReference {
	var returns DataAwsccQuicksightTopicDataSetsColumnsDefaultFormattingOutputReference
	_jsii_.Get(
		j,
		"defaultFormatting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) DisableIndexing() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableIndexing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) InternalValue() *DataAwsccQuicksightTopicDataSetsColumns {
	var returns *DataAwsccQuicksightTopicDataSetsColumns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) IsIncludedInTopic() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isIncludedInTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) NeverAggregateInFilter() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"neverAggregateInFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) NonAdditive() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"nonAdditive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) NotAllowedAggregations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notAllowedAggregations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) SemanticType() DataAwsccQuicksightTopicDataSetsColumnsSemanticTypeOutputReference {
	var returns DataAwsccQuicksightTopicDataSetsColumnsSemanticTypeOutputReference
	_jsii_.Get(
		j,
		"semanticType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) TimeGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeGranularity",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightTopicDataSetsColumnsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccQuicksightTopicDataSetsColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightTopicDataSetsColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightTopic.DataAwsccQuicksightTopicDataSetsColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightTopicDataSetsColumnsOutputReference_Override(d DataAwsccQuicksightTopicDataSetsColumnsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightTopic.DataAwsccQuicksightTopicDataSetsColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference)SetInternalValue(val *DataAwsccQuicksightTopicDataSetsColumns) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightTopicDataSetsColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

