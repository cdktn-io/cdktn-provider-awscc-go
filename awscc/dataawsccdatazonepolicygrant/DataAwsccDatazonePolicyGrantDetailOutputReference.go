// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdatazonepolicygrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdatazonepolicygrant/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDatazonePolicyGrantDetailOutputReference interface {
	cdktn.ComplexObject
	AddToProjectMemberPool() DataAwsccDatazonePolicyGrantDetailAddToProjectMemberPoolOutputReference
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
	CreateAssetType() DataAwsccDatazonePolicyGrantDetailCreateAssetTypeOutputReference
	CreateDomainUnit() DataAwsccDatazonePolicyGrantDetailCreateDomainUnitOutputReference
	CreateEnvironment() *string
	CreateEnvironmentFromBlueprint() *string
	CreateEnvironmentProfile() DataAwsccDatazonePolicyGrantDetailCreateEnvironmentProfileOutputReference
	CreateFormType() DataAwsccDatazonePolicyGrantDetailCreateFormTypeOutputReference
	CreateGlossary() DataAwsccDatazonePolicyGrantDetailCreateGlossaryOutputReference
	CreateProject() DataAwsccDatazonePolicyGrantDetailCreateProjectOutputReference
	CreateProjectFromProjectProfile() DataAwsccDatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DelegateCreateEnvironmentProfile() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccDatazonePolicyGrantDetail
	SetInternalValue(val *DataAwsccDatazonePolicyGrantDetail)
	OverrideDomainUnitOwners() DataAwsccDatazonePolicyGrantDetailOverrideDomainUnitOwnersOutputReference
	OverrideProjectOwners() DataAwsccDatazonePolicyGrantDetailOverrideProjectOwnersOutputReference
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

// The jsii proxy struct for DataAwsccDatazonePolicyGrantDetailOutputReference
type jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) AddToProjectMemberPool() DataAwsccDatazonePolicyGrantDetailAddToProjectMemberPoolOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailAddToProjectMemberPoolOutputReference
	_jsii_.Get(
		j,
		"addToProjectMemberPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateAssetType() DataAwsccDatazonePolicyGrantDetailCreateAssetTypeOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateAssetTypeOutputReference
	_jsii_.Get(
		j,
		"createAssetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateDomainUnit() DataAwsccDatazonePolicyGrantDetailCreateDomainUnitOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateDomainUnitOutputReference
	_jsii_.Get(
		j,
		"createDomainUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateEnvironmentFromBlueprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createEnvironmentFromBlueprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateEnvironmentProfile() DataAwsccDatazonePolicyGrantDetailCreateEnvironmentProfileOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateEnvironmentProfileOutputReference
	_jsii_.Get(
		j,
		"createEnvironmentProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateFormType() DataAwsccDatazonePolicyGrantDetailCreateFormTypeOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateFormTypeOutputReference
	_jsii_.Get(
		j,
		"createFormType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateGlossary() DataAwsccDatazonePolicyGrantDetailCreateGlossaryOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateGlossaryOutputReference
	_jsii_.Get(
		j,
		"createGlossary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateProject() DataAwsccDatazonePolicyGrantDetailCreateProjectOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateProjectOutputReference
	_jsii_.Get(
		j,
		"createProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreateProjectFromProjectProfile() DataAwsccDatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference
	_jsii_.Get(
		j,
		"createProjectFromProjectProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) DelegateCreateEnvironmentProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegateCreateEnvironmentProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) InternalValue() *DataAwsccDatazonePolicyGrantDetail {
	var returns *DataAwsccDatazonePolicyGrantDetail
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) OverrideDomainUnitOwners() DataAwsccDatazonePolicyGrantDetailOverrideDomainUnitOwnersOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailOverrideDomainUnitOwnersOutputReference
	_jsii_.Get(
		j,
		"overrideDomainUnitOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) OverrideProjectOwners() DataAwsccDatazonePolicyGrantDetailOverrideProjectOwnersOutputReference {
	var returns DataAwsccDatazonePolicyGrantDetailOverrideProjectOwnersOutputReference
	_jsii_.Get(
		j,
		"overrideProjectOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccDatazonePolicyGrantDetailOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDatazonePolicyGrantDetailOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDatazonePolicyGrantDetailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDatazonePolicyGrant.DataAwsccDatazonePolicyGrantDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDatazonePolicyGrantDetailOutputReference_Override(d DataAwsccDatazonePolicyGrantDetailOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDatazonePolicyGrant.DataAwsccDatazonePolicyGrantDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference)SetInternalValue(val *DataAwsccDatazonePolicyGrantDetail) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDatazonePolicyGrantDetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

