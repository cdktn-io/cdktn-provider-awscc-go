// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptors struct {
	// The A2A agent card descriptor, populated when the record type is AGENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#a2_a_agent_card AgentregistryRegistryRecord#a2_a_agent_card}
	A2AAgentCard *AgentregistryRegistryRecordDescriptorsA2AAgentCard `field:"optional" json:"a2AAgentCard" yaml:"a2AAgentCard"`
	// The agent skills definition descriptor, populated when the record type is SKILL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#agent_skills_definition AgentregistryRegistryRecord#agent_skills_definition}
	AgentSkillsDefinition *AgentregistryRegistryRecordDescriptorsAgentSkillsDefinition `field:"optional" json:"agentSkillsDefinition" yaml:"agentSkillsDefinition"`
	// The custom descriptor, populated when the record type is CUSTOM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#custom AgentregistryRegistryRecord#custom}
	Custom *AgentregistryRegistryRecordDescriptorsCustom `field:"optional" json:"custom" yaml:"custom"`
	// The MCP server descriptor, populated when the record type is MCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#mcp_server AgentregistryRegistryRecord#mcp_server}
	McpServer *AgentregistryRegistryRecordDescriptorsMcpServer `field:"optional" json:"mcpServer" yaml:"mcpServer"`
}

