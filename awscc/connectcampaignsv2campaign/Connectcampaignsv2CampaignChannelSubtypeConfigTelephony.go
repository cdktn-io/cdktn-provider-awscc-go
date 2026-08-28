// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfigTelephony struct {
	// Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connectcampaignsv2_campaign#capacity Connectcampaignsv2Campaign#capacity}
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The queue for the call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connectcampaignsv2_campaign#connect_queue_id Connectcampaignsv2Campaign#connect_queue_id}
	ConnectQueueId *string `field:"optional" json:"connectQueueId" yaml:"connectQueueId"`
	// Default Telephone Outbound config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connectcampaignsv2_campaign#default_outbound_config Connectcampaignsv2Campaign#default_outbound_config}
	DefaultOutboundConfig *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig `field:"optional" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// Telephony Outbound Mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connectcampaignsv2_campaign#outbound_mode Connectcampaignsv2Campaign#outbound_mode}
	OutboundMode *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundMode `field:"optional" json:"outboundMode" yaml:"outboundMode"`
}

