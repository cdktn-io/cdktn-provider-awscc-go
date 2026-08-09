// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccec2spotfleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference interface {
	cdktn.ComplexObject
	AssociatePublicIpAddress() cdktn.IResolvable
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
	DeleteOnTermination() cdktn.IResolvable
	Description() *string
	DeviceIndex() *float64
	// Experimental.
	Fqn() *string
	Groups() *[]*string
	InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfaces
	SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfaces)
	Ipv6AddressCount() *float64
	Ipv6Addresses() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesIpv6AddressesList
	NetworkInterfaceId() *string
	PrivateIpAddresses() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesPrivateIpAddressesList
	SecondaryPrivateIpAddressCount() *float64
	SubnetId() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference
type jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) AssociatePublicIpAddress() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) DeleteOnTermination() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) InternalValue() *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfaces {
	var returns *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfaces
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) Ipv6Addresses() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesIpv6AddressesList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesIpv6AddressesList
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) PrivateIpAddresses() DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesPrivateIpAddressesList {
	var returns DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesPrivateIpAddressesList
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) SecondaryPrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference_Override(d DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEc2SpotFleet.DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference)SetInternalValue(val *DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfaces) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

