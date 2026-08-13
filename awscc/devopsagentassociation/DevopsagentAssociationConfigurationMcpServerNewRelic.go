// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationMcpServerNewRelic struct {
	// New Relic Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_association#account_id DevopsagentAssociation#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// MCP server endpoint URL (e.g., https://mcp.newrelic.com/mcp/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsagent_association#endpoint DevopsagentAssociation#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

