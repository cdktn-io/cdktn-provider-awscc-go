// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_harness#agent_core_runtime_environment BedrockagentcoreHarness#agent_core_runtime_environment}.
	AgentCoreRuntimeEnvironment *BedrockagentcoreHarnessEnvironmentAgentCoreRuntimeEnvironment `field:"optional" json:"agentCoreRuntimeEnvironment" yaml:"agentCoreRuntimeEnvironment"`
}

