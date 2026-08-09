// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference interface {
	cdktn.ComplexObject
	BufferingHints() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHintsOutputReference
	BufferingHintsInput() interface{}
	CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptionsOutputReference
	CloudwatchLoggingOptionsInput() interface{}
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
	HecAcknowledgmentTimeoutInSeconds() *float64
	SetHecAcknowledgmentTimeoutInSeconds(val *float64)
	HecAcknowledgmentTimeoutInSecondsInput() *float64
	HecEndpoint() *string
	SetHecEndpoint(val *string)
	HecEndpointInput() *string
	HecEndpointType() *string
	SetHecEndpointType(val *string)
	HecEndpointTypeInput() *string
	HecToken() *string
	SetHecToken(val *string)
	HecTokenInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProcessingConfiguration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() interface{}
	RetryOptions() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptionsOutputReference
	RetryOptionsInput() interface{}
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
	S3Configuration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() interface{}
	SecretsManagerConfiguration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationSecretsManagerConfigurationOutputReference
	SecretsManagerConfigurationInput() interface{}
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
	PutBufferingHints(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHints)
	PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptions)
	PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfiguration)
	PutRetryOptions(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptions)
	PutS3Configuration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3Configuration)
	PutSecretsManagerConfiguration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationSecretsManagerConfiguration)
	ResetBufferingHints()
	ResetCloudwatchLoggingOptions()
	ResetHecAcknowledgmentTimeoutInSeconds()
	ResetHecEndpoint()
	ResetHecEndpointType()
	ResetHecToken()
	ResetProcessingConfiguration()
	ResetRetryOptions()
	ResetS3BackupMode()
	ResetS3Configuration()
	ResetSecretsManagerConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) BufferingHints() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHintsOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHintsOutputReference
	_jsii_.Get(
		j,
		"bufferingHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) BufferingHintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferingHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) CloudwatchLoggingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecAcknowledgmentTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hecAcknowledgmentTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecAcknowledgmentTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hecAcknowledgmentTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecEndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) HecTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ProcessingConfiguration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ProcessingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) RetryOptions() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptionsOutputReference
	_jsii_.Get(
		j,
		"retryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) RetryOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) S3Configuration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) SecretsManagerConfiguration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationSecretsManagerConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationSecretsManagerConfigurationOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) SecretsManagerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetHecAcknowledgmentTimeoutInSeconds(val *float64) {
	if err := j.validateSetHecAcknowledgmentTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecAcknowledgmentTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetHecEndpoint(val *string) {
	if err := j.validateSetHecEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecEndpoint",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetHecEndpointType(val *string) {
	if err := j.validateSetHecEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecEndpointType",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetHecToken(val *string) {
	if err := j.validateSetHecTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecToken",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) PutBufferingHints(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHints) {
	if err := k.validatePutBufferingHintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putBufferingHints",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) PutRetryOptions(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptions) {
	if err := k.validatePutRetryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRetryOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) PutS3Configuration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3Configuration) {
	if err := k.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) PutSecretsManagerConfiguration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationSecretsManagerConfiguration) {
	if err := k.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetBufferingHints() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingHints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetHecAcknowledgmentTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetHecAcknowledgmentTimeoutInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetHecEndpoint() {
	_jsii_.InvokeVoid(
		k,
		"resetHecEndpoint",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetHecEndpointType() {
	_jsii_.InvokeVoid(
		k,
		"resetHecEndpointType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetHecToken() {
	_jsii_.InvokeVoid(
		k,
		"resetHecToken",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetRetryOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

