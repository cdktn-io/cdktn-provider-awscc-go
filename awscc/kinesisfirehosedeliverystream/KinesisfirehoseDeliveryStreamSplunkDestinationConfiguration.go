// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamSplunkDestinationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#buffering_hints KinesisfirehoseDeliveryStream#buffering_hints}.
	BufferingHints *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#cloudwatch_logging_options KinesisfirehoseDeliveryStream#cloudwatch_logging_options}.
	CloudwatchLoggingOptions *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#hec_acknowledgment_timeout_in_seconds KinesisfirehoseDeliveryStream#hec_acknowledgment_timeout_in_seconds}.
	HecAcknowledgmentTimeoutInSeconds *float64 `field:"optional" json:"hecAcknowledgmentTimeoutInSeconds" yaml:"hecAcknowledgmentTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#hec_endpoint KinesisfirehoseDeliveryStream#hec_endpoint}.
	HecEndpoint *string `field:"optional" json:"hecEndpoint" yaml:"hecEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#hec_endpoint_type KinesisfirehoseDeliveryStream#hec_endpoint_type}.
	HecEndpointType *string `field:"optional" json:"hecEndpointType" yaml:"hecEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#hec_token KinesisfirehoseDeliveryStream#hec_token}.
	HecToken *string `field:"optional" json:"hecToken" yaml:"hecToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#processing_configuration KinesisfirehoseDeliveryStream#processing_configuration}.
	ProcessingConfiguration *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#retry_options KinesisfirehoseDeliveryStream#retry_options}.
	RetryOptions *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#s3_backup_mode KinesisfirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#s3_configuration KinesisfirehoseDeliveryStream#s3_configuration}.
	S3Configuration *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#secrets_manager_configuration KinesisfirehoseDeliveryStream#secrets_manager_configuration}.
	SecretsManagerConfiguration *KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationSecretsManagerConfiguration `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
}

