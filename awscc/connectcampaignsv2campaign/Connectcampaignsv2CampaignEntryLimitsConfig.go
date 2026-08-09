// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignEntryLimitsConfig struct {
	// Maximum number of entries per participant. 0 indicates unlimited entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#max_entry_count Connectcampaignsv2Campaign#max_entry_count}
	MaxEntryCount *float64 `field:"optional" json:"maxEntryCount" yaml:"maxEntryCount"`
	// Minimum time interval between entries for the same participant in ISO 8601 duration format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#min_entry_interval Connectcampaignsv2Campaign#min_entry_interval}
	MinEntryInterval *string `field:"optional" json:"minEntryInterval" yaml:"minEntryInterval"`
}

