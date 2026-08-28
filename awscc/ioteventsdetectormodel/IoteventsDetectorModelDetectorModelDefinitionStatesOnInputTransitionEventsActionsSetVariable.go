// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariable struct {
	// The new value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#value IoteventsDetectorModel#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#variable_name IoteventsDetectorModel#variable_name}
	VariableName *string `field:"optional" json:"variableName" yaml:"variableName"`
}

