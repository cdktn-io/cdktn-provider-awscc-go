// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dlmlifecyclepolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference interface {
	cdktn.ComplexObject
	ArchiveRule() DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleOutputReference
	ArchiveRuleInput() interface{}
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
	CopyTags() interface{}
	SetCopyTags(val interface{})
	CopyTagsInput() interface{}
	CreateRule() DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleOutputReference
	CreateRuleInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrossRegionCopyRules() DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesList
	CrossRegionCopyRulesInput() interface{}
	DeprecateRule() DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRuleOutputReference
	DeprecateRuleInput() interface{}
	FastRestoreRule() DlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRuleOutputReference
	FastRestoreRuleInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	RetainRule() DlmLifecyclePolicyPolicyDetailsSchedulesRetainRuleOutputReference
	RetainRuleInput() interface{}
	ShareRules() DlmLifecyclePolicyPolicyDetailsSchedulesShareRulesList
	ShareRulesInput() interface{}
	TagsToAdd() DlmLifecyclePolicyPolicyDetailsSchedulesTagsToAddList
	TagsToAddInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VariableTags() DlmLifecyclePolicyPolicyDetailsSchedulesVariableTagsList
	VariableTagsInput() interface{}
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
	PutArchiveRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRule)
	PutCreateRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesCreateRule)
	PutCrossRegionCopyRules(value interface{})
	PutDeprecateRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRule)
	PutFastRestoreRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRule)
	PutRetainRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesRetainRule)
	PutShareRules(value interface{})
	PutTagsToAdd(value interface{})
	PutVariableTags(value interface{})
	ResetArchiveRule()
	ResetCopyTags()
	ResetCreateRule()
	ResetCrossRegionCopyRules()
	ResetDeprecateRule()
	ResetFastRestoreRule()
	ResetName()
	ResetRetainRule()
	ResetShareRules()
	ResetTagsToAdd()
	ResetVariableTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference
type jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ArchiveRule() DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleOutputReference
	_jsii_.Get(
		j,
		"archiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ArchiveRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CreateRule() DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesCreateRuleOutputReference
	_jsii_.Get(
		j,
		"createRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CreateRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CrossRegionCopyRules() DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesList {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesList
	_jsii_.Get(
		j,
		"crossRegionCopyRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) CrossRegionCopyRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionCopyRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) DeprecateRule() DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRuleOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) DeprecateRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deprecateRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) FastRestoreRule() DlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRuleOutputReference
	_jsii_.Get(
		j,
		"fastRestoreRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) FastRestoreRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastRestoreRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) RetainRule() DlmLifecyclePolicyPolicyDetailsSchedulesRetainRuleOutputReference {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesRetainRuleOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) RetainRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ShareRules() DlmLifecyclePolicyPolicyDetailsSchedulesShareRulesList {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesShareRulesList
	_jsii_.Get(
		j,
		"shareRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ShareRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TagsToAdd() DlmLifecyclePolicyPolicyDetailsSchedulesTagsToAddList {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesTagsToAddList
	_jsii_.Get(
		j,
		"tagsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TagsToAddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) VariableTags() DlmLifecyclePolicyPolicyDetailsSchedulesVariableTagsList {
	var returns DlmLifecyclePolicyPolicyDetailsSchedulesVariableTagsList
	_jsii_.Get(
		j,
		"variableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) VariableTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableTagsInput",
		&returns,
	)
	return returns
}


func NewDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference {
	_init_.Initialize()

	if err := validateNewDlmLifecyclePolicyPolicyDetailsSchedulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDlmLifecyclePolicyPolicyDetailsSchedulesOutputReference_Override(d DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dlmLifecyclePolicy.DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutArchiveRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRule) {
	if err := d.validatePutArchiveRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putArchiveRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutCreateRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesCreateRule) {
	if err := d.validatePutCreateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutCrossRegionCopyRules(value interface{}) {
	if err := d.validatePutCrossRegionCopyRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCrossRegionCopyRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutDeprecateRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesDeprecateRule) {
	if err := d.validatePutDeprecateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeprecateRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutFastRestoreRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesFastRestoreRule) {
	if err := d.validatePutFastRestoreRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFastRestoreRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutRetainRule(value *DlmLifecyclePolicyPolicyDetailsSchedulesRetainRule) {
	if err := d.validatePutRetainRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRetainRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutShareRules(value interface{}) {
	if err := d.validatePutShareRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putShareRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutTagsToAdd(value interface{}) {
	if err := d.validatePutTagsToAddParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTagsToAdd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) PutVariableTags(value interface{}) {
	if err := d.validatePutVariableTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVariableTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetArchiveRule() {
	_jsii_.InvokeVoid(
		d,
		"resetArchiveRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetCreateRule() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetCrossRegionCopyRules() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossRegionCopyRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetDeprecateRule() {
	_jsii_.InvokeVoid(
		d,
		"resetDeprecateRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetFastRestoreRule() {
	_jsii_.InvokeVoid(
		d,
		"resetFastRestoreRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetRetainRule() {
	_jsii_.InvokeVoid(
		d,
		"resetRetainRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetShareRules() {
	_jsii_.InvokeVoid(
		d,
		"resetShareRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetTagsToAdd() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsToAdd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ResetVariableTags() {
	_jsii_.InvokeVoid(
		d,
		"resetVariableTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

