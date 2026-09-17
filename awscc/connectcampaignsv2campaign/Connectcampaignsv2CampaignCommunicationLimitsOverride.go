// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationLimitsOverride struct {
	// Communication limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#all_channels_subtypes Connectcampaignsv2Campaign#all_channels_subtypes}
	AllChannelsSubtypes *Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypes `field:"optional" json:"allChannelsSubtypes" yaml:"allChannelsSubtypes"`
	// Enumeration of Instance Limits handling in a Campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#instance_limits_handling Connectcampaignsv2Campaign#instance_limits_handling}
	InstanceLimitsHandling *string `field:"optional" json:"instanceLimitsHandling" yaml:"instanceLimitsHandling"`
}

