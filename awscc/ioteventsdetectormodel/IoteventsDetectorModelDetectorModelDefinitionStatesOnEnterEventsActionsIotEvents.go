// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEvents struct {
	// The name of the ITE input where the data is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#input_name IoteventsDetectorModel#input_name}
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// You can configure the action payload when you send a message to an ITE input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotEventsPayload `field:"optional" json:"payload" yaml:"payload"`
}

