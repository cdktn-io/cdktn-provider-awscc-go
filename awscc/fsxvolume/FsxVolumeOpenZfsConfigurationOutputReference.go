// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/fsxvolume/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxVolumeOpenZfsConfigurationOutputReference interface {
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
	CopyTagsToSnapshots() interface{}
	SetCopyTagsToSnapshots(val interface{})
	CopyTagsToSnapshotsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NfsExports() FsxVolumeOpenZfsConfigurationNfsExportsList
	NfsExportsInput() interface{}
	Options() *[]*string
	SetOptions(val *[]*string)
	OptionsInput() *[]*string
	OriginSnapshot() FsxVolumeOpenZfsConfigurationOriginSnapshotOutputReference
	OriginSnapshotInput() interface{}
	ParentVolumeId() *string
	SetParentVolumeId(val *string)
	ParentVolumeIdInput() *string
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	RecordSizeKiB() *float64
	SetRecordSizeKiB(val *float64)
	RecordSizeKiBInput() *float64
	StorageCapacityQuotaGiB() *float64
	SetStorageCapacityQuotaGiB(val *float64)
	StorageCapacityQuotaGiBInput() *float64
	StorageCapacityReservationGiB() *float64
	SetStorageCapacityReservationGiB(val *float64)
	StorageCapacityReservationGiBInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserAndGroupQuotas() FsxVolumeOpenZfsConfigurationUserAndGroupQuotasList
	UserAndGroupQuotasInput() interface{}
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
	PutNfsExports(value interface{})
	PutOriginSnapshot(value *FsxVolumeOpenZfsConfigurationOriginSnapshot)
	PutUserAndGroupQuotas(value interface{})
	ResetCopyTagsToSnapshots()
	ResetDataCompressionType()
	ResetNfsExports()
	ResetOptions()
	ResetOriginSnapshot()
	ResetParentVolumeId()
	ResetReadOnly()
	ResetRecordSizeKiB()
	ResetStorageCapacityQuotaGiB()
	ResetStorageCapacityReservationGiB()
	ResetUserAndGroupQuotas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxVolumeOpenZfsConfigurationOutputReference
type jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) NfsExports() FsxVolumeOpenZfsConfigurationNfsExportsList {
	var returns FsxVolumeOpenZfsConfigurationNfsExportsList
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) NfsExportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) Options() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) OptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) OriginSnapshot() FsxVolumeOpenZfsConfigurationOriginSnapshotOutputReference {
	var returns FsxVolumeOpenZfsConfigurationOriginSnapshotOutputReference
	_jsii_.Get(
		j,
		"originSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) OriginSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ParentVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ParentVolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) RecordSizeKiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) RecordSizeKiBInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) StorageCapacityQuotaGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) StorageCapacityQuotaGiBInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) StorageCapacityReservationGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) StorageCapacityReservationGiBInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGiBInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) UserAndGroupQuotas() FsxVolumeOpenZfsConfigurationUserAndGroupQuotasList {
	var returns FsxVolumeOpenZfsConfigurationUserAndGroupQuotasList
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}


func NewFsxVolumeOpenZfsConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FsxVolumeOpenZfsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxVolumeOpenZfsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOpenZfsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxVolumeOpenZfsConfigurationOutputReference_Override(f FsxVolumeOpenZfsConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.fsxVolume.FsxVolumeOpenZfsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetCopyTagsToSnapshots(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetOptions(val *[]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetParentVolumeId(val *string) {
	if err := j.validateSetParentVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentVolumeId",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetRecordSizeKiB(val *float64) {
	if err := j.validateSetRecordSizeKiBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSizeKiB",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetStorageCapacityQuotaGiB(val *float64) {
	if err := j.validateSetStorageCapacityQuotaGiBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacityQuotaGiB",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetStorageCapacityReservationGiB(val *float64) {
	if err := j.validateSetStorageCapacityReservationGiBParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacityReservationGiB",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) PutNfsExports(value interface{}) {
	if err := f.validatePutNfsExportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) PutOriginSnapshot(value *FsxVolumeOpenZfsConfigurationOriginSnapshot) {
	if err := f.validatePutOriginSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putOriginSnapshot",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) PutUserAndGroupQuotas(value interface{}) {
	if err := f.validatePutUserAndGroupQuotasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putUserAndGroupQuotas",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetNfsExports() {
	_jsii_.InvokeVoid(
		f,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		f,
		"resetOptions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetOriginSnapshot() {
	_jsii_.InvokeVoid(
		f,
		"resetOriginSnapshot",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetParentVolumeId() {
	_jsii_.InvokeVoid(
		f,
		"resetParentVolumeId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetRecordSizeKiB() {
	_jsii_.InvokeVoid(
		f,
		"resetRecordSizeKiB",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetStorageCapacityQuotaGiB() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacityQuotaGiB",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetStorageCapacityReservationGiB() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacityReservationGiB",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		f,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (f *jsiiProxy_FsxVolumeOpenZfsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

