// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2instance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2InstanceNetworkInterfacesOutputReference interface {
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
	DeviceIndex() *string
	SetDeviceIndex(val *string)
	DeviceIndexInput() *string
	EnaSrdSpecification() Ec2InstanceNetworkInterfacesEnaSrdSpecificationOutputReference
	EnaSrdSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	GroupSet() *[]*string
	SetGroupSet(val *[]*string)
	GroupSetInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6AddressCount() *float64
	SetIpv6AddressCount(val *float64)
	Ipv6AddressCountInput() *float64
	Ipv6Addresses() Ec2InstanceNetworkInterfacesIpv6AddressesList
	Ipv6AddressesInput() interface{}
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddresses() Ec2InstanceNetworkInterfacesPrivateIpAddressesList
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
	PutEnaSrdSpecification(value *Ec2InstanceNetworkInterfacesEnaSrdSpecification)
	PutIpv6Addresses(value interface{})
	PutPrivateIpAddresses(value interface{})
	ResetAssociateCarrierIpAddress()
	ResetAssociatePublicIpAddress()
	ResetDeleteOnTermination()
	ResetDescription()
	ResetDeviceIndex()
	ResetEnaSrdSpecification()
	ResetGroupSet()
	ResetIpv6AddressCount()
	ResetIpv6Addresses()
	ResetNetworkInterfaceId()
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

// The jsii proxy struct for Ec2InstanceNetworkInterfacesOutputReference
type jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) AssociateCarrierIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateCarrierIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) AssociateCarrierIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateCarrierIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) DeleteOnTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) DeleteOnTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) DeviceIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) DeviceIndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) EnaSrdSpecification() Ec2InstanceNetworkInterfacesEnaSrdSpecificationOutputReference {
	var returns Ec2InstanceNetworkInterfacesEnaSrdSpecificationOutputReference
	_jsii_.Get(
		j,
		"enaSrdSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) EnaSrdSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSrdSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GroupSet() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GroupSetInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Ipv6Addresses() Ec2InstanceNetworkInterfacesIpv6AddressesList {
	var returns Ec2InstanceNetworkInterfacesIpv6AddressesList
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Ipv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PrivateIpAddresses() Ec2InstanceNetworkInterfacesPrivateIpAddressesList {
	var returns Ec2InstanceNetworkInterfacesPrivateIpAddressesList
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PrivateIpAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) SecondaryPrivateIpAddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) SecondaryPrivateIpAddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryPrivateIpAddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2InstanceNetworkInterfacesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2InstanceNetworkInterfacesOutputReference {
	_init_.Initialize()

	if err := validateNewEc2InstanceNetworkInterfacesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Instance.Ec2InstanceNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2InstanceNetworkInterfacesOutputReference_Override(e Ec2InstanceNetworkInterfacesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Instance.Ec2InstanceNetworkInterfacesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetAssociateCarrierIpAddress(val interface{}) {
	if err := j.validateSetAssociateCarrierIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateCarrierIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetDeleteOnTermination(val interface{}) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetDeviceIndex(val *string) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetGroupSet(val *[]*string) {
	if err := j.validateSetGroupSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupSet",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetSecondaryPrivateIpAddressCount(val *float64) {
	if err := j.validateSetSecondaryPrivateIpAddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIpAddressCount",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PutEnaSrdSpecification(value *Ec2InstanceNetworkInterfacesEnaSrdSpecification) {
	if err := e.validatePutEnaSrdSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnaSrdSpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PutIpv6Addresses(value interface{}) {
	if err := e.validatePutIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIpv6Addresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) PutPrivateIpAddresses(value interface{}) {
	if err := e.validatePutPrivateIpAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrivateIpAddresses",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetAssociateCarrierIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetAssociateCarrierIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		e,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetEnaSrdSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetEnaSrdSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetGroupSet() {
	_jsii_.InvokeVoid(
		e,
		"resetGroupSet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		e,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetPrivateIpAddresses() {
	_jsii_.InvokeVoid(
		e,
		"resetPrivateIpAddresses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetSecondaryPrivateIpAddressCount() {
	_jsii_.InvokeVoid(
		e,
		"resetSecondaryPrivateIpAddressCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2InstanceNetworkInterfacesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

