// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2ec2fleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference interface {
	cdktn.ComplexObject
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	AvailabilityZoneIdInput() *string
	AvailabilityZoneInput() *string
	BlockDeviceMappings() Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsList
	BlockDeviceMappingsInput() interface{}
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
	IamInstanceProfile() Ec2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfileOutputReference
	IamInstanceProfileInput() interface{}
	InstanceRequirements() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference
	InstanceRequirementsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	MaxPrice() *string
	SetMaxPrice(val *string)
	MaxPriceInput() *string
	MetadataOptions() Ec2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptionsOutputReference
	MetadataOptionsInput() interface{}
	NetworkInterfaces() Ec2Ec2FleetLaunchTemplateConfigsOverridesNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	Placement() Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference
	PlacementInput() interface{}
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	WeightedCapacity() *float64
	SetWeightedCapacity(val *float64)
	WeightedCapacityInput() *float64
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
	PutIamInstanceProfile(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfile)
	PutInstanceRequirements(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirements)
	PutMetadataOptions(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptions)
	PutNetworkInterfaces(value interface{})
	PutPlacement(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacement)
	ResetAvailabilityZone()
	ResetAvailabilityZoneId()
	ResetBlockDeviceMappings()
	ResetIamInstanceProfile()
	ResetInstanceRequirements()
	ResetInstanceType()
	ResetKeyName()
	ResetMaxPrice()
	ResetMetadataOptions()
	ResetNetworkInterfaces()
	ResetPlacement()
	ResetPriority()
	ResetSubnetId()
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference
type jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) BlockDeviceMappings() Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsList {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) IamInstanceProfile() Ec2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfileOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfileOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) IamInstanceProfileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InstanceRequirements() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InstanceRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) MaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) MaxPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) MetadataOptions() Ec2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptionsOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) MetadataOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) NetworkInterfaces() Ec2Ec2FleetLaunchTemplateConfigsOverridesNetworkInterfacesList {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Placement() Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) WeightedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) WeightedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}


func NewEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2Ec2FleetLaunchTemplateConfigsOverridesOutputReference_Override(e Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetMaxPrice(val *string) {
	if err := j.validateSetMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference)SetWeightedCapacity(val *float64) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PutBlockDeviceMappings(value interface{}) {
	if err := e.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PutIamInstanceProfile(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesIamInstanceProfile) {
	if err := e.validatePutIamInstanceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIamInstanceProfile",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PutInstanceRequirements(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirements) {
	if err := e.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PutMetadataOptions(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesMetadataOptions) {
	if err := e.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := e.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) PutPlacement(value *Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacement) {
	if err := e.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacement",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetAvailabilityZoneId() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZoneId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		e,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetMaxPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacement",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		e,
		"resetPriority",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

