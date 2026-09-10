// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnExitEvents struct {
	// The actions to be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#actions IoteventsDetectorModel#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Optional.
	//
	// The Boolean expression that, when TRUE, causes the ``actions`` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#condition IoteventsDetectorModel#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The name of the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#event_name IoteventsDetectorModel#event_name}
	EventName *string `field:"optional" json:"eventName" yaml:"eventName"`
}

