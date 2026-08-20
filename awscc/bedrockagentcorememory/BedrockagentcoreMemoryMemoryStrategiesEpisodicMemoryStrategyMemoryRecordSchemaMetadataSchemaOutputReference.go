// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcorememory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference interface {
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
	ExtractionConfig() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfigOutputReference
	ExtractionConfigInput() interface{}
	ExtractionType() *string
	SetExtractionType(val *string)
	ExtractionTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
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
	PutExtractionConfig(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfig)
	ResetExtractionConfig()
	ResetExtractionType()
	ResetKey()
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

// The jsii proxy struct for BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference
type jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ExtractionConfig() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfigOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfigOutputReference
	_jsii_.Get(
		j,
		"extractionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ExtractionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ExtractionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ExtractionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference_Override(b BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetExtractionType(val *string) {
	if err := j.validateSetExtractionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extractionType",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) PutExtractionConfig(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaExtractionConfig) {
	if err := b.validatePutExtractionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExtractionConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ResetExtractionConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetExtractionConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ResetExtractionType() {
	_jsii_.InvokeVoid(
		b,
		"resetExtractionType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		b,
		"resetKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyMemoryRecordSchemaMetadataSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

