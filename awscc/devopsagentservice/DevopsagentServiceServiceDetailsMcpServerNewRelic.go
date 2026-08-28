// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerNewRelic struct {
	// New Relic authorization configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#authorization_config DevopsagentService#authorization_config}
	AuthorizationConfig *DevopsagentServiceServiceDetailsMcpServerNewRelicAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

