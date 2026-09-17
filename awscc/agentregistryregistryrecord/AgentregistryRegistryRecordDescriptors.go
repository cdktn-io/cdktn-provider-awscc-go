// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptors struct {
	// The A2A agent card descriptor, populated when the record type is AGENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#a2_a_agent_card AgentregistryRegistryRecord#a2_a_agent_card}
	A2AAgentCard *AgentregistryRegistryRecordDescriptorsA2AAgentCard `field:"optional" json:"a2AAgentCard" yaml:"a2AAgentCard"`
	// The agent skills definition descriptor, populated when the record type is SKILL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#agent_skills_definition AgentregistryRegistryRecord#agent_skills_definition}
	AgentSkillsDefinition *AgentregistryRegistryRecordDescriptorsAgentSkillsDefinition `field:"optional" json:"agentSkillsDefinition" yaml:"agentSkillsDefinition"`
	// The AG-UI (Agent-User Interaction) descriptor, populated for records detected from an AG-UI protocol source.
	//
	// This descriptor is source-only: its content is synchronized from the configured source URL rather than supplied inline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#agui AgentregistryRegistryRecord#agui}
	Agui *AgentregistryRegistryRecordDescriptorsAgui `field:"optional" json:"agui" yaml:"agui"`
	// The custom descriptor, populated when the record type is CUSTOM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#custom AgentregistryRegistryRecord#custom}
	Custom *AgentregistryRegistryRecordDescriptorsCustom `field:"optional" json:"custom" yaml:"custom"`
	// The HTTP descriptor, populated for records detected from an HTTP protocol source.
	//
	// This descriptor is source-only: its content is synchronized from the configured source URL rather than supplied inline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#http AgentregistryRegistryRecord#http}
	Http *AgentregistryRegistryRecordDescriptorsHttp `field:"optional" json:"http" yaml:"http"`
	// The MCP server descriptor, populated when the record type is MCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#mcp_server AgentregistryRegistryRecord#mcp_server}
	McpServer *AgentregistryRegistryRecordDescriptorsMcpServer `field:"optional" json:"mcpServer" yaml:"mcpServer"`
}

