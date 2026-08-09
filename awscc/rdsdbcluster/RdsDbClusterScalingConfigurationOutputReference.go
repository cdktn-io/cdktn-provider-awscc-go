// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rdsdbcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbClusterScalingConfigurationOutputReference interface {
	cdktn.ComplexObject
	AutoPause() interface{}
	SetAutoPause(val interface{})
	AutoPauseInput() interface{}
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
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MinCapacity() *float64
	SetMinCapacity(val *float64)
	MinCapacityInput() *float64
	SecondsBeforeTimeout() *float64
	SetSecondsBeforeTimeout(val *float64)
	SecondsBeforeTimeoutInput() *float64
	SecondsUntilAutoPause() *float64
	SetSecondsUntilAutoPause(val *float64)
	SecondsUntilAutoPauseInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutAction() *string
	SetTimeoutAction(val *string)
	TimeoutActionInput() *string
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
	ResetAutoPause()
	ResetMaxCapacity()
	ResetMinCapacity()
	ResetSecondsBeforeTimeout()
	ResetSecondsUntilAutoPause()
	ResetTimeoutAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RdsDbClusterScalingConfigurationOutputReference
type jsiiProxy_RdsDbClusterScalingConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) AutoPause() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) AutoPauseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) SecondsBeforeTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsBeforeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) SecondsBeforeTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsBeforeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) SecondsUntilAutoPause() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsUntilAutoPause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) SecondsUntilAutoPauseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsUntilAutoPauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) TimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) TimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutActionInput",
		&returns,
	)
	return returns
}


func NewRdsDbClusterScalingConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RdsDbClusterScalingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewRdsDbClusterScalingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDbClusterScalingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbClusterScalingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRdsDbClusterScalingConfigurationOutputReference_Override(r RdsDbClusterScalingConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbClusterScalingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetAutoPause(val interface{}) {
	if err := j.validateSetAutoPauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoPause",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetSecondsBeforeTimeout(val *float64) {
	if err := j.validateSetSecondsBeforeTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsBeforeTimeout",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetSecondsUntilAutoPause(val *float64) {
	if err := j.validateSetSecondsUntilAutoPauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsUntilAutoPause",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference)SetTimeoutAction(val *string) {
	if err := j.validateSetTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutAction",
		val,
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ResetAutoPause() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoPause",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ResetMinCapacity() {
	_jsii_.InvokeVoid(
		r,
		"resetMinCapacity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ResetSecondsBeforeTimeout() {
	_jsii_.InvokeVoid(
		r,
		"resetSecondsBeforeTimeout",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ResetSecondsUntilAutoPause() {
	_jsii_.InvokeVoid(
		r,
		"resetSecondsUntilAutoPause",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ResetTimeoutAction() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeoutAction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbClusterScalingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

