// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyexperiment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/evidentlyexperiment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvidentlyExperimentMetricGoalsOutputReference interface {
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
	DesiredChange() *string
	SetDesiredChange(val *string)
	DesiredChangeInput() *string
	EntityIdKey() *string
	SetEntityIdKey(val *string)
	EntityIdKeyInput() *string
	EventPattern() *string
	SetEventPattern(val *string)
	EventPatternInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnitLabel() *string
	SetUnitLabel(val *string)
	UnitLabelInput() *string
	ValueKey() *string
	SetValueKey(val *string)
	ValueKeyInput() *string
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
	ResetEventPattern()
	ResetUnitLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EvidentlyExperimentMetricGoalsOutputReference
type jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) DesiredChange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) DesiredChangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) EntityIdKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) EntityIdKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) EventPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) EventPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) UnitLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) UnitLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ValueKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ValueKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueKeyInput",
		&returns,
	)
	return returns
}


func NewEvidentlyExperimentMetricGoalsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EvidentlyExperimentMetricGoalsOutputReference {
	_init_.Initialize()

	if err := validateNewEvidentlyExperimentMetricGoalsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.evidentlyExperiment.EvidentlyExperimentMetricGoalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEvidentlyExperimentMetricGoalsOutputReference_Override(e EvidentlyExperimentMetricGoalsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.evidentlyExperiment.EvidentlyExperimentMetricGoalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetDesiredChange(val *string) {
	if err := j.validateSetDesiredChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredChange",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetEntityIdKey(val *string) {
	if err := j.validateSetEntityIdKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityIdKey",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetEventPattern(val *string) {
	if err := j.validateSetEventPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventPattern",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetUnitLabel(val *string) {
	if err := j.validateSetUnitLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitLabel",
		val,
	)
}

func (j *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference)SetValueKey(val *string) {
	if err := j.validateSetValueKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueKey",
		val,
	)
}

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ResetEventPattern() {
	_jsii_.InvokeVoid(
		e,
		"resetEventPattern",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ResetUnitLabel() {
	_jsii_.InvokeVoid(
		e,
		"resetUnitLabel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EvidentlyExperimentMetricGoalsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

