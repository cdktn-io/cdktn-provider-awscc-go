// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2eventsource


type Eventsv2EventSourceConfigurationPartnerEventsConfigurationOnFailureConfiguration struct {
	// The ARN of the Amazon SQS standard queue that receives events that could not be forwarded.
	//
	// FIFO queues are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_event_source#arn Eventsv2EventSource#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

