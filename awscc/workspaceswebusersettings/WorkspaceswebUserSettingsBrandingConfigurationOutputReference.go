// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspaceswebusersettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceswebUserSettingsBrandingConfigurationOutputReference interface {
	cdktn.ComplexObject
	ColorTheme() *string
	SetColorTheme(val *string)
	ColorThemeInput() *string
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
	Favicon() *string
	SetFavicon(val *string)
	FaviconInput() *string
	FaviconMetadata() WorkspaceswebUserSettingsBrandingConfigurationFaviconMetadataOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalizedStrings() WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap
	LocalizedStringsInput() interface{}
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	LogoMetadata() WorkspaceswebUserSettingsBrandingConfigurationLogoMetadataOutputReference
	TermsOfService() *string
	SetTermsOfService(val *string)
	TermsOfServiceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Wallpaper() *string
	SetWallpaper(val *string)
	WallpaperInput() *string
	WallpaperMetadata() WorkspaceswebUserSettingsBrandingConfigurationWallpaperMetadataOutputReference
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
	PutLocalizedStrings(value interface{})
	ResetColorTheme()
	ResetFavicon()
	ResetLocalizedStrings()
	ResetLogo()
	ResetTermsOfService()
	ResetWallpaper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceswebUserSettingsBrandingConfigurationOutputReference
type jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ColorTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ColorThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorThemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) Favicon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"favicon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) FaviconInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faviconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) FaviconMetadata() WorkspaceswebUserSettingsBrandingConfigurationFaviconMetadataOutputReference {
	var returns WorkspaceswebUserSettingsBrandingConfigurationFaviconMetadataOutputReference
	_jsii_.Get(
		j,
		"faviconMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) LocalizedStrings() WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap {
	var returns WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap
	_jsii_.Get(
		j,
		"localizedStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) LocalizedStringsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localizedStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) LogoMetadata() WorkspaceswebUserSettingsBrandingConfigurationLogoMetadataOutputReference {
	var returns WorkspaceswebUserSettingsBrandingConfigurationLogoMetadataOutputReference
	_jsii_.Get(
		j,
		"logoMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) TermsOfService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) TermsOfServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) Wallpaper() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wallpaper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) WallpaperInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wallpaperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) WallpaperMetadata() WorkspaceswebUserSettingsBrandingConfigurationWallpaperMetadataOutputReference {
	var returns WorkspaceswebUserSettingsBrandingConfigurationWallpaperMetadataOutputReference
	_jsii_.Get(
		j,
		"wallpaperMetadata",
		&returns,
	)
	return returns
}


func NewWorkspaceswebUserSettingsBrandingConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceswebUserSettingsBrandingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceswebUserSettingsBrandingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettingsBrandingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceswebUserSettingsBrandingConfigurationOutputReference_Override(w WorkspaceswebUserSettingsBrandingConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettingsBrandingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetColorTheme(val *string) {
	if err := j.validateSetColorThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorTheme",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetFavicon(val *string) {
	if err := j.validateSetFaviconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"favicon",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetTermsOfService(val *string) {
	if err := j.validateSetTermsOfServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"termsOfService",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetWallpaper(val *string) {
	if err := j.validateSetWallpaperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wallpaper",
		val,
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) PutLocalizedStrings(value interface{}) {
	if err := w.validatePutLocalizedStringsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLocalizedStrings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ResetColorTheme() {
	_jsii_.InvokeVoid(
		w,
		"resetColorTheme",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ResetFavicon() {
	_jsii_.InvokeVoid(
		w,
		"resetFavicon",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ResetLocalizedStrings() {
	_jsii_.InvokeVoid(
		w,
		"resetLocalizedStrings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ResetLogo() {
	_jsii_.InvokeVoid(
		w,
		"resetLogo",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ResetTermsOfService() {
	_jsii_.InvokeVoid(
		w,
		"resetTermsOfService",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ResetWallpaper() {
	_jsii_.InvokeVoid(
		w,
		"resetWallpaper",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

