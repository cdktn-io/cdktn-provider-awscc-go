// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttheme/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference interface {
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
	FontSize() QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontSizeOutputReference
	FontSizeInput() interface{}
	FontStyle() *string
	SetFontStyle(val *string)
	FontStyleInput() *string
	FontWeight() QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontWeightOutputReference
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
	PutFontSize(value *QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontSize)
	PutFontWeight(value *QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontWeight)
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

// The jsii proxy struct for QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference
type jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontDecoration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontDecoration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontDecorationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontDecorationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontSize() QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontSizeOutputReference {
	var returns QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontSizeOutputReference
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontWeight() QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontWeightOutputReference {
	var returns QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontWeightOutputReference
	_jsii_.Get(
		j,
		"fontWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) FontWeightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference_Override(q QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetFontColor(val *string) {
	if err := j.validateSetFontColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontColor",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetFontDecoration(val *string) {
	if err := j.validateSetFontDecorationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontDecoration",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetFontStyle(val *string) {
	if err := j.validateSetFontStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontStyle",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) PutFontSize(value *QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontSize) {
	if err := q.validatePutFontSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFontSize",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) PutFontWeight(value *QuicksightThemeConfigurationTypographyDataLabelFontConfigurationFontWeight) {
	if err := q.validatePutFontWeightParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFontWeight",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ResetFontColor() {
	_jsii_.InvokeVoid(
		q,
		"resetFontColor",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ResetFontDecoration() {
	_jsii_.InvokeVoid(
		q,
		"resetFontDecoration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		q,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		q,
		"resetFontSize",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ResetFontStyle() {
	_jsii_.InvokeVoid(
		q,
		"resetFontStyle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ResetFontWeight() {
	_jsii_.InvokeVoid(
		q,
		"resetFontWeight",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

