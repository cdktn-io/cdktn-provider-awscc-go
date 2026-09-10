// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesisfirehose_delivery_stream#enabled KinesisfirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesisfirehose_delivery_stream#input_format_configuration KinesisfirehoseDeliveryStream#input_format_configuration}.
	InputFormatConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfiguration `field:"optional" json:"inputFormatConfiguration" yaml:"inputFormatConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesisfirehose_delivery_stream#output_format_configuration KinesisfirehoseDeliveryStream#output_format_configuration}.
	OutputFormatConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfiguration `field:"optional" json:"outputFormatConfiguration" yaml:"outputFormatConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesisfirehose_delivery_stream#schema_configuration KinesisfirehoseDeliveryStream#schema_configuration}.
	SchemaConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationSchemaConfiguration `field:"optional" json:"schemaConfiguration" yaml:"schemaConfiguration"`
}

