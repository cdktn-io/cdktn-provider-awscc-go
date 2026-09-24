// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotversion


type LexBotVersionBotVersionLocaleSpecification struct {
	// The version of a bot used for a bot locale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot_version#bot_version_locale_details LexBotVersion#bot_version_locale_details}
	BotVersionLocaleDetails *LexBotVersionBotVersionLocaleSpecificationBotVersionLocaleDetails `field:"required" json:"botVersionLocaleDetails" yaml:"botVersionLocaleDetails"`
	// The identifier of the language and locale that the bot will be used in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lex_bot_version#locale_id LexBotVersion#locale_id}
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

