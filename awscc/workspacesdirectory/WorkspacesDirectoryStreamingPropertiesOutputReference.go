// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspacesdirectory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesDirectoryStreamingPropertiesOutputReference interface {
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
	GlobalAccelerator() WorkspacesDirectoryStreamingPropertiesGlobalAcceleratorOutputReference
	GlobalAcceleratorInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageConnectors() WorkspacesDirectoryStreamingPropertiesStorageConnectorsList
	StorageConnectorsInput() interface{}
	StreamingExperiencePreferredProtocol() *string
	SetStreamingExperiencePreferredProtocol(val *string)
	StreamingExperiencePreferredProtocolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserSettings() WorkspacesDirectoryStreamingPropertiesUserSettingsList
	UserSettingsInput() interface{}
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
	PutGlobalAccelerator(value *WorkspacesDirectoryStreamingPropertiesGlobalAccelerator)
	PutStorageConnectors(value interface{})
	PutUserSettings(value interface{})
	ResetGlobalAccelerator()
	ResetStorageConnectors()
	ResetStreamingExperiencePreferredProtocol()
	ResetUserSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesDirectoryStreamingPropertiesOutputReference
type jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GlobalAccelerator() WorkspacesDirectoryStreamingPropertiesGlobalAcceleratorOutputReference {
	var returns WorkspacesDirectoryStreamingPropertiesGlobalAcceleratorOutputReference
	_jsii_.Get(
		j,
		"globalAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GlobalAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) StorageConnectors() WorkspacesDirectoryStreamingPropertiesStorageConnectorsList {
	var returns WorkspacesDirectoryStreamingPropertiesStorageConnectorsList
	_jsii_.Get(
		j,
		"storageConnectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) StorageConnectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConnectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) StreamingExperiencePreferredProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingExperiencePreferredProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) StreamingExperiencePreferredProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamingExperiencePreferredProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) UserSettings() WorkspacesDirectoryStreamingPropertiesUserSettingsList {
	var returns WorkspacesDirectoryStreamingPropertiesUserSettingsList
	_jsii_.Get(
		j,
		"userSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) UserSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSettingsInput",
		&returns,
	)
	return returns
}


func NewWorkspacesDirectoryStreamingPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspacesDirectoryStreamingPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesDirectoryStreamingPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesDirectory.WorkspacesDirectoryStreamingPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesDirectoryStreamingPropertiesOutputReference_Override(w WorkspacesDirectoryStreamingPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesDirectory.WorkspacesDirectoryStreamingPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference)SetStreamingExperiencePreferredProtocol(val *string) {
	if err := j.validateSetStreamingExperiencePreferredProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingExperiencePreferredProtocol",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) PutGlobalAccelerator(value *WorkspacesDirectoryStreamingPropertiesGlobalAccelerator) {
	if err := w.validatePutGlobalAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGlobalAccelerator",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) PutStorageConnectors(value interface{}) {
	if err := w.validatePutStorageConnectorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStorageConnectors",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) PutUserSettings(value interface{}) {
	if err := w.validatePutUserSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUserSettings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ResetGlobalAccelerator() {
	_jsii_.InvokeVoid(
		w,
		"resetGlobalAccelerator",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ResetStorageConnectors() {
	_jsii_.InvokeVoid(
		w,
		"resetStorageConnectors",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ResetStreamingExperiencePreferredProtocol() {
	_jsii_.InvokeVoid(
		w,
		"resetStreamingExperiencePreferredProtocol",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ResetUserSettings() {
	_jsii_.InvokeVoid(
		w,
		"resetUserSettings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspacesDirectoryStreamingPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

