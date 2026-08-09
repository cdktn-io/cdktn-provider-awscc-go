// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference interface {
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
	IndexArn() *string
	SetIndexArn(val *string)
	IndexArnInput() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
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
	VectorBucketArn() *string
	SetVectorBucketArn(val *string)
	VectorBucketArnInput() *string
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
	ResetIndexArn()
	ResetIndexName()
	ResetVectorBucketArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference
type jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) IndexArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) IndexArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) VectorBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) VectorBucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorBucketArnInput",
		&returns,
	)
	return returns
}


func NewBedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference_Override(b BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetIndexArn(val *string) {
	if err := j.validateSetIndexArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexArn",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference)SetVectorBucketArn(val *string) {
	if err := j.validateSetVectorBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vectorBucketArn",
		val,
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ResetIndexArn() {
	_jsii_.InvokeVoid(
		b,
		"resetIndexArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		b,
		"resetIndexName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ResetVectorBucketArn() {
	_jsii_.InvokeVoid(
		b,
		"resetVectorBucketArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

