// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationSlack struct {
	// Optional bidirectional communication configuration.
	//
	// Supply this configuration and set Enabled to true so you can interact with the agent directly from Slack
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#bidirectional DevopsagentAssociation#bidirectional}
	Bidirectional *DevopsagentAssociationConfigurationSlackBidirectional `field:"optional" json:"bidirectional" yaml:"bidirectional"`
	// Transmission targets for agent notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#transmission_target DevopsagentAssociation#transmission_target}
	TransmissionTarget *DevopsagentAssociationConfigurationSlackTransmissionTarget `field:"optional" json:"transmissionTarget" yaml:"transmissionTarget"`
	// Associated Slack workspace ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#workspace_id DevopsagentAssociation#workspace_id}
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
	// Associated Slack workspace name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#workspace_name DevopsagentAssociation#workspace_name}
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

