// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/efsfilesystem/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EfsFileSystemLifecyclePoliciesOutputReference interface {
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransitionToArchive() *string
	SetTransitionToArchive(val *string)
	TransitionToArchiveInput() *string
	TransitionToIa() *string
	SetTransitionToIa(val *string)
	TransitionToIaInput() *string
	TransitionToPrimaryStorageClass() *string
	SetTransitionToPrimaryStorageClass(val *string)
	TransitionToPrimaryStorageClassInput() *string
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
	ResetTransitionToArchive()
	ResetTransitionToIa()
	ResetTransitionToPrimaryStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EfsFileSystemLifecyclePoliciesOutputReference
type jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TransitionToArchive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TransitionToArchiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToArchiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TransitionToIa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToIa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TransitionToIaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToIaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TransitionToPrimaryStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToPrimaryStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) TransitionToPrimaryStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitionToPrimaryStorageClassInput",
		&returns,
	)
	return returns
}


func NewEfsFileSystemLifecyclePoliciesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EfsFileSystemLifecyclePoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewEfsFileSystemLifecyclePoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.efsFileSystem.EfsFileSystemLifecyclePoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEfsFileSystemLifecyclePoliciesOutputReference_Override(e EfsFileSystemLifecyclePoliciesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.efsFileSystem.EfsFileSystemLifecyclePoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetTransitionToArchive(val *string) {
	if err := j.validateSetTransitionToArchiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitionToArchive",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetTransitionToIa(val *string) {
	if err := j.validateSetTransitionToIaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitionToIa",
		val,
	)
}

func (j *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference)SetTransitionToPrimaryStorageClass(val *string) {
	if err := j.validateSetTransitionToPrimaryStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitionToPrimaryStorageClass",
		val,
	)
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ResetTransitionToArchive() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitionToArchive",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ResetTransitionToIa() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitionToIa",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ResetTransitionToPrimaryStorageClass() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitionToPrimaryStorageClass",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EfsFileSystemLifecyclePoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

