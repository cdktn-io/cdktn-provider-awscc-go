// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2ec2fleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference interface {
	cdktn.ComplexObject
	AvailabilityZone() *string
	AvailabilityZoneId() *string
	BlockDeviceMappings() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsList
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
	IamInstanceProfile() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfileOutputReference
	InstanceRequirements() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference
	InstanceType() *string
	InternalValue() *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverrides
	SetInternalValue(val *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverrides)
	KeyName() *string
	MaxPrice() *string
	MetadataOptions() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptionsOutputReference
	NetworkInterfaces() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesNetworkInterfacesList
	Placement() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference
	Priority() *float64
	SubnetId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference
type jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) BlockDeviceMappings() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsList {
	var returns DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) IamInstanceProfile() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfileOutputReference {
	var returns DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InstanceRequirements() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference {
	var returns DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InternalValue() *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverrides {
	var returns *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) MaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) MetadataOptions() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptionsOutputReference {
	var returns DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) NetworkInterfaces() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesNetworkInterfacesList {
	var returns DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Placement() DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference {
	var returns DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) WeightedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2Ec2Fleet.DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference_Override(d DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2Ec2Fleet.DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetInternalValue(val *DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

