// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerdomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference interface {
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
	EfsFileSystemConfig() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference
	EfsFileSystemConfigInput() interface{}
	// Experimental.
	Fqn() *string
	FsxLustreFileSystemConfig() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference
	FsxLustreFileSystemConfigInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3FileSystemConfig() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference
	S3FileSystemConfigInput() interface{}
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
	PutEfsFileSystemConfig(value *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfig)
	PutFsxLustreFileSystemConfig(value *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig)
	PutS3FileSystemConfig(value *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfig)
	ResetEfsFileSystemConfig()
	ResetFsxLustreFileSystemConfig()
	ResetS3FileSystemConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference
type jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) EfsFileSystemConfig() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfigOutputReference
	_jsii_.Get(
		j,
		"efsFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) EfsFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) FsxLustreFileSystemConfig() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfigOutputReference
	_jsii_.Get(
		j,
		"fsxLustreFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) FsxLustreFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fsxLustreFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) S3FileSystemConfig() SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference {
	var returns SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfigOutputReference
	_jsii_.Get(
		j,
		"s3FileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) S3FileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3FileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference_Override(s SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerDomain.SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) PutEfsFileSystemConfig(value *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfig) {
	if err := s.validatePutEfsFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEfsFileSystemConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) PutFsxLustreFileSystemConfig(value *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig) {
	if err := s.validatePutFsxLustreFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFsxLustreFileSystemConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) PutS3FileSystemConfig(value *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfig) {
	if err := s.validatePutS3FileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putS3FileSystemConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ResetEfsFileSystemConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetEfsFileSystemConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ResetFsxLustreFileSystemConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetFsxLustreFileSystemConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ResetS3FileSystemConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetS3FileSystemConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

