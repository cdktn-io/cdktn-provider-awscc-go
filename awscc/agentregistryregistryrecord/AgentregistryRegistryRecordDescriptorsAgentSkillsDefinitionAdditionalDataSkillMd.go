// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsAgentSkillsDefinitionAdditionalDataSkillMd struct {
	// Descriptor payload data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/agentregistry_registry_record#data AgentregistryRegistryRecord#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the descriptor type schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/agentregistry_registry_record#data_schema_version AgentregistryRegistryRecord#data_schema_version}
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
	// Source configuration for a SkillMd document. Unlike MCP/A2A sources, SkillMd does not support credential providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/agentregistry_registry_record#source AgentregistryRegistryRecord#source}
	Source *AgentregistryRegistryRecordDescriptorsAgentSkillsDefinitionAdditionalDataSkillMdSource `field:"optional" json:"source" yaml:"source"`
}

