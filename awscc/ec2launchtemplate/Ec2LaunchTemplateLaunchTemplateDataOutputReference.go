// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2launchtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2LaunchTemplateLaunchTemplateDataOutputReference interface {
	cdktn.ComplexObject
	BlockDeviceMappings() Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
	CapacityReservationSpecification() Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationOutputReference
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
	CpuOptions() Ec2LaunchTemplateLaunchTemplateDataCpuOptionsOutputReference
	CpuOptionsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreditSpecification() Ec2LaunchTemplateLaunchTemplateDataCreditSpecificationOutputReference
	CreditSpecificationInput() interface{}
	DisableApiStop() interface{}
	SetDisableApiStop(val interface{})
	DisableApiStopInput() interface{}
	DisableApiTermination() interface{}
	SetDisableApiTermination(val interface{})
	DisableApiTerminationInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	ElasticGpuSpecifications() Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsList
	ElasticGpuSpecificationsInput() interface{}
	ElasticInferenceAccelerators() Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsList
	ElasticInferenceAcceleratorsInput() interface{}
	EnclaveOptions() Ec2LaunchTemplateLaunchTemplateDataEnclaveOptionsOutputReference
	EnclaveOptionsInput() interface{}
	// Experimental.
	Fqn() *string
	HibernationOptions() Ec2LaunchTemplateLaunchTemplateDataHibernationOptionsOutputReference
	HibernationOptionsInput() interface{}
	IamInstanceProfile() Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfileOutputReference
	IamInstanceProfileInput() interface{}
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceInitiatedShutdownBehavior() *string
	SetInstanceInitiatedShutdownBehavior(val *string)
	InstanceInitiatedShutdownBehaviorInput() *string
	InstanceMarketOptions() Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsOutputReference
	InstanceMarketOptionsInput() interface{}
	InstanceRequirements() Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference
	InstanceRequirementsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KernelId() *string
	SetKernelId(val *string)
	KernelIdInput() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	LicenseSpecifications() Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsList
	LicenseSpecificationsInput() interface{}
	MaintenanceOptions() Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptionsOutputReference
	MaintenanceOptionsInput() interface{}
	MetadataOptions() Ec2LaunchTemplateLaunchTemplateDataMetadataOptionsOutputReference
	MetadataOptionsInput() interface{}
	Monitoring() Ec2LaunchTemplateLaunchTemplateDataMonitoringOutputReference
	MonitoringInput() interface{}
	NetworkInterfaces() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	Placement() Ec2LaunchTemplateLaunchTemplateDataPlacementOutputReference
	PlacementInput() interface{}
	PrivateDnsNameOptions() Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptionsOutputReference
	PrivateDnsNameOptionsInput() interface{}
	RamDiskId() *string
	SetRamDiskId(val *string)
	RamDiskIdInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	TagSpecifications() Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsList
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
	PutCapacityReservationSpecification(value *Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification)
	PutCpuOptions(value *Ec2LaunchTemplateLaunchTemplateDataCpuOptions)
	PutCreditSpecification(value *Ec2LaunchTemplateLaunchTemplateDataCreditSpecification)
	PutElasticGpuSpecifications(value interface{})
	PutElasticInferenceAccelerators(value interface{})
	PutEnclaveOptions(value *Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions)
	PutHibernationOptions(value *Ec2LaunchTemplateLaunchTemplateDataHibernationOptions)
	PutIamInstanceProfile(value *Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile)
	PutInstanceMarketOptions(value *Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions)
	PutInstanceRequirements(value *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements)
	PutLicenseSpecifications(value interface{})
	PutMaintenanceOptions(value *Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions)
	PutMetadataOptions(value *Ec2LaunchTemplateLaunchTemplateDataMetadataOptions)
	PutMonitoring(value *Ec2LaunchTemplateLaunchTemplateDataMonitoring)
	PutNetworkInterfaces(value interface{})
	PutPlacement(value *Ec2LaunchTemplateLaunchTemplateDataPlacement)
	PutPrivateDnsNameOptions(value *Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions)
	PutTagSpecifications(value interface{})
	ResetBlockDeviceMappings()
	ResetCapacityReservationSpecification()
	ResetCpuOptions()
	ResetCreditSpecification()
	ResetDisableApiStop()
	ResetDisableApiTermination()
	ResetEbsOptimized()
	ResetElasticGpuSpecifications()
	ResetElasticInferenceAccelerators()
	ResetEnclaveOptions()
	ResetHibernationOptions()
	ResetIamInstanceProfile()
	ResetImageId()
	ResetInstanceInitiatedShutdownBehavior()
	ResetInstanceMarketOptions()
	ResetInstanceRequirements()
	ResetInstanceType()
	ResetKernelId()
	ResetKeyName()
	ResetLicenseSpecifications()
	ResetMaintenanceOptions()
	ResetMetadataOptions()
	ResetMonitoring()
	ResetNetworkInterfaces()
	ResetPlacement()
	ResetPrivateDnsNameOptions()
	ResetRamDiskId()
	ResetSecurityGroupIds()
	ResetSecurityGroups()
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

// The jsii proxy struct for Ec2LaunchTemplateLaunchTemplateDataOutputReference
type jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) BlockDeviceMappings() Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsList {
	var returns Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CapacityReservationSpecification() Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CapacityReservationSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CpuOptions() Ec2LaunchTemplateLaunchTemplateDataCpuOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CpuOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CreditSpecification() Ec2LaunchTemplateLaunchTemplateDataCreditSpecificationOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) CreditSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ElasticGpuSpecifications() Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsList {
	var returns Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsList
	_jsii_.Get(
		j,
		"elasticGpuSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ElasticGpuSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticGpuSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ElasticInferenceAccelerators() Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsList {
	var returns Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsList
	_jsii_.Get(
		j,
		"elasticInferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ElasticInferenceAcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticInferenceAcceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) EnclaveOptions() Ec2LaunchTemplateLaunchTemplateDataEnclaveOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) EnclaveOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) HibernationOptions() Ec2LaunchTemplateLaunchTemplateDataHibernationOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataHibernationOptionsOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) HibernationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) IamInstanceProfile() Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfileOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) IamInstanceProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceMarketOptions() Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceMarketOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMarketOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceRequirements() Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) LicenseSpecifications() Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsList {
	var returns Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsList
	_jsii_.Get(
		j,
		"licenseSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) LicenseSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) MaintenanceOptions() Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptionsOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) MaintenanceOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) MetadataOptions() Ec2LaunchTemplateLaunchTemplateDataMetadataOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) MetadataOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) Monitoring() Ec2LaunchTemplateLaunchTemplateDataMonitoringOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) NetworkInterfaces() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesList {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) Placement() Ec2LaunchTemplateLaunchTemplateDataPlacementOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PrivateDnsNameOptions() Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptionsOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PrivateDnsNameOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) RamDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) TagSpecifications() Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsList {
	var returns Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}


func NewEc2LaunchTemplateLaunchTemplateDataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2LaunchTemplateLaunchTemplateDataOutputReference {
	_init_.Initialize()

	if err := validateNewEc2LaunchTemplateLaunchTemplateDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2LaunchTemplateLaunchTemplateDataOutputReference_Override(e Ec2LaunchTemplateLaunchTemplateDataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetRamDiskId(val *string) {
	if err := j.validateSetRamDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutBlockDeviceMappings(value interface{}) {
	if err := e.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutCapacityReservationSpecification(value *Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification) {
	if err := e.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutCpuOptions(value *Ec2LaunchTemplateLaunchTemplateDataCpuOptions) {
	if err := e.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutCreditSpecification(value *Ec2LaunchTemplateLaunchTemplateDataCreditSpecification) {
	if err := e.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutElasticGpuSpecifications(value interface{}) {
	if err := e.validatePutElasticGpuSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putElasticGpuSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutElasticInferenceAccelerators(value interface{}) {
	if err := e.validatePutElasticInferenceAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putElasticInferenceAccelerators",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutEnclaveOptions(value *Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions) {
	if err := e.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutHibernationOptions(value *Ec2LaunchTemplateLaunchTemplateDataHibernationOptions) {
	if err := e.validatePutHibernationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHibernationOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutIamInstanceProfile(value *Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile) {
	if err := e.validatePutIamInstanceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIamInstanceProfile",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutInstanceMarketOptions(value *Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions) {
	if err := e.validatePutInstanceMarketOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceMarketOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutInstanceRequirements(value *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements) {
	if err := e.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutLicenseSpecifications(value interface{}) {
	if err := e.validatePutLicenseSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLicenseSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutMaintenanceOptions(value *Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions) {
	if err := e.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutMetadataOptions(value *Ec2LaunchTemplateLaunchTemplateDataMetadataOptions) {
	if err := e.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutMonitoring(value *Ec2LaunchTemplateLaunchTemplateDataMonitoring) {
	if err := e.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := e.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutPlacement(value *Ec2LaunchTemplateLaunchTemplateDataPlacement) {
	if err := e.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacement",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutPrivateDnsNameOptions(value *Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions) {
	if err := e.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) PutTagSpecifications(value interface{}) {
	if err := e.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		e,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		e,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetElasticGpuSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticGpuSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetElasticInferenceAccelerators() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticInferenceAccelerators",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetHibernationOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetHibernationOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		e,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetImageId() {
	_jsii_.InvokeVoid(
		e,
		"resetImageId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetInstanceMarketOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceMarketOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetKernelId() {
	_jsii_.InvokeVoid(
		e,
		"resetKernelId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetLicenseSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetLicenseSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		e,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetRamDiskId() {
	_jsii_.InvokeVoid(
		e,
		"resetRamDiskId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		e,
		"resetUserData",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

