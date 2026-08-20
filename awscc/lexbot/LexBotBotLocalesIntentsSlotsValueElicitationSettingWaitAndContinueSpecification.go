// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecification struct {
	// The response that Amazon Lex sends to indicate that the bot is ready to continue the conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot#continue_response LexBot#continue_response}
	ContinueResponse *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationContinueResponse `field:"optional" json:"continueResponse" yaml:"continueResponse"`
	// Specifies whether the bot will wait for a user to respond.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot#is_active LexBot#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The response that Amazon Lex sends periodically to the user to indicate that the bot is still waiting for input from the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot#still_waiting_response LexBot#still_waiting_response}
	StillWaitingResponse *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponse `field:"optional" json:"stillWaitingResponse" yaml:"stillWaitingResponse"`
	// The response that Amazon Lex sends to indicate that the bot is waiting for the conversation to continue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot#waiting_response LexBot#waiting_response}
	WaitingResponse *LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationWaitingResponse `field:"optional" json:"waitingResponse" yaml:"waitingResponse"`
}

