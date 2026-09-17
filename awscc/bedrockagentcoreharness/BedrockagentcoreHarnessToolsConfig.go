// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreharness


type BedrockagentcoreHarnessToolsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#agent_core_browser BedrockagentcoreHarness#agent_core_browser}.
	AgentCoreBrowser *BedrockagentcoreHarnessToolsConfigAgentCoreBrowser `field:"optional" json:"agentCoreBrowser" yaml:"agentCoreBrowser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#agent_core_code_interpreter BedrockagentcoreHarness#agent_core_code_interpreter}.
	AgentCoreCodeInterpreter *BedrockagentcoreHarnessToolsConfigAgentCoreCodeInterpreter `field:"optional" json:"agentCoreCodeInterpreter" yaml:"agentCoreCodeInterpreter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#agent_core_gateway BedrockagentcoreHarness#agent_core_gateway}.
	AgentCoreGateway *BedrockagentcoreHarnessToolsConfigAgentCoreGateway `field:"optional" json:"agentCoreGateway" yaml:"agentCoreGateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#inline_function BedrockagentcoreHarness#inline_function}.
	InlineFunction *BedrockagentcoreHarnessToolsConfigInlineFunction `field:"optional" json:"inlineFunction" yaml:"inlineFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_harness#remote_mcp BedrockagentcoreHarness#remote_mcp}.
	RemoteMcp *BedrockagentcoreHarnessToolsConfigRemoteMcp `field:"optional" json:"remoteMcp" yaml:"remoteMcp"`
}

