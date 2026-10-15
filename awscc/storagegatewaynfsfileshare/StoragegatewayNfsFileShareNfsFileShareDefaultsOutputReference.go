// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagegatewaynfsfileshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/storagegatewaynfsfileshare/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference interface {
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
	DirectoryMode() *string
	SetDirectoryMode(val *string)
	DirectoryModeInput() *string
	FileMode() *string
	SetFileMode(val *string)
	FileModeInput() *string
	// Experimental.
	Fqn() *string
	GroupId() *float64
	SetGroupId(val *float64)
	GroupIdInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OwnerId() *float64
	SetOwnerId(val *float64)
	OwnerIdInput() *float64
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
	ResetDirectoryMode()
	ResetFileMode()
	ResetGroupId()
	ResetOwnerId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference
type jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) DirectoryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) DirectoryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) FileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) FileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GroupId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GroupIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) OwnerId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) OwnerIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ownerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewStoragegatewayNfsFileShareNfsFileShareDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.storagegatewayNfsFileShare.StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference_Override(s StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.storagegatewayNfsFileShare.StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetDirectoryMode(val *string) {
	if err := j.validateSetDirectoryModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryMode",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetFileMode(val *string) {
	if err := j.validateSetFileModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileMode",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetGroupId(val *float64) {
	if err := j.validateSetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetOwnerId(val *float64) {
	if err := j.validateSetOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetDirectoryMode() {
	_jsii_.InvokeVoid(
		s,
		"resetDirectoryMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetFileMode() {
	_jsii_.InvokeVoid(
		s,
		"resetFileMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetOwnerId() {
	_jsii_.InvokeVoid(
		s,
		"resetOwnerId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

