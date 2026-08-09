// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotsitewiseassetmodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseAssetModelAssetModelCompositeModelsOutputReference interface {
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
	ComposedAssetModelId() *string
	SetComposedAssetModelId(val *string)
	ComposedAssetModelIdInput() *string
	CompositeModelProperties() IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesList
	CompositeModelPropertiesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	ParentAssetModelCompositeModelExternalId() *string
	SetParentAssetModelCompositeModelExternalId(val *string)
	ParentAssetModelCompositeModelExternalIdInput() *string
	Path() *[]*string
	SetPath(val *[]*string)
	PathInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutCompositeModelProperties(value interface{})
	ResetComposedAssetModelId()
	ResetCompositeModelProperties()
	ResetDescription()
	ResetExternalId()
	ResetId()
	ResetName()
	ResetParentAssetModelCompositeModelExternalId()
	ResetPath()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotsitewiseAssetModelAssetModelCompositeModelsOutputReference
type jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ComposedAssetModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composedAssetModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ComposedAssetModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composedAssetModelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) CompositeModelProperties() IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesList {
	var returns IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesList
	_jsii_.Get(
		j,
		"compositeModelProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) CompositeModelPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compositeModelPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ParentAssetModelCompositeModelExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentAssetModelCompositeModelExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ParentAssetModelCompositeModelExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentAssetModelCompositeModelExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Path() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) PathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewIotsitewiseAssetModelAssetModelCompositeModelsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotsitewiseAssetModelAssetModelCompositeModelsOutputReference {
	_init_.Initialize()

	if err := validateNewIotsitewiseAssetModelAssetModelCompositeModelsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelCompositeModelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotsitewiseAssetModelAssetModelCompositeModelsOutputReference_Override(i IotsitewiseAssetModelAssetModelCompositeModelsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelCompositeModelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetComposedAssetModelId(val *string) {
	if err := j.validateSetComposedAssetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"composedAssetModelId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetParentAssetModelCompositeModelExternalId(val *string) {
	if err := j.validateSetParentAssetModelCompositeModelExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentAssetModelCompositeModelExternalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetPath(val *[]*string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) PutCompositeModelProperties(value interface{}) {
	if err := i.validatePutCompositeModelPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCompositeModelProperties",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetComposedAssetModelId() {
	_jsii_.InvokeVoid(
		i,
		"resetComposedAssetModelId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetCompositeModelProperties() {
	_jsii_.InvokeVoid(
		i,
		"resetCompositeModelProperties",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		i,
		"resetExternalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetParentAssetModelCompositeModelExternalId() {
	_jsii_.InvokeVoid(
		i,
		"resetParentAssetModelCompositeModelExternalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		i,
		"resetPath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		i,
		"resetType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelCompositeModelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

