// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfigurationSnsParameters struct {
	// Custom message attributes to attach to each message; Amazon SNS subscription filter policies can match on them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#message_attributes Eventsv2Subscriber#message_attributes}
	MessageAttributes interface{} `field:"optional" json:"messageAttributes" yaml:"messageAttributes"`
	// The message deduplication ID to use when the target is a FIFO topic.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#message_deduplication_id Eventsv2Subscriber#message_deduplication_id}
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The message group ID to use when the target is a FIFO topic.
	//
	// Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#message_group_id Eventsv2Subscriber#message_group_id}
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
	// Set to json to send a different message per delivery protocol. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#message_structure Eventsv2Subscriber#message_structure}
	MessageStructure *string `field:"optional" json:"messageStructure" yaml:"messageStructure"`
	// The subject line to use for email-protocol subscriptions. Accepts a literal value or a JSONata expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/eventsv2_subscriber#subject Eventsv2Subscriber#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

