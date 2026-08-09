// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kafkaconnectconnector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference interface {
	cdktn.ComplexObject
	CloudwatchLogs() KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogsOutputReference
	CloudwatchLogsInput() interface{}
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
	Firehose() KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehoseOutputReference
	FirehoseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3() KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3OutputReference
	S3Input() interface{}
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
	PutCloudwatchLogs(value *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs)
	PutFirehose(value *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehose)
	PutS3(value *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3)
	ResetCloudwatchLogs()
	ResetFirehose()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference
type jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) CloudwatchLogs() KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogsOutputReference {
	var returns KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) Firehose() KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehoseOutputReference {
	var returns KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehoseOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) S3() KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3OutputReference {
	var returns KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference {
	_init_.Initialize()

	if err := validateNewKafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kafkaconnectConnector.KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference_Override(k KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kafkaconnectConnector.KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) PutCloudwatchLogs(value *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryCloudwatchLogs) {
	if err := k.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) PutFirehose(value *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehose) {
	if err := k.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putFirehose",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) PutS3(value *KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3) {
	if err := k.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		k,
		"resetFirehose",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		k,
		"resetS3",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KafkaconnectConnectorLogDeliveryWorkerLogDeliveryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

