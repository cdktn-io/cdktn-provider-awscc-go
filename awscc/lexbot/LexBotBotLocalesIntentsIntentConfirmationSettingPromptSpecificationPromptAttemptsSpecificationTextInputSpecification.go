// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecification struct {
	// Time for which a bot waits before re-prompting a customer for text input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lex_bot#start_timeout_ms LexBot#start_timeout_ms}
	StartTimeoutMs *float64 `field:"optional" json:"startTimeoutMs" yaml:"startTimeoutMs"`
}

