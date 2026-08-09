// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference interface {
	cdktn.ComplexObject
	BufferingHints() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHintsOutputReference
	BufferingHintsInput() interface{}
	CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptionsOutputReference
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
	EndpointConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfigurationOutputReference
	EndpointConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProcessingConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() interface{}
	RequestConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationOutputReference
	RequestConfigurationInput() interface{}
	RetryOptions() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptionsOutputReference
	RetryOptionsInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
	S3Configuration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() interface{}
	SecretsManagerConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationSecretsManagerConfigurationOutputReference
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
	PutBufferingHints(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHints)
	PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptions)
	PutEndpointConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfiguration)
	PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfiguration)
	PutRequestConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfiguration)
	PutRetryOptions(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptions)
	PutS3Configuration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3Configuration)
	PutSecretsManagerConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationSecretsManagerConfiguration)
	ResetBufferingHints()
	ResetCloudwatchLoggingOptions()
	ResetEndpointConfiguration()
	ResetProcessingConfiguration()
	ResetRequestConfiguration()
	ResetRetryOptions()
	ResetRoleArn()
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

// The jsii proxy struct for KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) BufferingHints() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHintsOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHintsOutputReference
	_jsii_.Get(
		j,
		"bufferingHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) BufferingHintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferingHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) CloudwatchLoggingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) EndpointConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfigurationOutputReference
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) EndpointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ProcessingConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ProcessingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) RequestConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfigurationOutputReference
	_jsii_.Get(
		j,
		"requestConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) RequestConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) RetryOptions() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptionsOutputReference
	_jsii_.Get(
		j,
		"retryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) RetryOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) S3Configuration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) SecretsManagerConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationSecretsManagerConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationSecretsManagerConfigurationOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) SecretsManagerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutBufferingHints(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationBufferingHints) {
	if err := k.validatePutBufferingHintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putBufferingHints",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutEndpointConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationEndpointConfiguration) {
	if err := k.validatePutEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutRequestConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRequestConfiguration) {
	if err := k.validatePutRequestConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRequestConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutRetryOptions(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationRetryOptions) {
	if err := k.validatePutRetryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRetryOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutS3Configuration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationS3Configuration) {
	if err := k.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) PutSecretsManagerConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationSecretsManagerConfiguration) {
	if err := k.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetBufferingHints() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingHints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetRequestConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetRequestConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetRetryOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetRetryOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		k,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

