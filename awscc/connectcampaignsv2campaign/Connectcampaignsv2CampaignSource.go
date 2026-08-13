// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignSource struct {
	// Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connectcampaignsv2_campaign#customer_profiles_segment_arn Connectcampaignsv2Campaign#customer_profiles_segment_arn}
	CustomerProfilesSegmentArn *string `field:"optional" json:"customerProfilesSegmentArn" yaml:"customerProfilesSegmentArn"`
	// The event trigger of the campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connectcampaignsv2_campaign#event_trigger Connectcampaignsv2Campaign#event_trigger}
	EventTrigger *Connectcampaignsv2CampaignSourceEventTrigger `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
}

