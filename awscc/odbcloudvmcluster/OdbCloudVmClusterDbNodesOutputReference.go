// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/odbcloudvmcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OdbCloudVmClusterDbNodesOutputReference interface {
	cdktn.ComplexObject
	BackupIpId() *string
	SetBackupIpId(val *string)
	BackupIpIdInput() *string
	BackupVnic2Id() *string
	SetBackupVnic2Id(val *string)
	BackupVnic2IdInput() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DbNodeArn() *string
	SetDbNodeArn(val *string)
	DbNodeArnInput() *string
	DbNodeId() *string
	SetDbNodeId(val *string)
	DbNodeIdInput() *string
	DbNodeStorageSizeInGBs() *float64
	SetDbNodeStorageSizeInGBs(val *float64)
	DbNodeStorageSizeInGBsInput() *float64
	DbServerId() *string
	SetDbServerId(val *string)
	DbServerIdInput() *string
	DbSystemId() *string
	SetDbSystemId(val *string)
	DbSystemIdInput() *string
	// Experimental.
	Fqn() *string
	HostIpId() *string
	SetHostIpId(val *string)
	HostIpIdInput() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MemorySizeInGBs() *float64
	SetMemorySizeInGBs(val *float64)
	MemorySizeInGBsInput() *float64
	Ocid() *string
	SetOcid(val *string)
	OcidInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tags() OdbCloudVmClusterDbNodesTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vnic2Id() *string
	SetVnic2Id(val *string)
	Vnic2IdInput() *string
	VnicId() *string
	SetVnicId(val *string)
	VnicIdInput() *string
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
	PutTags(value interface{})
	ResetBackupIpId()
	ResetBackupVnic2Id()
	ResetCpuCoreCount()
	ResetDbNodeArn()
	ResetDbNodeId()
	ResetDbNodeStorageSizeInGBs()
	ResetDbServerId()
	ResetDbSystemId()
	ResetHostIpId()
	ResetHostname()
	ResetMemorySizeInGBs()
	ResetOcid()
	ResetStatus()
	ResetTags()
	ResetVnic2Id()
	ResetVnicId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OdbCloudVmClusterDbNodesOutputReference
type jsiiProxy_OdbCloudVmClusterDbNodesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) BackupIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) BackupIpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIpIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) BackupVnic2Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVnic2Id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) BackupVnic2IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVnic2IdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbNodeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNodeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbNodeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNodeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbNodeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNodeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbNodeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNodeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbNodeStorageSizeInGBsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGBsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) HostIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) HostIpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIpIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) MemorySizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) MemorySizeInGBsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGBsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) OcidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Tags() OdbCloudVmClusterDbNodesTagsList {
	var returns OdbCloudVmClusterDbNodesTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Vnic2Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnic2Id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Vnic2IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnic2IdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) VnicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) VnicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnicIdInput",
		&returns,
	)
	return returns
}


func NewOdbCloudVmClusterDbNodesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OdbCloudVmClusterDbNodesOutputReference {
	_init_.Initialize()

	if err := validateNewOdbCloudVmClusterDbNodesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbCloudVmClusterDbNodesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.odbCloudVmCluster.OdbCloudVmClusterDbNodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOdbCloudVmClusterDbNodesOutputReference_Override(o OdbCloudVmClusterDbNodesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.odbCloudVmCluster.OdbCloudVmClusterDbNodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetBackupIpId(val *string) {
	if err := j.validateSetBackupIpIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupIpId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetBackupVnic2Id(val *string) {
	if err := j.validateSetBackupVnic2IdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVnic2Id",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetDbNodeArn(val *string) {
	if err := j.validateSetDbNodeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeArn",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetDbNodeId(val *string) {
	if err := j.validateSetDbNodeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetDbNodeStorageSizeInGBs(val *float64) {
	if err := j.validateSetDbNodeStorageSizeInGBsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeStorageSizeInGBs",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetDbServerId(val *string) {
	if err := j.validateSetDbServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServerId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetHostIpId(val *string) {
	if err := j.validateSetHostIpIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostIpId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetMemorySizeInGBs(val *float64) {
	if err := j.validateSetMemorySizeInGBsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeInGBs",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetOcid(val *string) {
	if err := j.validateSetOcidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocid",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetVnic2Id(val *string) {
	if err := j.validateSetVnic2IdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnic2Id",
		val,
	)
}

func (j *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference)SetVnicId(val *string) {
	if err := j.validateSetVnicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnicId",
		val,
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) PutTags(value interface{}) {
	if err := o.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetBackupIpId() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupIpId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetBackupVnic2Id() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupVnic2Id",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetCpuCoreCount() {
	_jsii_.InvokeVoid(
		o,
		"resetCpuCoreCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetDbNodeArn() {
	_jsii_.InvokeVoid(
		o,
		"resetDbNodeArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetDbNodeId() {
	_jsii_.InvokeVoid(
		o,
		"resetDbNodeId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetDbNodeStorageSizeInGBs() {
	_jsii_.InvokeVoid(
		o,
		"resetDbNodeStorageSizeInGBs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetDbServerId() {
	_jsii_.InvokeVoid(
		o,
		"resetDbServerId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		o,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetHostIpId() {
	_jsii_.InvokeVoid(
		o,
		"resetHostIpId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		o,
		"resetHostname",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetMemorySizeInGBs() {
	_jsii_.InvokeVoid(
		o,
		"resetMemorySizeInGBs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetOcid() {
	_jsii_.InvokeVoid(
		o,
		"resetOcid",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		o,
		"resetStatus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetVnic2Id() {
	_jsii_.InvokeVoid(
		o,
		"resetVnic2Id",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ResetVnicId() {
	_jsii_.InvokeVoid(
		o,
		"resetVnicId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudVmClusterDbNodesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

