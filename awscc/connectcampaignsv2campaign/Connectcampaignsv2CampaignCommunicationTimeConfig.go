// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfig struct {
	// Time window config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#email Connectcampaignsv2Campaign#email}
	Email *Connectcampaignsv2CampaignCommunicationTimeConfigEmail `field:"optional" json:"email" yaml:"email"`
	// Local time zone config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#local_time_zone_config Connectcampaignsv2Campaign#local_time_zone_config}
	LocalTimeZoneConfig *Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfig `field:"optional" json:"localTimeZoneConfig" yaml:"localTimeZoneConfig"`
	// Time window config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#sms Connectcampaignsv2Campaign#sms}
	Sms *Connectcampaignsv2CampaignCommunicationTimeConfigSms `field:"optional" json:"sms" yaml:"sms"`
	// Time window config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#telephony Connectcampaignsv2Campaign#telephony}
	Telephony *Connectcampaignsv2CampaignCommunicationTimeConfigTelephony `field:"optional" json:"telephony" yaml:"telephony"`
	// Time window config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaignsv2_campaign#whats_app Connectcampaignsv2Campaign#whats_app}
	WhatsApp *Connectcampaignsv2CampaignCommunicationTimeConfigWhatsApp `field:"optional" json:"whatsApp" yaml:"whatsApp"`
}

