// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfigWhatsApp struct {
	// Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#capacity Connectcampaignsv2Campaign#capacity}
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Default WhatsApp outbound config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#default_outbound_config Connectcampaignsv2Campaign#default_outbound_config}
	DefaultOutboundConfig *Connectcampaignsv2CampaignChannelSubtypeConfigWhatsAppDefaultOutboundConfig `field:"optional" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// WhatsApp Outbound Mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#outbound_mode Connectcampaignsv2Campaign#outbound_mode}
	OutboundMode *Connectcampaignsv2CampaignChannelSubtypeConfigWhatsAppOutboundMode `field:"optional" json:"outboundMode" yaml:"outboundMode"`
}

