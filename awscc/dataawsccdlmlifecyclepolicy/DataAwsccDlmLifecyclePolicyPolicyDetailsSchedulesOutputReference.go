// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdlmlifecyclepolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference interface {
	cdktn.ComplexObject
	ArchiveRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleOutputReference
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
	CopyTags() cdktn.IResolvable
	CreateRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrossRegionCopyRules() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesList
	DeprecateRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRuleOutputReference
	FastRestoreRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRuleOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedules
	SetInternalValue(val *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedules)
	Name() *string
	RetainRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesRetainRuleOutputReference
	ShareRules() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesShareRulesList
	TagsToAdd() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesTagsToAddList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VariableTags() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesVariableTagsList
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

// The jsii proxy struct for DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference
type jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ArchiveRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleOutputReference
	_jsii_.Get(
		j,
		"archiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CopyTags() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CreateRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleOutputReference
	_jsii_.Get(
		j,
		"createRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CrossRegionCopyRules() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesList
	_jsii_.Get(
		j,
		"crossRegionCopyRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) DeprecateRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRuleOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRuleOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) FastRestoreRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRuleOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRuleOutputReference
	_jsii_.Get(
		j,
		"fastRestoreRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) InternalValue() *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedules {
	var returns *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) RetainRule() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesRetainRuleOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesRetainRuleOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ShareRules() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesShareRulesList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesShareRulesList
	_jsii_.Get(
		j,
		"shareRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TagsToAdd() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesTagsToAddList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesTagsToAddList
	_jsii_.Get(
		j,
		"tagsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) VariableTags() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesVariableTagsList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesVariableTagsList
	_jsii_.Get(
		j,
		"variableTags",
		&returns,
	)
	return returns
}


func NewDataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDlmLifecyclePolicy.DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference_Override(d DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDlmLifecyclePolicy.DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetInternalValue(val *DataAwsccDlmLifecyclePolicyPolicyDetailsSchedules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

