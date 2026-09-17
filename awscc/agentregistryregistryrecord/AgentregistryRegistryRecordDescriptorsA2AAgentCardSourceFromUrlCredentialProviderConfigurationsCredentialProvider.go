// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProvider struct {
	// IAM credential provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#iam_credential_provider AgentregistryRegistryRecord#iam_credential_provider}
	IamCredentialProvider *AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderIamCredentialProvider `field:"optional" json:"iamCredentialProvider" yaml:"iamCredentialProvider"`
	// OAuth credential provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#oauth_credential_provider AgentregistryRegistryRecord#oauth_credential_provider}
	OauthCredentialProvider *AgentregistryRegistryRecordDescriptorsA2AAgentCardSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProvider `field:"optional" json:"oauthCredentialProvider" yaml:"oauthCredentialProvider"`
}

