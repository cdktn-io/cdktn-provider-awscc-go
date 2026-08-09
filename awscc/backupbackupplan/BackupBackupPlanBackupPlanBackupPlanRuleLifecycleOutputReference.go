// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/backupbackupplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference interface {
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
	DeleteAfterDays() *float64
	SetDeleteAfterDays(val *float64)
	DeleteAfterDaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MoveToColdStorageAfterDays() *float64
	SetMoveToColdStorageAfterDays(val *float64)
	MoveToColdStorageAfterDaysInput() *float64
	OptInToArchiveForSupportedResources() interface{}
	SetOptInToArchiveForSupportedResources(val interface{})
	OptInToArchiveForSupportedResourcesInput() interface{}
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
	ResetDeleteAfterDays()
	ResetMoveToColdStorageAfterDays()
	ResetOptInToArchiveForSupportedResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference
type jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) DeleteAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) DeleteAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) MoveToColdStorageAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"moveToColdStorageAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) MoveToColdStorageAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"moveToColdStorageAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) OptInToArchiveForSupportedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optInToArchiveForSupportedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) OptInToArchiveForSupportedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optInToArchiveForSupportedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference {
	_init_.Initialize()

	if err := validateNewBackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference_Override(b BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.backupBackupPlan.BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetDeleteAfterDays(val *float64) {
	if err := j.validateSetDeleteAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAfterDays",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetMoveToColdStorageAfterDays(val *float64) {
	if err := j.validateSetMoveToColdStorageAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"moveToColdStorageAfterDays",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetOptInToArchiveForSupportedResources(val interface{}) {
	if err := j.validateSetOptInToArchiveForSupportedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optInToArchiveForSupportedResources",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ResetDeleteAfterDays() {
	_jsii_.InvokeVoid(
		b,
		"resetDeleteAfterDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ResetMoveToColdStorageAfterDays() {
	_jsii_.InvokeVoid(
		b,
		"resetMoveToColdStorageAfterDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ResetOptInToArchiveForSupportedResources() {
	_jsii_.InvokeVoid(
		b,
		"resetOptInToArchiveForSupportedResources",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupBackupPlanBackupPlanBackupPlanRuleLifecycleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

