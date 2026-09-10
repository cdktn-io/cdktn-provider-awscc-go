// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace


type GrafanaWorkspaceSamlConfigurationAssertionAttributes struct {
	// Name of the attribute within the SAML assert to use as the users email in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/grafana_workspace#email GrafanaWorkspace#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Name of the attribute within the SAML assert to use as the users groups in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/grafana_workspace#groups GrafanaWorkspace#groups}
	Groups *string `field:"optional" json:"groups" yaml:"groups"`
	// Name of the attribute within the SAML assert to use as the users login handle in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/grafana_workspace#login GrafanaWorkspace#login}
	Login *string `field:"optional" json:"login" yaml:"login"`
	// Name of the attribute within the SAML assert to use as the users name in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/grafana_workspace#name GrafanaWorkspace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Name of the attribute within the SAML assert to use as the users organizations in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/grafana_workspace#org GrafanaWorkspace#org}
	Org *string `field:"optional" json:"org" yaml:"org"`
	// Name of the attribute within the SAML assert to use as the users roles in Grafana.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/grafana_workspace#role GrafanaWorkspace#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

