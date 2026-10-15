// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableTableInputStorageDescriptorSkewedInfoOutputReference interface {
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
	SkewedColumnNames() *[]*string
	SetSkewedColumnNames(val *[]*string)
	SkewedColumnNamesInput() *[]*string
	SkewedColumnValueLocationMaps() *string
	SetSkewedColumnValueLocationMaps(val *string)
	SkewedColumnValueLocationMapsInput() *string
	SkewedColumnValues() *[]*string
	SetSkewedColumnValues(val *[]*string)
	SkewedColumnValuesInput() *[]*string
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
	ResetSkewedColumnNames()
	ResetSkewedColumnValueLocationMaps()
	ResetSkewedColumnValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableTableInputStorageDescriptorSkewedInfoOutputReference
type jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) SkewedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) SkewedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMaps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) SkewedColumnValueLocationMapsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) SkewedColumnValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) SkewedColumnValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueTableTableInputStorageDescriptorSkewedInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueTableTableInputStorageDescriptorSkewedInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableTableInputStorageDescriptorSkewedInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueTableTableInputStorageDescriptorSkewedInfoOutputReference_Override(g GlueTableTableInputStorageDescriptorSkewedInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputStorageDescriptorSkewedInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetSkewedColumnNames(val *[]*string) {
	if err := j.validateSetSkewedColumnNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnNames",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetSkewedColumnValueLocationMaps(val *string) {
	if err := j.validateSetSkewedColumnValueLocationMapsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnValueLocationMaps",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetSkewedColumnValues(val *[]*string) {
	if err := j.validateSetSkewedColumnValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnValues",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnNames() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValueLocationMaps() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValueLocationMaps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ResetSkewedColumnValues() {
	_jsii_.InvokeVoid(
		g,
		"resetSkewedColumnValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueTableTableInputStorageDescriptorSkewedInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

