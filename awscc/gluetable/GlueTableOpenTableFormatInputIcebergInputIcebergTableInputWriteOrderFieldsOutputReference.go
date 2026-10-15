// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference interface {
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
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NullOrder() *string
	SetNullOrder(val *string)
	NullOrderInput() *string
	SourceId() *float64
	SetSourceId(val *float64)
	SourceIdInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Transform() *string
	SetTransform(val *string)
	TransformInput() *string
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
	ResetDirection()
	ResetNullOrder()
	ResetSourceId()
	ResetTransform()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference
type jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) NullOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) NullOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) SourceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) SourceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) Transform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) TransformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}


func NewGlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference_Override(g GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetNullOrder(val *string) {
	if err := j.validateSetNullOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullOrder",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetSourceId(val *float64) {
	if err := j.validateSetSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference)SetTransform(val *string) {
	if err := j.validateSetTransformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transform",
		val,
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		g,
		"resetDirection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ResetNullOrder() {
	_jsii_.InvokeVoid(
		g,
		"resetNullOrder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ResetSourceId() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ResetTransform() {
	_jsii_.InvokeVoid(
		g,
		"resetTransform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueTableOpenTableFormatInputIcebergInputIcebergTableInputWriteOrderFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

