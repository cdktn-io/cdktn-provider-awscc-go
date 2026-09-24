// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace


type GrafanaWorkspaceSamlConfigurationRoleValues struct {
	// List of SAML roles which will be mapped into the Grafana Admin role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/grafana_workspace#admin GrafanaWorkspace#admin}
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// List of SAML roles which will be mapped into the Grafana Editor role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/grafana_workspace#editor GrafanaWorkspace#editor}
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

