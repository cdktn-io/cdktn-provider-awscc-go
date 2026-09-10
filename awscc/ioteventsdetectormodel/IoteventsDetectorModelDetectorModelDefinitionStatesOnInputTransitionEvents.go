// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEvents struct {
	// The actions to be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#actions IoteventsDetectorModel#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Required. A Boolean expression that when TRUE causes the actions to be performed and the ``nextState`` to be entered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#condition IoteventsDetectorModel#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The name of the transition event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#event_name IoteventsDetectorModel#event_name}
	EventName *string `field:"optional" json:"eventName" yaml:"eventName"`
	// The next state to enter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotevents_detector_model#next_state IoteventsDetectorModel#next_state}
	NextState *string `field:"optional" json:"nextState" yaml:"nextState"`
}

