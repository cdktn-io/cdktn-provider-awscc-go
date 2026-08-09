// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspaceswebbrowsersettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspaceswebbrowsersettings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference interface {
	cdktn.ComplexObject
	AllowedUrls() *[]*string
	SetAllowedUrls(val *[]*string)
	AllowedUrlsInput() *[]*string
	BlockedCategories() *[]*string
	SetBlockedCategories(val *[]*string)
	BlockedCategoriesInput() *[]*string
	BlockedUrls() *[]*string
	SetBlockedUrls(val *[]*string)
	BlockedUrlsInput() *[]*string
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
	ResetAllowedUrls()
	ResetBlockedCategories()
	ResetBlockedUrls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference
type jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) AllowedUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) AllowedUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) BlockedCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) BlockedCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) BlockedUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) BlockedUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebBrowserSettings.WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference_Override(w WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspaceswebBrowserSettings.WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetAllowedUrls(val *[]*string) {
	if err := j.validateSetAllowedUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUrls",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetBlockedCategories(val *[]*string) {
	if err := j.validateSetBlockedCategoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedCategories",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetBlockedUrls(val *[]*string) {
	if err := j.validateSetBlockedUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedUrls",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ResetAllowedUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ResetBlockedCategories() {
	_jsii_.InvokeVoid(
		w,
		"resetBlockedCategories",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ResetBlockedUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetBlockedUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspaceswebBrowserSettingsWebContentFilteringPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

