// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigSmsRestrictedPeriodsRestrictedPeriodListStruct struct {
	// Date in ISO 8601 format, e.g. 2024-01-01.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#end_date Connectcampaignsv2Campaign#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The name of a restricted period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#name Connectcampaignsv2Campaign#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Date in ISO 8601 format, e.g. 2024-01-01.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#start_date Connectcampaignsv2Campaign#start_date}
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
}

