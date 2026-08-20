// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockenforcedguardrailconfiguration


type BedrockEnforcedGuardrailConfigurationModelEnforcement struct {
	// Models to exclude from enforcement. If a model is in both lists, it is excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_enforced_guardrail_configuration#excluded_models BedrockEnforcedGuardrailConfiguration#excluded_models}
	ExcludedModels *[]*string `field:"optional" json:"excludedModels" yaml:"excludedModels"`
	// Models to enforce the guardrail on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrock_enforced_guardrail_configuration#included_models BedrockEnforcedGuardrailConfiguration#included_models}
	IncludedModels *[]*string `field:"optional" json:"includedModels" yaml:"includedModels"`
}

