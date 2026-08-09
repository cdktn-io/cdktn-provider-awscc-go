// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressdirectorybucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3expressdirectorybucket/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference interface {
	cdktn.ComplexObject
	AbortIncompleteMultipartUpload() S3ExpressDirectoryBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
	AbortIncompleteMultipartUploadInput() interface{}
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
	ExpirationInDays() *float64
	SetExpirationInDays(val *float64)
	ExpirationInDaysInput() *float64
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ObjectSizeGreaterThan() *string
	SetObjectSizeGreaterThan(val *string)
	ObjectSizeGreaterThanInput() *string
	ObjectSizeLessThan() *string
	SetObjectSizeLessThan(val *string)
	ObjectSizeLessThanInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	PutAbortIncompleteMultipartUpload(value *S3ExpressDirectoryBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload)
	ResetAbortIncompleteMultipartUpload()
	ResetExpirationInDays()
	ResetId()
	ResetObjectSizeGreaterThan()
	ResetObjectSizeLessThan()
	ResetPrefix()
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference
type jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUpload() S3ExpressDirectoryBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference {
	var returns S3ExpressDirectoryBucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ObjectSizeGreaterThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ObjectSizeGreaterThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ObjectSizeLessThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ObjectSizeLessThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewS3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3ExpressDirectoryBucket.S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference_Override(s S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3ExpressDirectoryBucket.S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetExpirationInDays(val *float64) {
	if err := j.validateSetExpirationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationInDays",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetObjectSizeGreaterThan(val *string) {
	if err := j.validateSetObjectSizeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectSizeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetObjectSizeLessThan(val *string) {
	if err := j.validateSetObjectSizeLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectSizeLessThan",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) PutAbortIncompleteMultipartUpload(value *S3ExpressDirectoryBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload) {
	if err := s.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAbortIncompleteMultipartUpload",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetExpirationInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetExpirationInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetObjectSizeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectSizeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetObjectSizeLessThan() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectSizeLessThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ExpressDirectoryBucketLifecycleConfigurationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

