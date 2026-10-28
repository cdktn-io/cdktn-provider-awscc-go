// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/bedrockagentcoreevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference interface {
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
	DimensionInstruction() *string
	SetDimensionInstruction(val *string)
	DimensionInstructionInput() *string
	DimensionName() *string
	SetDimensionName(val *string)
	DimensionNameInput() *string
	DimensionRatingScale() BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScaleOutputReference
	DimensionRatingScaleInput() interface{}
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
	PutDimensionRatingScale(value *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScale)
	ResetDimensionInstruction()
	ResetDimensionName()
	ResetDimensionRatingScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference
type jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) DimensionInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) DimensionInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) DimensionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) DimensionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) DimensionRatingScale() BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScaleOutputReference {
	var returns BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScaleOutputReference
	_jsii_.Get(
		j,
		"dimensionRatingScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) DimensionRatingScaleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionRatingScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreEvaluator.BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference_Override(b BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreEvaluator.BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetDimensionInstruction(val *string) {
	if err := j.validateSetDimensionInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionInstruction",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetDimensionName(val *string) {
	if err := j.validateSetDimensionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) PutDimensionRatingScale(value *BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalDimensionRatingScale) {
	if err := b.validatePutDimensionRatingScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDimensionRatingScale",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ResetDimensionInstruction() {
	_jsii_.InvokeVoid(
		b,
		"resetDimensionInstruction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ResetDimensionName() {
	_jsii_.InvokeVoid(
		b,
		"resetDimensionName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ResetDimensionRatingScale() {
	_jsii_.InvokeVoid(
		b,
		"resetDimensionRatingScale",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreEvaluatorEvaluatorConfigLlmAsAJudgeRatingScaleMultiDimensionalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

