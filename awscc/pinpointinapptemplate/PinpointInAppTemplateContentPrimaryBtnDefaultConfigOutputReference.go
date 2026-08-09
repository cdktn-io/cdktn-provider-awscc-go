// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointinapptemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pinpointinapptemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference interface {
	cdktn.ComplexObject
	BackgroundColor() *string
	SetBackgroundColor(val *string)
	BackgroundColorInput() *string
	BorderRadius() *float64
	SetBorderRadius(val *float64)
	BorderRadiusInput() *float64
	ButtonAction() *string
	SetButtonAction(val *string)
	ButtonActionInput() *string
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
	Link() *string
	SetLink(val *string)
	LinkInput() *string
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
	TextColor() *string
	SetTextColor(val *string)
	TextColorInput() *string
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
	ResetBackgroundColor()
	ResetBorderRadius()
	ResetButtonAction()
	ResetLink()
	ResetText()
	ResetTextColor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference
type jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) BorderRadius() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"borderRadius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) BorderRadiusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"borderRadiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ButtonAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ButtonActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buttonActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) Link() *string {
	var returns *string
	_jsii_.Get(
		j,
		"link",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) LinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) TextColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) TextColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewPinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pinpointInAppTemplate.PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference_Override(p PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pinpointInAppTemplate.PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetBorderRadius(val *float64) {
	if err := j.validateSetBorderRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"borderRadius",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetButtonAction(val *string) {
	if err := j.validateSetButtonActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buttonAction",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetLink(val *string) {
	if err := j.validateSetLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"link",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference)SetTextColor(val *string) {
	if err := j.validateSetTextColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textColor",
		val,
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		p,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ResetBorderRadius() {
	_jsii_.InvokeVoid(
		p,
		"resetBorderRadius",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ResetButtonAction() {
	_jsii_.InvokeVoid(
		p,
		"resetButtonAction",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ResetLink() {
	_jsii_.InvokeVoid(
		p,
		"resetLink",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		p,
		"resetText",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ResetTextColor() {
	_jsii_.InvokeVoid(
		p,
		"resetTextColor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PinpointInAppTemplateContentPrimaryBtnDefaultConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

