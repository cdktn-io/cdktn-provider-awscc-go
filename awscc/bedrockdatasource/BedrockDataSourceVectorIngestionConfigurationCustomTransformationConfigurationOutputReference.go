// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference interface {
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
	IntermediateStorage() BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorageOutputReference
	IntermediateStorageInput() interface{}
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
	Transformations() BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsList
	TransformationsInput() interface{}
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
	PutIntermediateStorage(value *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorage)
	PutTransformations(value interface{})
	ResetIntermediateStorage()
	ResetTransformations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference
type jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) IntermediateStorage() BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorageOutputReference {
	var returns BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorageOutputReference
	_jsii_.Get(
		j,
		"intermediateStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) IntermediateStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intermediateStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) Transformations() BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsList {
	var returns BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsList
	_jsii_.Get(
		j,
		"transformations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) TransformationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformationsInput",
		&returns,
	)
	return returns
}


func NewBedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference_Override(b BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataSource.BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) PutIntermediateStorage(value *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorage) {
	if err := b.validatePutIntermediateStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putIntermediateStorage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) PutTransformations(value interface{}) {
	if err := b.validatePutTransformationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTransformations",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) ResetIntermediateStorage() {
	_jsii_.InvokeVoid(
		b,
		"resetIntermediateStorage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) ResetTransformations() {
	_jsii_.InvokeVoid(
		b,
		"resetTransformations",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

