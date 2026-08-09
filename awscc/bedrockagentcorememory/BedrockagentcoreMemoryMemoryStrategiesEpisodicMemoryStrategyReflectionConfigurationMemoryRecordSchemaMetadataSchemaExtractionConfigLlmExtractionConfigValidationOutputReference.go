// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorememory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockagentcorememory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference interface {
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
	NumberValidation() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationNumberValidationOutputReference
	NumberValidationInput() interface{}
	StringListValidation() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringListValidationOutputReference
	StringListValidationInput() interface{}
	StringValidation() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringValidationOutputReference
	StringValidationInput() interface{}
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
	PutNumberValidation(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationNumberValidation)
	PutStringListValidation(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringListValidation)
	PutStringValidation(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringValidation)
	ResetNumberValidation()
	ResetStringListValidation()
	ResetStringValidation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference
type jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) NumberValidation() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationNumberValidationOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationNumberValidationOutputReference
	_jsii_.Get(
		j,
		"numberValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) NumberValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) StringListValidation() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringListValidationOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringListValidationOutputReference
	_jsii_.Get(
		j,
		"stringListValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) StringListValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringListValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) StringValidation() BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringValidationOutputReference {
	var returns BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringValidationOutputReference
	_jsii_.Get(
		j,
		"stringValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) StringValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference_Override(b BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockagentcoreMemory.BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) PutNumberValidation(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationNumberValidation) {
	if err := b.validatePutNumberValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNumberValidation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) PutStringListValidation(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringListValidation) {
	if err := b.validatePutStringListValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringListValidation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) PutStringValidation(value *BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationStringValidation) {
	if err := b.validatePutStringValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStringValidation",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ResetNumberValidation() {
	_jsii_.InvokeVoid(
		b,
		"resetNumberValidation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ResetStringListValidation() {
	_jsii_.InvokeVoid(
		b,
		"resetStringListValidation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ResetStringValidation() {
	_jsii_.InvokeVoid(
		b,
		"resetStringValidation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockagentcoreMemoryMemoryStrategiesEpisodicMemoryStrategyReflectionConfigurationMemoryRecordSchemaMetadataSchemaExtractionConfigLlmExtractionConfigValidationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

