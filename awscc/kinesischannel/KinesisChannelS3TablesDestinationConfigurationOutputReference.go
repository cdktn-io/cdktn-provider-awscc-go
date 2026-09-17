// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesischannel/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisChannelS3TablesDestinationConfigurationOutputReference interface {
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
	DeadLetterQueueS3Configuration() KinesisChannelS3TablesDestinationConfigurationDeadLetterQueueS3ConfigurationOutputReference
	DeadLetterQueueS3ConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3TablesConfigurationList() KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList
	S3TablesConfigurationListInput() interface{}
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
	PutDeadLetterQueueS3Configuration(value *KinesisChannelS3TablesDestinationConfigurationDeadLetterQueueS3Configuration)
	PutS3TablesConfigurationList(value interface{})
	ResetDataFreshnessInSeconds()
	ResetDeadLetterQueueS3Configuration()
	ResetS3TablesConfigurationList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisChannelS3TablesDestinationConfigurationOutputReference
type jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) DataFreshnessInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) DataFreshnessInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFreshnessInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) DeadLetterQueueS3Configuration() KinesisChannelS3TablesDestinationConfigurationDeadLetterQueueS3ConfigurationOutputReference {
	var returns KinesisChannelS3TablesDestinationConfigurationDeadLetterQueueS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"deadLetterQueueS3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) DeadLetterQueueS3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueS3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) S3TablesConfigurationList() KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList {
	var returns KinesisChannelS3TablesDestinationConfigurationS3TablesConfigurationListStructList
	_jsii_.Get(
		j,
		"s3TablesConfigurationList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) S3TablesConfigurationListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TablesConfigurationListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisChannelS3TablesDestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisChannelS3TablesDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisChannelS3TablesDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3TablesDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisChannelS3TablesDestinationConfigurationOutputReference_Override(k KinesisChannelS3TablesDestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisChannel.KinesisChannelS3TablesDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference)SetDataFreshnessInSeconds(val *float64) {
	if err := j.validateSetDataFreshnessInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFreshnessInSeconds",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) PutDeadLetterQueueS3Configuration(value *KinesisChannelS3TablesDestinationConfigurationDeadLetterQueueS3Configuration) {
	if err := k.validatePutDeadLetterQueueS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDeadLetterQueueS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) PutS3TablesConfigurationList(value interface{}) {
	if err := k.validatePutS3TablesConfigurationListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3TablesConfigurationList",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ResetDataFreshnessInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetDataFreshnessInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ResetDeadLetterQueueS3Configuration() {
	_jsii_.InvokeVoid(
		k,
		"resetDeadLetterQueueS3Configuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ResetS3TablesConfigurationList() {
	_jsii_.InvokeVoid(
		k,
		"resetS3TablesConfigurationList",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisChannelS3TablesDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

