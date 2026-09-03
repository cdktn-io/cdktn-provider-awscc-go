// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakercluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference interface {
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
	EbsVolumeConfig() SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfigOutputReference
	EbsVolumeConfigInput() interface{}
	// Experimental.
	Fqn() *string
	FsxLustreConfig() SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxLustreConfigOutputReference
	FsxLustreConfigInput() interface{}
	FsxOpenZfsConfig() SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfigOutputReference
	FsxOpenZfsConfigInput() interface{}
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
	PutEbsVolumeConfig(value *SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfig)
	PutFsxLustreConfig(value *SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxLustreConfig)
	PutFsxOpenZfsConfig(value *SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfig)
	ResetEbsVolumeConfig()
	ResetFsxLustreConfig()
	ResetFsxOpenZfsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference
type jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) EbsVolumeConfig() SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfigOutputReference
	_jsii_.Get(
		j,
		"ebsVolumeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) EbsVolumeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolumeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) FsxLustreConfig() SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxLustreConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxLustreConfigOutputReference
	_jsii_.Get(
		j,
		"fsxLustreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) FsxLustreConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fsxLustreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) FsxOpenZfsConfig() SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfigOutputReference {
	var returns SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfigOutputReference
	_jsii_.Get(
		j,
		"fsxOpenZfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) FsxOpenZfsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fsxOpenZfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference_Override(s SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerCluster.SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) PutEbsVolumeConfig(value *SagemakerClusterInstanceGroupsInstanceStorageConfigsEbsVolumeConfig) {
	if err := s.validatePutEbsVolumeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEbsVolumeConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) PutFsxLustreConfig(value *SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxLustreConfig) {
	if err := s.validatePutFsxLustreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFsxLustreConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) PutFsxOpenZfsConfig(value *SagemakerClusterInstanceGroupsInstanceStorageConfigsFsxOpenZfsConfig) {
	if err := s.validatePutFsxOpenZfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFsxOpenZfsConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ResetEbsVolumeConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetEbsVolumeConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ResetFsxLustreConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetFsxLustreConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ResetFsxOpenZfsConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetFsxOpenZfsConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerClusterInstanceGroupsInstanceStorageConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

