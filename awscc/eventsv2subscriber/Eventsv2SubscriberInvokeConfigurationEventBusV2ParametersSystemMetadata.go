// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationEventBusV2ParametersSystemMetadata struct {
	// The deduplication ID for FIFO deduplication on the downstream event bus. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#deduplication_id Eventsv2Subscriber#deduplication_id}
	DeduplicationId *string `field:"optional" json:"deduplicationId" yaml:"deduplicationId"`
	// The event group ID for FIFO ordering on the downstream event bus.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#event_group_id Eventsv2Subscriber#event_group_id}
	EventGroupId *string `field:"optional" json:"eventGroupId" yaml:"eventGroupId"`
}

