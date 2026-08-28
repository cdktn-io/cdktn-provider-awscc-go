// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigTelephony struct {
	// Open Hours config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connectcampaignsv2_campaign#open_hours Connectcampaignsv2Campaign#open_hours}
	OpenHours *Connectcampaignsv2CampaignCommunicationTimeConfigTelephonyOpenHours `field:"optional" json:"openHours" yaml:"openHours"`
	// Restricted period config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connectcampaignsv2_campaign#restricted_periods Connectcampaignsv2Campaign#restricted_periods}
	RestrictedPeriods *Connectcampaignsv2CampaignCommunicationTimeConfigTelephonyRestrictedPeriods `field:"optional" json:"restrictedPeriods" yaml:"restrictedPeriods"`
}

