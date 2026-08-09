// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcorememory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreMemoryMemoryStrategiesOutputReference interface {
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
	CustomMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyOutputReference
	CustomMemoryStrategyInput() interface{}
	EpisodicMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyOutputReference
	EpisodicMemoryStrategyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SemanticMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategyOutputReference
	SemanticMemoryStrategyInput() interface{}
	SummaryMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesSummaryMemoryStrategyOutputReference
	SummaryMemoryStrategyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserPreferenceMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesUserPreferenceMemoryStrategyOutputReference
	UserPreferenceMemoryStrategyInput() interface{}
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
	PutCustomMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategy)
	PutEpisodicMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategy)
	PutSemanticMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategy)
	PutSummaryMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesSummaryMemoryStrategy)
	PutUserPreferenceMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesUserPreferenceMemoryStrategy)
	ResetCustomMemoryStrategy()
	ResetEpisodicMemoryStrategy()
	ResetSemanticMemoryStrategy()
	ResetSummaryMemoryStrategy()
	ResetUserPreferenceMemoryStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreMemoryMemoryStrategiesOutputReference
type jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) CustomMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategyOutputReference
	_jsii_.Get(
		j,
		"customMemoryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) CustomMemoryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMemoryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) EpisodicMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyOutputReference
	_jsii_.Get(
		j,
		"episodicMemoryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) EpisodicMemoryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"episodicMemoryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) SemanticMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategyOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategyOutputReference
	_jsii_.Get(
		j,
		"semanticMemoryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) SemanticMemoryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticMemoryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) SummaryMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesSummaryMemoryStrategyOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesSummaryMemoryStrategyOutputReference
	_jsii_.Get(
		j,
		"summaryMemoryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) SummaryMemoryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summaryMemoryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) UserPreferenceMemoryStrategy() BedrockagentcoreMemoryMemoryStrategiesUserPreferenceMemoryStrategyOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesUserPreferenceMemoryStrategyOutputReference
	_jsii_.Get(
		j,
		"userPreferenceMemoryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) UserPreferenceMemoryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPreferenceMemoryStrategyInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreMemoryMemoryStrategiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreMemoryMemoryStrategiesOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreMemoryMemoryStrategiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreMemoryMemoryStrategiesOutputReference_Override(b BedrockagentcoreMemoryMemoryStrategiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) PutCustomMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesCustomMemoryStrategy) {
	if err := b.validatePutCustomMemoryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCustomMemoryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) PutEpisodicMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategy) {
	if err := b.validatePutEpisodicMemoryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEpisodicMemoryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) PutSemanticMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesSemanticMemoryStrategy) {
	if err := b.validatePutSemanticMemoryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSemanticMemoryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) PutSummaryMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesSummaryMemoryStrategy) {
	if err := b.validatePutSummaryMemoryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSummaryMemoryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) PutUserPreferenceMemoryStrategy(value *BedrockagentcoreMemoryMemoryStrategiesUserPreferenceMemoryStrategy) {
	if err := b.validatePutUserPreferenceMemoryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserPreferenceMemoryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ResetCustomMemoryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomMemoryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ResetEpisodicMemoryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetEpisodicMemoryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ResetSemanticMemoryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetSemanticMemoryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ResetSummaryMemoryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetSummaryMemoryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ResetUserPreferenceMemoryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetUserPreferenceMemoryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

