// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2ec2fleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2Ec2FleetSpotOptionsOutputReference interface {
	cdktn.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
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
	InstanceInterruptionBehavior() *string
	SetInstanceInterruptionBehavior(val *string)
	InstanceInterruptionBehaviorInput() *string
	InstancePoolsToUseCount() *float64
	SetInstancePoolsToUseCount(val *float64)
	InstancePoolsToUseCountInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaintenanceStrategies() Ec2Ec2FleetSpotOptionsMaintenanceStrategiesOutputReference
	MaintenanceStrategiesInput() interface{}
	MaxTotalPrice() *string
	SetMaxTotalPrice(val *string)
	MaxTotalPriceInput() *string
	MinTargetCapacity() *float64
	SetMinTargetCapacity(val *float64)
	MinTargetCapacityInput() *float64
	SingleAvailabilityZone() interface{}
	SetSingleAvailabilityZone(val interface{})
	SingleAvailabilityZoneInput() interface{}
	SingleInstanceType() interface{}
	SetSingleInstanceType(val interface{})
	SingleInstanceTypeInput() interface{}
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
	PutMaintenanceStrategies(value *Ec2Ec2FleetSpotOptionsMaintenanceStrategies)
	ResetAllocationStrategy()
	ResetInstanceInterruptionBehavior()
	ResetInstancePoolsToUseCount()
	ResetMaintenanceStrategies()
	ResetMaxTotalPrice()
	ResetMinTargetCapacity()
	ResetSingleAvailabilityZone()
	ResetSingleInstanceType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2Ec2FleetSpotOptionsOutputReference
type jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InstanceInterruptionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InstanceInterruptionBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InstancePoolsToUseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) MaintenanceStrategies() Ec2Ec2FleetSpotOptionsMaintenanceStrategiesOutputReference {
	var returns Ec2Ec2FleetSpotOptionsMaintenanceStrategiesOutputReference
	_jsii_.Get(
		j,
		"maintenanceStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) MaintenanceStrategiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceStrategiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) MaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) MaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) MinTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) MinTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) SingleAvailabilityZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) SingleAvailabilityZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) SingleInstanceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) SingleInstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2Ec2FleetSpotOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2Ec2FleetSpotOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetSpotOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetSpotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2Ec2FleetSpotOptionsOutputReference_Override(e Ec2Ec2FleetSpotOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2Ec2Fleet.Ec2Ec2FleetSpotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetInstanceInterruptionBehavior(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehavior",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetInstancePoolsToUseCount(val *float64) {
	if err := j.validateSetInstancePoolsToUseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolsToUseCount",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetMaxTotalPrice(val *string) {
	if err := j.validateSetMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetMinTargetCapacity(val *float64) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetSingleAvailabilityZone(val interface{}) {
	if err := j.validateSetSingleAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetSingleInstanceType(val interface{}) {
	if err := j.validateSetSingleInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleInstanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) PutMaintenanceStrategies(value *Ec2Ec2FleetSpotOptionsMaintenanceStrategies) {
	if err := e.validatePutMaintenanceStrategiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMaintenanceStrategies",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetInstanceInterruptionBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceInterruptionBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetInstancePoolsToUseCount() {
	_jsii_.InvokeVoid(
		e,
		"resetInstancePoolsToUseCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetMaintenanceStrategies() {
	_jsii_.InvokeVoid(
		e,
		"resetMaintenanceStrategies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetMaxTotalPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxTotalPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetSingleAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetSingleAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ResetSingleInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetSingleInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetSpotOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

