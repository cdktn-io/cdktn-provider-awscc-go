// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2 struct {
	// Information needed to configure the payload.
	//
	// By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ``contentExpression``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2Payload `field:"optional" json:"payload" yaml:"payload"`
	// The name of the DynamoDB table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#table_name IoteventsAlarmModel#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

