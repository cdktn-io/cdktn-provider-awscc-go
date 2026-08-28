// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActionsSns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#message_format IotTopicRule#message_format}.
	MessageFormat *string `field:"optional" json:"messageFormat" yaml:"messageFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#target_arn IotTopicRule#target_arn}.
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

