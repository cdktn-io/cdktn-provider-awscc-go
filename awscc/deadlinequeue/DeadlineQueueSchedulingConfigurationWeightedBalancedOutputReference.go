// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinequeue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlinequeue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference interface {
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
	ErrorWeight() *float64
	SetErrorWeight(val *float64)
	ErrorWeightInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxPriorityOverride() DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverrideOutputReference
	MaxPriorityOverrideInput() interface{}
	MinPriorityOverride() DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverrideOutputReference
	MinPriorityOverrideInput() interface{}
	PriorityWeight() *float64
	SetPriorityWeight(val *float64)
	PriorityWeightInput() *float64
	RenderingTaskBuffer() *float64
	SetRenderingTaskBuffer(val *float64)
	RenderingTaskBufferInput() *float64
	RenderingTaskWeight() *float64
	SetRenderingTaskWeight(val *float64)
	RenderingTaskWeightInput() *float64
	SubmissionTimeWeight() *float64
	SetSubmissionTimeWeight(val *float64)
	SubmissionTimeWeightInput() *float64
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
	PutMaxPriorityOverride(value *DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverride)
	PutMinPriorityOverride(value *DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverride)
	ResetErrorWeight()
	ResetMaxPriorityOverride()
	ResetMinPriorityOverride()
	ResetPriorityWeight()
	ResetRenderingTaskBuffer()
	ResetRenderingTaskWeight()
	ResetSubmissionTimeWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference
type jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ErrorWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ErrorWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) MaxPriorityOverride() DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverrideOutputReference {
	var returns DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverrideOutputReference
	_jsii_.Get(
		j,
		"maxPriorityOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) MaxPriorityOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxPriorityOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) MinPriorityOverride() DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverrideOutputReference {
	var returns DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverrideOutputReference
	_jsii_.Get(
		j,
		"minPriorityOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) MinPriorityOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minPriorityOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) PriorityWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) PriorityWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) RenderingTaskBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renderingTaskBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) RenderingTaskBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renderingTaskBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) RenderingTaskWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renderingTaskWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) RenderingTaskWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renderingTaskWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) SubmissionTimeWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"submissionTimeWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) SubmissionTimeWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"submissionTimeWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference {
	_init_.Initialize()

	if err := validateNewDeadlineQueueSchedulingConfigurationWeightedBalancedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineQueue.DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference_Override(d DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineQueue.DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetErrorWeight(val *float64) {
	if err := j.validateSetErrorWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorWeight",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetPriorityWeight(val *float64) {
	if err := j.validateSetPriorityWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priorityWeight",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetRenderingTaskBuffer(val *float64) {
	if err := j.validateSetRenderingTaskBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renderingTaskBuffer",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetRenderingTaskWeight(val *float64) {
	if err := j.validateSetRenderingTaskWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renderingTaskWeight",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetSubmissionTimeWeight(val *float64) {
	if err := j.validateSetSubmissionTimeWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"submissionTimeWeight",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) PutMaxPriorityOverride(value *DeadlineQueueSchedulingConfigurationWeightedBalancedMaxPriorityOverride) {
	if err := d.validatePutMaxPriorityOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMaxPriorityOverride",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) PutMinPriorityOverride(value *DeadlineQueueSchedulingConfigurationWeightedBalancedMinPriorityOverride) {
	if err := d.validatePutMinPriorityOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMinPriorityOverride",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetErrorWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetErrorWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetMaxPriorityOverride() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxPriorityOverride",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetMinPriorityOverride() {
	_jsii_.InvokeVoid(
		d,
		"resetMinPriorityOverride",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetPriorityWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetPriorityWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetRenderingTaskBuffer() {
	_jsii_.InvokeVoid(
		d,
		"resetRenderingTaskBuffer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetRenderingTaskWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetRenderingTaskWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ResetSubmissionTimeWeight() {
	_jsii_.InvokeVoid(
		d,
		"resetSubmissionTimeWeight",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeadlineQueueSchedulingConfigurationWeightedBalancedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

