// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationLimitsOverrideAllChannelsSubtypesCommunicationLimitListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#frequency Connectcampaignsv2Campaign#frequency}.
	Frequency *float64 `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#max_count_per_recipient Connectcampaignsv2Campaign#max_count_per_recipient}.
	MaxCountPerRecipient *float64 `field:"optional" json:"maxCountPerRecipient" yaml:"maxCountPerRecipient"`
	// The communication limit time unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#unit Connectcampaignsv2Campaign#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

