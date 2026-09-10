// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsA2AAgentCardSource struct {
	// URL-based descriptor source configuration, with credential provider configurations for authenticated URL retrieval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/agentregistry_registry_record#from_url AgentregistryRegistryRecord#from_url}
	FromUrl *AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrl `field:"optional" json:"fromUrl" yaml:"fromUrl"`
}

