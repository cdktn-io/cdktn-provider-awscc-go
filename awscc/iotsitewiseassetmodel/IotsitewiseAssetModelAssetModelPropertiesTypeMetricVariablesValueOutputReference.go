// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotsitewiseassetmodel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference interface {
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
	PropertyPath() IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValuePropertyPathList
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

// The jsii proxy struct for IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference
type jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) HierarchyExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) HierarchyExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) HierarchyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) HierarchyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) HierarchyLogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLogicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) HierarchyLogicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchyLogicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyLogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLogicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyLogicalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLogicalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyPath() IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValuePropertyPathList {
	var returns IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValuePropertyPathList
	_jsii_.Get(
		j,
		"propertyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PropertyPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference {
	_init_.Initialize()

	if err := validateNewIotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference_Override(i IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotsitewiseAssetModel.IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetHierarchyExternalId(val *string) {
	if err := j.validateSetHierarchyExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyExternalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetHierarchyId(val *string) {
	if err := j.validateSetHierarchyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetHierarchyLogicalId(val *string) {
	if err := j.validateSetHierarchyLogicalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hierarchyLogicalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetPropertyExternalId(val *string) {
	if err := j.validateSetPropertyExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyExternalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetPropertyId(val *string) {
	if err := j.validateSetPropertyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetPropertyLogicalId(val *string) {
	if err := j.validateSetPropertyLogicalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyLogicalId",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) PutPropertyPath(value interface{}) {
	if err := i.validatePutPropertyPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPropertyPath",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetHierarchyExternalId() {
	_jsii_.InvokeVoid(
		i,
		"resetHierarchyExternalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetHierarchyId() {
	_jsii_.InvokeVoid(
		i,
		"resetHierarchyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetHierarchyLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetHierarchyLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetPropertyExternalId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyExternalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetPropertyId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetPropertyLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ResetPropertyPath() {
	_jsii_.InvokeVoid(
		i,
		"resetPropertyPath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotsitewiseAssetModelAssetModelPropertiesTypeMetricVariablesValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

