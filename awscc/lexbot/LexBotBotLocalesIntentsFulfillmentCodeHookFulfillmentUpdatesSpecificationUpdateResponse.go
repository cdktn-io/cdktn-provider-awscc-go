// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationUpdateResponse struct {
	// Determines whether the user can interrupt an update message while it is playing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#allow_interrupt LexBot#allow_interrupt}
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// The frequency that a message is sent to the user.
	//
	// When the period ends, Amazon Lex chooses a message from the message groups and plays it to the user. If the fulfillment Lambda returns before the first period ends, an update message is not played to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#frequency_in_seconds LexBot#frequency_in_seconds}
	FrequencyInSeconds *float64 `field:"optional" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// One to 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#message_groups LexBot#message_groups}
	MessageGroups interface{} `field:"optional" json:"messageGroups" yaml:"messageGroups"`
}

