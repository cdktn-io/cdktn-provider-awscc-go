// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsAgentSkillsDefinition struct {
	// Additional data associated with an agent skills definition descriptor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/agentregistry_registry_record#additional_data AgentregistryRegistryRecord#additional_data}
	AdditionalData *AgentregistryRegistryRecordDescriptorsAgentSkillsDefinitionAdditionalData `field:"optional" json:"additionalData" yaml:"additionalData"`
	// Descriptor payload data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/agentregistry_registry_record#data AgentregistryRegistryRecord#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Version of the descriptor type schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/agentregistry_registry_record#data_schema_version AgentregistryRegistryRecord#data_schema_version}
	DataSchemaVersion *string `field:"optional" json:"dataSchemaVersion" yaml:"dataSchemaVersion"`
}

