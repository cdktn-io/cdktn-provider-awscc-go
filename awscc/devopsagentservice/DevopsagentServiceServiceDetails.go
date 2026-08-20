// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetails struct {
	// Azure Identity service configuration for federated identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#azure_identity DevopsagentService#azure_identity}
	AzureIdentity *DevopsagentServiceServiceDetailsAzureIdentity `field:"optional" json:"azureIdentity" yaml:"azureIdentity"`
	// Dynatrace service configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#dynatrace DevopsagentService#dynatrace}
	Dynatrace *DevopsagentServiceServiceDetailsDynatrace `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// GitLab service configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#git_lab DevopsagentService#git_lab}
	GitLab *DevopsagentServiceServiceDetailsGitLab `field:"optional" json:"gitLab" yaml:"gitLab"`
	// MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#mcp_server DevopsagentService#mcp_server}
	McpServer *DevopsagentServiceServiceDetailsMcpServer `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// Grafana MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#mcp_server_grafana DevopsagentService#mcp_server_grafana}
	McpServerGrafana *DevopsagentServiceServiceDetailsMcpServerGrafana `field:"optional" json:"mcpServerGrafana" yaml:"mcpServerGrafana"`
	// New Relic service configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#mcp_server_new_relic DevopsagentService#mcp_server_new_relic}
	McpServerNewRelic *DevopsagentServiceServiceDetailsMcpServerNewRelic `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// SigV4-authenticated MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#mcp_server_sig_v4 DevopsagentService#mcp_server_sig_v4}
	McpServerSigV4 *DevopsagentServiceServiceDetailsMcpServerSigV4 `field:"optional" json:"mcpServerSigV4" yaml:"mcpServerSigV4"`
	// Splunk MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#mcp_server_splunk DevopsagentService#mcp_server_splunk}
	McpServerSplunk *DevopsagentServiceServiceDetailsMcpServerSplunk `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// PagerDuty service configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#pager_duty DevopsagentService#pager_duty}
	PagerDuty *DevopsagentServiceServiceDetailsPagerDuty `field:"optional" json:"pagerDuty" yaml:"pagerDuty"`
	// ServiceNow service configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsagent_service#service_now DevopsagentService#service_now}
	ServiceNow *DevopsagentServiceServiceDetailsServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}

