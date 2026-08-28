// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehose struct {
	// The name of the Kinesis Data Firehose delivery stream where the data is written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#delivery_stream_name IoteventsDetectorModel#delivery_stream_name}
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// You can configure the action payload when you send a message to an Amazon Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehosePayload `field:"optional" json:"payload" yaml:"payload"`
	// A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotevents_detector_model#separator IoteventsDetectorModel#separator}
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

