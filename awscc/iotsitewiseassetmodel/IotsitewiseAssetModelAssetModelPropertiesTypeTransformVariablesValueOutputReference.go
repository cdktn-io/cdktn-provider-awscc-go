// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotsitewiseassetmodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference interface {
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
	HierarchyExternalId() *string
	SetHierarchyExternalId(val *string)
	HierarchyExternalIdInput() *string
	HierarchyId() *string
	SetHierarchyId(val *string)
	HierarchyIdInput() *string
	HierarchyLogicalId() *string
	SetHierarchyLogicalId(val *string)
	HierarchyLogicalIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PropertyExternalId() *string
	SetPropertyExternalId(val *string)
	PropertyExternalIdInput() *string
	PropertyId() *string
	SetPropertyId(val *string)
	PropertyIdInput() *string
	PropertyLogicalId() *string
	SetPropertyLogicalId(val *string)
	PropertyLogicalIdInput() *string
	PropertyPath() IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValuePropertyPathList
	PropertyPathInput() interface{}
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
	PutPropertyPath(value interface{})
	ResetHierarchyExternalId()
	ResetHierarchyId()
	ResetHierarchyLogicalId()
	ResetPropertyExternalId()
	ResetPropertyId()
	ResetPropertyLogicalId()
	ResetPropertyPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference
type jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) HierarchyExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) HierarchyExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) HierarchyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) HierarchyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) HierarchyLogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLogicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) HierarchyLogicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLogicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyLogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLogicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyLogicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLogicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyPath() IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValuePropertyPathList {
	var returns IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValuePropertyPathList
	_jsii_.Get(
		j,
		"propertyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PropertyPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference {
	_init_.Initialize()

	if err := validateNewIotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference_Override(i IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetHierarchyExternalId(val *string) {
	if err := j.validateSetHierarchyExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyExternalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetHierarchyId(val *string) {
	if err := j.validateSetHierarchyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetHierarchyLogicalId(val *string) {
	if err := j.validateSetHierarchyLogicalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyLogicalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetPropertyExternalId(val *string) {
	if err := j.validateSetPropertyExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyExternalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetPropertyId(val *string) {
	if err := j.validateSetPropertyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetPropertyLogicalId(val *string) {
	if err := j.validateSetPropertyLogicalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyLogicalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) PutPropertyPath(value interface{}) {
	if err := i.validatePutPropertyPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPropertyPath",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetHierarchyExternalId() {
	_jsii_.InvokeVoid(
		i,
		"resetHierarchyExternalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetHierarchyId() {
	_jsii_.InvokeVoid(
		i,
		"resetHierarchyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetHierarchyLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetHierarchyLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetPropertyExternalId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyExternalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetPropertyId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetPropertyLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ResetPropertyPath() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyPath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeTransformVariablesValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

