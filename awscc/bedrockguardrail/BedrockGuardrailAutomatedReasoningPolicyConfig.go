// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailAutomatedReasoningPolicyConfig struct {
	// The confidence threshold for triggering guardrail actions based on Automated Reasoning policy violations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#confidence_threshold BedrockGuardrail#confidence_threshold}
	ConfidenceThreshold *float64 `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// The list of Automated Reasoning policy ARNs to include in the guardrail configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrock_guardrail#policies BedrockGuardrail#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

