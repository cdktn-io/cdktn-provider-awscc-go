// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableTableInputViewDefinitionRepresentationsOutputReference interface {
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
	Dialect() *string
	SetDialect(val *string)
	DialectInput() *string
	DialectVersion() *string
	SetDialectVersion(val *string)
	DialectVersionInput() *string
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
	ValidationConnection() *string
	SetValidationConnection(val *string)
	ValidationConnectionInput() *string
	ViewExpandedText() *string
	SetViewExpandedText(val *string)
	ViewExpandedTextInput() *string
	ViewOriginalText() *string
	SetViewOriginalText(val *string)
	ViewOriginalTextInput() *string
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
	ResetDialect()
	ResetDialectVersion()
	ResetValidationConnection()
	ResetViewExpandedText()
	ResetViewOriginalText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableTableInputViewDefinitionRepresentationsOutputReference
type jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) Dialect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) DialectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) DialectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) DialectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ValidationConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ValidationConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ViewExpandedText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ViewExpandedTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ViewOriginalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ViewOriginalTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalTextInput",
		&returns,
	)
	return returns
}


func NewGlueTableTableInputViewDefinitionRepresentationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GlueTableTableInputViewDefinitionRepresentationsOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableTableInputViewDefinitionRepresentationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputViewDefinitionRepresentationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGlueTableTableInputViewDefinitionRepresentationsOutputReference_Override(g GlueTableTableInputViewDefinitionRepresentationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputViewDefinitionRepresentationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetDialect(val *string) {
	if err := j.validateSetDialectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialect",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetDialectVersion(val *string) {
	if err := j.validateSetDialectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialectVersion",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetValidationConnection(val *string) {
	if err := j.validateSetValidationConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationConnection",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetViewExpandedText(val *string) {
	if err := j.validateSetViewExpandedTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewExpandedText",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference)SetViewOriginalText(val *string) {
	if err := j.validateSetViewOriginalTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewOriginalText",
		val,
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ResetDialect() {
	_jsii_.InvokeVoid(
		g,
		"resetDialect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ResetDialectVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetDialectVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ResetValidationConnection() {
	_jsii_.InvokeVoid(
		g,
		"resetValidationConnection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ResetViewExpandedText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewExpandedText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ResetViewOriginalText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewOriginalText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionRepresentationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

