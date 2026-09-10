// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponse struct {
	// Indicates whether the user can interrupt a speech prompt from the bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#allow_interrupt LexBot#allow_interrupt}
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// How often a message should be sent to the user in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#frequency_in_seconds LexBot#frequency_in_seconds}
	FrequencyInSeconds *float64 `field:"optional" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// One to 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#message_groups_list LexBot#message_groups_list}
	MessageGroupsList interface{} `field:"optional" json:"messageGroupsList" yaml:"messageGroupsList"`
	// If Amazon Lex waits longer than this length of time in seconds for a response, it will stop sending messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/lex_bot#timeout_in_seconds LexBot#timeout_in_seconds}
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

