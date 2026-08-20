// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsServiceNow struct {
	// ServiceNow OAuth authorization configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#authorization_config DevopsagentService#authorization_config}
	AuthorizationConfig *DevopsagentServiceServiceDetailsServiceNowAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// ServiceNow instance URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#instance_url DevopsagentService#instance_url}
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
}

