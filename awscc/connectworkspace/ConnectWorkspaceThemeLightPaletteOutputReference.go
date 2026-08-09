// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectWorkspaceThemeLightPaletteOutputReference interface {
	cdktn.ComplexObject
	Canvas() ConnectWorkspaceThemeLightPaletteCanvasOutputReference
	CanvasInput() interface{}
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
	Header() ConnectWorkspaceThemeLightPaletteHeaderOutputReference
	HeaderInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Navigation() ConnectWorkspaceThemeLightPaletteNavigationOutputReference
	NavigationInput() interface{}
	Primary() ConnectWorkspaceThemeLightPalettePrimaryOutputReference
	PrimaryInput() interface{}
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
	PutCanvas(value *ConnectWorkspaceThemeLightPaletteCanvas)
	PutHeader(value *ConnectWorkspaceThemeLightPaletteHeader)
	PutNavigation(value *ConnectWorkspaceThemeLightPaletteNavigation)
	PutPrimary(value *ConnectWorkspaceThemeLightPalettePrimary)
	ResetCanvas()
	ResetHeader()
	ResetNavigation()
	ResetPrimary()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectWorkspaceThemeLightPaletteOutputReference
type jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) Canvas() ConnectWorkspaceThemeLightPaletteCanvasOutputReference {
	var returns ConnectWorkspaceThemeLightPaletteCanvasOutputReference
	_jsii_.Get(
		j,
		"canvas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) CanvasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canvasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) Header() ConnectWorkspaceThemeLightPaletteHeaderOutputReference {
	var returns ConnectWorkspaceThemeLightPaletteHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) Navigation() ConnectWorkspaceThemeLightPaletteNavigationOutputReference {
	var returns ConnectWorkspaceThemeLightPaletteNavigationOutputReference
	_jsii_.Get(
		j,
		"navigation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) NavigationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"navigationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) Primary() ConnectWorkspaceThemeLightPalettePrimaryOutputReference {
	var returns ConnectWorkspaceThemeLightPalettePrimaryOutputReference
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectWorkspaceThemeLightPaletteOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectWorkspaceThemeLightPaletteOutputReference {
	_init_.Initialize()

	if err := validateNewConnectWorkspaceThemeLightPaletteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectWorkspace.ConnectWorkspaceThemeLightPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectWorkspaceThemeLightPaletteOutputReference_Override(c ConnectWorkspaceThemeLightPaletteOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectWorkspace.ConnectWorkspaceThemeLightPaletteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) PutCanvas(value *ConnectWorkspaceThemeLightPaletteCanvas) {
	if err := c.validatePutCanvasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCanvas",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) PutHeader(value *ConnectWorkspaceThemeLightPaletteHeader) {
	if err := c.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) PutNavigation(value *ConnectWorkspaceThemeLightPaletteNavigation) {
	if err := c.validatePutNavigationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNavigation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) PutPrimary(value *ConnectWorkspaceThemeLightPalettePrimary) {
	if err := c.validatePutPrimaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPrimary",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ResetCanvas() {
	_jsii_.InvokeVoid(
		c,
		"resetCanvas",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ResetNavigation() {
	_jsii_.InvokeVoid(
		c,
		"resetNavigation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		c,
		"resetPrimary",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

