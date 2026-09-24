// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttheme/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference interface {
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
	FontColor() *string
	SetFontColor(val *string)
	FontColorInput() *string
	FontDecoration() *string
	SetFontDecoration(val *string)
	FontDecorationInput() *string
	FontFamily() *string
	SetFontFamily(val *string)
	FontFamilyInput() *string
	FontSize() QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontSizeOutputReference
	FontSizeInput() interface{}
	FontStyle() *string
	SetFontStyle(val *string)
	FontStyleInput() *string
	FontWeight() QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontWeightOutputReference
	FontWeightInput() interface{}
	// Experimental.
	Fqn() *string
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
	PutFontSize(value *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontSize)
	PutFontWeight(value *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontWeight)
	ResetFontColor()
	ResetFontDecoration()
	ResetFontFamily()
	ResetFontSize()
	ResetFontStyle()
	ResetFontWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference
type jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontDecoration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontDecoration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontDecorationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontDecorationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontSize() QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontSizeOutputReference {
	var returns QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontSizeOutputReference
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontWeight() QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontWeightOutputReference {
	var returns QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontWeightOutputReference
	_jsii_.Get(
		j,
		"fontWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) FontWeightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference_Override(q QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetFontColor(val *string) {
	if err := j.validateSetFontColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontColor",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetFontDecoration(val *string) {
	if err := j.validateSetFontDecorationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontDecoration",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetFontStyle(val *string) {
	if err := j.validateSetFontStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontStyle",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) PutFontSize(value *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontSize) {
	if err := q.validatePutFontSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFontSize",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) PutFontWeight(value *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationFontWeight) {
	if err := q.validatePutFontWeightParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFontWeight",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		q,
		"resetFontColor",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ResetFontDecoration() {
	_jsii_.InvokeVoid(
		q,
		"resetFontDecoration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		q,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		q,
		"resetFontSize",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ResetFontStyle() {
	_jsii_.InvokeVoid(
		q,
		"resetFontStyle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ResetFontWeight() {
	_jsii_.InvokeVoid(
		q,
		"resetFontWeight",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationFontConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

