// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification struct {
	// Specifies the audio input specifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#audio_specification LexBot#audio_specification}
	AudioSpecification *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationAudioSpecification `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// Specifies the settings on DTMF input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#dtmf_specification LexBot#dtmf_specification}
	DtmfSpecification *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecification `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
	// Time for which a bot waits before assuming that the customer isn't going to speak or press a key.
	//
	// This timeout is shared between Audio and DTMF inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot#start_timeout_ms LexBot#start_timeout_ms}
	StartTimeoutMs *float64 `field:"optional" json:"startTimeoutMs" yaml:"startTimeoutMs"`
}

