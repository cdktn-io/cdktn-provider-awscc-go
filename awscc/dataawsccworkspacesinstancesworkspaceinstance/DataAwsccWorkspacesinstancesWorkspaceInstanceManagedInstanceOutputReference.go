// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccworkspacesinstancesworkspaceinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccworkspacesinstancesworkspaceinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference interface {
	cdktn.ComplexObject
	BlockDeviceMappings() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsList
	CapacityReservationSpecification() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationOutputReference
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
	CpuOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptionsOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreditSpecification() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecificationOutputReference
	DisableApiStop() cdktn.IResolvable
	EbsOptimized() cdktn.IResolvable
	EnablePrimaryIpv6() cdktn.IResolvable
	EnclaveOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptionsOutputReference
	// Experimental.
	Fqn() *string
	HibernationOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptionsOutputReference
	IamInstanceProfile() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfileOutputReference
	ImageId() *string
	InstanceType() *string
	InternalValue() *DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstance
	SetInternalValue(val *DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstance)
	Ipv6AddressCount() *float64
	KeyName() *string
	LicenseSpecifications() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceLicenseSpecificationsList
	MaintenanceOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptionsOutputReference
	MetadataOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference
	Monitoring() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoringOutputReference
	NetworkInterfaces() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfacesList
	NetworkPerformanceOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptionsOutputReference
	Placement() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstancePlacementOutputReference
	PrivateDnsNameOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptionsOutputReference
	SubnetId() *string
	TagSpecifications() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceTagSpecificationsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserData() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference
type jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) BlockDeviceMappings() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsList {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CapacityReservationSpecification() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCapacityReservationSpecificationOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CpuOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptionsOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) CreditSpecification() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecificationOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecificationOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) DisableApiStop() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EbsOptimized() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EnablePrimaryIpv6() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enablePrimaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) EnclaveOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptionsOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) HibernationOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptionsOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) IamInstanceProfile() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfileOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InternalValue() *DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstance {
	var returns *DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) LicenseSpecifications() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceLicenseSpecificationsList {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceLicenseSpecificationsList
	_jsii_.Get(
		j,
		"licenseSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MaintenanceOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptionsOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) MetadataOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Monitoring() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoringOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) NetworkInterfaces() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfacesList {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) NetworkPerformanceOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptionsOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Placement() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstancePlacementOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstancePlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) PrivateDnsNameOptions() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptionsOutputReference {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptionsOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TagSpecifications() DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceTagSpecificationsList {
	var returns DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


func NewDataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWorkspacesinstancesWorkspaceInstance.DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference_Override(d DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccWorkspacesinstancesWorkspaceInstance.DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetInternalValue(val *DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccWorkspacesinstancesWorkspaceInstanceManagedInstanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

