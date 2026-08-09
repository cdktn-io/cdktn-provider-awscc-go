// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcoreharness/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference interface {
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
	RelevanceScore() *float64
	SetRelevanceScore(val *float64)
	RelevanceScoreInput() *float64
	StrategyId() *string
	SetStrategyId(val *string)
	StrategyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TopK() *float64
	SetTopK(val *float64)
	TopKInput() *float64
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
	ResetRelevanceScore()
	ResetStrategyId()
	ResetTopK()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference
type jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) RelevanceScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"relevanceScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) RelevanceScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"relevanceScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) StrategyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) StrategyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) TopK() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topK",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) TopKInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topKInput",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreHarness.BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference_Override(b BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreHarness.BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetRelevanceScore(val *float64) {
	if err := j.validateSetRelevanceScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relevanceScore",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetStrategyId(val *string) {
	if err := j.validateSetStrategyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategyId",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference)SetTopK(val *float64) {
	if err := j.validateSetTopKParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topK",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ResetRelevanceScore() {
	_jsii_.InvokeVoid(
		b,
		"resetRelevanceScore",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ResetStrategyId() {
	_jsii_.InvokeVoid(
		b,
		"resetStrategyId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ResetTopK() {
	_jsii_.InvokeVoid(
		b,
		"resetTopK",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

