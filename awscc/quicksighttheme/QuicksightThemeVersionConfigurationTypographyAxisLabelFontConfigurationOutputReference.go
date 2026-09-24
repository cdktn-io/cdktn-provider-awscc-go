// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttheme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksighttheme/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference interface {
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
	FontDecoration() *string
	FontFamily() *string
	FontSize() QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationFontSizeOutputReference
	FontStyle() *string
	FontWeight() QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationFontWeightOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfiguration
	SetInternalValue(val *QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfiguration)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference
type jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) FontColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) FontDecoration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontDecoration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) FontSize() QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationFontSizeOutputReference {
	var returns QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationFontSizeOutputReference
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) FontStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) FontWeight() QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationFontWeightOutputReference {
	var returns QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationFontWeightOutputReference
	_jsii_.Get(
		j,
		"fontWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) InternalValue() *QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfiguration {
	var returns *QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference_Override(q QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightTheme.QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference)SetInternalValue(val *QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightThemeVersionConfigurationTypographyAxisLabelFontConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

