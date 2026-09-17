// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference interface {
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
	ContentPath() *string
	SetContentPath(val *string)
	ContentPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HeaderKey() *string
	SetHeaderKey(val *string)
	HeaderKeyInput() *string
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
	ResetContentPath()
	ResetHeaderKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ContentPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ContentPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) HeaderKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) HeaderKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference_Override(g GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetContentPath(val *string) {
	if err := j.validateSetContentPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentPath",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetHeaderKey(val *string) {
	if err := j.validateSetHeaderKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerKey",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ResetContentPath() {
	_jsii_.InvokeVoid(
		g,
		"resetContentPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ResetHeaderKey() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

