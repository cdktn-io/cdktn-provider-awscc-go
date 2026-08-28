// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationOrchestrationAiAgentConfigurationToolConfigurationsInstruction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#examples WisdomAiAgent#examples}.
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wisdom_ai_agent#instruction WisdomAiAgent#instruction}.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
}

