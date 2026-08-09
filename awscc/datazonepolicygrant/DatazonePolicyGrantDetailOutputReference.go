// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonepolicygrant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazonepolicygrant/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazonePolicyGrantDetailOutputReference interface {
	cdktn.ComplexObject
	AddToProjectMemberPool() DatazonePolicyGrantDetailAddToProjectMemberPoolOutputReference
	AddToProjectMemberPoolInput() interface{}
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
	CreateAssetType() DatazonePolicyGrantDetailCreateAssetTypeOutputReference
	CreateAssetTypeInput() interface{}
	CreateDomainUnit() DatazonePolicyGrantDetailCreateDomainUnitOutputReference
	CreateDomainUnitInput() interface{}
	CreateEnvironment() *string
	SetCreateEnvironment(val *string)
	CreateEnvironmentFromBlueprint() *string
	SetCreateEnvironmentFromBlueprint(val *string)
	CreateEnvironmentFromBlueprintInput() *string
	CreateEnvironmentInput() *string
	CreateEnvironmentProfile() DatazonePolicyGrantDetailCreateEnvironmentProfileOutputReference
	CreateEnvironmentProfileInput() interface{}
	CreateFormType() DatazonePolicyGrantDetailCreateFormTypeOutputReference
	CreateFormTypeInput() interface{}
	CreateGlossary() DatazonePolicyGrantDetailCreateGlossaryOutputReference
	CreateGlossaryInput() interface{}
	CreateProject() DatazonePolicyGrantDetailCreateProjectOutputReference
	CreateProjectFromProjectProfile() DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference
	CreateProjectFromProjectProfileInput() interface{}
	CreateProjectInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DelegateCreateEnvironmentProfile() *string
	SetDelegateCreateEnvironmentProfile(val *string)
	DelegateCreateEnvironmentProfileInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OverrideDomainUnitOwners() DatazonePolicyGrantDetailOverrideDomainUnitOwnersOutputReference
	OverrideDomainUnitOwnersInput() interface{}
	OverrideProjectOwners() DatazonePolicyGrantDetailOverrideProjectOwnersOutputReference
	OverrideProjectOwnersInput() interface{}
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
	PutAddToProjectMemberPool(value *DatazonePolicyGrantDetailAddToProjectMemberPool)
	PutCreateAssetType(value *DatazonePolicyGrantDetailCreateAssetType)
	PutCreateDomainUnit(value *DatazonePolicyGrantDetailCreateDomainUnit)
	PutCreateEnvironmentProfile(value *DatazonePolicyGrantDetailCreateEnvironmentProfile)
	PutCreateFormType(value *DatazonePolicyGrantDetailCreateFormType)
	PutCreateGlossary(value *DatazonePolicyGrantDetailCreateGlossary)
	PutCreateProject(value *DatazonePolicyGrantDetailCreateProject)
	PutCreateProjectFromProjectProfile(value *DatazonePolicyGrantDetailCreateProjectFromProjectProfile)
	PutOverrideDomainUnitOwners(value *DatazonePolicyGrantDetailOverrideDomainUnitOwners)
	PutOverrideProjectOwners(value *DatazonePolicyGrantDetailOverrideProjectOwners)
	ResetAddToProjectMemberPool()
	ResetCreateAssetType()
	ResetCreateDomainUnit()
	ResetCreateEnvironment()
	ResetCreateEnvironmentFromBlueprint()
	ResetCreateEnvironmentProfile()
	ResetCreateFormType()
	ResetCreateGlossary()
	ResetCreateProject()
	ResetCreateProjectFromProjectProfile()
	ResetDelegateCreateEnvironmentProfile()
	ResetOverrideDomainUnitOwners()
	ResetOverrideProjectOwners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazonePolicyGrantDetailOutputReference
type jsiiProxy_DatazonePolicyGrantDetailOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) AddToProjectMemberPool() DatazonePolicyGrantDetailAddToProjectMemberPoolOutputReference {
	var returns DatazonePolicyGrantDetailAddToProjectMemberPoolOutputReference
	_jsii_.Get(
		j,
		"addToProjectMemberPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) AddToProjectMemberPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addToProjectMemberPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateAssetType() DatazonePolicyGrantDetailCreateAssetTypeOutputReference {
	var returns DatazonePolicyGrantDetailCreateAssetTypeOutputReference
	_jsii_.Get(
		j,
		"createAssetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateAssetTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createAssetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateDomainUnit() DatazonePolicyGrantDetailCreateDomainUnitOutputReference {
	var returns DatazonePolicyGrantDetailCreateDomainUnitOutputReference
	_jsii_.Get(
		j,
		"createDomainUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateDomainUnitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDomainUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateEnvironmentFromBlueprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createEnvironmentFromBlueprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateEnvironmentFromBlueprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createEnvironmentFromBlueprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateEnvironmentProfile() DatazonePolicyGrantDetailCreateEnvironmentProfileOutputReference {
	var returns DatazonePolicyGrantDetailCreateEnvironmentProfileOutputReference
	_jsii_.Get(
		j,
		"createEnvironmentProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateEnvironmentProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createEnvironmentProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateFormType() DatazonePolicyGrantDetailCreateFormTypeOutputReference {
	var returns DatazonePolicyGrantDetailCreateFormTypeOutputReference
	_jsii_.Get(
		j,
		"createFormType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateFormTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createFormTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateGlossary() DatazonePolicyGrantDetailCreateGlossaryOutputReference {
	var returns DatazonePolicyGrantDetailCreateGlossaryOutputReference
	_jsii_.Get(
		j,
		"createGlossary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateGlossaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createGlossaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateProject() DatazonePolicyGrantDetailCreateProjectOutputReference {
	var returns DatazonePolicyGrantDetailCreateProjectOutputReference
	_jsii_.Get(
		j,
		"createProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateProjectFromProjectProfile() DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference {
	var returns DatazonePolicyGrantDetailCreateProjectFromProjectProfileOutputReference
	_jsii_.Get(
		j,
		"createProjectFromProjectProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateProjectFromProjectProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createProjectFromProjectProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreateProjectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) DelegateCreateEnvironmentProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegateCreateEnvironmentProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) DelegateCreateEnvironmentProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegateCreateEnvironmentProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) OverrideDomainUnitOwners() DatazonePolicyGrantDetailOverrideDomainUnitOwnersOutputReference {
	var returns DatazonePolicyGrantDetailOverrideDomainUnitOwnersOutputReference
	_jsii_.Get(
		j,
		"overrideDomainUnitOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) OverrideDomainUnitOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideDomainUnitOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) OverrideProjectOwners() DatazonePolicyGrantDetailOverrideProjectOwnersOutputReference {
	var returns DatazonePolicyGrantDetailOverrideProjectOwnersOutputReference
	_jsii_.Get(
		j,
		"overrideProjectOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) OverrideProjectOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideProjectOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatazonePolicyGrantDetailOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazonePolicyGrantDetailOutputReference {
	_init_.Initialize()

	if err := validateNewDatazonePolicyGrantDetailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazonePolicyGrantDetailOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazonePolicyGrant.DatazonePolicyGrantDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazonePolicyGrantDetailOutputReference_Override(d DatazonePolicyGrantDetailOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazonePolicyGrant.DatazonePolicyGrantDetailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetCreateEnvironment(val *string) {
	if err := j.validateSetCreateEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createEnvironment",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetCreateEnvironmentFromBlueprint(val *string) {
	if err := j.validateSetCreateEnvironmentFromBlueprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createEnvironmentFromBlueprint",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetDelegateCreateEnvironmentProfile(val *string) {
	if err := j.validateSetDelegateCreateEnvironmentProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegateCreateEnvironmentProfile",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazonePolicyGrantDetailOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutAddToProjectMemberPool(value *DatazonePolicyGrantDetailAddToProjectMemberPool) {
	if err := d.validatePutAddToProjectMemberPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAddToProjectMemberPool",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateAssetType(value *DatazonePolicyGrantDetailCreateAssetType) {
	if err := d.validatePutCreateAssetTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateAssetType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateDomainUnit(value *DatazonePolicyGrantDetailCreateDomainUnit) {
	if err := d.validatePutCreateDomainUnitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateDomainUnit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateEnvironmentProfile(value *DatazonePolicyGrantDetailCreateEnvironmentProfile) {
	if err := d.validatePutCreateEnvironmentProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateEnvironmentProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateFormType(value *DatazonePolicyGrantDetailCreateFormType) {
	if err := d.validatePutCreateFormTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateFormType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateGlossary(value *DatazonePolicyGrantDetailCreateGlossary) {
	if err := d.validatePutCreateGlossaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateGlossary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateProject(value *DatazonePolicyGrantDetailCreateProject) {
	if err := d.validatePutCreateProjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateProject",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutCreateProjectFromProjectProfile(value *DatazonePolicyGrantDetailCreateProjectFromProjectProfile) {
	if err := d.validatePutCreateProjectFromProjectProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCreateProjectFromProjectProfile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutOverrideDomainUnitOwners(value *DatazonePolicyGrantDetailOverrideDomainUnitOwners) {
	if err := d.validatePutOverrideDomainUnitOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOverrideDomainUnitOwners",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) PutOverrideProjectOwners(value *DatazonePolicyGrantDetailOverrideProjectOwners) {
	if err := d.validatePutOverrideProjectOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOverrideProjectOwners",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetAddToProjectMemberPool() {
	_jsii_.InvokeVoid(
		d,
		"resetAddToProjectMemberPool",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateAssetType() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateAssetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateDomainUnit() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateDomainUnit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateEnvironment() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateEnvironment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateEnvironmentFromBlueprint() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateEnvironmentFromBlueprint",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateEnvironmentProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateEnvironmentProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateFormType() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateFormType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateGlossary() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateGlossary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateProject() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetCreateProjectFromProjectProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetCreateProjectFromProjectProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetDelegateCreateEnvironmentProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetDelegateCreateEnvironmentProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetOverrideDomainUnitOwners() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideDomainUnitOwners",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ResetOverrideProjectOwners() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideProjectOwners",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazonePolicyGrantDetailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

