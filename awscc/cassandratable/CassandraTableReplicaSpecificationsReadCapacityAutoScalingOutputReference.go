// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cassandratable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference interface {
	cdktn.ComplexObject
	AutoScalingDisabled() interface{}
	SetAutoScalingDisabled(val interface{})
	AutoScalingDisabledInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumUnits() *float64
	SetMaximumUnits(val *float64)
	MaximumUnitsInput() *float64
	MinimumUnits() *float64
	SetMinimumUnits(val *float64)
	MinimumUnitsInput() *float64
	ScalingPolicy() CassandraTableReplicaSpecificationsReadCapacityAutoScalingScalingPolicyOutputReference
	ScalingPolicyInput() interface{}
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
	PutScalingPolicy(value *CassandraTableReplicaSpecificationsReadCapacityAutoScalingScalingPolicy)
	ResetAutoScalingDisabled()
	ResetMaximumUnits()
	ResetMinimumUnits()
	ResetScalingPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference
type jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) AutoScalingDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) AutoScalingDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) MaximumUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) MaximumUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) MinimumUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) MinimumUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ScalingPolicy() CassandraTableReplicaSpecificationsReadCapacityAutoScalingScalingPolicyOutputReference {
	var returns CassandraTableReplicaSpecificationsReadCapacityAutoScalingScalingPolicyOutputReference
	_jsii_.Get(
		j,
		"scalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ScalingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewCassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cassandraTable.CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference_Override(c CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cassandraTable.CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetAutoScalingDisabled(val interface{}) {
	if err := j.validateSetAutoScalingDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingDisabled",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetMaximumUnits(val *float64) {
	if err := j.validateSetMaximumUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumUnits",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetMinimumUnits(val *float64) {
	if err := j.validateSetMinimumUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumUnits",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) PutScalingPolicy(value *CassandraTableReplicaSpecificationsReadCapacityAutoScalingScalingPolicy) {
	if err := c.validatePutScalingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScalingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ResetAutoScalingDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ResetMaximumUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ResetMinimumUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ResetScalingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetScalingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableReplicaSpecificationsReadCapacityAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

