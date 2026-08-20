// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign


type ConnectcampaignsCampaignDialerConfigProgressiveDialerConfig struct {
	// The bandwidth allocation of a queue resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connectcampaigns_campaign#bandwidth_allocation ConnectcampaignsCampaign#bandwidth_allocation}
	BandwidthAllocation *float64 `field:"optional" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
	// Allocates dialing capacity for this campaign between multiple active campaigns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connectcampaigns_campaign#dialing_capacity ConnectcampaignsCampaign#dialing_capacity}
	DialingCapacity *float64 `field:"optional" json:"dialingCapacity" yaml:"dialingCapacity"`
}

