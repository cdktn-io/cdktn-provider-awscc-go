// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsMcpServerSourceFromUrl struct {
	// The credential providers used to authenticate when fetching descriptor content from the source URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry_record#credential_provider_configurations AgentregistryRegistryRecord#credential_provider_configurations}
	CredentialProviderConfigurations interface{} `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// URL source for descriptor content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/agentregistry_registry_record#url AgentregistryRegistryRecord#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

