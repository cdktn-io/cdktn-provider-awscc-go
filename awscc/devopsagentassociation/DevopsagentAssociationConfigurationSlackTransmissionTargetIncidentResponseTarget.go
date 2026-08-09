// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationSlackTransmissionTargetIncidentResponseTarget struct {
	// Slack channel ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_association#channel_id DevopsagentAssociation#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// Slack channel name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_association#channel_name DevopsagentAssociation#channel_name}
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
}

