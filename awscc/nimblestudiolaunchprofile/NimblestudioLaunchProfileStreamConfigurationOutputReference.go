// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiolaunchprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/nimblestudiolaunchprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NimblestudioLaunchProfileStreamConfigurationOutputReference interface {
	cdktn.ComplexObject
	AutomaticTerminationMode() *string
	SetAutomaticTerminationMode(val *string)
	AutomaticTerminationModeInput() *string
	ClipboardMode() *string
	SetClipboardMode(val *string)
	ClipboardModeInput() *string
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
	Ec2InstanceTypes() *[]*string
	SetEc2InstanceTypes(val *[]*string)
	Ec2InstanceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxSessionLengthInMinutes() *float64
	SetMaxSessionLengthInMinutes(val *float64)
	MaxSessionLengthInMinutesInput() *float64
	MaxStoppedSessionLengthInMinutes() *float64
	SetMaxStoppedSessionLengthInMinutes(val *float64)
	MaxStoppedSessionLengthInMinutesInput() *float64
	SessionBackup() NimblestudioLaunchProfileStreamConfigurationSessionBackupOutputReference
	SessionBackupInput() interface{}
	SessionPersistenceMode() *string
	SetSessionPersistenceMode(val *string)
	SessionPersistenceModeInput() *string
	SessionStorage() NimblestudioLaunchProfileStreamConfigurationSessionStorageOutputReference
	SessionStorageInput() interface{}
	StreamingImageIds() *[]*string
	SetStreamingImageIds(val *[]*string)
	StreamingImageIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VolumeConfiguration() NimblestudioLaunchProfileStreamConfigurationVolumeConfigurationOutputReference
	VolumeConfigurationInput() interface{}
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
	PutSessionBackup(value *NimblestudioLaunchProfileStreamConfigurationSessionBackup)
	PutSessionStorage(value *NimblestudioLaunchProfileStreamConfigurationSessionStorage)
	PutVolumeConfiguration(value *NimblestudioLaunchProfileStreamConfigurationVolumeConfiguration)
	ResetAutomaticTerminationMode()
	ResetMaxSessionLengthInMinutes()
	ResetMaxStoppedSessionLengthInMinutes()
	ResetSessionBackup()
	ResetSessionPersistenceMode()
	ResetSessionStorage()
	ResetVolumeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NimblestudioLaunchProfileStreamConfigurationOutputReference
type jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) AutomaticTerminationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticTerminationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) AutomaticTerminationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticTerminationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ClipboardMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clipboardMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ClipboardModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clipboardModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) Ec2InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2InstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) Ec2InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2InstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) MaxSessionLengthInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSessionLengthInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) MaxSessionLengthInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSessionLengthInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) MaxStoppedSessionLengthInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStoppedSessionLengthInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) MaxStoppedSessionLengthInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStoppedSessionLengthInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) SessionBackup() NimblestudioLaunchProfileStreamConfigurationSessionBackupOutputReference {
	var returns NimblestudioLaunchProfileStreamConfigurationSessionBackupOutputReference
	_jsii_.Get(
		j,
		"sessionBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) SessionBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) SessionPersistenceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPersistenceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) SessionPersistenceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPersistenceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) SessionStorage() NimblestudioLaunchProfileStreamConfigurationSessionStorageOutputReference {
	var returns NimblestudioLaunchProfileStreamConfigurationSessionStorageOutputReference
	_jsii_.Get(
		j,
		"sessionStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) SessionStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) StreamingImageIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamingImageIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) StreamingImageIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamingImageIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) VolumeConfiguration() NimblestudioLaunchProfileStreamConfigurationVolumeConfigurationOutputReference {
	var returns NimblestudioLaunchProfileStreamConfigurationVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"volumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) VolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeConfigurationInput",
		&returns,
	)
	return returns
}


func NewNimblestudioLaunchProfileStreamConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NimblestudioLaunchProfileStreamConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNimblestudioLaunchProfileStreamConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.nimblestudioLaunchProfile.NimblestudioLaunchProfileStreamConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNimblestudioLaunchProfileStreamConfigurationOutputReference_Override(n NimblestudioLaunchProfileStreamConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.nimblestudioLaunchProfile.NimblestudioLaunchProfileStreamConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetAutomaticTerminationMode(val *string) {
	if err := j.validateSetAutomaticTerminationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticTerminationMode",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetClipboardMode(val *string) {
	if err := j.validateSetClipboardModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clipboardMode",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetEc2InstanceTypes(val *[]*string) {
	if err := j.validateSetEc2InstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceTypes",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetMaxSessionLengthInMinutes(val *float64) {
	if err := j.validateSetMaxSessionLengthInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSessionLengthInMinutes",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetMaxStoppedSessionLengthInMinutes(val *float64) {
	if err := j.validateSetMaxStoppedSessionLengthInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStoppedSessionLengthInMinutes",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetSessionPersistenceMode(val *string) {
	if err := j.validateSetSessionPersistenceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPersistenceMode",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetStreamingImageIds(val *[]*string) {
	if err := j.validateSetStreamingImageIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingImageIds",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) PutSessionBackup(value *NimblestudioLaunchProfileStreamConfigurationSessionBackup) {
	if err := n.validatePutSessionBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSessionBackup",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) PutSessionStorage(value *NimblestudioLaunchProfileStreamConfigurationSessionStorage) {
	if err := n.validatePutSessionStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putSessionStorage",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) PutVolumeConfiguration(value *NimblestudioLaunchProfileStreamConfigurationVolumeConfiguration) {
	if err := n.validatePutVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putVolumeConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetAutomaticTerminationMode() {
	_jsii_.InvokeVoid(
		n,
		"resetAutomaticTerminationMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetMaxSessionLengthInMinutes() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxSessionLengthInMinutes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetMaxStoppedSessionLengthInMinutes() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxStoppedSessionLengthInMinutes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetSessionBackup() {
	_jsii_.InvokeVoid(
		n,
		"resetSessionBackup",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetSessionPersistenceMode() {
	_jsii_.InvokeVoid(
		n,
		"resetSessionPersistenceMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetSessionStorage() {
	_jsii_.InvokeVoid(
		n,
		"resetSessionStorage",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ResetVolumeConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetVolumeConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NimblestudioLaunchProfileStreamConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

