// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreconsentportal


type BedrockagentcoreConsentPortalIdpConfig struct {
	// The ARN of an AWS::BedrockAgentCore::OAuth2CredentialProvider resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#credential_provider_arn BedrockagentcoreConsentPortal#credential_provider_arn}
	CredentialProviderArn *string `field:"required" json:"credentialProviderArn" yaml:"credentialProviderArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#scopes BedrockagentcoreConsentPortal#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/bedrockagentcore_consent_portal#audience BedrockagentcoreConsentPortal#audience}.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

