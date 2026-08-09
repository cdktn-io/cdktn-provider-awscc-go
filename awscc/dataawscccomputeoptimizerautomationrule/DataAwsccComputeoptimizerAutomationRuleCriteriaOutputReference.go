// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscccomputeoptimizerautomationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscccomputeoptimizerautomationrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference interface {
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
	EbsVolumeSizeInGib() DataAwsccComputeoptimizerAutomationRuleCriteriaEbsVolumeSizeInGibList
	EbsVolumeType() DataAwsccComputeoptimizerAutomationRuleCriteriaEbsVolumeTypeList
	EstimatedMonthlySavings() DataAwsccComputeoptimizerAutomationRuleCriteriaEstimatedMonthlySavingsList
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccComputeoptimizerAutomationRuleCriteria
	SetInternalValue(val *DataAwsccComputeoptimizerAutomationRuleCriteria)
	LookBackPeriodInDays() DataAwsccComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDaysList
	Region() DataAwsccComputeoptimizerAutomationRuleCriteriaRegionList
	ResourceArn() DataAwsccComputeoptimizerAutomationRuleCriteriaResourceArnList
	ResourceTag() DataAwsccComputeoptimizerAutomationRuleCriteriaResourceTagList
	RestartNeeded() DataAwsccComputeoptimizerAutomationRuleCriteriaRestartNeededList
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

// The jsii proxy struct for DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference
type jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) EbsVolumeSizeInGib() DataAwsccComputeoptimizerAutomationRuleCriteriaEbsVolumeSizeInGibList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaEbsVolumeSizeInGibList
	_jsii_.Get(
		j,
		"ebsVolumeSizeInGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) EbsVolumeType() DataAwsccComputeoptimizerAutomationRuleCriteriaEbsVolumeTypeList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaEbsVolumeTypeList
	_jsii_.Get(
		j,
		"ebsVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) EstimatedMonthlySavings() DataAwsccComputeoptimizerAutomationRuleCriteriaEstimatedMonthlySavingsList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaEstimatedMonthlySavingsList
	_jsii_.Get(
		j,
		"estimatedMonthlySavings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) InternalValue() *DataAwsccComputeoptimizerAutomationRuleCriteria {
	var returns *DataAwsccComputeoptimizerAutomationRuleCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) LookBackPeriodInDays() DataAwsccComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDaysList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaLookBackPeriodInDaysList
	_jsii_.Get(
		j,
		"lookBackPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) Region() DataAwsccComputeoptimizerAutomationRuleCriteriaRegionList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaRegionList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) ResourceArn() DataAwsccComputeoptimizerAutomationRuleCriteriaResourceArnList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaResourceArnList
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) ResourceTag() DataAwsccComputeoptimizerAutomationRuleCriteriaResourceTagList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaResourceTagList
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) RestartNeeded() DataAwsccComputeoptimizerAutomationRuleCriteriaRestartNeededList {
	var returns DataAwsccComputeoptimizerAutomationRuleCriteriaRestartNeededList
	_jsii_.Get(
		j,
		"restartNeeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccComputeoptimizerAutomationRuleCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccComputeoptimizerAutomationRule.DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference_Override(d DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccComputeoptimizerAutomationRule.DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference)SetInternalValue(val *DataAwsccComputeoptimizerAutomationRuleCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccComputeoptimizerAutomationRuleCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

