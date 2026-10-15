// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationMcpServerGrafana struct {
	// The mutative MCP tools enabled for this integration, by tool name.
	//
	// Only tools the service classifies as MUTATIVE may be listed; other names are rejected. Tool names can be found in the console or from the association's current configuration. The list fully replaces the stored value on update: removing a name disables that tool, and removing the property disables all elevated tools. Independent of Tools, which selects the toolsets available to the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_association#enabled_elevated_tools DevopsagentAssociation#enabled_elevated_tools}
	EnabledElevatedTools *[]*string `field:"optional" json:"enabledElevatedTools" yaml:"enabledElevatedTools"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_association#enable_webhook_updates DevopsagentAssociation#enable_webhook_updates}
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// MCP server endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_association#endpoint DevopsagentAssociation#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// List of tool categories to enable for the Grafana MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_association#tools DevopsagentAssociation#tools}
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

