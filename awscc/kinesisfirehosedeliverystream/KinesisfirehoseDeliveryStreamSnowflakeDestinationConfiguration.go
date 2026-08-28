// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamSnowflakeDestinationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#account_url KinesisfirehoseDeliveryStream#account_url}.
	AccountUrl *string `field:"optional" json:"accountUrl" yaml:"accountUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#buffering_hints KinesisfirehoseDeliveryStream#buffering_hints}.
	BufferingHints *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#cloudwatch_logging_options KinesisfirehoseDeliveryStream#cloudwatch_logging_options}.
	CloudwatchLoggingOptions *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#content_column_name KinesisfirehoseDeliveryStream#content_column_name}.
	ContentColumnName *string `field:"optional" json:"contentColumnName" yaml:"contentColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#database KinesisfirehoseDeliveryStream#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#data_loading_option KinesisfirehoseDeliveryStream#data_loading_option}.
	DataLoadingOption *string `field:"optional" json:"dataLoadingOption" yaml:"dataLoadingOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#key_passphrase KinesisfirehoseDeliveryStream#key_passphrase}.
	KeyPassphrase *string `field:"optional" json:"keyPassphrase" yaml:"keyPassphrase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#meta_data_column_name KinesisfirehoseDeliveryStream#meta_data_column_name}.
	MetaDataColumnName *string `field:"optional" json:"metaDataColumnName" yaml:"metaDataColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#private_key KinesisfirehoseDeliveryStream#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#processing_configuration KinesisfirehoseDeliveryStream#processing_configuration}.
	ProcessingConfiguration *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#retry_options KinesisfirehoseDeliveryStream#retry_options}.
	RetryOptions *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#role_arn KinesisfirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#s3_backup_mode KinesisfirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#s3_configuration KinesisfirehoseDeliveryStream#s3_configuration}.
	S3Configuration *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#schema KinesisfirehoseDeliveryStream#schema}.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#secrets_manager_configuration KinesisfirehoseDeliveryStream#secrets_manager_configuration}.
	SecretsManagerConfiguration *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSecretsManagerConfiguration `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#snowflake_role_configuration KinesisfirehoseDeliveryStream#snowflake_role_configuration}.
	SnowflakeRoleConfiguration *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeRoleConfiguration `field:"optional" json:"snowflakeRoleConfiguration" yaml:"snowflakeRoleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#snowflake_vpc_configuration KinesisfirehoseDeliveryStream#snowflake_vpc_configuration}.
	SnowflakeVpcConfiguration *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationSnowflakeVpcConfiguration `field:"optional" json:"snowflakeVpcConfiguration" yaml:"snowflakeVpcConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#table KinesisfirehoseDeliveryStream#table}.
	Table *string `field:"optional" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kinesisfirehose_delivery_stream#user KinesisfirehoseDeliveryStream#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

