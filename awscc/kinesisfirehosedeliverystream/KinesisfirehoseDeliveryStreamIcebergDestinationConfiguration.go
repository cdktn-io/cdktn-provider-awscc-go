// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamIcebergDestinationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#append_only KinesisfirehoseDeliveryStream#append_only}.
	AppendOnly interface{} `field:"optional" json:"appendOnly" yaml:"appendOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#buffering_hints KinesisfirehoseDeliveryStream#buffering_hints}.
	BufferingHints *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#catalog_configuration KinesisfirehoseDeliveryStream#catalog_configuration}.
	CatalogConfiguration *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationCatalogConfiguration `field:"optional" json:"catalogConfiguration" yaml:"catalogConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#cloudwatch_logging_options KinesisfirehoseDeliveryStream#cloudwatch_logging_options}.
	CloudwatchLoggingOptions *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#destination_table_configuration_list KinesisfirehoseDeliveryStream#destination_table_configuration_list}.
	DestinationTableConfigurationList interface{} `field:"optional" json:"destinationTableConfigurationList" yaml:"destinationTableConfigurationList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#processing_configuration KinesisfirehoseDeliveryStream#processing_configuration}.
	ProcessingConfiguration *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#retry_options KinesisfirehoseDeliveryStream#retry_options}.
	RetryOptions *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#role_arn KinesisfirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#s_3_backup_mode KinesisfirehoseDeliveryStream#s_3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#s3_configuration KinesisfirehoseDeliveryStream#s3_configuration}.
	S3Configuration *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#schema_evolution_configuration KinesisfirehoseDeliveryStream#schema_evolution_configuration}.
	SchemaEvolutionConfiguration *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationSchemaEvolutionConfiguration `field:"optional" json:"schemaEvolutionConfiguration" yaml:"schemaEvolutionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#table_creation_configuration KinesisfirehoseDeliveryStream#table_creation_configuration}.
	TableCreationConfiguration *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationTableCreationConfiguration `field:"optional" json:"tableCreationConfiguration" yaml:"tableCreationConfiguration"`
}

