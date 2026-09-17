// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelTopicConfigurationListStruct struct {
	// Record converter configuration for a topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#record_converter MskChannel#record_converter}
	RecordConverter *MskChannelTopicConfigurationListRecordConverter `field:"required" json:"recordConverter" yaml:"recordConverter"`
	// The Amazon Resource Name (ARN) that uniquely identifies the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#topic_arn MskChannel#topic_arn}
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Record schema configuration for a topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_channel#record_schema MskChannel#record_schema}
	RecordSchema *MskChannelTopicConfigurationListRecordSchema `field:"optional" json:"recordSchema" yaml:"recordSchema"`
}

