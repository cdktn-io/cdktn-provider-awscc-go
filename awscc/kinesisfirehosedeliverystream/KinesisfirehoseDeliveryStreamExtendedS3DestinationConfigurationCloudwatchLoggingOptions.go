// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#enabled KinesisfirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#log_group_name KinesisfirehoseDeliveryStream#log_group_name}.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisfirehose_delivery_stream#log_stream_name KinesisfirehoseDeliveryStream#log_stream_name}.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

