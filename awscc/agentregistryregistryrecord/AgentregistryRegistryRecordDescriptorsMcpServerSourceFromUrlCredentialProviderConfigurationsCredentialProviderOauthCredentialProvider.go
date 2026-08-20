// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord


type AgentregistryRegistryRecordDescriptorsMcpServerSourceFromUrlCredentialProviderConfigurationsCredentialProviderOauthCredentialProvider struct {
	// Additional custom parameters for the OAuth flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#custom_parameters AgentregistryRegistryRecord#custom_parameters}
	CustomParameters *map[string]*string `field:"optional" json:"customParameters" yaml:"customParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#grant_type AgentregistryRegistryRecord#grant_type}.
	GrantType *string `field:"optional" json:"grantType" yaml:"grantType"`
	// The ARN of the OAuth credential provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#provider_arn AgentregistryRegistryRecord#provider_arn}
	ProviderArn *string `field:"optional" json:"providerArn" yaml:"providerArn"`
	// OAuth scopes to request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/agentregistry_registry_record#scopes AgentregistryRegistryRecord#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

