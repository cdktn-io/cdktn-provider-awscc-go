// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspaceswebusersettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference interface {
	cdktn.ComplexObject
	BrowserTabTitle() *string
	SetBrowserTabTitle(val *string)
	BrowserTabTitleInput() *string
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
	ContactButtonText() *string
	SetContactButtonText(val *string)
	ContactButtonTextInput() *string
	ContactLink() *string
	SetContactLink(val *string)
	ContactLinkInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadingText() *string
	SetLoadingText(val *string)
	LoadingTextInput() *string
	LoginButtonText() *string
	SetLoginButtonText(val *string)
	LoginButtonTextInput() *string
	LoginDescription() *string
	SetLoginDescription(val *string)
	LoginDescriptionInput() *string
	LoginTitle() *string
	SetLoginTitle(val *string)
	LoginTitleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WelcomeText() *string
	SetWelcomeText(val *string)
	WelcomeTextInput() *string
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
	ResetBrowserTabTitle()
	ResetContactButtonText()
	ResetContactLink()
	ResetLoadingText()
	ResetLoginButtonText()
	ResetLoginDescription()
	ResetLoginTitle()
	ResetWelcomeText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference
type jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) BrowserTabTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserTabTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) BrowserTabTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserTabTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ContactButtonText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactButtonText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ContactButtonTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactButtonTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ContactLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ContactLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoadingText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadingText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoadingTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadingTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoginButtonText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginButtonText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoginButtonTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginButtonTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoginDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoginDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoginTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) LoginTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) WelcomeText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"welcomeText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) WelcomeTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"welcomeTextInput",
		&returns,
	)
	return returns
}


func NewWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectKey); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		&j,
	)

	return &j
}

func NewWorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference_Override(w WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectKey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebUserSettings.WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectKey},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetBrowserTabTitle(val *string) {
	if err := j.validateSetBrowserTabTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserTabTitle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetContactButtonText(val *string) {
	if err := j.validateSetContactButtonTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactButtonText",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetContactLink(val *string) {
	if err := j.validateSetContactLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactLink",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetLoadingText(val *string) {
	if err := j.validateSetLoadingTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadingText",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetLoginButtonText(val *string) {
	if err := j.validateSetLoginButtonTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginButtonText",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetLoginDescription(val *string) {
	if err := j.validateSetLoginDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginDescription",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetLoginTitle(val *string) {
	if err := j.validateSetLoginTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginTitle",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference)SetWelcomeText(val *string) {
	if err := j.validateSetWelcomeTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"welcomeText",
		val,
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetBrowserTabTitle() {
	_jsii_.InvokeVoid(
		w,
		"resetBrowserTabTitle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetContactButtonText() {
	_jsii_.InvokeVoid(
		w,
		"resetContactButtonText",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetContactLink() {
	_jsii_.InvokeVoid(
		w,
		"resetContactLink",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetLoadingText() {
	_jsii_.InvokeVoid(
		w,
		"resetLoadingText",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetLoginButtonText() {
	_jsii_.InvokeVoid(
		w,
		"resetLoginButtonText",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetLoginDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetLoginDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetLoginTitle() {
	_jsii_.InvokeVoid(
		w,
		"resetLoginTitle",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ResetWelcomeText() {
	_jsii_.InvokeVoid(
		w,
		"resetWelcomeText",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspaceswebUserSettingsBrandingConfigurationLocalizedStringsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

