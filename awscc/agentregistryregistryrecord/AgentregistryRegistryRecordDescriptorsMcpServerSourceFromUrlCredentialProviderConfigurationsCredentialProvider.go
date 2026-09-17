// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsMcpServerSourceFromUrlCredentialProviderConfigurationsCredentialProvider struct {
	// IAM credential provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#iam_credential_provider AgentregistryRegistryRecord#iam_credential_provider}
	IamCredentialProvider *AgentregistryRegistryRecordDescriptorsMcpServerSourceFromUrlCredentialProviderConfigurationsCredentialProviderIamCredentialProvider `field:"optional" json:"iamCredentialProvider" yaml:"iamCredentialProvider"`
	// OAuth credential provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/agentregistry_registry_record#oauth_credential_provider AgentregistryRegistryRecord#oauth_credential_provider}
	OauthCredentialProvider *AgentregistryRegistryRecordDescriptorsMcpServerSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProvider `field:"optional" json:"oauthCredentialProvider" yaml:"oauthCredentialProvider"`
}

