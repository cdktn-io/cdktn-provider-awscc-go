// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference interface {
	cdktn.ComplexObject
	BetweenConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference
	BetweenConfigurationInput() interface{}
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
	DateTimeFormat() *string
	SetDateTimeFormat(val *string)
	DateTimeFormatInput() *string
	FilterMode() *string
	SetFilterMode(val *string)
	FilterModeInput() *string
	FilterStringConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationFilterStringConfigurationOutputReference
	FilterStringConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OperatorMappings() *map[string]*string
	SetOperatorMappings(val *map[string]*string)
	OperatorMappingsInput() *map[string]*string
	StripQuotes() interface{}
	SetStripQuotes(val interface{})
	StripQuotesInput() interface{}
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
	PutBetweenConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfiguration)
	PutFilterStringConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationFilterStringConfiguration)
	ResetBetweenConfiguration()
	ResetDateTimeFormat()
	ResetFilterMode()
	ResetFilterStringConfiguration()
	ResetOperatorMappings()
	ResetStripQuotes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) BetweenConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfigurationOutputReference
	_jsii_.Get(
		j,
		"betweenConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) BetweenConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"betweenConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) DateTimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateTimeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) DateTimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateTimeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) FilterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) FilterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) FilterStringConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationFilterStringConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationFilterStringConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterStringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) FilterStringConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterStringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) OperatorMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"operatorMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) OperatorMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"operatorMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) StripQuotes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripQuotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) StripQuotesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripQuotesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference_Override(g GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetDateTimeFormat(val *string) {
	if err := j.validateSetDateTimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateTimeFormat",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetFilterMode(val *string) {
	if err := j.validateSetFilterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterMode",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetOperatorMappings(val *map[string]*string) {
	if err := j.validateSetOperatorMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatorMappings",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetStripQuotes(val interface{}) {
	if err := j.validateSetStripQuotesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripQuotes",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) PutBetweenConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationBetweenConfiguration) {
	if err := g.validatePutBetweenConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBetweenConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) PutFilterStringConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationFilterStringConfiguration) {
	if err := g.validatePutFilterStringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilterStringConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ResetBetweenConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetBetweenConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ResetDateTimeFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetDateTimeFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ResetFilterMode() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ResetFilterStringConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterStringConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ResetOperatorMappings() {
	_jsii_.InvokeVoid(
		g,
		"resetOperatorMappings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ResetStripQuotes() {
	_jsii_.InvokeVoid(
		g,
		"resetStripQuotes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

