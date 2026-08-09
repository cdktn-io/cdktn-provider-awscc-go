// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/connectworkspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectWorkspaceThemeLightPaletteCanvasOutputReference interface {
	cdktn.ComplexObject
	ActiveBackground() *string
	SetActiveBackground(val *string)
	ActiveBackgroundInput() *string
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
	ContainerBackground() *string
	SetContainerBackground(val *string)
	ContainerBackgroundInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PageBackground() *string
	SetPageBackground(val *string)
	PageBackgroundInput() *string
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
	ResetActiveBackground()
	ResetContainerBackground()
	ResetPageBackground()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectWorkspaceThemeLightPaletteCanvasOutputReference
type jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ActiveBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ActiveBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ContainerBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ContainerBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) PageBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) PageBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pageBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectWorkspaceThemeLightPaletteCanvasOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConnectWorkspaceThemeLightPaletteCanvasOutputReference {
	_init_.Initialize()

	if err := validateNewConnectWorkspaceThemeLightPaletteCanvasOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.connectWorkspace.ConnectWorkspaceThemeLightPaletteCanvasOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectWorkspaceThemeLightPaletteCanvasOutputReference_Override(c ConnectWorkspaceThemeLightPaletteCanvasOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.connectWorkspace.ConnectWorkspaceThemeLightPaletteCanvasOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetActiveBackground(val *string) {
	if err := j.validateSetActiveBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeBackground",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetContainerBackground(val *string) {
	if err := j.validateSetContainerBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerBackground",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetPageBackground(val *string) {
	if err := j.validateSetPageBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pageBackground",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ResetActiveBackground() {
	_jsii_.InvokeVoid(
		c,
		"resetActiveBackground",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ResetContainerBackground() {
	_jsii_.InvokeVoid(
		c,
		"resetContainerBackground",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ResetPageBackground() {
	_jsii_.InvokeVoid(
		c,
		"resetPageBackground",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ConnectWorkspaceThemeLightPaletteCanvasOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

