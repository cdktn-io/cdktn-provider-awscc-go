// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsIntentClosingSetting struct {
	// A list of message groups that Amazon Lex uses to respond the user input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot#closing_response LexBot#closing_response}
	ClosingResponse *LexBotBotLocalesIntentsIntentClosingSettingClosingResponse `field:"optional" json:"closingResponse" yaml:"closingResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot#is_active LexBot#is_active}.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

