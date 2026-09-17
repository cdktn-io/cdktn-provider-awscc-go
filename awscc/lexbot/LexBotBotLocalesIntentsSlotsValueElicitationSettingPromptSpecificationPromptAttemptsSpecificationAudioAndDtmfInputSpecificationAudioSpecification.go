// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationAudioSpecification struct {
	// Time for which a bot waits after the customer stops speaking to assume the utterance is finished.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#end_timeout_ms LexBot#end_timeout_ms}
	EndTimeoutMs *float64 `field:"optional" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Time for how long Amazon Lex waits before speech input is truncated and the speech is returned to application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lex_bot#max_length_ms LexBot#max_length_ms}
	MaxLengthMs *float64 `field:"optional" json:"maxLengthMs" yaml:"maxLengthMs"`
}

