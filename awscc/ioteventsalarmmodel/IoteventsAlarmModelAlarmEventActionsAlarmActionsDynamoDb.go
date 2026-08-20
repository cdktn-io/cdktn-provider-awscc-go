// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDb struct {
	// The name of the hash key (also called the partition key).
	//
	// The ``hashKeyField`` value must match the partition key of the target DynamoDB table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#hash_key_field IoteventsAlarmModel#hash_key_field}
	HashKeyField *string `field:"optional" json:"hashKeyField" yaml:"hashKeyField"`
	// The data type for the hash key (also called the partition key).
	//
	// You can specify the following values:
	//   +  ``'STRING'`` - The hash key is a string.
	//   +  ``'NUMBER'`` - The hash key is a number.
	//
	//  If you don't specify ``hashKeyType``, the default value is ``'STRING'``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#hash_key_type IoteventsAlarmModel#hash_key_type}
	HashKeyType *string `field:"optional" json:"hashKeyType" yaml:"hashKeyType"`
	// The value of the hash key (also called the partition key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#hash_key_value IoteventsAlarmModel#hash_key_value}
	HashKeyValue *string `field:"optional" json:"hashKeyValue" yaml:"hashKeyValue"`
	// The type of operation to perform.
	//
	// You can specify the following values:
	//   +  ``'INSERT'`` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.
	//   +  ``'UPDATE'`` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	//   +  ``'DELETE'`` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.
	//
	//  If you don't specify this parameter, ITE triggers the ``'INSERT'`` operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#operation IoteventsAlarmModel#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Information needed to configure the payload.
	//
	// By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ``contentExpression``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDbPayload `field:"optional" json:"payload" yaml:"payload"`
	// The name of the DynamoDB column that receives the action payload.
	//
	// If you don't specify this parameter, the name of the DynamoDB column is ``payload``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#payload_field IoteventsAlarmModel#payload_field}
	PayloadField *string `field:"optional" json:"payloadField" yaml:"payloadField"`
	// The name of the range key (also called the sort key).
	//
	// The ``rangeKeyField`` value must match the sort key of the target DynamoDB table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#range_key_field IoteventsAlarmModel#range_key_field}
	RangeKeyField *string `field:"optional" json:"rangeKeyField" yaml:"rangeKeyField"`
	// The data type for the range key (also called the sort key), You can specify the following values:   +  ``'STRING'`` - The range key is a string.
	//
	// +  ``'NUMBER'`` - The range key is number.
	//
	//  If you don't specify ``rangeKeyField``, the default value is ``'STRING'``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#range_key_type IoteventsAlarmModel#range_key_type}
	RangeKeyType *string `field:"optional" json:"rangeKeyType" yaml:"rangeKeyType"`
	// The value of the range key (also called the sort key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#range_key_value IoteventsAlarmModel#range_key_value}
	RangeKeyValue *string `field:"optional" json:"rangeKeyValue" yaml:"rangeKeyValue"`
	// The name of the DynamoDB table. The ``tableName`` value must match the table name of the target DynamoDB table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_alarm_model#table_name IoteventsAlarmModel#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

