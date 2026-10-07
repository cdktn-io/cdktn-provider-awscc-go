// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/applicationsignalsinstrumentationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxCollectionDepth() *float64
	SetMaxCollectionDepth(val *float64)
	MaxCollectionDepthInput() *float64
	MaxCollectionWidth() *float64
	SetMaxCollectionWidth(val *float64)
	MaxCollectionWidthInput() *float64
	MaxFieldsPerObject() *float64
	SetMaxFieldsPerObject(val *float64)
	MaxFieldsPerObjectInput() *float64
	MaxHits() *float64
	SetMaxHits(val *float64)
	MaxHitsInput() *float64
	MaxObjectDepth() *float64
	SetMaxObjectDepth(val *float64)
	MaxObjectDepthInput() *float64
	MaxStackFrames() *float64
	SetMaxStackFrames(val *float64)
	MaxStackFramesInput() *float64
	MaxStackTraceSize() *float64
	SetMaxStackTraceSize(val *float64)
	MaxStackTraceSizeInput() *float64
	MaxStringLength() *float64
	SetMaxStringLength(val *float64)
	MaxStringLengthInput() *float64
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
	ResetMaxCollectionDepth()
	ResetMaxCollectionWidth()
	ResetMaxFieldsPerObject()
	ResetMaxHits()
	ResetMaxObjectDepth()
	ResetMaxStackFrames()
	ResetMaxStackTraceSize()
	ResetMaxStringLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference
type jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxCollectionDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCollectionDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxCollectionDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCollectionDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxCollectionWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCollectionWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxCollectionWidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCollectionWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxFieldsPerObject() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFieldsPerObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxFieldsPerObjectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFieldsPerObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxHits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxHitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxObjectDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxObjectDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxObjectDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxObjectDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxStackFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxStackFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxStackTraceSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackTraceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxStackTraceSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackTraceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxStringLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStringLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) MaxStringLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStringLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsInstrumentationConfig.ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference_Override(a ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsInstrumentationConfig.ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxCollectionDepth(val *float64) {
	if err := j.validateSetMaxCollectionDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCollectionDepth",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxCollectionWidth(val *float64) {
	if err := j.validateSetMaxCollectionWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCollectionWidth",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxFieldsPerObject(val *float64) {
	if err := j.validateSetMaxFieldsPerObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFieldsPerObject",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxHits(val *float64) {
	if err := j.validateSetMaxHitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHits",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxObjectDepth(val *float64) {
	if err := j.validateSetMaxObjectDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxObjectDepth",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxStackFrames(val *float64) {
	if err := j.validateSetMaxStackFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStackFrames",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxStackTraceSize(val *float64) {
	if err := j.validateSetMaxStackTraceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStackTraceSize",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetMaxStringLength(val *float64) {
	if err := j.validateSetMaxStringLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStringLength",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxCollectionDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCollectionDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxCollectionWidth() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCollectionWidth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxFieldsPerObject() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxFieldsPerObject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxHits() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxHits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxObjectDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxObjectDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxStackFrames() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxStackFrames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxStackTraceSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxStackTraceSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ResetMaxStringLength() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxStringLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

