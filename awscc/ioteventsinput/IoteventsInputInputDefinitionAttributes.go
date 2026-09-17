// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsinput


type IoteventsInputInputDefinitionAttributes struct {
	// An expression that specifies an attribute-value pair in a JSON structure.
	//
	// Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to ITE (``BatchPutMessage``). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the ``condition`` expressions used by detectors.
	//  Syntax: ``<field-name>.<field-name>...``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_input#json_path IoteventsInput#json_path}
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
}

