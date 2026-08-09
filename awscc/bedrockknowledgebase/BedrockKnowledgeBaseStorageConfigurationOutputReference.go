// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockKnowledgeBaseStorageConfigurationOutputReference interface {
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
	MongoDbAtlasConfiguration() BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference
	MongoDbAtlasConfigurationInput() interface{}
	NeptuneAnalyticsConfiguration() BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationOutputReference
	NeptuneAnalyticsConfigurationInput() interface{}
	OpensearchManagedClusterConfiguration() BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationOutputReference
	OpensearchManagedClusterConfigurationInput() interface{}
	OpensearchServerlessConfiguration() BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationOutputReference
	OpensearchServerlessConfigurationInput() interface{}
	PineconeConfiguration() BedrockKnowledgeBaseStorageConfigurationPineconeConfigurationOutputReference
	PineconeConfigurationInput() interface{}
	RdsConfiguration() BedrockKnowledgeBaseStorageConfigurationRdsConfigurationOutputReference
	RdsConfigurationInput() interface{}
	S3VectorsConfiguration() BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference
	S3VectorsConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutMongoDbAtlasConfiguration(value *BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfiguration)
	PutNeptuneAnalyticsConfiguration(value *BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfiguration)
	PutOpensearchManagedClusterConfiguration(value *BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfiguration)
	PutOpensearchServerlessConfiguration(value *BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration)
	PutPineconeConfiguration(value *BedrockKnowledgeBaseStorageConfigurationPineconeConfiguration)
	PutRdsConfiguration(value *BedrockKnowledgeBaseStorageConfigurationRdsConfiguration)
	PutS3VectorsConfiguration(value *BedrockKnowledgeBaseStorageConfigurationS3VectorsConfiguration)
	ResetMongoDbAtlasConfiguration()
	ResetNeptuneAnalyticsConfiguration()
	ResetOpensearchManagedClusterConfiguration()
	ResetOpensearchServerlessConfiguration()
	ResetPineconeConfiguration()
	ResetRdsConfiguration()
	ResetS3VectorsConfiguration()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockKnowledgeBaseStorageConfigurationOutputReference
type jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) MongoDbAtlasConfiguration() BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference
	_jsii_.Get(
		j,
		"mongoDbAtlasConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) MongoDbAtlasConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbAtlasConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) NeptuneAnalyticsConfiguration() BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationOutputReference
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) NeptuneAnalyticsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) OpensearchManagedClusterConfiguration() BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationOutputReference
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) OpensearchManagedClusterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) OpensearchServerlessConfiguration() BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationOutputReference
	_jsii_.Get(
		j,
		"opensearchServerlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) OpensearchServerlessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchServerlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PineconeConfiguration() BedrockKnowledgeBaseStorageConfigurationPineconeConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationPineconeConfigurationOutputReference
	_jsii_.Get(
		j,
		"pineconeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PineconeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pineconeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) RdsConfiguration() BedrockKnowledgeBaseStorageConfigurationRdsConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationRdsConfigurationOutputReference
	_jsii_.Get(
		j,
		"rdsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) RdsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) S3VectorsConfiguration() BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3VectorsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) S3VectorsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3VectorsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBedrockKnowledgeBaseStorageConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockKnowledgeBaseStorageConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockKnowledgeBaseStorageConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockKnowledgeBaseStorageConfigurationOutputReference_Override(b BedrockKnowledgeBaseStorageConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutMongoDbAtlasConfiguration(value *BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfiguration) {
	if err := b.validatePutMongoDbAtlasConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMongoDbAtlasConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutNeptuneAnalyticsConfiguration(value *BedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfiguration) {
	if err := b.validatePutNeptuneAnalyticsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNeptuneAnalyticsConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutOpensearchManagedClusterConfiguration(value *BedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfiguration) {
	if err := b.validatePutOpensearchManagedClusterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOpensearchManagedClusterConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutOpensearchServerlessConfiguration(value *BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration) {
	if err := b.validatePutOpensearchServerlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOpensearchServerlessConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutPineconeConfiguration(value *BedrockKnowledgeBaseStorageConfigurationPineconeConfiguration) {
	if err := b.validatePutPineconeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPineconeConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutRdsConfiguration(value *BedrockKnowledgeBaseStorageConfigurationRdsConfiguration) {
	if err := b.validatePutRdsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRdsConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) PutS3VectorsConfiguration(value *BedrockKnowledgeBaseStorageConfigurationS3VectorsConfiguration) {
	if err := b.validatePutS3VectorsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3VectorsConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetMongoDbAtlasConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetMongoDbAtlasConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetNeptuneAnalyticsConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetNeptuneAnalyticsConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetOpensearchManagedClusterConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetOpensearchManagedClusterConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetOpensearchServerlessConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetOpensearchServerlessConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetPineconeConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetPineconeConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetRdsConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetRdsConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetS3VectorsConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetS3VectorsConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

