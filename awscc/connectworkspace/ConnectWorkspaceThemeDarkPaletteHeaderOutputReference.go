// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectWorkspaceThemeDarkPaletteHeaderOutputReference interface {
	cdktn.ComplexObject
	Background() *string
	SetBackground(val *string)
	BackgroundInput() *string
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
	InvertActionsColors() interface{}
	SetInvertActionsColors(val interface{})
	InvertActionsColorsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Text() *string
	SetText(val *string)
	TextHover() *string
	SetTextHover(val *string)
	TextHoverInput() *string
	TextInput() *string
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
	ResetBackground()
	ResetInvertActionsColors()
	ResetText()
	ResetTextHover()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectWorkspaceThemeDarkPaletteHeaderOutputReference
type jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) Background() *string {
	var returns *string
	_jsii_.Get(
		j,
		"background",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) BackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) InvertActionsColors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertActionsColors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) InvertActionsColorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertActionsColorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) TextHover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textHover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) TextHoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textHoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewConnectWorkspaceThemeDarkPaletteHeaderOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectWorkspaceThemeDarkPaletteHeaderOutputReference {
	_init_.Initialize()

	if err := validateNewConnectWorkspaceThemeDarkPaletteHeaderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectWorkspace.ConnectWorkspaceThemeDarkPaletteHeaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectWorkspaceThemeDarkPaletteHeaderOutputReference_Override(c ConnectWorkspaceThemeDarkPaletteHeaderOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectWorkspace.ConnectWorkspaceThemeDarkPaletteHeaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetBackground(val *string) {
	if err := j.validateSetBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"background",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetInvertActionsColors(val interface{}) {
	if err := j.validateSetInvertActionsColorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invertActionsColors",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference)SetTextHover(val *string) {
	if err := j.validateSetTextHoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textHover",
		val,
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ResetBackground() {
	_jsii_.InvokeVoid(
		c,
		"resetBackground",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ResetInvertActionsColors() {
	_jsii_.InvokeVoid(
		c,
		"resetInvertActionsColors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		c,
		"resetText",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ResetTextHover() {
	_jsii_.InvokeVoid(
		c,
		"resetTextHover",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeDarkPaletteHeaderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

