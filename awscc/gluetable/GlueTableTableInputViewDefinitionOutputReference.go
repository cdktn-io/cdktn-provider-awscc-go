// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/gluetable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueTableTableInputViewDefinitionOutputReference interface {
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
	Definer() *string
	SetDefiner(val *string)
	DefinerInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsProtected() interface{}
	SetIsProtected(val interface{})
	IsProtectedInput() interface{}
	Representations() GlueTableTableInputViewDefinitionRepresentationsList
	RepresentationsInput() interface{}
	SubObjects() *[]*string
	SetSubObjects(val *[]*string)
	SubObjectsInput() *[]*string
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
	PutRepresentations(value interface{})
	ResetDefiner()
	ResetIsProtected()
	ResetRepresentations()
	ResetSubObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueTableTableInputViewDefinitionOutputReference
type jsiiProxy_GlueTableTableInputViewDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) Definer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) DefinerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) IsProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) IsProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) Representations() GlueTableTableInputViewDefinitionRepresentationsList {
	var returns GlueTableTableInputViewDefinitionRepresentationsList
	_jsii_.Get(
		j,
		"representations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) RepresentationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"representationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) SubObjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) SubObjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueTableTableInputViewDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueTableTableInputViewDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewGlueTableTableInputViewDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueTableTableInputViewDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputViewDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueTableTableInputViewDefinitionOutputReference_Override(g GlueTableTableInputViewDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueTable.GlueTableTableInputViewDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetDefiner(val *string) {
	if err := j.validateSetDefinerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definer",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetIsProtected(val interface{}) {
	if err := j.validateSetIsProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isProtected",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetSubObjects(val *[]*string) {
	if err := j.validateSetSubObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subObjects",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) PutRepresentations(value interface{}) {
	if err := g.validatePutRepresentationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRepresentations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ResetDefiner() {
	_jsii_.InvokeVoid(
		g,
		"resetDefiner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ResetIsProtected() {
	_jsii_.InvokeVoid(
		g,
		"resetIsProtected",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ResetRepresentations() {
	_jsii_.InvokeVoid(
		g,
		"resetRepresentations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ResetSubObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetSubObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueTableTableInputViewDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

