// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationMcpServerSigV4 struct {
	// List of MCP tools available for the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#tools DevopsagentAssociation#tools}
	Tools *[]*string `field:"optional" json:"tools" yaml:"tools"`
}

