// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsAguiSource struct {
	// URL-based source configuration for a source-only descriptor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/agentregistry_registry_record#from_url AgentregistryRegistryRecord#from_url}
	FromUrl *AgentregistryRegistryRecordDescriptorsAguiSourceFromUrl `field:"optional" json:"fromUrl" yaml:"fromUrl"`
}

