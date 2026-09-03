// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference interface {
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
	PerUnitStorageThroughput() *float64
	SetPerUnitStorageThroughput(val *float64)
	PerUnitStorageThroughputInput() *float64
	SizeInGiB() *float64
	SetSizeInGiB(val *float64)
	SizeInGiBInput() *float64
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
	ResetPerUnitStorageThroughput()
	ResetSizeInGiB()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference
type jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) SizeInGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) SizeInGiBInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference_Override(s SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetPerUnitStorageThroughput(val *float64) {
	if err := j.validateSetPerUnitStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetSizeInGiB(val *float64) {
	if err := j.validateSetSizeInGiBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInGiB",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) ResetPerUnitStorageThroughput() {
	_jsii_.InvokeVoid(
		s,
		"resetPerUnitStorageThroughput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) ResetSizeInGiB() {
	_jsii_.InvokeVoid(
		s,
		"resetSizeInGiB",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterRestrictedInstanceGroupsConfigSharedEnvironmentConfigFsxLustreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

