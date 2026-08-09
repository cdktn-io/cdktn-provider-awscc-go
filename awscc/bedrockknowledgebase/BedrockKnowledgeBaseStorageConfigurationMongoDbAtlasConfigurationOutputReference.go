// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference interface {
	cdktn.ComplexObject
	CollectionName() *string
	SetCollectionName(val *string)
	CollectionNameInput() *string
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
	CredentialsSecretArn() *string
	SetCredentialsSecretArn(val *string)
	CredentialsSecretArnInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	EndpointServiceName() *string
	SetEndpointServiceName(val *string)
	EndpointServiceNameInput() *string
	FieldMapping() BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationFieldMappingOutputReference
	FieldMappingInput() interface{}
	// Experimental.
	Fqn() *string
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
	TextIndexName() *string
	SetTextIndexName(val *string)
	TextIndexNameInput() *string
	VectorIndexName() *string
	SetVectorIndexName(val *string)
	VectorIndexNameInput() *string
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
	PutFieldMapping(value *BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationFieldMapping)
	ResetCollectionName()
	ResetCredentialsSecretArn()
	ResetDatabaseName()
	ResetEndpoint()
	ResetEndpointServiceName()
	ResetFieldMapping()
	ResetTextIndexName()
	ResetVectorIndexName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference
type jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) CollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) CollectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) CredentialsSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) CredentialsSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) EndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) EndpointServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) FieldMapping() BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationFieldMappingOutputReference {
	var returns BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationFieldMappingOutputReference
	_jsii_.Get(
		j,
		"fieldMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) FieldMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) TextIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textIndexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) TextIndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textIndexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) VectorIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) VectorIndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexNameInput",
		&returns,
	)
	return returns
}


func NewBedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference_Override(b BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetCollectionName(val *string) {
	if err := j.validateSetCollectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionName",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetCredentialsSecretArn(val *string) {
	if err := j.validateSetCredentialsSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsSecretArn",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetEndpointServiceName(val *string) {
	if err := j.validateSetEndpointServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointServiceName",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetTextIndexName(val *string) {
	if err := j.validateSetTextIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textIndexName",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference)SetVectorIndexName(val *string) {
	if err := j.validateSetVectorIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vectorIndexName",
		val,
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) PutFieldMapping(value *BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationFieldMapping) {
	if err := b.validatePutFieldMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFieldMapping",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetCollectionName() {
	_jsii_.InvokeVoid(
		b,
		"resetCollectionName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetCredentialsSecretArn() {
	_jsii_.InvokeVoid(
		b,
		"resetCredentialsSecretArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		b,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetEndpointServiceName() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpointServiceName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetFieldMapping() {
	_jsii_.InvokeVoid(
		b,
		"resetFieldMapping",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetTextIndexName() {
	_jsii_.InvokeVoid(
		b,
		"resetTextIndexName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ResetVectorIndexName() {
	_jsii_.InvokeVoid(
		b,
		"resetVectorIndexName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

