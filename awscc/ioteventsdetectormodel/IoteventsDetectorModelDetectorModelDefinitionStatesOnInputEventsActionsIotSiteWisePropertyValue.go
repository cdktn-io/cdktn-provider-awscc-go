// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWisePropertyValue struct {
	// The quality of the asset property value. The value must be ``'GOOD'``, ``'BAD'``, or ``'UNCERTAIN'``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#quality IoteventsDetectorModel#quality}
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// The timestamp associated with the asset property value. The default is the current event time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#timestamp IoteventsDetectorModel#timestamp}
	Timestamp *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWisePropertyValueTimestamp `field:"optional" json:"timestamp" yaml:"timestamp"`
	// The value to send to an asset property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#value IoteventsDetectorModel#value}
	Value *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWisePropertyValueValue `field:"optional" json:"value" yaml:"value"`
}

