// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationSqsParameters struct {
	// The delay in seconds for the message, written as a string. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#delay_seconds Eventsv2Subscriber#delay_seconds}
	DelaySeconds *string `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// Custom message attributes to attach to each message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#message_attributes Eventsv2Subscriber#message_attributes}
	MessageAttributes interface{} `field:"optional" json:"messageAttributes" yaml:"messageAttributes"`
	// The message deduplication ID to use when the target is a FIFO queue.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#message_deduplication_id Eventsv2Subscriber#message_deduplication_id}
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The message group ID to use when the target is a FIFO queue.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#message_group_id Eventsv2Subscriber#message_group_id}
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
	// Message system attributes to attach to each message, such as AWSTraceHeader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#message_system_attributes Eventsv2Subscriber#message_system_attributes}
	MessageSystemAttributes interface{} `field:"optional" json:"messageSystemAttributes" yaml:"messageSystemAttributes"`
}

