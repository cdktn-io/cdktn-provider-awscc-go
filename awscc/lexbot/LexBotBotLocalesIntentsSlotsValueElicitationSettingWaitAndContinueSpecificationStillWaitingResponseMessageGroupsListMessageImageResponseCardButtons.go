// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroupsListMessageImageResponseCardButtons struct {
	// The text that appears on the button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#text LexBot#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#value LexBot#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

