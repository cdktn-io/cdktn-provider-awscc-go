// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerSplunk struct {
	// MCP server splunk authorization configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#authorization_config DevopsagentService#authorization_config}
	AuthorizationConfig *DevopsagentServiceServiceDetailsMcpServerSplunkAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Optional description for the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#description DevopsagentService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// MCP server endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#endpoint DevopsagentService#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// MCP server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_service#name DevopsagentService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

