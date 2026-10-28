// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentintegration


type SecurityagentIntegrationProviderInputConfluence struct {
	// The OAuth 2.0 authorization code returned from the consent redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#code SecurityagentIntegration#code}
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The Atlassian installation identifier, available from the Atlassian administration console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#installation_id SecurityagentIntegration#installation_id}
	InstallationId *string `field:"optional" json:"installationId" yaml:"installationId"`
	// The Confluence Cloud site URL, for example https://mysite.atlassian.net.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#site_url SecurityagentIntegration#site_url}
	SiteUrl *string `field:"optional" json:"siteUrl" yaml:"siteUrl"`
	// The CSRF state token echoed back from the OAuth redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#state SecurityagentIntegration#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

