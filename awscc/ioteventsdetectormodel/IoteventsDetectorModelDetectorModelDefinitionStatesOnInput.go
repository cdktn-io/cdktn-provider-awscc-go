// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInput struct {
	// Specifies the actions performed when the ``condition`` evaluates to TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#events IoteventsDetectorModel#events}
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// Specifies the actions performed, and the next state entered, when a ``condition`` evaluates to TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#transition_events IoteventsDetectorModel#transition_events}
	TransitionEvents interface{} `field:"optional" json:"transitionEvents" yaml:"transitionEvents"`
}

