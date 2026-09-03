// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOverrideInputValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/wisdom_ai_agent#json_path WisdomAiAgent#json_path}.
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/wisdom_ai_agent#value WisdomAiAgent#value}.
	Value *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOverrideInputValuesValue `field:"optional" json:"value" yaml:"value"`
}

