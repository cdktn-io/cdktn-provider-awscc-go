// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccworkspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccworkspaceswebusersettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference interface {
	cdktn.ComplexObject
	ColorTheme() *string
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
	FaviconMetadata() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationFaviconMetadataOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccWorkspaceswebUserSettingsBrandingConfiguration
	SetInternalValue(val *DataAwsccWorkspaceswebUserSettingsBrandingConfiguration)
	LocalizedStrings() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap
	Logo() *string
	LogoMetadata() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLogoMetadataOutputReference
	TermsOfService() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Wallpaper() *string
	WallpaperMetadata() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationWallpaperMetadataOutputReference
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

// The jsii proxy struct for DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference
type jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) ColorTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) Favicon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"favicon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) FaviconMetadata() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationFaviconMetadataOutputReference {
	var returns DataAwsccWorkspaceswebUserSettingsBrandingConfigurationFaviconMetadataOutputReference
	_jsii_.Get(
		j,
		"faviconMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) InternalValue() *DataAwsccWorkspaceswebUserSettingsBrandingConfiguration {
	var returns *DataAwsccWorkspaceswebUserSettingsBrandingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) LocalizedStrings() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap {
	var returns DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsMap
	_jsii_.Get(
		j,
		"localizedStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) LogoMetadata() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLogoMetadataOutputReference {
	var returns DataAwsccWorkspaceswebUserSettingsBrandingConfigurationLogoMetadataOutputReference
	_jsii_.Get(
		j,
		"logoMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) TermsOfService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"termsOfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) Wallpaper() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wallpaper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) WallpaperMetadata() DataAwsccWorkspaceswebUserSettingsBrandingConfigurationWallpaperMetadataOutputReference {
	var returns DataAwsccWorkspaceswebUserSettingsBrandingConfigurationWallpaperMetadataOutputReference
	_jsii_.Get(
		j,
		"wallpaperMetadata",
		&returns,
	)
	return returns
}


func NewDataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWorkspaceswebUserSettings.DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference_Override(d DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWorkspaceswebUserSettings.DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetInternalValue(val *DataAwsccWorkspaceswebUserSettingsBrandingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspaceswebUserSettingsBrandingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

