// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscccleanroomsconfiguredtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscccleanroomsconfiguredtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference interface {
	cdktn.ComplexObject
	AllowedAggregateExpressionType() *string
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
	InternalValue() *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholds
	SetInternalValue(val *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholds)
	MinimumIdentityCount() *float64
	OutputColumnThresholds() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholdsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
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

// The jsii proxy struct for DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference
type jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) AllowedAggregateExpressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedAggregateExpressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) IdentityColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) InternalValue() *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholds {
	var returns *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholds
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) MinimumIdentityCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumIdentityCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) OutputColumnThresholds() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholdsList {
	var returns DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputColumnThresholdsList
	_jsii_.Get(
		j,
		"outputColumnThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCleanroomsConfiguredTable.DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference_Override(d DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCleanroomsConfiguredTable.DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetInternalValue(val *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholds) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

