// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignSchedule struct {
	// Timestamp with no UTC offset or timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#end_time Connectcampaignsv2Campaign#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Time duration in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#refresh_frequency Connectcampaignsv2Campaign#refresh_frequency}
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
	// Timestamp with no UTC offset or timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#start_time Connectcampaignsv2Campaign#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

