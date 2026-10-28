// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentintegration


type SecurityagentIntegrationProviderInputGithub struct {
	// The OAuth authorization code received from GitHub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#code SecurityagentIntegration#code}
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The installation identifier provided by GitHub Enterprise Server on the install callback.
	//
	// Required for GitHub Enterprise Server integrations and ignored for GitHub.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#installation_id SecurityagentIntegration#installation_id}
	InstallationId *string `field:"optional" json:"installationId" yaml:"installationId"`
	// The name of the GitHub organization to integrate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#organization_name SecurityagentIntegration#organization_name}
	OrganizationName *string `field:"optional" json:"organizationName" yaml:"organizationName"`
	// The CSRF state token for validating the OAuth flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#state SecurityagentIntegration#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The HTTPS URL of a self-hosted GitHub Enterprise Server instance. Omit this value for GitHub.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#target_url SecurityagentIntegration#target_url}
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}

