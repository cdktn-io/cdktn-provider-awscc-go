// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsinput


type IoteventsInputInputDefinition struct {
	// The attributes from the JSON payload that are made available by the input.
	//
	// Inputs are derived from messages sent to the ITE system using ``BatchPutMessage``. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the ``condition`` expressions used by detectors that monitor this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_input#attributes IoteventsInput#attributes}
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
}

