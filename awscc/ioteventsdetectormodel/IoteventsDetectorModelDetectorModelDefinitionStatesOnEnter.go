// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnter struct {
	// Specifies the actions that are performed when the state is entered and the ``condition`` is ``TRUE``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotevents_detector_model#events IoteventsDetectorModel#events}
	Events interface{} `field:"optional" json:"events" yaml:"events"`
}

