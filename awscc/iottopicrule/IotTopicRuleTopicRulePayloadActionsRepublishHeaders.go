// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActionsRepublishHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule#content_type IotTopicRule#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule#correlation_data IotTopicRule#correlation_data}.
	CorrelationData *string `field:"optional" json:"correlationData" yaml:"correlationData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule#message_expiry IotTopicRule#message_expiry}.
	MessageExpiry *string `field:"optional" json:"messageExpiry" yaml:"messageExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule#payload_format_indicator IotTopicRule#payload_format_indicator}.
	PayloadFormatIndicator *string `field:"optional" json:"payloadFormatIndicator" yaml:"payloadFormatIndicator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule#response_topic IotTopicRule#response_topic}.
	ResponseTopic *string `field:"optional" json:"responseTopic" yaml:"responseTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule#user_properties IotTopicRule#user_properties}.
	UserProperties interface{} `field:"optional" json:"userProperties" yaml:"userProperties"`
}

