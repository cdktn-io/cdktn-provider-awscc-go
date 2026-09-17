// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigWhatsAppOpenHoursDailyHoursValue struct {
	// Time in ISO 8601 format, e.g. T23:11.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#end_time Connectcampaignsv2Campaign#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Time in ISO 8601 format, e.g. T23:11.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#start_time Connectcampaignsv2Campaign#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

