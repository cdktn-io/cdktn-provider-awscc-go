// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsHttp struct {
	// Source configuration for a source-only descriptor. Unlike mcpServer/a2aAgentCard sources, source-only descriptors do not support credential providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/agentregistry_registry_record#source AgentregistryRegistryRecord#source}
	Source *AgentregistryRegistryRecordDescriptorsHttpSource `field:"optional" json:"source" yaml:"source"`
}

