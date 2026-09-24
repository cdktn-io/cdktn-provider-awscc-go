// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSns struct {
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsSnsPayload `field:"optional" json:"payload" yaml:"payload"`
	// The ARN of the Amazon SNS target where the message is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotevents_detector_model#target_arn IoteventsDetectorModel#target_arn}
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

