// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasBotAliasLocaleSettings struct {
	// You can use this parameter to specify a specific Lambda function to run different functions in different locales.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#bot_alias_locale_setting LexBotAlias#bot_alias_locale_setting}
	BotAliasLocaleSetting *LexBotAliasBotAliasLocaleSettingsBotAliasLocaleSetting `field:"optional" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// A string used to identify the locale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lex_bot_alias#locale_id LexBotAlias#locale_id}
	LocaleId *string `field:"optional" json:"localeId" yaml:"localeId"`
}

