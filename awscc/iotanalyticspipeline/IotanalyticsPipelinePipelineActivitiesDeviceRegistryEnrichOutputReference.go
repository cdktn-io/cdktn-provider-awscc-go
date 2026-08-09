// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticspipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotanalyticspipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference interface {
	cdktn.ComplexObject
	Attribute() *string
	SetAttribute(val *string)
	AttributeInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Next() *string
	SetNext(val *string)
	NextInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThingName() *string
	SetThingName(val *string)
	ThingNameInput() *string
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
	ResetAttribute()
	ResetName()
	ResetNext()
	ResetRoleArn()
	ResetThingName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference
type jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) Next() *string {
	var returns *string
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) NextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ThingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ThingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingNameInput",
		&returns,
	)
	return returns
}


func NewIotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference {
	_init_.Initialize()

	if err := validateNewIotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotanalyticsPipeline.IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference_Override(i IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotanalyticsPipeline.IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetNext(val *string) {
	if err := j.validateSetNextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"next",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference)SetThingName(val *string) {
	if err := j.validateSetThingNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thingName",
		val,
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ResetNext() {
	_jsii_.InvokeVoid(
		i,
		"resetNext",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		i,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ResetThingName() {
	_jsii_.InvokeVoid(
		i,
		"resetThingName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

