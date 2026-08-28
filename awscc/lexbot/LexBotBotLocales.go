// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocales struct {
	// A custom vocabulary is a list of specific phrases that you want Amazon Lex V2 to recognize in the audio input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#custom_vocabulary LexBot#custom_vocabulary}
	CustomVocabulary *LexBotBotLocalesCustomVocabulary `field:"optional" json:"customVocabulary" yaml:"customVocabulary"`
	// A description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#description LexBot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of intents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#intents LexBot#intents}
	Intents interface{} `field:"optional" json:"intents" yaml:"intents"`
	// The identifier of the language and locale that the bot will be used in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#locale_id LexBot#locale_id}
	LocaleId *string `field:"optional" json:"localeId" yaml:"localeId"`
	// The specified confidence threshold for inserting the AMAZON.FallbackIntent and AMAZON.KendraSearchIntent intents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#nlu_confidence_threshold LexBot#nlu_confidence_threshold}
	NluConfidenceThreshold *float64 `field:"optional" json:"nluConfidenceThreshold" yaml:"nluConfidenceThreshold"`
	// List of SlotTypes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#slot_types LexBot#slot_types}
	SlotTypes interface{} `field:"optional" json:"slotTypes" yaml:"slotTypes"`
	// Settings for using an Amazon Polly voice to communicate with a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#voice_settings LexBot#voice_settings}
	VoiceSettings *LexBotBotLocalesVoiceSettings `field:"optional" json:"voiceSettings" yaml:"voiceSettings"`
}

