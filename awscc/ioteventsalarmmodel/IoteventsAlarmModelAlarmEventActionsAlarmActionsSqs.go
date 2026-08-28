// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsSqs struct {
	// You can configure the action payload when you send a message to an Amazon SQS queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsSqsPayload `field:"optional" json:"payload" yaml:"payload"`
	// The URL of the SQS queue where the data is written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#queue_url IoteventsAlarmModel#queue_url}
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
	// Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue.
	//
	// Otherwise, set this to FALSE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_alarm_model#use_base_64 IoteventsAlarmModel#use_base_64}
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

