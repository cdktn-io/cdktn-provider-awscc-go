// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#alarm_name IotTopicRule#alarm_name}.
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#state_reason IotTopicRule#state_reason}.
	StateReason *string `field:"optional" json:"stateReason" yaml:"stateReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#state_value IotTopicRule#state_value}.
	StateValue *string `field:"optional" json:"stateValue" yaml:"stateValue"`
}

