// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspacesdirectory/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesDirectorySelfservicePermissionsOutputReference interface {
	cdktn.ComplexObject
	ChangeComputeType() *string
	SetChangeComputeType(val *string)
	ChangeComputeTypeInput() *string
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
	IncreaseVolumeSize() *string
	SetIncreaseVolumeSize(val *string)
	IncreaseVolumeSizeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RebuildWorkspace() *string
	SetRebuildWorkspace(val *string)
	RebuildWorkspaceInput() *string
	RestartWorkspace() *string
	SetRestartWorkspace(val *string)
	RestartWorkspaceInput() *string
	SwitchRunningMode() *string
	SetSwitchRunningMode(val *string)
	SwitchRunningModeInput() *string
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
	ResetChangeComputeType()
	ResetIncreaseVolumeSize()
	ResetRebuildWorkspace()
	ResetRestartWorkspace()
	ResetSwitchRunningMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesDirectorySelfservicePermissionsOutputReference
type jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ChangeComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ChangeComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeComputeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) IncreaseVolumeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"increaseVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) IncreaseVolumeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"increaseVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) RebuildWorkspace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebuildWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) RebuildWorkspaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebuildWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) RestartWorkspace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) RestartWorkspaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) SwitchRunningMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"switchRunningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) SwitchRunningModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"switchRunningModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWorkspacesDirectorySelfservicePermissionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspacesDirectorySelfservicePermissionsOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesDirectorySelfservicePermissionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesDirectory.WorkspacesDirectorySelfservicePermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesDirectorySelfservicePermissionsOutputReference_Override(w WorkspacesDirectorySelfservicePermissionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesDirectory.WorkspacesDirectorySelfservicePermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetChangeComputeType(val *string) {
	if err := j.validateSetChangeComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeComputeType",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetIncreaseVolumeSize(val *string) {
	if err := j.validateSetIncreaseVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"increaseVolumeSize",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetRebuildWorkspace(val *string) {
	if err := j.validateSetRebuildWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebuildWorkspace",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetRestartWorkspace(val *string) {
	if err := j.validateSetRestartWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartWorkspace",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetSwitchRunningMode(val *string) {
	if err := j.validateSetSwitchRunningModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchRunningMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ResetChangeComputeType() {
	_jsii_.InvokeVoid(
		w,
		"resetChangeComputeType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ResetIncreaseVolumeSize() {
	_jsii_.InvokeVoid(
		w,
		"resetIncreaseVolumeSize",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ResetRebuildWorkspace() {
	_jsii_.InvokeVoid(
		w,
		"resetRebuildWorkspace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ResetRestartWorkspace() {
	_jsii_.InvokeVoid(
		w,
		"resetRestartWorkspace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ResetSwitchRunningMode() {
	_jsii_.InvokeVoid(
		w,
		"resetSwitchRunningMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspacesDirectorySelfservicePermissionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

