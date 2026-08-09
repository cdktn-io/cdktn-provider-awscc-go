// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccs3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccs3bucket/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccS3BucketLifecycleConfigurationRulesOutputReference interface {
	cdktn.ComplexObject
	AbortIncompleteMultipartUpload() DataAwsccS3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
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
	ExpirationInDays() *float64
	ExpiredObjectDeleteMarker() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataAwsccS3BucketLifecycleConfigurationRules
	SetInternalValue(val *DataAwsccS3BucketLifecycleConfigurationRules)
	NoncurrentVersionExpiration() DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionExpirationOutputReference
	NoncurrentVersionExpirationInDays() *float64
	NoncurrentVersionTransition() DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionOutputReference
	NoncurrentVersionTransitions() DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionsList
	ObjectSizeGreaterThan() *string
	ObjectSizeLessThan() *string
	Prefix() *string
	Status() *string
	TagFilters() DataAwsccS3BucketLifecycleConfigurationRulesTagFiltersList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Transition() DataAwsccS3BucketLifecycleConfigurationRulesTransitionOutputReference
	Transitions() DataAwsccS3BucketLifecycleConfigurationRulesTransitionsList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccS3BucketLifecycleConfigurationRulesOutputReference
type jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) AbortIncompleteMultipartUpload() DataAwsccS3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUploadOutputReference
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ExpiredObjectDeleteMarker() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"expiredObjectDeleteMarker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) InternalValue() *DataAwsccS3BucketLifecycleConfigurationRules {
	var returns *DataAwsccS3BucketLifecycleConfigurationRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionExpiration() DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionExpirationOutputReference {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionExpirationOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionTransition() DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionOutputReference {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) NoncurrentVersionTransitions() DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionsList {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesNoncurrentVersionTransitionsList
	_jsii_.Get(
		j,
		"noncurrentVersionTransitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ObjectSizeGreaterThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ObjectSizeLessThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectSizeLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) TagFilters() DataAwsccS3BucketLifecycleConfigurationRulesTagFiltersList {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesTagFiltersList
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Transition() DataAwsccS3BucketLifecycleConfigurationRulesTransitionOutputReference {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesTransitionOutputReference
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Transitions() DataAwsccS3BucketLifecycleConfigurationRulesTransitionsList {
	var returns DataAwsccS3BucketLifecycleConfigurationRulesTransitionsList
	_jsii_.Get(
		j,
		"transitions",
		&returns,
	)
	return returns
}


func NewDataAwsccS3BucketLifecycleConfigurationRulesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccS3BucketLifecycleConfigurationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccS3BucketLifecycleConfigurationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3BucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccS3BucketLifecycleConfigurationRulesOutputReference_Override(d DataAwsccS3BucketLifecycleConfigurationRulesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3BucketLifecycleConfigurationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference)SetInternalValue(val *DataAwsccS3BucketLifecycleConfigurationRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3BucketLifecycleConfigurationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

