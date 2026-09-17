// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference interface {
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
	HighBoundKey() *string
	SetHighBoundKey(val *string)
	HighBoundKeyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LowBoundKey() *string
	SetLowBoundKey(val *string)
	LowBoundKeyInput() *string
	Template() *string
	SetTemplate(val *string)
	TemplateInput() *string
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
	ResetHighBoundKey()
	ResetLowBoundKey()
	ResetTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) HighBoundKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highBoundKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) HighBoundKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"highBoundKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) LowBoundKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lowBoundKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) LowBoundKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lowBoundKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) Template() *string {
	var returns *string
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) TemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference_Override(g GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetHighBoundKey(val *string) {
	if err := j.validateSetHighBoundKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highBoundKey",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetLowBoundKey(val *string) {
	if err := j.validateSetLowBoundKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lowBoundKey",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetTemplate(val *string) {
	if err := j.validateSetTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"template",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ResetHighBoundKey() {
	_jsii_.InvokeVoid(
		g,
		"resetHighBoundKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ResetLowBoundKey() {
	_jsii_.InvokeVoid(
		g,
		"resetLowBoundKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ResetTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

