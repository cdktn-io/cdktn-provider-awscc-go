// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#json_path WisdomAiAgent#json_path}.
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#output_configuration WisdomAiAgent#output_configuration}.
	OutputConfiguration *WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsOutputFiltersOutputConfiguration `field:"optional" json:"outputConfiguration" yaml:"outputConfiguration"`
}

