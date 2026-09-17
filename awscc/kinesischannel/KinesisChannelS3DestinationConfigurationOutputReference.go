// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesischannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisChannelS3DestinationConfigurationOutputReference interface {
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
	DataFreshnessInSeconds() *float64
	SetDataFreshnessInSeconds(val *float64)
	DataFreshnessInSecondsInput() *float64
	DeadLetterQueueS3Configuration() KinesisChannelS3DestinationConfigurationDeadLetterQueueS3ConfigurationOutputReference
	DeadLetterQueueS3ConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageConfiguration() KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference
	StorageConfigurationInput() interface{}
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
	PutDeadLetterQueueS3Configuration(value *KinesisChannelS3DestinationConfigurationDeadLetterQueueS3Configuration)
	PutStorageConfiguration(value *KinesisChannelS3DestinationConfigurationStorageConfiguration)
	ResetDataFreshnessInSeconds()
	ResetDeadLetterQueueS3Configuration()
	ResetStorageConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisChannelS3DestinationConfigurationOutputReference
type jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) DataFreshnessInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) DataFreshnessInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) DeadLetterQueueS3Configuration() KinesisChannelS3DestinationConfigurationDeadLetterQueueS3ConfigurationOutputReference {
	var returns KinesisChannelS3DestinationConfigurationDeadLetterQueueS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"deadLetterQueueS3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) DeadLetterQueueS3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueS3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) StorageConfiguration() KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference {
	var returns KinesisChannelS3DestinationConfigurationStorageConfigurationOutputReference
	_jsii_.Get(
		j,
		"storageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) StorageConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisChannelS3DestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisChannelS3DestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisChannelS3DestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3DestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisChannelS3DestinationConfigurationOutputReference_Override(k KinesisChannelS3DestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3DestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference)SetDataFreshnessInSeconds(val *float64) {
	if err := j.validateSetDataFreshnessInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFreshnessInSeconds",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) PutDeadLetterQueueS3Configuration(value *KinesisChannelS3DestinationConfigurationDeadLetterQueueS3Configuration) {
	if err := k.validatePutDeadLetterQueueS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDeadLetterQueueS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) PutStorageConfiguration(value *KinesisChannelS3DestinationConfigurationStorageConfiguration) {
	if err := k.validatePutStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putStorageConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ResetDataFreshnessInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetDataFreshnessInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ResetDeadLetterQueueS3Configuration() {
	_jsii_.InvokeVoid(
		k,
		"resetDeadLetterQueueS3Configuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ResetStorageConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetStorageConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisChannelS3DestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

