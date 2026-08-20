// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsLambda struct {
	// The ARN of the Lambda function that is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#function_arn IoteventsAlarmModel#function_arn}
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
	// You can configure the action payload when you send a message to a Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsLambdaPayload `field:"optional" json:"payload" yaml:"payload"`
}

