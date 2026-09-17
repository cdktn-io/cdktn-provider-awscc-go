// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeRestConfigurationOutputReference interface {
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
	EntityConfigurations() GlueConnectionTypeRestConfigurationEntityConfigurationsMap
	EntityConfigurationsInput() interface{}
	// Experimental.
	Fqn() *string
	GlobalSourceConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference
	GlobalSourceConfigurationInput() interface{}
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
	ValidationEndpointConfiguration() GlueConnectionTypeRestConfigurationValidationEndpointConfigurationOutputReference
	ValidationEndpointConfigurationInput() interface{}
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
	PutEntityConfigurations(value interface{})
	PutGlobalSourceConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfiguration)
	PutValidationEndpointConfiguration(value *GlueConnectionTypeRestConfigurationValidationEndpointConfiguration)
	ResetEntityConfigurations()
	ResetGlobalSourceConfiguration()
	ResetValidationEndpointConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeRestConfigurationOutputReference
type jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) EntityConfigurations() GlueConnectionTypeRestConfigurationEntityConfigurationsMap {
	var returns GlueConnectionTypeRestConfigurationEntityConfigurationsMap
	_jsii_.Get(
		j,
		"entityConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) EntityConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GlobalSourceConfiguration() GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationGlobalSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"globalSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GlobalSourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ValidationEndpointConfiguration() GlueConnectionTypeRestConfigurationValidationEndpointConfigurationOutputReference {
	var returns GlueConnectionTypeRestConfigurationValidationEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"validationEndpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ValidationEndpointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationEndpointConfigurationInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeRestConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeRestConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeRestConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeRestConfigurationOutputReference_Override(g GlueConnectionTypeRestConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeRestConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) PutEntityConfigurations(value interface{}) {
	if err := g.validatePutEntityConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEntityConfigurations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) PutGlobalSourceConfiguration(value *GlueConnectionTypeRestConfigurationGlobalSourceConfiguration) {
	if err := g.validatePutGlobalSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGlobalSourceConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) PutValidationEndpointConfiguration(value *GlueConnectionTypeRestConfigurationValidationEndpointConfiguration) {
	if err := g.validatePutValidationEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putValidationEndpointConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ResetEntityConfigurations() {
	_jsii_.InvokeVoid(
		g,
		"resetEntityConfigurations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ResetGlobalSourceConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetGlobalSourceConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ResetValidationEndpointConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetValidationEndpointConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

