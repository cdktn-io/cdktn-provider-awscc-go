// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference interface {
	cdktn.ComplexObject
	BucketArn() *string
	SetBucketArn(val *string)
	BucketArnInput() *string
	BufferingHints() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHintsOutputReference
	BufferingHintsInput() interface{}
	CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference
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
	CompressionFormat() *string
	SetCompressionFormat(val *string)
	CompressionFormatInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomTimeZone() *string
	SetCustomTimeZone(val *string)
	CustomTimeZoneInput() *string
	DataFormatConversionConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputReference
	DataFormatConversionConfigurationInput() interface{}
	DynamicPartitioningConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationOutputReference
	DynamicPartitioningConfigurationInput() interface{}
	EncryptionConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() interface{}
	ErrorOutputPrefix() *string
	SetErrorOutputPrefix(val *string)
	ErrorOutputPrefixInput() *string
	FileExtension() *string
	SetFileExtension(val *string)
	FileExtensionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	ProcessingConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationOutputReference
	ProcessingConfigurationInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BackupConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationOutputReference
	S3BackupConfigurationInput() interface{}
	S3BackupMode() *string
	SetS3BackupMode(val *string)
	S3BackupModeInput() *string
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
	PutBufferingHints(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHints)
	PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptions)
	PutDataFormatConversionConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfiguration)
	PutDynamicPartitioningConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfiguration)
	PutEncryptionConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfiguration)
	PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfiguration)
	PutS3BackupConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfiguration)
	ResetBucketArn()
	ResetBufferingHints()
	ResetCloudwatchLoggingOptions()
	ResetCompressionFormat()
	ResetCustomTimeZone()
	ResetDataFormatConversionConfiguration()
	ResetDynamicPartitioningConfiguration()
	ResetEncryptionConfiguration()
	ResetErrorOutputPrefix()
	ResetFileExtension()
	ResetPrefix()
	ResetProcessingConfiguration()
	ResetRoleArn()
	ResetS3BackupConfiguration()
	ResetS3BackupMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference
type jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) BufferingHints() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHintsOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHintsOutputReference
	_jsii_.Get(
		j,
		"bufferingHints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) BufferingHintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferingHintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CloudwatchLoggingOptions() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CloudwatchLoggingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CustomTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) CustomTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) DataFormatConversionConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataFormatConversionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) DataFormatConversionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataFormatConversionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) DynamicPartitioningConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationOutputReference
	_jsii_.Get(
		j,
		"dynamicPartitioningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) DynamicPartitioningConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPartitioningConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) EncryptionConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) EncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ErrorOutputPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) FileExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ProcessingConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfigurationOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ProcessingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) S3BackupConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3BackupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) S3BackupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BackupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference_Override(k KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetBucketArn(val *string) {
	if err := j.validateSetBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketArn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetCustomTimeZone(val *string) {
	if err := j.validateSetCustomTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTimeZone",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetErrorOutputPrefix(val *string) {
	if err := j.validateSetErrorOutputPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOutputPrefix",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetFileExtension(val *string) {
	if err := j.validateSetFileExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtension",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutBufferingHints(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHints) {
	if err := k.validatePutBufferingHintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putBufferingHints",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutCloudwatchLoggingOptions(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptions) {
	if err := k.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutDataFormatConversionConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfiguration) {
	if err := k.validatePutDataFormatConversionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDataFormatConversionConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutDynamicPartitioningConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfiguration) {
	if err := k.validatePutDynamicPartitioningConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDynamicPartitioningConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutEncryptionConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfiguration) {
	if err := k.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutProcessingConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfiguration) {
	if err := k.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) PutS3BackupConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfiguration) {
	if err := k.validatePutS3BackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3BackupConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetBucketArn() {
	_jsii_.InvokeVoid(
		k,
		"resetBucketArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetBufferingHints() {
	_jsii_.InvokeVoid(
		k,
		"resetBufferingHints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		k,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetCustomTimeZone() {
	_jsii_.InvokeVoid(
		k,
		"resetCustomTimeZone",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetDataFormatConversionConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDataFormatConversionConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetDynamicPartitioningConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDynamicPartitioningConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetErrorOutputPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetErrorOutputPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetFileExtension() {
	_jsii_.InvokeVoid(
		k,
		"resetFileExtension",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		k,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetS3BackupConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		k,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

