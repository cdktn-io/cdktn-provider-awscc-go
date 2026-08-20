// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotSiteWisePropertyValueValue struct {
	// The asset property value is a Boolean value that must be ``'TRUE'`` or ``'FALSE'``.
	//
	// You must use an expression, and the evaluated result should be a Boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#boolean_value IoteventsDetectorModel#boolean_value}
	BooleanValue *string `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The asset property value is a double.
	//
	// You must use an expression, and the evaluated result should be a double.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#double_value IoteventsDetectorModel#double_value}
	DoubleValue *string `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The asset property value is an integer.
	//
	// You must use an expression, and the evaluated result should be an integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#integer_value IoteventsDetectorModel#integer_value}
	IntegerValue *string `field:"optional" json:"integerValue" yaml:"integerValue"`
	// The asset property value is a string.
	//
	// You must use an expression, and the evaluated result should be a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotevents_detector_model#string_value IoteventsDetectorModel#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

