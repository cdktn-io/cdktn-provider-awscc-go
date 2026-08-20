// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsDynamoDBv2Payload struct {
	// The content of the payload.
	//
	// You can use a string expression that includes quoted strings (``'<string>'``), variables (``$variable.<variable-name>``), input values (``$input.<input-name>.<path-to-datum>``), string concatenations, and quoted strings that contain ``${}`` as the content. The recommended maximum size of a content expression is 1 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#content_expression IoteventsDetectorModel#content_expression}
	ContentExpression *string `field:"optional" json:"contentExpression" yaml:"contentExpression"`
	// The value of the payload type can be either ``STRING`` or ``JSON``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#type IoteventsDetectorModel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

