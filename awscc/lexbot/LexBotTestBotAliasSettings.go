// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotTestBotAliasSettings struct {
	// A list of bot alias locale settings to add to the bot alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#bot_alias_locale_settings LexBot#bot_alias_locale_settings}
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// Contains information about code hooks that Amazon Lex calls during a conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#conversation_log_settings LexBot#conversation_log_settings}
	ConversationLogSettings *LexBotTestBotAliasSettingsConversationLogSettings `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// A description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#description LexBot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#sentiment_analysis_settings LexBot#sentiment_analysis_settings}
	SentimentAnalysisSettings *LexBotTestBotAliasSettingsSentimentAnalysisSettings `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}

