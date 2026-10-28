// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsGitHubAppVerification struct {
	// Permission variant of the DevOps Agent GitHub App that must already be installed on the organization or account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_service#app_variant DevopsagentService#app_variant}
	AppVariant *string `field:"optional" json:"appVariant" yaml:"appVariant"`
	// GitHub organization to scope the service to.
	//
	// The token must be an active member of the organization and only that organization's repositories are accessible. Omit to register the token's own user account
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_service#organization_name DevopsagentService#organization_name}
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// GitHub instance URL.
	//
	// Omit for github.com. Provide the instance host for GitHub Enterprise Server or GitHub Enterprise Cloud with data residency (e.g. https://github.example.com)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_service#target_url DevopsagentService#target_url}
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
	// GitHub personal access token (PAT) value. Accepts classic PATs (ghp_...), fine-grained PATs (github_pat_...), and GitHub Enterprise Server token formats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_service#token_value DevopsagentService#token_value}
	TokenValue *string `field:"optional" json:"tokenValue" yaml:"tokenValue"`
}

