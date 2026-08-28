// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListVariations struct {
	// A message in a custom format defined by the client application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#custom_payload LexBot#custom_payload}
	CustomPayload *LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListVariationsCustomPayload `field:"optional" json:"customPayload" yaml:"customPayload"`
	// A message that defines a response card that the client application can show to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#image_response_card LexBot#image_response_card}
	ImageResponseCard *LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListVariationsImageResponseCard `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// A message in plain text format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#plain_text_message LexBot#plain_text_message}
	PlainTextMessage *LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListVariationsPlainTextMessage `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// A message in Speech Synthesis Markup Language (SSML).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#ssml_message LexBot#ssml_message}
	SsmlMessage *LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponseMessageGroupsListVariationsSsmlMessage `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

