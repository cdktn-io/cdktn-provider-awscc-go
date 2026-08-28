// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfiguration struct {
	// AWS association for 'monitor' account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#aws DevopsagentAssociation#aws}
	Aws *DevopsagentAssociationConfigurationAws `field:"optional" json:"aws" yaml:"aws"`
	// Azure subscription integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#azure DevopsagentAssociation#azure}
	Azure *DevopsagentAssociationConfigurationAzure `field:"optional" json:"azure" yaml:"azure"`
	// Dynatrace monitoring configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#dynatrace DevopsagentAssociation#dynatrace}
	Dynatrace *DevopsagentAssociationConfigurationDynatrace `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// EventChannelconfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#event_channel DevopsagentAssociation#event_channel}
	EventChannel *DevopsagentAssociationConfigurationEventChannel `field:"optional" json:"eventChannel" yaml:"eventChannel"`
	// GitHub repository integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#git_hub DevopsagentAssociation#git_hub}
	GitHub *DevopsagentAssociationConfigurationGitHub `field:"optional" json:"gitHub" yaml:"gitHub"`
	// GitLab project integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#git_lab DevopsagentAssociation#git_lab}
	GitLab *DevopsagentAssociationConfigurationGitLab `field:"optional" json:"gitLab" yaml:"gitLab"`
	// MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#mcp_server DevopsagentAssociation#mcp_server}
	McpServer *DevopsagentAssociationConfigurationMcpServer `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// Datadog MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#mcp_server_datadog DevopsagentAssociation#mcp_server_datadog}
	McpServerDatadog *DevopsagentAssociationConfigurationMcpServerDatadog `field:"optional" json:"mcpServerDatadog" yaml:"mcpServerDatadog"`
	// Grafana MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#mcp_server_grafana DevopsagentAssociation#mcp_server_grafana}
	McpServerGrafana *DevopsagentAssociationConfigurationMcpServerGrafana `field:"optional" json:"mcpServerGrafana" yaml:"mcpServerGrafana"`
	// NewRelic MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#mcp_server_new_relic DevopsagentAssociation#mcp_server_new_relic}
	McpServerNewRelic *DevopsagentAssociationConfigurationMcpServerNewRelic `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// SigV4-authenticated MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#mcp_server_sig_v4 DevopsagentAssociation#mcp_server_sig_v4}
	McpServerSigV4 *DevopsagentAssociationConfigurationMcpServerSigV4 `field:"optional" json:"mcpServerSigV4" yaml:"mcpServerSigV4"`
	// Splunk MCP server configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#mcp_server_splunk DevopsagentAssociation#mcp_server_splunk}
	McpServerSplunk *DevopsagentAssociationConfigurationMcpServerSplunk `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// PagerDuty integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#pager_duty DevopsagentAssociation#pager_duty}
	PagerDuty *DevopsagentAssociationConfigurationPagerDuty `field:"optional" json:"pagerDuty" yaml:"pagerDuty"`
	// ServiceNow integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#service_now DevopsagentAssociation#service_now}
	ServiceNow *DevopsagentAssociationConfigurationServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Slack workspace integration configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#slack DevopsagentAssociation#slack}
	Slack *DevopsagentAssociationConfigurationSlack `field:"optional" json:"slack" yaml:"slack"`
	// AWS association for 'source' account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsagent_association#source_aws DevopsagentAssociation#source_aws}
	SourceAws *DevopsagentAssociationConfigurationSourceAws `field:"optional" json:"sourceAws" yaml:"sourceAws"`
}

