// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotfleetwisecampaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotfleetwiseCampaignSignalsToFetchOutputReference interface {
	cdktn.ComplexObject
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
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
	ConditionLanguageVersion() *float64
	SetConditionLanguageVersion(val *float64)
	ConditionLanguageVersionInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FullyQualifiedName() *string
	SetFullyQualifiedName(val *string)
	FullyQualifiedNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SignalFetchConfig() IotfleetwiseCampaignSignalsToFetchSignalFetchConfigOutputReference
	SignalFetchConfigInput() interface{}
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
	PutSignalFetchConfig(value *IotfleetwiseCampaignSignalsToFetchSignalFetchConfig)
	ResetActions()
	ResetConditionLanguageVersion()
	ResetFullyQualifiedName()
	ResetSignalFetchConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotfleetwiseCampaignSignalsToFetchOutputReference
type jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ConditionLanguageVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"conditionLanguageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ConditionLanguageVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"conditionLanguageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) FullyQualifiedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) SignalFetchConfig() IotfleetwiseCampaignSignalsToFetchSignalFetchConfigOutputReference {
	var returns IotfleetwiseCampaignSignalsToFetchSignalFetchConfigOutputReference
	_jsii_.Get(
		j,
		"signalFetchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) SignalFetchConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signalFetchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotfleetwiseCampaignSignalsToFetchOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotfleetwiseCampaignSignalsToFetchOutputReference {
	_init_.Initialize()

	if err := validateNewIotfleetwiseCampaignSignalsToFetchOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaignSignalsToFetchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotfleetwiseCampaignSignalsToFetchOutputReference_Override(i IotfleetwiseCampaignSignalsToFetchOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseCampaign.IotfleetwiseCampaignSignalsToFetchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetConditionLanguageVersion(val *float64) {
	if err := j.validateSetConditionLanguageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionLanguageVersion",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetFullyQualifiedName(val *string) {
	if err := j.validateSetFullyQualifiedNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullyQualifiedName",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) PutSignalFetchConfig(value *IotfleetwiseCampaignSignalsToFetchSignalFetchConfig) {
	if err := i.validatePutSignalFetchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSignalFetchConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		i,
		"resetActions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ResetConditionLanguageVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetConditionLanguageVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ResetFullyQualifiedName() {
	_jsii_.InvokeVoid(
		i,
		"resetFullyQualifiedName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ResetSignalFetchConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetSignalFetchConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotfleetwiseCampaignSignalsToFetchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

