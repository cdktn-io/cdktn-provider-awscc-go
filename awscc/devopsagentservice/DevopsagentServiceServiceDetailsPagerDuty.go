// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsPagerDuty struct {
	// PagerDuty OAuth authorization configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_service#authorization_config DevopsagentService#authorization_config}
	AuthorizationConfig *DevopsagentServiceServiceDetailsPagerDutyAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// PagerDuty scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_service#scopes DevopsagentService#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

