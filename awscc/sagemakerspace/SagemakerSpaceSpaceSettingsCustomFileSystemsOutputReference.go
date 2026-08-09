// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference interface {
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
	EfsFileSystem() SagemakerSpaceSpaceSettingsCustomFileSystemsEfsFileSystemOutputReference
	EfsFileSystemInput() interface{}
	// Experimental.
	Fqn() *string
	FsxLustreFileSystem() SagemakerSpaceSpaceSettingsCustomFileSystemsFsxLustreFileSystemOutputReference
	FsxLustreFileSystemInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3FileSystem() SagemakerSpaceSpaceSettingsCustomFileSystemsS3FileSystemOutputReference
	S3FileSystemInput() interface{}
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
	PutEfsFileSystem(value *SagemakerSpaceSpaceSettingsCustomFileSystemsEfsFileSystem)
	PutFsxLustreFileSystem(value *SagemakerSpaceSpaceSettingsCustomFileSystemsFsxLustreFileSystem)
	PutS3FileSystem(value *SagemakerSpaceSpaceSettingsCustomFileSystemsS3FileSystem)
	ResetEfsFileSystem()
	ResetFsxLustreFileSystem()
	ResetS3FileSystem()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference
type jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) EfsFileSystem() SagemakerSpaceSpaceSettingsCustomFileSystemsEfsFileSystemOutputReference {
	var returns SagemakerSpaceSpaceSettingsCustomFileSystemsEfsFileSystemOutputReference
	_jsii_.Get(
		j,
		"efsFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) EfsFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) FsxLustreFileSystem() SagemakerSpaceSpaceSettingsCustomFileSystemsFsxLustreFileSystemOutputReference {
	var returns SagemakerSpaceSpaceSettingsCustomFileSystemsFsxLustreFileSystemOutputReference
	_jsii_.Get(
		j,
		"fsxLustreFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) FsxLustreFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fsxLustreFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) S3FileSystem() SagemakerSpaceSpaceSettingsCustomFileSystemsS3FileSystemOutputReference {
	var returns SagemakerSpaceSpaceSettingsCustomFileSystemsS3FileSystemOutputReference
	_jsii_.Get(
		j,
		"s3FileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) S3FileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3FileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerSpaceSpaceSettingsCustomFileSystemsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerSpace.SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference_Override(s SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerSpace.SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) PutEfsFileSystem(value *SagemakerSpaceSpaceSettingsCustomFileSystemsEfsFileSystem) {
	if err := s.validatePutEfsFileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEfsFileSystem",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) PutFsxLustreFileSystem(value *SagemakerSpaceSpaceSettingsCustomFileSystemsFsxLustreFileSystem) {
	if err := s.validatePutFsxLustreFileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFsxLustreFileSystem",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) PutS3FileSystem(value *SagemakerSpaceSpaceSettingsCustomFileSystemsS3FileSystem) {
	if err := s.validatePutS3FileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putS3FileSystem",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ResetEfsFileSystem() {
	_jsii_.InvokeVoid(
		s,
		"resetEfsFileSystem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ResetFsxLustreFileSystem() {
	_jsii_.InvokeVoid(
		s,
		"resetFsxLustreFileSystem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ResetS3FileSystem() {
	_jsii_.InvokeVoid(
		s,
		"resetS3FileSystem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerSpaceSpaceSettingsCustomFileSystemsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

