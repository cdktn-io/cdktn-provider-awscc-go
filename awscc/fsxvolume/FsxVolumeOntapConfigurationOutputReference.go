// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/fsxvolume/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxVolumeOntapConfigurationOutputReference interface {
	cdktn.ComplexObject
	AggregateConfiguration() FsxVolumeOntapConfigurationAggregateConfigurationOutputReference
	AggregateConfigurationInput() interface{}
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
	CopyTagsToBackups() *string
	SetCopyTagsToBackups(val *string)
	CopyTagsToBackupsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JunctionPath() *string
	SetJunctionPath(val *string)
	JunctionPathInput() *string
	OntapVolumeType() *string
	SetOntapVolumeType(val *string)
	OntapVolumeTypeInput() *string
	SecurityStyle() *string
	SetSecurityStyle(val *string)
	SecurityStyleInput() *string
	SizeInBytes() *string
	SetSizeInBytes(val *string)
	SizeInBytesInput() *string
	SizeInMegabytes() *string
	SetSizeInMegabytes(val *string)
	SizeInMegabytesInput() *string
	SnaplockConfiguration() FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference
	SnaplockConfigurationInput() interface{}
	SnapshotPolicy() *string
	SetSnapshotPolicy(val *string)
	SnapshotPolicyInput() *string
	StorageEfficiencyEnabled() *string
	SetStorageEfficiencyEnabled(val *string)
	StorageEfficiencyEnabledInput() *string
	StorageVirtualMachineId() *string
	SetStorageVirtualMachineId(val *string)
	StorageVirtualMachineIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TieringPolicy() FsxVolumeOntapConfigurationTieringPolicyOutputReference
	TieringPolicyInput() interface{}
	VolumeStyle() *string
	SetVolumeStyle(val *string)
	VolumeStyleInput() *string
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
	PutAggregateConfiguration(value *FsxVolumeOntapConfigurationAggregateConfiguration)
	PutSnaplockConfiguration(value *FsxVolumeOntapConfigurationSnaplockConfiguration)
	PutTieringPolicy(value *FsxVolumeOntapConfigurationTieringPolicy)
	ResetAggregateConfiguration()
	ResetCopyTagsToBackups()
	ResetJunctionPath()
	ResetOntapVolumeType()
	ResetSecurityStyle()
	ResetSizeInBytes()
	ResetSizeInMegabytes()
	ResetSnaplockConfiguration()
	ResetSnapshotPolicy()
	ResetStorageEfficiencyEnabled()
	ResetStorageVirtualMachineId()
	ResetTieringPolicy()
	ResetVolumeStyle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxVolumeOntapConfigurationOutputReference
type jsiiProxy_FsxVolumeOntapConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) AggregateConfiguration() FsxVolumeOntapConfigurationAggregateConfigurationOutputReference {
	var returns FsxVolumeOntapConfigurationAggregateConfigurationOutputReference
	_jsii_.Get(
		j,
		"aggregateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) AggregateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) CopyTagsToBackups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) CopyTagsToBackupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) JunctionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) JunctionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) OntapVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) OntapVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SizeInBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SizeInBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SizeInMegabytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInMegabytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SizeInMegabytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInMegabytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SnaplockConfiguration() FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference {
	var returns FsxVolumeOntapConfigurationSnaplockConfigurationOutputReference
	_jsii_.Get(
		j,
		"snaplockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SnaplockConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snaplockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SnapshotPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) SnapshotPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) StorageEfficiencyEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEfficiencyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) StorageEfficiencyEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEfficiencyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) StorageVirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) StorageVirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) TieringPolicy() FsxVolumeOntapConfigurationTieringPolicyOutputReference {
	var returns FsxVolumeOntapConfigurationTieringPolicyOutputReference
	_jsii_.Get(
		j,
		"tieringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) TieringPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tieringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) VolumeStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) VolumeStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStyleInput",
		&returns,
	)
	return returns
}


func NewFsxVolumeOntapConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FsxVolumeOntapConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxVolumeOntapConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxVolumeOntapConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOntapConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxVolumeOntapConfigurationOutputReference_Override(f FsxVolumeOntapConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOntapConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetCopyTagsToBackups(val *string) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetJunctionPath(val *string) {
	if err := j.validateSetJunctionPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"junctionPath",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetOntapVolumeType(val *string) {
	if err := j.validateSetOntapVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ontapVolumeType",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetSizeInBytes(val *string) {
	if err := j.validateSetSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInBytes",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetSizeInMegabytes(val *string) {
	if err := j.validateSetSizeInMegabytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetSnapshotPolicy(val *string) {
	if err := j.validateSetSnapshotPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotPolicy",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetStorageEfficiencyEnabled(val *string) {
	if err := j.validateSetStorageEfficiencyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEfficiencyEnabled",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetStorageVirtualMachineId(val *string) {
	if err := j.validateSetStorageVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageVirtualMachineId",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOntapConfigurationOutputReference)SetVolumeStyle(val *string) {
	if err := j.validateSetVolumeStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeStyle",
		val,
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) PutAggregateConfiguration(value *FsxVolumeOntapConfigurationAggregateConfiguration) {
	if err := f.validatePutAggregateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAggregateConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) PutSnaplockConfiguration(value *FsxVolumeOntapConfigurationSnaplockConfiguration) {
	if err := f.validatePutSnaplockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSnaplockConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) PutTieringPolicy(value *FsxVolumeOntapConfigurationTieringPolicy) {
	if err := f.validatePutTieringPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTieringPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetAggregateConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetAggregateConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetJunctionPath() {
	_jsii_.InvokeVoid(
		f,
		"resetJunctionPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetOntapVolumeType() {
	_jsii_.InvokeVoid(
		f,
		"resetOntapVolumeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetSizeInBytes() {
	_jsii_.InvokeVoid(
		f,
		"resetSizeInBytes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetSizeInMegabytes() {
	_jsii_.InvokeVoid(
		f,
		"resetSizeInMegabytes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetSnaplockConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetSnaplockConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetSnapshotPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetSnapshotPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetStorageEfficiencyEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageEfficiencyEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetStorageVirtualMachineId() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageVirtualMachineId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetTieringPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetTieringPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ResetVolumeStyle() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeStyle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOntapConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

