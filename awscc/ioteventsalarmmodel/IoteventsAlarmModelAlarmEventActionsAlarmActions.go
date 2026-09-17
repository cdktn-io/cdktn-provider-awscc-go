// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActions struct {
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.
	//  You must use expressions for all parameters in ``DynamoDBAction``. The expressions accept literals, operators, functions, references, and substitution templates.
	//   **Examples**
	//  +  For literal values, the expressions must contain single quotes. For example, the value for the ``hashKeyType`` parameter can be ``'STRING'``.
	//   +  For references, you must specify either variables or input values. For example, the value for the ``hashKeyField`` parameter can be ``$input.GreenhouseInput.name``.
	//   +  For a substitution template, you must use ``${}``, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//  In the following example, the value for the ``hashKeyValue`` parameter uses a substitution template.
	//   ``'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'``
	//   +  For a string concatenation, you must use ``+``. A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
	//  In the following example, the value for the ``tableName`` parameter uses a string concatenation.
	//   ``'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date``
	//
	//  For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.
	//  If the defined payload type is a string, ``DynamoDBAction`` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the ``payloadField`` parameter is ``<payload-field>_raw``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#dynamo_db IoteventsAlarmModel#dynamo_db}
	DynamoDb *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
	//  You must use expressions for all parameters in ``DynamoDBv2Action``. The expressions accept literals, operators, functions, references, and substitution templates.
	//   **Examples**
	//  +  For literal values, the expressions must contain single quotes. For example, the value for the ``tableName`` parameter can be ``'GreenhouseTemperatureTable'``.
	//   +  For references, you must specify either variables or input values. For example, the value for the ``tableName`` parameter can be ``$variable.ddbtableName``.
	//   +  For a substitution template, you must use ``${}``, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//  In the following example, the value for the ``contentExpression`` parameter in ``Payload`` uses a substitution template.
	//   ``'{\"sensorID\": \"${$input.GreenhouseInput.sensor_id}\", \"temperature\": \"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\"}'``
	//   +  For a string concatenation, you must use ``+``. A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.
	//  In the following example, the value for the ``tableName`` parameter uses a string concatenation.
	//   ``'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date``
	//
	//  For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.
	//  The value for the ``type`` parameter in ``Payload`` must be ``JSON``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#dynamo_d_bv_2 IoteventsAlarmModel#dynamo_d_bv_2}
	DynamoDBv2 *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2 `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#firehose IoteventsAlarmModel#firehose}
	Firehose *IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends an ITE input, passing in information about the detector model instance and the event that triggered the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#iot_events IoteventsAlarmModel#iot_events}
	IotEvents *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to a specified asset property in ITSW.
	//
	// You must use expressions for all parameters in ``IotSiteWiseAction``. The expressions accept literals, operators, functions, references, and substitutions templates.
	//   **Examples**
	//  +  For literal values, the expressions must contain single quotes. For example, the value for the ``propertyAlias`` parameter can be ``'/company/windfarm/3/turbine/7/temperature'``.
	//   +  For references, you must specify either variables or input values. For example, the value for the ``assetId`` parameter can be ``$input.TurbineInput.assetId1``.
	//   +  For a substitution template, you must use ``${}``, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
	//  In the following example, the value for the ``propertyAlias`` parameter uses a substitution template.
	//   ``'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'``
	//
	//  You must specify either ``propertyAlias`` or both ``assetId`` and ``propertyId`` to identify the target asset property in ITSW.
	//  For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#iot_site_wise IoteventsAlarmModel#iot_site_wise}
	IotSiteWise *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWise `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Information required to publish the MQTT message through the IoT message broker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#iot_topic_publish IoteventsAlarmModel#iot_topic_publish}
	IotTopicPublish *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublish `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#lambda IoteventsAlarmModel#lambda}
	Lambda *IoteventsAlarmModelAlarmEventActionsAlarmActionsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Information required to publish the Amazon SNS message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#sns IoteventsAlarmModel#sns}
	Sns *IoteventsAlarmModelAlarmEventActionsAlarmActionsSns `field:"optional" json:"sns" yaml:"sns"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_alarm_model#sqs IoteventsAlarmModel#sqs}
	Sqs *IoteventsAlarmModelAlarmEventActionsAlarmActionsSqs `field:"optional" json:"sqs" yaml:"sqs"`
}

