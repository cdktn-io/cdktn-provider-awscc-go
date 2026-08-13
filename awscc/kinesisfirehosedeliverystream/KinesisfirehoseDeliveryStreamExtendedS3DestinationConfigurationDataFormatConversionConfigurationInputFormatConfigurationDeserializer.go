// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#hive_json_ser_de KinesisfirehoseDeliveryStream#hive_json_ser_de}.
	HiveJsonSerDe *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerHiveJsonSerDe `field:"optional" json:"hiveJsonSerDe" yaml:"hiveJsonSerDe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kinesisfirehose_delivery_stream#open_x_json_ser_de KinesisfirehoseDeliveryStream#open_x_json_ser_de}.
	OpenXJsonSerDe *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationInputFormatConfigurationDeserializerOpenXJsonSerDe `field:"optional" json:"openXJsonSerDe" yaml:"openXJsonSerDe"`
}

