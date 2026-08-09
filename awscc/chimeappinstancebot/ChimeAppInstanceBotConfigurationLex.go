// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimeappinstancebot


type ChimeAppInstanceBotConfigurationLex struct {
	// The ARN of the Amazon Lex V2 bot's alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chime_app_instance_bot#lex_bot_alias_arn ChimeAppInstanceBot#lex_bot_alias_arn}
	LexBotAliasArn *string `field:"required" json:"lexBotAliasArn" yaml:"lexBotAliasArn"`
	// Identifies the Amazon Lex V2 bot's language and locale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chime_app_instance_bot#locale_id ChimeAppInstanceBot#locale_id}
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
	// Specifies the type of message that triggers a bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chime_app_instance_bot#invoked_by ChimeAppInstanceBot#invoked_by}
	InvokedBy *ChimeAppInstanceBotConfigurationLexInvokedBy `field:"optional" json:"invokedBy" yaml:"invokedBy"`
	// Determines whether the Amazon Lex V2 bot responds to all standard messages. Control messages are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chime_app_instance_bot#responds_to ChimeAppInstanceBot#responds_to}
	RespondsTo *string `field:"optional" json:"respondsTo" yaml:"respondsTo"`
	// The name of the welcome intent configured in the Amazon Lex V2 bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chime_app_instance_bot#welcome_intent ChimeAppInstanceBot#welcome_intent}
	WelcomeIntent *string `field:"optional" json:"welcomeIntent" yaml:"welcomeIntent"`
}

