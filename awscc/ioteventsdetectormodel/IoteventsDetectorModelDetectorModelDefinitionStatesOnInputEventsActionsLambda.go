// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambda struct {
	// The ARN of the Lambda function that is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#function_arn IoteventsDetectorModel#function_arn}
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
	// You can configure the action payload when you send a message to a Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambdaPayload `field:"optional" json:"payload" yaml:"payload"`
}

