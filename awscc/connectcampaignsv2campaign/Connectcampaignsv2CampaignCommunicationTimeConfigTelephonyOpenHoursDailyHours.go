// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigTelephonyOpenHoursDailyHours struct {
	// Day of week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connectcampaignsv2_campaign#key Connectcampaignsv2Campaign#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// List of time range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connectcampaignsv2_campaign#value Connectcampaignsv2Campaign#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

