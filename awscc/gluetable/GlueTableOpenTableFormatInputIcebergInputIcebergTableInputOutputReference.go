// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference interface {
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	PartitionSpec() GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecOutputReference
	PartitionSpecInput() interface{}
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	Schema() GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference
	SchemaInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WriteOrder() GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderOutputReference
	WriteOrderInput() interface{}
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
	PutPartitionSpec(value *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec)
	PutSchema(value *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchema)
	PutWriteOrder(value *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrder)
	ResetLocation()
	ResetPartitionSpec()
	ResetProperties()
	ResetSchema()
	ResetWriteOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference
type jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PartitionSpec() GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecOutputReference {
	var returns GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecOutputReference
	_jsii_.Get(
		j,
		"partitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PartitionSpecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Schema() GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference {
	var returns GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchemaOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) SchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) WriteOrder() GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderOutputReference {
	var returns GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderOutputReference
	_jsii_.Get(
		j,
		"writeOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) WriteOrderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeOrderInput",
		&returns,
	)
	return returns
}


func NewGlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference_Override(g GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PutPartitionSpec(value *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputPartitionSpec) {
	if err := g.validatePutPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitionSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PutSchema(value *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputSchema) {
	if err := g.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchema",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) PutWriteOrder(value *GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrder) {
	if err := g.validatePutWriteOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWriteOrder",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetPartitionSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ResetWriteOrder() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteOrder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

