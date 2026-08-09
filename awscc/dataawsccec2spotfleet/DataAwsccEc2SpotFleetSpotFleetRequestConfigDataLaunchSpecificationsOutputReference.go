// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2spotfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference interface {
	cdktn.ComplexObject
	BlockDeviceMappings() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsBlockDeviceMappingsList
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
	EbsOptimized() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	IamInstanceProfile() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsIamInstanceProfileOutputReference
	ImageId() *string
	InstanceRequirements() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference
	InstanceType() *string
	InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecifications
	SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecifications)
	KernelId() *string
	KeyName() *string
	Monitoring() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsMonitoringOutputReference
	NetworkInterfaces() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesList
	Placement() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsPlacementOutputReference
	RamdiskId() *string
	SecurityGroups() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsSecurityGroupsList
	SpotPrice() *string
	SubnetId() *string
	TagSpecifications() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserData() *string
	WeightedCapacity() *float64
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

// The jsii proxy struct for DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference
type jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) BlockDeviceMappings() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsBlockDeviceMappingsList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) EbsOptimized() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) IamInstanceProfile() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsIamInstanceProfileOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) InstanceRequirements() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecifications {
	var returns *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) Monitoring() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsMonitoringOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) NetworkInterfaces() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) Placement() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsPlacementOutputReference {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) RamdiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) SecurityGroups() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsSecurityGroupsList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsSecurityGroupsList
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) TagSpecifications() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) WeightedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference_Override(d DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference)SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecifications) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

