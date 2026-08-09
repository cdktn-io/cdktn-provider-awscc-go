// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference interface {
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
	MetadataField() *string
	SetMetadataField(val *string)
	MetadataFieldInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextField() *string
	SetTextField(val *string)
	TextFieldInput() *string
	VectorField() *string
	SetVectorField(val *string)
	VectorFieldInput() *string
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
	ResetMetadataField()
	ResetTextField()
	ResetVectorField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference
type jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) MetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) MetadataFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) TextField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) TextFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) VectorField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) VectorFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorFieldInput",
		&returns,
	)
	return returns
}


func NewBedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference_Override(b BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockKnowledgeBase.BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetMetadataField(val *string) {
	if err := j.validateSetMetadataFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataField",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetTextField(val *string) {
	if err := j.validateSetTextFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textField",
		val,
	)
}

func (j *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference)SetVectorField(val *string) {
	if err := j.validateSetVectorFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vectorField",
		val,
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ResetMetadataField() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadataField",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ResetTextField() {
	_jsii_.InvokeVoid(
		b,
		"resetTextField",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ResetVectorField() {
	_jsii_.InvokeVoid(
		b,
		"resetVectorField",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

