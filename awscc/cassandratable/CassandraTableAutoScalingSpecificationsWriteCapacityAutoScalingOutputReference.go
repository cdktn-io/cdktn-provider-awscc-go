// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cassandratable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference interface {
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
	ScalingPolicy() CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyOutputReference
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
	PutScalingPolicy(value *CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicy)
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

// The jsii proxy struct for CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference
type jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) AutoScalingDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) AutoScalingDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) MaximumUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) MaximumUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) MinimumUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) MinimumUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ScalingPolicy() CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyOutputReference {
	var returns CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicyOutputReference
	_jsii_.Get(
		j,
		"scalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ScalingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cassandraTable.CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference_Override(c CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cassandraTable.CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetAutoScalingDisabled(val interface{}) {
	if err := j.validateSetAutoScalingDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingDisabled",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetMaximumUnits(val *float64) {
	if err := j.validateSetMaximumUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumUnits",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetMinimumUnits(val *float64) {
	if err := j.validateSetMinimumUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumUnits",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) PutScalingPolicy(value *CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingScalingPolicy) {
	if err := c.validatePutScalingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScalingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ResetAutoScalingDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ResetMaximumUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ResetMinimumUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ResetScalingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetScalingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CassandraTableAutoScalingSpecificationsWriteCapacityAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

