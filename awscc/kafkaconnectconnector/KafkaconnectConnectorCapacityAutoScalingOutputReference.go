// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kafkaconnectconnector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KafkaconnectConnectorCapacityAutoScalingOutputReference interface {
	cdktn.ComplexObject
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
	MaxAutoscalingTaskCount() *float64
	SetMaxAutoscalingTaskCount(val *float64)
	MaxAutoscalingTaskCountInput() *float64
	MaxWorkerCount() *float64
	SetMaxWorkerCount(val *float64)
	MaxWorkerCountInput() *float64
	McuCount() *float64
	SetMcuCount(val *float64)
	McuCountInput() *float64
	MinWorkerCount() *float64
	SetMinWorkerCount(val *float64)
	MinWorkerCountInput() *float64
	ScaleInPolicy() KafkaconnectConnectorCapacityAutoScalingScaleInPolicyOutputReference
	ScaleInPolicyInput() interface{}
	ScaleOutPolicy() KafkaconnectConnectorCapacityAutoScalingScaleOutPolicyOutputReference
	ScaleOutPolicyInput() interface{}
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
	PutScaleInPolicy(value *KafkaconnectConnectorCapacityAutoScalingScaleInPolicy)
	PutScaleOutPolicy(value *KafkaconnectConnectorCapacityAutoScalingScaleOutPolicy)
	ResetMaxAutoscalingTaskCount()
	ResetMaxWorkerCount()
	ResetMcuCount()
	ResetMinWorkerCount()
	ResetScaleInPolicy()
	ResetScaleOutPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KafkaconnectConnectorCapacityAutoScalingOutputReference
type jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) MaxAutoscalingTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAutoscalingTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) MaxAutoscalingTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAutoscalingTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) MaxWorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) MaxWorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) McuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mcuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) McuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mcuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) MinWorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkerCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) MinWorkerCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkerCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ScaleInPolicy() KafkaconnectConnectorCapacityAutoScalingScaleInPolicyOutputReference {
	var returns KafkaconnectConnectorCapacityAutoScalingScaleInPolicyOutputReference
	_jsii_.Get(
		j,
		"scaleInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ScaleInPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ScaleOutPolicy() KafkaconnectConnectorCapacityAutoScalingScaleOutPolicyOutputReference {
	var returns KafkaconnectConnectorCapacityAutoScalingScaleOutPolicyOutputReference
	_jsii_.Get(
		j,
		"scaleOutPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ScaleOutPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scaleOutPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKafkaconnectConnectorCapacityAutoScalingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KafkaconnectConnectorCapacityAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewKafkaconnectConnectorCapacityAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kafkaconnectConnector.KafkaconnectConnectorCapacityAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKafkaconnectConnectorCapacityAutoScalingOutputReference_Override(k KafkaconnectConnectorCapacityAutoScalingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kafkaconnectConnector.KafkaconnectConnectorCapacityAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetMaxAutoscalingTaskCount(val *float64) {
	if err := j.validateSetMaxAutoscalingTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAutoscalingTaskCount",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetMaxWorkerCount(val *float64) {
	if err := j.validateSetMaxWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkerCount",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetMcuCount(val *float64) {
	if err := j.validateSetMcuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mcuCount",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetMinWorkerCount(val *float64) {
	if err := j.validateSetMinWorkerCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWorkerCount",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) PutScaleInPolicy(value *KafkaconnectConnectorCapacityAutoScalingScaleInPolicy) {
	if err := k.validatePutScaleInPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putScaleInPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) PutScaleOutPolicy(value *KafkaconnectConnectorCapacityAutoScalingScaleOutPolicy) {
	if err := k.validatePutScaleOutPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putScaleOutPolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ResetMaxAutoscalingTaskCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxAutoscalingTaskCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ResetMaxWorkerCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxWorkerCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ResetMcuCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMcuCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ResetMinWorkerCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMinWorkerCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ResetScaleInPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleInPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ResetScaleOutPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleOutPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorCapacityAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

