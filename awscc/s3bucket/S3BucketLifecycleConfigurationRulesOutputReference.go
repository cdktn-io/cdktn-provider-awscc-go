// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/s3bucket/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3BucketLifecycleConfigurationRulesOutputReference interface {
	cdktn.ComplexObject
	AbortIncompleteMultipartUpload() S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
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
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	ExpirationInDays() *float64
	SetExpirationInDays(val *float64)
	ExpirationInDaysInput() *float64
	ExpiredObjectDeleteMarker() interface{}
	SetExpiredObjectDeleteMarker(val interface{})
	ExpiredObjectDeleteMarkerInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoncurrentVersionExpiration() S3BucketLifecycleConfigurationRulesNoncurrentVersionExpirationOutputReference
	NoncurrentVersionExpirationInDays() *float64
	SetNoncurrentVersionExpirationInDays(val *float64)
	NoncurrentVersionExpirationInDaysInput() *float64
	NoncurrentVersionExpirationInput() interface{}
	NoncurrentVersionTransition() S3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionOutputReference
	NoncurrentVersionTransitionInput() interface{}
	NoncurrentVersionTransitions() S3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionsList
	NoncurrentVersionTransitionsInput() interface{}
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
	TagFilters() S3BucketLifecycleConfigurationRulesTagFiltersList
	TagFiltersInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Transition() S3BucketLifecycleConfigurationRulesTransitionOutputReference
	TransitionInput() interface{}
	Transitions() S3BucketLifecycleConfigurationRulesTransitionsList
	TransitionsInput() interface{}
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
	PutAbortIncompleteMultipartUpload(value *S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload)
	PutNoncurrentVersionExpiration(value *S3BucketLifecycleConfigurationRulesNoncurrentVersionExpiration)
	PutNoncurrentVersionTransition(value *S3BucketLifecycleConfigurationRulesNoncurrentVersionTransition)
	PutNoncurrentVersionTransitions(value interface{})
	PutTagFilters(value interface{})
	PutTransition(value *S3BucketLifecycleConfigurationRulesTransition)
	PutTransitions(value interface{})
	ResetAbortIncompleteMultipartUpload()
	ResetExpirationDate()
	ResetExpirationInDays()
	ResetExpiredObjectDeleteMarker()
	ResetId()
	ResetNoncurrentVersionExpiration()
	ResetNoncurrentVersionExpirationInDays()
	ResetNoncurrentVersionTransition()
	ResetNoncurrentVersionTransitions()
	ResetObjectSizeGreaterThan()
	ResetObjectSizeLessThan()
	ResetPrefix()
	ResetStatus()
	ResetTagFilters()
	ResetTransition()
	ResetTransitions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketLifecycleConfigurationRulesOutputReference
type jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUpload() S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference {
	var returns S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ExpiredObjectDeleteMarker() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expiredObjectDeleteMarker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ExpiredObjectDeleteMarkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expiredObjectDeleteMarkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionExpiration() S3BucketLifecycleConfigurationRulesNoncurrentVersionExpirationOutputReference {
	var returns S3BucketLifecycleConfigurationRulesNoncurrentVersionExpirationOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionTransition() S3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionOutputReference {
	var returns S3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionTransitions() S3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionsList {
	var returns S3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionsList
	_jsii_.Get(
		j,
		"noncurrentVersionTransitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionTransitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ObjectSizeGreaterThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ObjectSizeGreaterThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ObjectSizeLessThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ObjectSizeLessThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) TagFilters() S3BucketLifecycleConfigurationRulesTagFiltersList {
	var returns S3BucketLifecycleConfigurationRulesTagFiltersList
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) TagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Transition() S3BucketLifecycleConfigurationRulesTransitionOutputReference {
	var returns S3BucketLifecycleConfigurationRulesTransitionOutputReference
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Transitions() S3BucketLifecycleConfigurationRulesTransitionsList {
	var returns S3BucketLifecycleConfigurationRulesTransitionsList
	_jsii_.Get(
		j,
		"transitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) TransitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionsInput",
		&returns,
	)
	return returns
}


func NewS3BucketLifecycleConfigurationRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3BucketLifecycleConfigurationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketLifecycleConfigurationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.s3Bucket.S3BucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3BucketLifecycleConfigurationRulesOutputReference_Override(s S3BucketLifecycleConfigurationRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.s3Bucket.S3BucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetExpirationDate(val *string) {
	if err := j.validateSetExpirationDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetExpirationInDays(val *float64) {
	if err := j.validateSetExpirationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationInDays",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetExpiredObjectDeleteMarker(val interface{}) {
	if err := j.validateSetExpiredObjectDeleteMarkerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiredObjectDeleteMarker",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetNoncurrentVersionExpirationInDays(val *float64) {
	if err := j.validateSetNoncurrentVersionExpirationInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noncurrentVersionExpirationInDays",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetObjectSizeGreaterThan(val *string) {
	if err := j.validateSetObjectSizeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectSizeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetObjectSizeLessThan(val *string) {
	if err := j.validateSetObjectSizeLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectSizeLessThan",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutAbortIncompleteMultipartUpload(value *S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload) {
	if err := s.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAbortIncompleteMultipartUpload",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutNoncurrentVersionExpiration(value *S3BucketLifecycleConfigurationRulesNoncurrentVersionExpiration) {
	if err := s.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutNoncurrentVersionTransition(value *S3BucketLifecycleConfigurationRulesNoncurrentVersionTransition) {
	if err := s.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutNoncurrentVersionTransitions(value interface{}) {
	if err := s.validatePutNoncurrentVersionTransitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionTransitions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutTagFilters(value interface{}) {
	if err := s.validatePutTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTagFilters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutTransition(value *S3BucketLifecycleConfigurationRulesTransition) {
	if err := s.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) PutTransitions(value interface{}) {
	if err := s.validatePutTransitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransitions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		s,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetExpirationInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetExpirationInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetExpiredObjectDeleteMarker() {
	_jsii_.InvokeVoid(
		s,
		"resetExpiredObjectDeleteMarker",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetNoncurrentVersionExpirationInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionExpirationInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetNoncurrentVersionTransitions() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionTransitions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetObjectSizeGreaterThan() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectSizeGreaterThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetObjectSizeLessThan() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectSizeLessThan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetTagFilters() {
	_jsii_.InvokeVoid(
		s,
		"resetTagFilters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ResetTransitions() {
	_jsii_.InvokeVoid(
		s,
		"resetTransitions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3BucketLifecycleConfigurationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

