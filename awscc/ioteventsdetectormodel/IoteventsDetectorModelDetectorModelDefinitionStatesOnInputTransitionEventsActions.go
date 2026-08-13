// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActions struct {
	// Information needed to clear the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#clear_timer IoteventsDetectorModel#clear_timer}
	ClearTimer *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsClearTimer `field:"optional" json:"clearTimer" yaml:"clearTimer"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#dynamo_db IoteventsDetectorModel#dynamo_db}
	DynamoDb *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#dynamo_d_bv_2 IoteventsDetectorModel#dynamo_d_bv_2}
	DynamoDBv2 *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2 `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#firehose IoteventsDetectorModel#firehose}
	Firehose *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends ITE input, which passes information about the detector model instance and the event that triggered the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#iot_events IoteventsDetectorModel#iot_events}
	IotEvents *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to an asset property in ITSW .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#iot_site_wise IoteventsDetectorModel#iot_site_wise}
	IotSiteWise *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWise `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Publishes an MQTT message with the given topic to the IoT message broker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#iot_topic_publish IoteventsDetectorModel#iot_topic_publish}
	IotTopicPublish *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublish `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#lambda IoteventsDetectorModel#lambda}
	Lambda *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Information needed to reset the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#reset_timer IoteventsDetectorModel#reset_timer}
	ResetTimer *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsResetTimer `field:"optional" json:"resetTimer" yaml:"resetTimer"`
	// Information needed to set the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#set_timer IoteventsDetectorModel#set_timer}
	SetTimer *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetTimer `field:"optional" json:"setTimer" yaml:"setTimer"`
	// Sets a variable to a specified value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#set_variable IoteventsDetectorModel#set_variable}
	SetVariable *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariable `field:"optional" json:"setVariable" yaml:"setVariable"`
	// Sends an Amazon SNS message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#sns IoteventsDetectorModel#sns}
	Sns *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSns `field:"optional" json:"sns" yaml:"sns"`
	// Sends an Amazon SNS message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotevents_detector_model#sqs IoteventsDetectorModel#sqs}
	Sqs *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqs `field:"optional" json:"sqs" yaml:"sqs"`
}

