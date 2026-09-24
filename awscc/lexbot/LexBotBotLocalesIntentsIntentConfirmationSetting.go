// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSetting struct {
	// A list of message groups that Amazon Lex uses to respond the user input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#declination_response LexBot#declination_response}
	DeclinationResponse *LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponse `field:"optional" json:"declinationResponse" yaml:"declinationResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#is_active LexBot#is_active}.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// Prompts the user to confirm the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot#prompt_specification LexBot#prompt_specification}
	PromptSpecification *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
}

