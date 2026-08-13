// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasBotAliasLocaleSettingsBotAliasLocaleSetting struct {
	// Contains information about code hooks that Amazon Lex calls during a conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#code_hook_specification LexBotAlias#code_hook_specification}
	CodeHookSpecification *LexBotAliasBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecification `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
	// Whether the Lambda code hook is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#enabled LexBotAlias#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

