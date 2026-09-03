// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStates struct {
	// The name of the state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#state_name IoteventsDetectorModel#state_name}
	StateName *string `field:"required" json:"stateName" yaml:"stateName"`
	// When entering this state, perform these ``actions`` if the ``condition`` is TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#on_enter IoteventsDetectorModel#on_enter}
	OnEnter *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnter `field:"optional" json:"onEnter" yaml:"onEnter"`
	// When exiting this state, perform these ``actions`` if the specified ``condition`` is ``TRUE``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#on_exit IoteventsDetectorModel#on_exit}
	OnExit *IoteventsDetectorModelDetectorModelDefinitionStatesOnExit `field:"optional" json:"onExit" yaml:"onExit"`
	// When an input is received and the ``condition`` is TRUE, perform the specified ``actions``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#on_input IoteventsDetectorModel#on_input}
	OnInput *IoteventsDetectorModelDetectorModelDefinitionStatesOnInput `field:"optional" json:"onInput" yaml:"onInput"`
}

