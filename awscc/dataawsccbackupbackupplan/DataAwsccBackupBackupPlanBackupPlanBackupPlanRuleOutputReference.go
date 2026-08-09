// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbackupbackupplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference interface {
	cdktn.ComplexObject
	CompletionWindowMinutes() *float64
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
	CopyActions() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableContinuousBackup() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	IndexActions() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleIndexActionsList
	InternalValue() *DataAwsccBackupBackupPlanBackupPlanBackupPlanRule
	SetInternalValue(val *DataAwsccBackupBackupPlanBackupPlanBackupPlanRule)
	Lifecycle() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference
	RecoveryPointTags() cdktn.StringMap
	RuleName() *string
	ScanActions() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleScanActionsList
	ScheduleExpression() *string
	ScheduleExpressionTimezone() *string
	StartWindowMinutes() *float64
	TargetBackupVault() *string
	TargetLogicallyAirGappedBackupVaultArn() *string
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

// The jsii proxy struct for DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference
type jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CompletionWindowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"completionWindowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CopyActions() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList {
	var returns DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleCopyActionsList
	_jsii_.Get(
		j,
		"copyActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) EnableContinuousBackup() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableContinuousBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) IndexActions() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleIndexActionsList {
	var returns DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleIndexActionsList
	_jsii_.Get(
		j,
		"indexActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) InternalValue() *DataAwsccBackupBackupPlanBackupPlanBackupPlanRule {
	var returns *DataAwsccBackupBackupPlanBackupPlanBackupPlanRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) Lifecycle() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference {
	var returns DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) RecoveryPointTags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"recoveryPointTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScanActions() DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleScanActionsList {
	var returns DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleScanActionsList
	_jsii_.Get(
		j,
		"scanActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ScheduleExpressionTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) StartWindowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startWindowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TargetBackupVault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetBackupVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TargetLogicallyAirGappedBackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetLogicallyAirGappedBackupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBackupBackupPlan.DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference_Override(d DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBackupBackupPlan.DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetInternalValue(val *DataAwsccBackupBackupPlanBackupPlanBackupPlanRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBackupBackupPlanBackupPlanBackupPlanRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

