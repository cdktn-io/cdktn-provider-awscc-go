// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerGrafanaAuthorizationConfig struct {
	// Bearer token authentication details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_service#bearer_token DevopsagentService#bearer_token}
	BearerToken *DevopsagentServiceServiceDetailsMcpServerGrafanaAuthorizationConfigBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

