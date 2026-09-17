// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnectiontype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference interface {
	cdktn.ComplexObject
	AllowedValues() *[]*string
	SetAllowedValues(val *[]*string)
	AllowedValuesInput() *[]*string
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
	DefaultValue() *string
	SetDefaultValue(val *string)
	DefaultValueInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyOverride() *string
	SetKeyOverride(val *string)
	KeyOverrideInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PropertyLocation() *string
	SetPropertyLocation(val *string)
	PropertyLocationInput() *string
	PropertyType() *string
	SetPropertyType(val *string)
	PropertyTypeInput() *string
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
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
	ResetAllowedValues()
	ResetDefaultValue()
	ResetKeyOverride()
	ResetName()
	ResetPropertyLocation()
	ResetPropertyType()
	ResetRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference
type jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) AllowedValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) AllowedValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) KeyOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) KeyOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) PropertyLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) PropertyLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) PropertyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) PropertyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference_Override(g GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnectionType.GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetAllowedValues(val *[]*string) {
	if err := j.validateSetAllowedValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedValues",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetKeyOverride(val *string) {
	if err := j.validateSetKeyOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyOverride",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetPropertyLocation(val *string) {
	if err := j.validateSetPropertyLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyLocation",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetPropertyType(val *string) {
	if err := j.validateSetPropertyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetAllowedValues() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetKeyOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetPropertyLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetPropertyLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetPropertyType() {
	_jsii_.InvokeVoid(
		g,
		"resetPropertyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionTypeConnectorAuthenticationConfigurationOAuth2PropertiesJwtBearerPropertiesTokenUrlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

