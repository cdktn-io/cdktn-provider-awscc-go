// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSetting struct {
	// A list of default values for a slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#default_value_specification LexBot#default_value_specification}
	DefaultValueSpecification *LexBotBotLocalesIntentsSlotsValueElicitationSettingDefaultValueSpecification `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// The prompt that Amazon Lex uses to elicit the slot value from the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#prompt_specification LexBot#prompt_specification}
	PromptSpecification *LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecification `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// If you know a specific pattern that users might respond to an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#sample_utterances LexBot#sample_utterances}
	SampleUtterances interface{} `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Specifies whether the slot is required or optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#slot_constraint LexBot#slot_constraint}
	SlotConstraint *string `field:"optional" json:"slotConstraint" yaml:"slotConstraint"`
	// Specifies the prompts that Amazon Lex uses while a bot is waiting for customer input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#wait_and_continue_specification LexBot#wait_and_continue_specification}
	WaitAndContinueSpecification *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecification `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

