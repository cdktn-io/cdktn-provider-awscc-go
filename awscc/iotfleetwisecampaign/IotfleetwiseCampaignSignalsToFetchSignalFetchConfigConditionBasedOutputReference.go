// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotfleetwisecampaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference interface {
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
	ConditionExpression() *string
	SetConditionExpression(val *string)
	ConditionExpressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TriggerMode() *string
	SetTriggerMode(val *string)
	TriggerModeInput() *string
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
	ResetConditionExpression()
	ResetTriggerMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference
type jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ConditionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ConditionExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) TriggerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) TriggerModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerModeInput",
		&returns,
	)
	return returns
}


func NewIotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference {
	_init_.Initialize()

	if err := validateNewIotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference_Override(i IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetConditionExpression(val *string) {
	if err := j.validateSetConditionExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionExpression",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference)SetTriggerMode(val *string) {
	if err := j.validateSetTriggerModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerMode",
		val,
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ResetConditionExpression() {
	_jsii_.InvokeVoid(
		i,
		"resetConditionExpression",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ResetTriggerMode() {
	_jsii_.InvokeVoid(
		i,
		"resetTriggerMode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBasedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

