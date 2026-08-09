// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinesswebexperience

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/qbusinesswebexperience/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessWebExperienceCustomizationConfigurationOutputReference interface {
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
	CustomCssUrl() *string
	SetCustomCssUrl(val *string)
	CustomCssUrlInput() *string
	FaviconUrl() *string
	SetFaviconUrl(val *string)
	FaviconUrlInput() *string
	FontUrl() *string
	SetFontUrl(val *string)
	FontUrlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogoUrl() *string
	SetLogoUrl(val *string)
	LogoUrlInput() *string
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
	ResetCustomCssUrl()
	ResetFaviconUrl()
	ResetFontUrl()
	ResetLogoUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QbusinessWebExperienceCustomizationConfigurationOutputReference
type jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) CustomCssUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCssUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) CustomCssUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCssUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) FaviconUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faviconUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) FaviconUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faviconUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) FontUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) FontUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) LogoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQbusinessWebExperienceCustomizationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QbusinessWebExperienceCustomizationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQbusinessWebExperienceCustomizationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessWebExperience.QbusinessWebExperienceCustomizationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQbusinessWebExperienceCustomizationConfigurationOutputReference_Override(q QbusinessWebExperienceCustomizationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.qbusinessWebExperience.QbusinessWebExperienceCustomizationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetCustomCssUrl(val *string) {
	if err := j.validateSetCustomCssUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCssUrl",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetFaviconUrl(val *string) {
	if err := j.validateSetFaviconUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"faviconUrl",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetFontUrl(val *string) {
	if err := j.validateSetFontUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontUrl",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetLogoUrl(val *string) {
	if err := j.validateSetLogoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUrl",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ResetCustomCssUrl() {
	_jsii_.InvokeVoid(
		q,
		"resetCustomCssUrl",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ResetFaviconUrl() {
	_jsii_.InvokeVoid(
		q,
		"resetFaviconUrl",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ResetFontUrl() {
	_jsii_.InvokeVoid(
		q,
		"resetFontUrl",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ResetLogoUrl() {
	_jsii_.InvokeVoid(
		q,
		"resetLogoUrl",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QbusinessWebExperienceCustomizationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

