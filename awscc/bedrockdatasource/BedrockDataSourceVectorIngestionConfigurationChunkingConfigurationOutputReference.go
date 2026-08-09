// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference interface {
	cdktn.ComplexObject
	ChunkingStrategy() *string
	SetChunkingStrategy(val *string)
	ChunkingStrategyInput() *string
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
	FixedSizeChunkingConfiguration() BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfigurationOutputReference
	FixedSizeChunkingConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	HierarchicalChunkingConfiguration() BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationOutputReference
	HierarchicalChunkingConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SemanticChunkingConfiguration() BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfigurationOutputReference
	SemanticChunkingConfigurationInput() interface{}
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
	PutFixedSizeChunkingConfiguration(value *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration)
	PutHierarchicalChunkingConfiguration(value *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration)
	PutSemanticChunkingConfiguration(value *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfiguration)
	ResetChunkingStrategy()
	ResetFixedSizeChunkingConfiguration()
	ResetHierarchicalChunkingConfiguration()
	ResetSemanticChunkingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference
type jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ChunkingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chunkingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ChunkingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chunkingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) FixedSizeChunkingConfiguration() BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfigurationOutputReference {
	var returns BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfigurationOutputReference
	_jsii_.Get(
		j,
		"fixedSizeChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) FixedSizeChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedSizeChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) HierarchicalChunkingConfiguration() BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationOutputReference {
	var returns BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfigurationOutputReference
	_jsii_.Get(
		j,
		"hierarchicalChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) HierarchicalChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hierarchicalChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) SemanticChunkingConfiguration() BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfigurationOutputReference {
	var returns BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfigurationOutputReference
	_jsii_.Get(
		j,
		"semanticChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) SemanticChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference_Override(b BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetChunkingStrategy(val *string) {
	if err := j.validateSetChunkingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chunkingStrategy",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) PutFixedSizeChunkingConfiguration(value *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration) {
	if err := b.validatePutFixedSizeChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFixedSizeChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) PutHierarchicalChunkingConfiguration(value *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration) {
	if err := b.validatePutHierarchicalChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putHierarchicalChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) PutSemanticChunkingConfiguration(value *BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationSemanticChunkingConfiguration) {
	if err := b.validatePutSemanticChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSemanticChunkingConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetChunkingStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetChunkingStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetFixedSizeChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetFixedSizeChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetHierarchicalChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetHierarchicalChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ResetSemanticChunkingConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetSemanticChunkingConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationChunkingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

