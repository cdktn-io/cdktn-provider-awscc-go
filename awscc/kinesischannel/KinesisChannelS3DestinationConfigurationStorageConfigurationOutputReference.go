// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesischannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference interface {
	cdktn.ComplexObject
	BucketArn() *string
	SetBucketArn(val *string)
	BucketArnInput() *string
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExpectedBucketOwner() *string
	SetExpectedBucketOwner(val *string)
	ExpectedBucketOwnerInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputKeyTemplate() *string
	SetOutputKeyTemplate(val *string)
	OutputKeyTemplateInput() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	ResetBucketArn()
	ResetCompressionType()
	ResetExpectedBucketOwner()
	ResetOutputKeyTemplate()
	ResetStorageClass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference
type jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) OutputKeyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputKeyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) OutputKeyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputKeyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisChannelS3DestinationConfigurationStorageConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference_Override(k KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetBucketArn(val *string) {
	if err := j.validateSetBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketArn",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetOutputKeyTemplate(val *string) {
	if err := j.validateSetOutputKeyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputKeyTemplate",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ResetBucketArn() {
	_jsii_.InvokeVoid(
		k,
		"resetBucketArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		k,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		k,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ResetOutputKeyTemplate() {
	_jsii_.InvokeVoid(
		k,
		"resetOutputKeyTemplate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		k,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

