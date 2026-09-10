// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdlmlifecyclepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdlmlifecyclepolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference interface {
	cdktn.ComplexObject
	Actions() DataAwsccDlmLifecyclePolicyPolicyDetailsActionsList
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
	CreateInterval() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrossRegionCopyTargets() DataAwsccDlmLifecyclePolicyPolicyDetailsCrossRegionCopyTargetsList
	EventSource() DataAwsccDlmLifecyclePolicyPolicyDetailsEventSourceOutputReference
	Exclusions() DataAwsccDlmLifecyclePolicyPolicyDetailsExclusionsOutputReference
	ExtendDeletion() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDlmLifecyclePolicyPolicyDetails
	SetInternalValue(val *DataAwsccDlmLifecyclePolicyPolicyDetails)
	Parameters() DataAwsccDlmLifecyclePolicyPolicyDetailsParametersOutputReference
	PolicyLanguage() *string
	PolicyType() *string
	ResourceLocations() *[]*string
	ResourceType() *string
	ResourceTypes() *[]*string
	RetainInterval() *float64
	Schedules() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesList
	TargetTags() DataAwsccDlmLifecyclePolicyPolicyDetailsTargetTagsList
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

// The jsii proxy struct for DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference
type jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) Actions() DataAwsccDlmLifecyclePolicyPolicyDetailsActionsList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsActionsList
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) CopyTags() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) CreateInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) CrossRegionCopyTargets() DataAwsccDlmLifecyclePolicyPolicyDetailsCrossRegionCopyTargetsList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsCrossRegionCopyTargetsList
	_jsii_.Get(
		j,
		"crossRegionCopyTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) EventSource() DataAwsccDlmLifecyclePolicyPolicyDetailsEventSourceOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsEventSourceOutputReference
	_jsii_.Get(
		j,
		"eventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) Exclusions() DataAwsccDlmLifecyclePolicyPolicyDetailsExclusionsOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsExclusionsOutputReference
	_jsii_.Get(
		j,
		"exclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ExtendDeletion() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"extendDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) InternalValue() *DataAwsccDlmLifecyclePolicyPolicyDetails {
	var returns *DataAwsccDlmLifecyclePolicyPolicyDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) Parameters() DataAwsccDlmLifecyclePolicyPolicyDetailsParametersOutputReference {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) PolicyLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ResourceLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) RetainInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) Schedules() DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsSchedulesList
	_jsii_.Get(
		j,
		"schedules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) TargetTags() DataAwsccDlmLifecyclePolicyPolicyDetailsTargetTagsList {
	var returns DataAwsccDlmLifecyclePolicyPolicyDetailsTargetTagsList
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDlmLifecyclePolicyPolicyDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDlmLifecyclePolicy.DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference_Override(d DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDlmLifecyclePolicy.DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference)SetInternalValue(val *DataAwsccDlmLifecyclePolicyPolicyDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDlmLifecyclePolicyPolicyDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

