// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluepartition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference interface {
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
	SchemaId() GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaIdOutputReference
	SchemaIdInput() interface{}
	SchemaVersionId() *string
	SetSchemaVersionId(val *string)
	SchemaVersionIdInput() *string
	SchemaVersionNumber() *float64
	SetSchemaVersionNumber(val *float64)
	SchemaVersionNumberInput() *float64
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
	PutSchemaId(value *GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaId)
	ResetSchemaId()
	ResetSchemaVersionId()
	ResetSchemaVersionNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference
type jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) SchemaId() GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaIdOutputReference {
	var returns GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaIdOutputReference
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) SchemaIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) SchemaVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) SchemaVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) SchemaVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) SchemaVersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference {
	_init_.Initialize()

	if err := validateNewGluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.gluePartition.GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference_Override(g GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.gluePartition.GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetSchemaVersionId(val *string) {
	if err := j.validateSetSchemaVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersionId",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetSchemaVersionNumber(val *float64) {
	if err := j.validateSetSchemaVersionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersionNumber",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) PutSchemaId(value *GluePartitionPartitionInputStorageDescriptorSchemaReferenceSchemaId) {
	if err := g.validatePutSchemaIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchemaId",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ResetSchemaVersionId() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaVersionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ResetSchemaVersionNumber() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaVersionNumber",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GluePartitionPartitionInputStorageDescriptorSchemaReferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

