// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotjobtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotJobTemplateAbortConfigCriteriaListStructOutputReference interface {
	cdktn.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	FailureType() *string
	SetFailureType(val *string)
	FailureTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinNumberOfExecutedThings() *float64
	SetMinNumberOfExecutedThings(val *float64)
	MinNumberOfExecutedThingsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThresholdPercentage() *float64
	SetThresholdPercentage(val *float64)
	ThresholdPercentageInput() *float64
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
	ResetAction()
	ResetFailureType()
	ResetMinNumberOfExecutedThings()
	ResetThresholdPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotJobTemplateAbortConfigCriteriaListStructOutputReference
type jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) FailureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) FailureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) MinNumberOfExecutedThings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfExecutedThings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) MinNumberOfExecutedThingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumberOfExecutedThingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ThresholdPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ThresholdPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdPercentageInput",
		&returns,
	)
	return returns
}


func NewIotJobTemplateAbortConfigCriteriaListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotJobTemplateAbortConfigCriteriaListStructOutputReference {
	_init_.Initialize()

	if err := validateNewIotJobTemplateAbortConfigCriteriaListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplateAbortConfigCriteriaListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotJobTemplateAbortConfigCriteriaListStructOutputReference_Override(i IotJobTemplateAbortConfigCriteriaListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplateAbortConfigCriteriaListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetFailureType(val *string) {
	if err := j.validateSetFailureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureType",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetMinNumberOfExecutedThings(val *float64) {
	if err := j.validateSetMinNumberOfExecutedThingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumberOfExecutedThings",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference)SetThresholdPercentage(val *float64) {
	if err := j.validateSetThresholdPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdPercentage",
		val,
	)
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		i,
		"resetAction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ResetFailureType() {
	_jsii_.InvokeVoid(
		i,
		"resetFailureType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ResetMinNumberOfExecutedThings() {
	_jsii_.InvokeVoid(
		i,
		"resetMinNumberOfExecutedThings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ResetThresholdPercentage() {
	_jsii_.InvokeVoid(
		i,
		"resetThresholdPercentage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplateAbortConfigCriteriaListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

