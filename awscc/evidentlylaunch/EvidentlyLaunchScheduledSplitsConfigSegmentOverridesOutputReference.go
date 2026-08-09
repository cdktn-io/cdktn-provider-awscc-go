// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/evidentlylaunch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference interface {
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
	EvaluationOrder() *float64
	SetEvaluationOrder(val *float64)
	EvaluationOrderInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Segment() *string
	SetSegment(val *string)
	SegmentInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Weights() EvidentlyLaunchScheduledSplitsConfigSegmentOverridesWeightsList
	WeightsInput() interface{}
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
	PutWeights(value interface{})
	ResetEvaluationOrder()
	ResetSegment()
	ResetWeights()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference
type jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) EvaluationOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) EvaluationOrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) Segment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) SegmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) Weights() EvidentlyLaunchScheduledSplitsConfigSegmentOverridesWeightsList {
	var returns EvidentlyLaunchScheduledSplitsConfigSegmentOverridesWeightsList
	_jsii_.Get(
		j,
		"weights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) WeightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightsInput",
		&returns,
	)
	return returns
}


func NewEvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewEvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.evidentlyLaunch.EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference_Override(e EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.evidentlyLaunch.EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetEvaluationOrder(val *float64) {
	if err := j.validateSetEvaluationOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationOrder",
		val,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetSegment(val *string) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) PutWeights(value interface{}) {
	if err := e.validatePutWeightsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putWeights",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ResetEvaluationOrder() {
	_jsii_.InvokeVoid(
		e,
		"resetEvaluationOrder",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ResetSegment() {
	_jsii_.InvokeVoid(
		e,
		"resetSegment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ResetWeights() {
	_jsii_.InvokeVoid(
		e,
		"resetWeights",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EvidentlyLaunchScheduledSplitsConfigSegmentOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

