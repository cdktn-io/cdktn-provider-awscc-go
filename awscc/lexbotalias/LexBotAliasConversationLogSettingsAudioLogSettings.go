// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasConversationLogSettingsAudioLogSettings struct {
	// The location of audio log files collected when conversation logging is enabled for a bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#destination LexBotAlias#destination}
	Destination *LexBotAliasConversationLogSettingsAudioLogSettingsDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#enabled LexBotAlias#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

