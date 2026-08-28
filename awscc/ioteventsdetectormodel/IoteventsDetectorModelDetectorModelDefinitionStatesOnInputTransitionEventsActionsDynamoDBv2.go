// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2 struct {
	// Information needed to configure the payload.
	//
	// By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ``contentExpression``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2Payload `field:"optional" json:"payload" yaml:"payload"`
	// The name of the DynamoDB table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#table_name IoteventsDetectorModel#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

