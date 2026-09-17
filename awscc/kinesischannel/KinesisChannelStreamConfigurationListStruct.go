// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelStreamConfigurationListStruct struct {
	// The configuration that describes how records on the source stream are encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#record_configuration KinesisChannel#record_configuration}
	RecordConfiguration *KinesisChannelStreamConfigurationListRecordConfiguration `field:"required" json:"recordConfiguration" yaml:"recordConfiguration"`
	// The Amazon resource name (ARN) of the Kinesis data stream that the channel reads from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesis_channel#stream_arn KinesisChannel#stream_arn}
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

