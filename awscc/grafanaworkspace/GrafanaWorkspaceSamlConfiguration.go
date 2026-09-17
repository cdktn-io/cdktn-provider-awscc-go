// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace


type GrafanaWorkspaceSamlConfiguration struct {
	// List of SAML organizations allowed to access Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/grafana_workspace#allowed_organizations GrafanaWorkspace#allowed_organizations}
	AllowedOrganizations *[]*string `field:"optional" json:"allowedOrganizations" yaml:"allowedOrganizations"`
	// Maps Grafana friendly names to the IdPs SAML attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/grafana_workspace#assertion_attributes GrafanaWorkspace#assertion_attributes}
	AssertionAttributes *GrafanaWorkspaceSamlConfigurationAssertionAttributes `field:"optional" json:"assertionAttributes" yaml:"assertionAttributes"`
	// IdP Metadata used to configure SAML authentication in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/grafana_workspace#idp_metadata GrafanaWorkspace#idp_metadata}
	IdpMetadata *GrafanaWorkspaceSamlConfigurationIdpMetadata `field:"optional" json:"idpMetadata" yaml:"idpMetadata"`
	// The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/grafana_workspace#login_validity_duration GrafanaWorkspace#login_validity_duration}
	LoginValidityDuration *float64 `field:"optional" json:"loginValidityDuration" yaml:"loginValidityDuration"`
	// Maps SAML roles to the Grafana Editor and Admin roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/grafana_workspace#role_values GrafanaWorkspace#role_values}
	RoleValues *GrafanaWorkspaceSamlConfigurationRoleValues `field:"optional" json:"roleValues" yaml:"roleValues"`
}

