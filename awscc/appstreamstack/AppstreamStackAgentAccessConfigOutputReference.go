// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/appstreamstack/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamStackAgentAccessConfigOutputReference interface {
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
	S3BucketArn() *string
	SetS3BucketArn(val *string)
	S3BucketArnInput() *string
	ScreenImageFormat() *string
	SetScreenImageFormat(val *string)
	ScreenImageFormatInput() *string
	ScreenResolution() *string
	SetScreenResolution(val *string)
	ScreenResolutionInput() *string
	ScreenshotsUploadEnabled() interface{}
	SetScreenshotsUploadEnabled(val interface{})
	ScreenshotsUploadEnabledInput() interface{}
	Settings() AppstreamStackAgentAccessConfigSettingsList
	SettingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserControlMode() *string
	SetUserControlMode(val *string)
	UserControlModeInput() *string
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
	PutSettings(value interface{})
	ResetS3BucketArn()
	ResetScreenImageFormat()
	ResetScreenResolution()
	ResetScreenshotsUploadEnabled()
	ResetSettings()
	ResetUserControlMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppstreamStackAgentAccessConfigOutputReference
type jsiiProxy_AppstreamStackAgentAccessConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) S3BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) S3BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ScreenImageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"screenImageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ScreenImageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"screenImageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ScreenResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"screenResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ScreenResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"screenResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ScreenshotsUploadEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"screenshotsUploadEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ScreenshotsUploadEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"screenshotsUploadEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) Settings() AppstreamStackAgentAccessConfigSettingsList {
	var returns AppstreamStackAgentAccessConfigSettingsList
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) SettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) UserControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) UserControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userControlModeInput",
		&returns,
	)
	return returns
}


func NewAppstreamStackAgentAccessConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppstreamStackAgentAccessConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppstreamStackAgentAccessConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppstreamStackAgentAccessConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appstreamStack.AppstreamStackAgentAccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppstreamStackAgentAccessConfigOutputReference_Override(a AppstreamStackAgentAccessConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appstreamStack.AppstreamStackAgentAccessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetS3BucketArn(val *string) {
	if err := j.validateSetS3BucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetScreenImageFormat(val *string) {
	if err := j.validateSetScreenImageFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"screenImageFormat",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetScreenResolution(val *string) {
	if err := j.validateSetScreenResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"screenResolution",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetScreenshotsUploadEnabled(val interface{}) {
	if err := j.validateSetScreenshotsUploadEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"screenshotsUploadEnabled",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference)SetUserControlMode(val *string) {
	if err := j.validateSetUserControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userControlMode",
		val,
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) PutSettings(value interface{}) {
	if err := a.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ResetS3BucketArn() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BucketArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ResetScreenImageFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetScreenImageFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ResetScreenResolution() {
	_jsii_.InvokeVoid(
		a,
		"resetScreenResolution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ResetScreenshotsUploadEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetScreenshotsUploadEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ResetUserControlMode() {
	_jsii_.InvokeVoid(
		a,
		"resetUserControlMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppstreamStackAgentAccessConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

