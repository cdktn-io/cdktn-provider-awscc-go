// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference interface {
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
	FilterConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference
	FilterConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PaginationConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOutputReference
	PaginationConfigurationInput() interface{}
	RequestMethod() *string
	SetRequestMethod(val *string)
	RequestMethodInput() *string
	RequestParameters() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationRequestParametersList
	RequestParametersInput() interface{}
	RequestPath() *string
	SetRequestPath(val *string)
	RequestPathInput() *string
	ResponseConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfigurationOutputReference
	ResponseConfigurationInput() interface{}
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
	PutFilterConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfiguration)
	PutPaginationConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfiguration)
	PutRequestParameters(value interface{})
	PutResponseConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfiguration)
	ResetFilterConfiguration()
	ResetPaginationConfiguration()
	ResetRequestMethod()
	ResetRequestParameters()
	ResetRequestPath()
	ResetResponseConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) FilterConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfigurationOutputReference
	_jsii_.Get(
		j,
		"filterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) FilterConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) PaginationConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOutputReference
	_jsii_.Get(
		j,
		"paginationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) PaginationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paginationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) RequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) RequestParameters() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationRequestParametersList {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationRequestParametersList
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) RequestParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) RequestPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) RequestPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResponseConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfigurationOutputReference
	_jsii_.Get(
		j,
		"responseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResponseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference_Override(g GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetRequestMethod(val *string) {
	if err := j.validateSetRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetRequestPath(val *string) {
	if err := j.validateSetRequestPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPath",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) PutFilterConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationFilterConfiguration) {
	if err := g.validatePutFilterConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilterConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) PutPaginationConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfiguration) {
	if err := g.validatePutPaginationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPaginationConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) PutRequestParameters(value interface{}) {
	if err := g.validatePutRequestParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) PutResponseConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationResponseConfiguration) {
	if err := g.validatePutResponseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResponseConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResetFilterConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResetPaginationConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetPaginationConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResetRequestMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResetRequestPath() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ResetResponseConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetResponseConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

