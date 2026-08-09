// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchcomputeenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchTemplateId() *string
	SetLaunchTemplateId(val *string)
	LaunchTemplateIdInput() *string
	LaunchTemplateName() *string
	SetLaunchTemplateName(val *string)
	LaunchTemplateNameInput() *string
	TargetInstanceTypes() *[]*string
	SetTargetInstanceTypes(val *[]*string)
	TargetInstanceTypesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserdataType() *string
	SetUserdataType(val *string)
	UserdataTypeInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetLaunchTemplateId()
	ResetLaunchTemplateName()
	ResetTargetInstanceTypes()
	ResetUserdataType()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) LaunchTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) LaunchTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) LaunchTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) LaunchTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) TargetInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) TargetInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) UserdataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) UserdataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewBatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchComputeEnvironment.BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference_Override(b BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchComputeEnvironment.BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetLaunchTemplateId(val *string) {
	if err := j.validateSetLaunchTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateId",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetLaunchTemplateName(val *string) {
	if err := j.validateSetLaunchTemplateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplateName",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetTargetInstanceTypes(val *[]*string) {
	if err := j.validateSetTargetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetUserdataType(val *string) {
	if err := j.validateSetUserdataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userdataType",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ResetLaunchTemplateId() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplateId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ResetLaunchTemplateName() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplateName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ResetTargetInstanceTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetTargetInstanceTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ResetUserdataType() {
	_jsii_.InvokeVoid(
		b,
		"resetUserdataType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

