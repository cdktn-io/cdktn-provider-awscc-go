// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticspipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotanalyticspipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotanalyticsPipelinePipelineActivitiesOutputReference interface {
	cdktn.ComplexObject
	AddAttributes() IotanalyticsPipelinePipelineActivitiesAddAttributesOutputReference
	AddAttributesInput() interface{}
	Channel() IotanalyticsPipelinePipelineActivitiesChannelOutputReference
	ChannelInput() interface{}
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
	Datastore() IotanalyticsPipelinePipelineActivitiesDatastoreOutputReference
	DatastoreInput() interface{}
	DeviceRegistryEnrich() IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference
	DeviceRegistryEnrichInput() interface{}
	DeviceShadowEnrich() IotanalyticsPipelinePipelineActivitiesDeviceShadowEnrichOutputReference
	DeviceShadowEnrichInput() interface{}
	Filter() IotanalyticsPipelinePipelineActivitiesFilterOutputReference
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lambda() IotanalyticsPipelinePipelineActivitiesLambdaOutputReference
	LambdaInput() interface{}
	Math() IotanalyticsPipelinePipelineActivitiesMathOutputReference
	MathInput() interface{}
	RemoveAttributes() IotanalyticsPipelinePipelineActivitiesRemoveAttributesOutputReference
	RemoveAttributesInput() interface{}
	SelectAttributes() IotanalyticsPipelinePipelineActivitiesSelectAttributesOutputReference
	SelectAttributesInput() interface{}
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
	PutAddAttributes(value *IotanalyticsPipelinePipelineActivitiesAddAttributes)
	PutChannel(value *IotanalyticsPipelinePipelineActivitiesChannel)
	PutDatastore(value *IotanalyticsPipelinePipelineActivitiesDatastore)
	PutDeviceRegistryEnrich(value *IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrich)
	PutDeviceShadowEnrich(value *IotanalyticsPipelinePipelineActivitiesDeviceShadowEnrich)
	PutFilter(value *IotanalyticsPipelinePipelineActivitiesFilter)
	PutLambda(value *IotanalyticsPipelinePipelineActivitiesLambda)
	PutMath(value *IotanalyticsPipelinePipelineActivitiesMath)
	PutRemoveAttributes(value *IotanalyticsPipelinePipelineActivitiesRemoveAttributes)
	PutSelectAttributes(value *IotanalyticsPipelinePipelineActivitiesSelectAttributes)
	ResetAddAttributes()
	ResetChannel()
	ResetDatastore()
	ResetDeviceRegistryEnrich()
	ResetDeviceShadowEnrich()
	ResetFilter()
	ResetLambda()
	ResetMath()
	ResetRemoveAttributes()
	ResetSelectAttributes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotanalyticsPipelinePipelineActivitiesOutputReference
type jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) AddAttributes() IotanalyticsPipelinePipelineActivitiesAddAttributesOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesAddAttributesOutputReference
	_jsii_.Get(
		j,
		"addAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) AddAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Channel() IotanalyticsPipelinePipelineActivitiesChannelOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesChannelOutputReference
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ChannelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Datastore() IotanalyticsPipelinePipelineActivitiesDatastoreOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesDatastoreOutputReference
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) DatastoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) DeviceRegistryEnrich() IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference
	_jsii_.Get(
		j,
		"deviceRegistryEnrich",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) DeviceRegistryEnrichInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceRegistryEnrichInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) DeviceShadowEnrich() IotanalyticsPipelinePipelineActivitiesDeviceShadowEnrichOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesDeviceShadowEnrichOutputReference
	_jsii_.Get(
		j,
		"deviceShadowEnrich",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) DeviceShadowEnrichInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceShadowEnrichInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Filter() IotanalyticsPipelinePipelineActivitiesFilterOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Lambda() IotanalyticsPipelinePipelineActivitiesLambdaOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Math() IotanalyticsPipelinePipelineActivitiesMathOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesMathOutputReference
	_jsii_.Get(
		j,
		"math",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) MathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) RemoveAttributes() IotanalyticsPipelinePipelineActivitiesRemoveAttributesOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesRemoveAttributesOutputReference
	_jsii_.Get(
		j,
		"removeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) RemoveAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) SelectAttributes() IotanalyticsPipelinePipelineActivitiesSelectAttributesOutputReference {
	var returns IotanalyticsPipelinePipelineActivitiesSelectAttributesOutputReference
	_jsii_.Get(
		j,
		"selectAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) SelectAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotanalyticsPipelinePipelineActivitiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotanalyticsPipelinePipelineActivitiesOutputReference {
	_init_.Initialize()

	if err := validateNewIotanalyticsPipelinePipelineActivitiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotanalyticsPipeline.IotanalyticsPipelinePipelineActivitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotanalyticsPipelinePipelineActivitiesOutputReference_Override(i IotanalyticsPipelinePipelineActivitiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotanalyticsPipeline.IotanalyticsPipelinePipelineActivitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutAddAttributes(value *IotanalyticsPipelinePipelineActivitiesAddAttributes) {
	if err := i.validatePutAddAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAddAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutChannel(value *IotanalyticsPipelinePipelineActivitiesChannel) {
	if err := i.validatePutChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putChannel",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutDatastore(value *IotanalyticsPipelinePipelineActivitiesDatastore) {
	if err := i.validatePutDatastoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDatastore",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutDeviceRegistryEnrich(value *IotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrich) {
	if err := i.validatePutDeviceRegistryEnrichParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDeviceRegistryEnrich",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutDeviceShadowEnrich(value *IotanalyticsPipelinePipelineActivitiesDeviceShadowEnrich) {
	if err := i.validatePutDeviceShadowEnrichParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDeviceShadowEnrich",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutFilter(value *IotanalyticsPipelinePipelineActivitiesFilter) {
	if err := i.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFilter",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutLambda(value *IotanalyticsPipelinePipelineActivitiesLambda) {
	if err := i.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLambda",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutMath(value *IotanalyticsPipelinePipelineActivitiesMath) {
	if err := i.validatePutMathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMath",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutRemoveAttributes(value *IotanalyticsPipelinePipelineActivitiesRemoveAttributes) {
	if err := i.validatePutRemoveAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRemoveAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) PutSelectAttributes(value *IotanalyticsPipelinePipelineActivitiesSelectAttributes) {
	if err := i.validatePutSelectAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSelectAttributes",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetAddAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetAddAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		i,
		"resetChannel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetDatastore() {
	_jsii_.InvokeVoid(
		i,
		"resetDatastore",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetDeviceRegistryEnrich() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceRegistryEnrich",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetDeviceShadowEnrich() {
	_jsii_.InvokeVoid(
		i,
		"resetDeviceShadowEnrich",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		i,
		"resetLambda",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetMath() {
	_jsii_.InvokeVoid(
		i,
		"resetMath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetRemoveAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetRemoveAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ResetSelectAttributes() {
	_jsii_.InvokeVoid(
		i,
		"resetSelectAttributes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IotanalyticsPipelinePipelineActivitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

