// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentintegration


type SecurityagentIntegrationProviderInputGitlab struct {
	// The GitLab access token used to authenticate. This can be a personal access token or a group access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#access_token SecurityagentIntegration#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The identifier of the GitLab group. Required when TokenType is GROUP and ignored for personal tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#group_id SecurityagentIntegration#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The HTTPS URL of a self-managed GitLab instance. Omit this value for GitLab SaaS (gitlab.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#target_url SecurityagentIntegration#target_url}
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// The type of GitLab access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#token_type SecurityagentIntegration#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

