// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsIotSiteWisePropertyValueTimestamp struct {
	// The nanosecond offset converted from ``timeInSeconds``. The valid range is between 0-999999999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_detector_model#offset_in_nanos IoteventsDetectorModel#offset_in_nanos}
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
	// The timestamp, in seconds, in the Unix epoch format. The valid range is between 1-31556889864403199.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotevents_detector_model#time_in_seconds IoteventsDetectorModel#time_in_seconds}
	TimeInSeconds *string `field:"optional" json:"timeInSeconds" yaml:"timeInSeconds"`
}

