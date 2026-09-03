// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscccleanroomsconfiguredtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscccleanroomsconfiguredtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference interface {
	cdktn.ComplexObject
	AdditionalAnalyses() *string
	AggregationThresholds() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsList
	AllowedAnalyses() *[]*string
	AllowedAnalysisProviders() *[]*string
	ComparisonControls() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomComparisonControlsOutputReference
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
	DifferentialPrivacy() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomDifferentialPrivacyOutputReference
	DisallowedOutputColumns() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1Custom
	SetInternalValue(val *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1Custom)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference
type jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) AdditionalAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) AggregationThresholds() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsList {
	var returns DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomAggregationThresholdsList
	_jsii_.Get(
		j,
		"aggregationThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) AllowedAnalyses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) AllowedAnalysisProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAnalysisProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) ComparisonControls() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomComparisonControlsOutputReference {
	var returns DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomComparisonControlsOutputReference
	_jsii_.Get(
		j,
		"comparisonControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) DifferentialPrivacy() DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomDifferentialPrivacyOutputReference {
	var returns DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomDifferentialPrivacyOutputReference
	_jsii_.Get(
		j,
		"differentialPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) DisallowedOutputColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedOutputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) InternalValue() *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1Custom {
	var returns *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1Custom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCleanroomsConfiguredTable.DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference_Override(d DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccCleanroomsConfiguredTable.DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference)SetInternalValue(val *DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1Custom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCleanroomsConfiguredTableAnalysisRulesPolicyV1CustomOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

