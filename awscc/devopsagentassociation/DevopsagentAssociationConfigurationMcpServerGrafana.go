// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationMcpServerGrafana struct {
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_association#enable_webhook_updates DevopsagentAssociation#enable_webhook_updates}
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// MCP server endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_association#endpoint DevopsagentAssociation#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// List of tool categories to enable for the Grafana MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_association#tools DevopsagentAssociation#tools}
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

