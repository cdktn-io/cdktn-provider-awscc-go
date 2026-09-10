// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerNewRelicAuthorizationConfig struct {
	// New Relic API key configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#api_key DevopsagentService#api_key}
	ApiKey *DevopsagentServiceServiceDetailsMcpServerNewRelicAuthorizationConfigApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
}

