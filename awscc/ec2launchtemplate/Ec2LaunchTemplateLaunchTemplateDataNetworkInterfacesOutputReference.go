// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2launchtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference interface {
	cdktn.ComplexObject
	AssociateCarrierIpAddress() interface{}
	SetAssociateCarrierIpAddress(val interface{})
	AssociateCarrierIpAddressInput() interface{}
	AssociatePublicIpAddress() interface{}
	SetAssociatePublicIpAddress(val interface{})
	AssociatePublicIpAddressInput() interface{}
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
	ConnectionTrackingSpecification() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecificationOutputReference
	ConnectionTrackingSpecificationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteOnTermination() interface{}
	SetDeleteOnTermination(val interface{})
	DeleteOnTerminationInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceIndex() *float64
	SetDeviceIndex(val *float64)
	DeviceIndexInput() *float64
	EnaSrdSpecification() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationOutputReference
	EnaSrdSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
	InterfaceType() *string
	SetInterfaceType(val *string)
	InterfaceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4PrefixCount() *float64
	SetIpv4PrefixCount(val *float64)
	Ipv4PrefixCountInput() *float64
	Ipv4Prefixes() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList
	Ipv4PrefixesInput() interface{}
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	Ipv6Addresses() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesList
	Ipv6AddressesInput() interface{}
	Ipv6PrefixCount() *float64
	SetIpv6PrefixCount(val *float64)
	Ipv6PrefixCountInput() *float64
	Ipv6Prefixes() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesList
	Ipv6PrefixesInput() interface{}
	NetworkCardIndex() *float64
	SetNetworkCardIndex(val *float64)
	NetworkCardIndexInput() *float64
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	PrimaryIpv6() interface{}
	SetPrimaryIpv6(val interface{})
	PrimaryIpv6Input() interface{}
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddresses() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesList
	PrivateIpAddressesInput() interface{}
	PrivateIpAddressInput() *string
	SecondaryPrivateIpAddressCount() *float64
	SetSecondaryPrivateIpAddressCount(val *float64)
	SecondaryPrivateIpAddressCountInput() *float64
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutConnectionTrackingSpecification(value *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification)
	PutEnaSrdSpecification(value *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecification)
	PutIpv4Prefixes(value interface{})
	PutIpv6Addresses(value interface{})
	PutIpv6Prefixes(value interface{})
	PutPrivateIpAddresses(value interface{})
	ResetAssociateCarrierIpAddress()
	ResetAssociatePublicIpAddress()
	ResetConnectionTrackingSpecification()
	ResetDeleteOnTermination()
	ResetDescription()
	ResetDeviceIndex()
	ResetEnaSrdSpecification()
	ResetGroups()
	ResetInterfaceType()
	ResetIpv4PrefixCount()
	ResetIpv4Prefixes()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetIpv6PrefixCount()
	ResetIpv6Prefixes()
	ResetNetworkCardIndex()
	ResetNetworkInterfaceId()
	ResetPrimaryIpv6()
	ResetPrivateIpAddress()
	ResetPrivateIpAddresses()
	ResetSecondaryPrivateIpAddressCount()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference
type jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) AssociateCarrierIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateCarrierIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) AssociateCarrierIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateCarrierIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ConnectionTrackingSpecification() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecificationOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecificationOutputReference
	_jsii_.Get(
		j,
		"connectionTrackingSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ConnectionTrackingSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTrackingSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) EnaSrdSpecification() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationOutputReference {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationOutputReference
	_jsii_.Get(
		j,
		"enaSrdSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) EnaSrdSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSrdSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv4PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv4PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv4Prefixes() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv4PrefixesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv4PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6Addresses() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesList {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesList
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6Prefixes() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesList {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesList
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Ipv6PrefixesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PrimaryIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PrimaryIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PrivateIpAddresses() Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesList {
	var returns Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesList
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PrivateIpAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) SecondaryPrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) SecondaryPrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference {
	_init_.Initialize()

	if err := validateNewEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference_Override(e Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetAssociateCarrierIpAddress(val interface{}) {
	if err := j.validateSetAssociateCarrierIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateCarrierIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetIpv4PrefixCount(val *float64) {
	if err := j.validateSetIpv4PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4PrefixCount",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetIpv6PrefixCount(val *float64) {
	if err := j.validateSetIpv6PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6PrefixCount",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetPrimaryIpv6(val interface{}) {
	if err := j.validateSetPrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryIpv6",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetSecondaryPrivateIpAddressCount(val *float64) {
	if err := j.validateSetSecondaryPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PutConnectionTrackingSpecification(value *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification) {
	if err := e.validatePutConnectionTrackingSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectionTrackingSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PutEnaSrdSpecification(value *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecification) {
	if err := e.validatePutEnaSrdSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnaSrdSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PutIpv4Prefixes(value interface{}) {
	if err := e.validatePutIpv4PrefixesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv4Prefixes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PutIpv6Addresses(value interface{}) {
	if err := e.validatePutIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6Addresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PutIpv6Prefixes(value interface{}) {
	if err := e.validatePutIpv6PrefixesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6Prefixes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) PutPrivateIpAddresses(value interface{}) {
	if err := e.validatePutPrivateIpAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrivateIpAddresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetAssociateCarrierIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetAssociateCarrierIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetConnectionTrackingSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectionTrackingSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		e,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetEnaSrdSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetEnaSrdSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetIpv4PrefixCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4PrefixCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetIpv4Prefixes() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv4Prefixes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetIpv6PrefixCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6PrefixCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetIpv6Prefixes() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Prefixes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetNetworkCardIndex() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkCardIndex",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetPrimaryIpv6() {
	_jsii_.InvokeVoid(
		e,
		"resetPrimaryIpv6",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetPrivateIpAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetSecondaryPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetSecondaryPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

