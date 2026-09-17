// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscctranscribecallanalyticscategory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscctranscribecallanalyticscategory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference interface {
	cdktn.ComplexObject
	AbsoluteTimeRange() DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterAbsoluteTimeRangeOutputReference
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
	InternalValue() *DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilter
	SetInternalValue(val *DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilter)
	Negate() cdktn.IResolvable
	ParticipantRole() *string
	RelativeTimeRange() DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterRelativeTimeRangeOutputReference
	Targets() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TranscriptFilterType() *string
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

// The jsii proxy struct for DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference
type jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) AbsoluteTimeRange() DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterAbsoluteTimeRangeOutputReference {
	var returns DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterAbsoluteTimeRangeOutputReference
	_jsii_.Get(
		j,
		"absoluteTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) InternalValue() *DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilter {
	var returns *DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) Negate() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) ParticipantRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"participantRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) RelativeTimeRange() DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterRelativeTimeRangeOutputReference {
	var returns DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterRelativeTimeRangeOutputReference
	_jsii_.Get(
		j,
		"relativeTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) Targets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) TranscriptFilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transcriptFilterType",
		&returns,
	)
	return returns
}


func NewDataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccTranscribeCallAnalyticsCategory.DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference_Override(d DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccTranscribeCallAnalyticsCategory.DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference)SetInternalValue(val *DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccTranscribeCallAnalyticsCategoryRulesTranscriptFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

