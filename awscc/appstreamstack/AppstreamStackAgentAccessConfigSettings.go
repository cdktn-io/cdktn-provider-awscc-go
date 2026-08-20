// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackAgentAccessConfigSettings struct {
	// The agent action to configure.
	//
	// Valid values are COMPUTER_VISION, COMPUTER_INPUT, and FORWARD_MCP_TOOLS. COMPUTER_VISION allows agents to take screenshots of the desktop. COMPUTER_INPUT allows agents to click, type, and scroll on the desktop and requires COMPUTER_VISION to also be enabled. FORWARD_MCP_TOOLS allows agents to interact with applications and the desktop operating system through direct MCP calls rather than using computer use tools. Forwards MCP tools configured on the WorkSpaces application session to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_stack#agent_action AppstreamStack#agent_action}
	AgentAction *string `field:"optional" json:"agentAction" yaml:"agentAction"`
	// Whether the agent action is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_stack#permission AppstreamStack#permission}
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

