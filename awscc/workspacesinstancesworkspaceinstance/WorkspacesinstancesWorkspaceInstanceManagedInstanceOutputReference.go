// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesworkspaceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspacesinstancesworkspaceinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference interface {
	cdktn.ComplexObject
	BlockDeviceMappings() WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	CapacityReservationSpecification() WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationOutputReference
	CapacityReservationSpecificationInput() interface{}
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
	CpuOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptionsOutputReference
	CpuOptionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreditSpecification() WorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecificationOutputReference
	CreditSpecificationInput() interface{}
	DisableApiStop() interface{}
	SetDisableApiStop(val interface{})
	DisableApiStopInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	EnablePrimaryIpv6() interface{}
	SetEnablePrimaryIpv6(val interface{})
	EnablePrimaryIpv6Input() interface{}
	EnclaveOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptionsOutputReference
	EnclaveOptionsInput() interface{}
	// Experimental.
	Fqn() *string
	HibernationOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptionsOutputReference
	HibernationOptionsInput() interface{}
	IamInstanceProfile() WorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfileOutputReference
	IamInstanceProfileInput() interface{}
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	LicenseSpecifications() WorkspacesinstancesWorkspaceInstanceManagedInstanceLicenseSpecificationsList
	LicenseSpecificationsInput() interface{}
	MaintenanceOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptionsOutputReference
	MaintenanceOptionsInput() interface{}
	MetadataOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference
	MetadataOptionsInput() interface{}
	Monitoring() WorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoringOutputReference
	MonitoringInput() interface{}
	NetworkInterfaces() WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	NetworkPerformanceOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptionsOutputReference
	NetworkPerformanceOptionsInput() interface{}
	Placement() WorkspacesinstancesWorkspaceInstanceManagedInstancePlacementOutputReference
	PlacementInput() interface{}
	PrivateDnsNameOptions() WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptionsOutputReference
	PrivateDnsNameOptionsInput() interface{}
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	TagSpecifications() WorkspacesinstancesWorkspaceInstanceManagedInstanceTagSpecificationsList
	TagSpecificationsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
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
	PutBlockDeviceMappings(value interface{})
	PutCapacityReservationSpecification(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecification)
	PutCpuOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptions)
	PutCreditSpecification(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecification)
	PutEnclaveOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptions)
	PutHibernationOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptions)
	PutIamInstanceProfile(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfile)
	PutLicenseSpecifications(value interface{})
	PutMaintenanceOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptions)
	PutMetadataOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptions)
	PutMonitoring(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoring)
	PutNetworkInterfaces(value interface{})
	PutNetworkPerformanceOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptions)
	PutPlacement(value *WorkspacesinstancesWorkspaceInstanceManagedInstancePlacement)
	PutPrivateDnsNameOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptions)
	PutTagSpecifications(value interface{})
	ResetBlockDeviceMappings()
	ResetCapacityReservationSpecification()
	ResetCpuOptions()
	ResetCreditSpecification()
	ResetDisableApiStop()
	ResetEbsOptimized()
	ResetEnablePrimaryIpv6()
	ResetEnclaveOptions()
	ResetHibernationOptions()
	ResetIamInstanceProfile()
	ResetImageId()
	ResetInstanceType()
	ResetIpv6AddressCount()
	ResetKeyName()
	ResetLicenseSpecifications()
	ResetMaintenanceOptions()
	ResetMetadataOptions()
	ResetMonitoring()
	ResetNetworkInterfaces()
	ResetNetworkPerformanceOptions()
	ResetPlacement()
	ResetPrivateDnsNameOptions()
	ResetSubnetId()
	ResetTagSpecifications()
	ResetUserData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference
type jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) BlockDeviceMappings() WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsList {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CapacityReservationSpecification() WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CapacityReservationSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CpuOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CpuOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CreditSpecification() WorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecificationOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CreditSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EnablePrimaryIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EnablePrimaryIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EnclaveOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EnclaveOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) HibernationOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptionsOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) HibernationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) IamInstanceProfile() WorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfileOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) IamInstanceProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) LicenseSpecifications() WorkspacesinstancesWorkspaceInstanceManagedInstanceLicenseSpecificationsList {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceLicenseSpecificationsList
	_jsii_.Get(
		j,
		"licenseSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) LicenseSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MaintenanceOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptionsOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MaintenanceOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MetadataOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MetadataOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Monitoring() WorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoringOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) NetworkInterfaces() WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfacesList {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) NetworkPerformanceOptions() WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptionsOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) NetworkPerformanceOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkPerformanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Placement() WorkspacesinstancesWorkspaceInstanceManagedInstancePlacementOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstancePlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PrivateDnsNameOptions() WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptionsOutputReference {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PrivateDnsNameOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TagSpecifications() WorkspacesinstancesWorkspaceInstanceManagedInstanceTagSpecificationsList {
	var returns WorkspacesinstancesWorkspaceInstanceManagedInstanceTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


func NewWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference {
	_init_.Initialize()

	if err := validateNewWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesinstancesWorkspaceInstance.WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference_Override(w WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesinstancesWorkspaceInstance.WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetEnablePrimaryIpv6(val interface{}) {
	if err := j.validateSetEnablePrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrimaryIpv6",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutBlockDeviceMappings(value interface{}) {
	if err := w.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutCapacityReservationSpecification(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecification) {
	if err := w.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutCpuOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptions) {
	if err := w.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutCreditSpecification(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecification) {
	if err := w.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutEnclaveOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptions) {
	if err := w.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutHibernationOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptions) {
	if err := w.validatePutHibernationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHibernationOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutIamInstanceProfile(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfile) {
	if err := w.validatePutIamInstanceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIamInstanceProfile",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutLicenseSpecifications(value interface{}) {
	if err := w.validatePutLicenseSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLicenseSpecifications",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutMaintenanceOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptions) {
	if err := w.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutMetadataOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptions) {
	if err := w.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutMonitoring(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoring) {
	if err := w.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := w.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutNetworkPerformanceOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptions) {
	if err := w.validatePutNetworkPerformanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNetworkPerformanceOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutPlacement(value *WorkspacesinstancesWorkspaceInstanceManagedInstancePlacement) {
	if err := w.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlacement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutPrivateDnsNameOptions(value *WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptions) {
	if err := w.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PutTagSpecifications(value interface{}) {
	if err := w.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		w,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		w,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		w,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		w,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		w,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetEnablePrimaryIpv6() {
	_jsii_.InvokeVoid(
		w,
		"resetEnablePrimaryIpv6",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetHibernationOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetHibernationOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		w,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetImageId() {
	_jsii_.InvokeVoid(
		w,
		"resetImageId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		w,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		w,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		w,
		"resetKeyName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetLicenseSpecifications() {
	_jsii_.InvokeVoid(
		w,
		"resetLicenseSpecifications",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		w,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		w,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetNetworkPerformanceOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetNetworkPerformanceOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		w,
		"resetPlacement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		w,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		w,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		w,
		"resetUserData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (w *jsiiProxy_WorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

