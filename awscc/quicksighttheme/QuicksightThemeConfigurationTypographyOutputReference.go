// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttheme/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightThemeConfigurationTypographyOutputReference interface {
	cdktn.ComplexObject
	AxisLabelFontConfiguration() QuicksightThemeConfigurationTypographyAxisLabelFontConfigurationOutputReference
	AxisLabelFontConfigurationInput() interface{}
	AxisTitleFontConfiguration() QuicksightThemeConfigurationTypographyAxisTitleFontConfigurationOutputReference
	AxisTitleFontConfigurationInput() interface{}
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
	DataLabelFontConfiguration() QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference
	DataLabelFontConfigurationInput() interface{}
	FontFamilies() QuicksightThemeConfigurationTypographyFontFamiliesList
	FontFamiliesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LegendTitleFontConfiguration() QuicksightThemeConfigurationTypographyLegendTitleFontConfigurationOutputReference
	LegendTitleFontConfigurationInput() interface{}
	LegendValueFontConfiguration() QuicksightThemeConfigurationTypographyLegendValueFontConfigurationOutputReference
	LegendValueFontConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VisualSubtitleFontConfiguration() QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationOutputReference
	VisualSubtitleFontConfigurationInput() interface{}
	VisualTitleFontConfiguration() QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationOutputReference
	VisualTitleFontConfigurationInput() interface{}
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
	PutAxisLabelFontConfiguration(value *QuicksightThemeConfigurationTypographyAxisLabelFontConfiguration)
	PutAxisTitleFontConfiguration(value *QuicksightThemeConfigurationTypographyAxisTitleFontConfiguration)
	PutDataLabelFontConfiguration(value *QuicksightThemeConfigurationTypographyDataLabelFontConfiguration)
	PutFontFamilies(value interface{})
	PutLegendTitleFontConfiguration(value *QuicksightThemeConfigurationTypographyLegendTitleFontConfiguration)
	PutLegendValueFontConfiguration(value *QuicksightThemeConfigurationTypographyLegendValueFontConfiguration)
	PutVisualSubtitleFontConfiguration(value *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfiguration)
	PutVisualTitleFontConfiguration(value *QuicksightThemeConfigurationTypographyVisualTitleFontConfiguration)
	ResetAxisLabelFontConfiguration()
	ResetAxisTitleFontConfiguration()
	ResetDataLabelFontConfiguration()
	ResetFontFamilies()
	ResetLegendTitleFontConfiguration()
	ResetLegendValueFontConfiguration()
	ResetVisualSubtitleFontConfiguration()
	ResetVisualTitleFontConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightThemeConfigurationTypographyOutputReference
type jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) AxisLabelFontConfiguration() QuicksightThemeConfigurationTypographyAxisLabelFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyAxisLabelFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"axisLabelFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) AxisLabelFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"axisLabelFontConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) AxisTitleFontConfiguration() QuicksightThemeConfigurationTypographyAxisTitleFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyAxisTitleFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"axisTitleFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) AxisTitleFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"axisTitleFontConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) DataLabelFontConfiguration() QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyDataLabelFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataLabelFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) DataLabelFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLabelFontConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) FontFamilies() QuicksightThemeConfigurationTypographyFontFamiliesList {
	var returns QuicksightThemeConfigurationTypographyFontFamiliesList
	_jsii_.Get(
		j,
		"fontFamilies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) FontFamiliesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fontFamiliesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) LegendTitleFontConfiguration() QuicksightThemeConfigurationTypographyLegendTitleFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyLegendTitleFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"legendTitleFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) LegendTitleFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendTitleFontConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) LegendValueFontConfiguration() QuicksightThemeConfigurationTypographyLegendValueFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyLegendValueFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"legendValueFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) LegendValueFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendValueFontConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) VisualSubtitleFontConfiguration() QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyVisualSubtitleFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"visualSubtitleFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) VisualSubtitleFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualSubtitleFontConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) VisualTitleFontConfiguration() QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationOutputReference {
	var returns QuicksightThemeConfigurationTypographyVisualTitleFontConfigurationOutputReference
	_jsii_.Get(
		j,
		"visualTitleFontConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) VisualTitleFontConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualTitleFontConfigurationInput",
		&returns,
	)
	return returns
}


func NewQuicksightThemeConfigurationTypographyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightThemeConfigurationTypographyOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeConfigurationTypographyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeConfigurationTypographyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeConfigurationTypographyOutputReference_Override(q QuicksightThemeConfigurationTypographyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeConfigurationTypographyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutAxisLabelFontConfiguration(value *QuicksightThemeConfigurationTypographyAxisLabelFontConfiguration) {
	if err := q.validatePutAxisLabelFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAxisLabelFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutAxisTitleFontConfiguration(value *QuicksightThemeConfigurationTypographyAxisTitleFontConfiguration) {
	if err := q.validatePutAxisTitleFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAxisTitleFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutDataLabelFontConfiguration(value *QuicksightThemeConfigurationTypographyDataLabelFontConfiguration) {
	if err := q.validatePutDataLabelFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataLabelFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutFontFamilies(value interface{}) {
	if err := q.validatePutFontFamiliesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFontFamilies",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutLegendTitleFontConfiguration(value *QuicksightThemeConfigurationTypographyLegendTitleFontConfiguration) {
	if err := q.validatePutLegendTitleFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putLegendTitleFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutLegendValueFontConfiguration(value *QuicksightThemeConfigurationTypographyLegendValueFontConfiguration) {
	if err := q.validatePutLegendValueFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putLegendValueFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutVisualSubtitleFontConfiguration(value *QuicksightThemeConfigurationTypographyVisualSubtitleFontConfiguration) {
	if err := q.validatePutVisualSubtitleFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putVisualSubtitleFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) PutVisualTitleFontConfiguration(value *QuicksightThemeConfigurationTypographyVisualTitleFontConfiguration) {
	if err := q.validatePutVisualTitleFontConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putVisualTitleFontConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetAxisLabelFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetAxisLabelFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetAxisTitleFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetAxisTitleFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetDataLabelFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetDataLabelFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetFontFamilies() {
	_jsii_.InvokeVoid(
		q,
		"resetFontFamilies",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetLegendTitleFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetLegendTitleFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetLegendValueFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetLegendValueFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetVisualSubtitleFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetVisualSubtitleFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ResetVisualTitleFontConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetVisualTitleFontConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightThemeConfigurationTypographyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

