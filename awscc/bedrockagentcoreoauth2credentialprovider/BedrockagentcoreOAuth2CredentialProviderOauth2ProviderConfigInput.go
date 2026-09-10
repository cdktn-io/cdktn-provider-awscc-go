// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInput struct {
	// Input configuration for an Atlassian OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#atlassian_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#atlassian_oauth_2_provider_config}
	AtlassianOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputAtlassianOauth2ProviderConfig `field:"optional" json:"atlassianOauth2ProviderConfig" yaml:"atlassianOauth2ProviderConfig"`
	// Input configuration for a custom OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#custom_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#custom_oauth_2_provider_config}
	CustomOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfig `field:"optional" json:"customOauth2ProviderConfig" yaml:"customOauth2ProviderConfig"`
	// Input configuration for a GitHub OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#github_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#github_oauth_2_provider_config}
	GithubOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGithubOauth2ProviderConfig `field:"optional" json:"githubOauth2ProviderConfig" yaml:"githubOauth2ProviderConfig"`
	// Input configuration for a Google OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#google_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#google_oauth_2_provider_config}
	GoogleOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputGoogleOauth2ProviderConfig `field:"optional" json:"googleOauth2ProviderConfig" yaml:"googleOauth2ProviderConfig"`
	// Input configuration for a supported non-custom OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#included_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#included_oauth_2_provider_config}
	IncludedOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputIncludedOauth2ProviderConfig `field:"optional" json:"includedOauth2ProviderConfig" yaml:"includedOauth2ProviderConfig"`
	// Input configuration for a LinkedIn OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#linkedin_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#linkedin_oauth_2_provider_config}
	LinkedinOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputLinkedinOauth2ProviderConfig `field:"optional" json:"linkedinOauth2ProviderConfig" yaml:"linkedinOauth2ProviderConfig"`
	// Input configuration for a Microsoft OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#microsoft_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#microsoft_oauth_2_provider_config}
	MicrosoftOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputMicrosoftOauth2ProviderConfig `field:"optional" json:"microsoftOauth2ProviderConfig" yaml:"microsoftOauth2ProviderConfig"`
	// Input configuration for a Salesforce OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#salesforce_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#salesforce_oauth_2_provider_config}
	SalesforceOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSalesforceOauth2ProviderConfig `field:"optional" json:"salesforceOauth2ProviderConfig" yaml:"salesforceOauth2ProviderConfig"`
	// Input configuration for a Slack OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#slack_oauth_2_provider_config BedrockagentcoreOAuth2CredentialProvider#slack_oauth_2_provider_config}
	SlackOauth2ProviderConfig *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputSlackOauth2ProviderConfig `field:"optional" json:"slackOauth2ProviderConfig" yaml:"slackOauth2ProviderConfig"`
}

