// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasConversationLogSettingsTextLogSettings struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot_alias#destination LexBotAlias#destination}
	Destination *LexBotAliasConversationLogSettingsTextLogSettingsDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lex_bot_alias#enabled LexBotAlias#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

