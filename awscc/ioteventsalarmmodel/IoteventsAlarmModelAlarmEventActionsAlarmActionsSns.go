// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsSns struct {
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsSnsPayload `field:"optional" json:"payload" yaml:"payload"`
	// The ARN of the Amazon SNS target where the message is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#target_arn IoteventsAlarmModel#target_arn}
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

