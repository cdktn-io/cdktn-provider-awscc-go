// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes struct {
	// Indicates whether audio input is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#allow_audio_input LexBot#allow_audio_input}
	AllowAudioInput interface{} `field:"optional" json:"allowAudioInput" yaml:"allowAudioInput"`
	// Indicates whether DTMF input is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#allow_dtmf_input LexBot#allow_dtmf_input}
	AllowDtmfInput interface{} `field:"optional" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}

