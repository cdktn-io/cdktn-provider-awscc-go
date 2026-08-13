// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfigEmailDefaultOutboundConfig struct {
	// Email address used for Email messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connectcampaignsv2_campaign#connect_source_email_address Connectcampaignsv2Campaign#connect_source_email_address}
	ConnectSourceEmailAddress *string `field:"optional" json:"connectSourceEmailAddress" yaml:"connectSourceEmailAddress"`
	// The name of the source email address display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connectcampaignsv2_campaign#source_email_address_display_name Connectcampaignsv2Campaign#source_email_address_display_name}
	SourceEmailAddressDisplayName *string `field:"optional" json:"sourceEmailAddressDisplayName" yaml:"sourceEmailAddressDisplayName"`
	// Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connectcampaignsv2_campaign#wisdom_template_arn Connectcampaignsv2Campaign#wisdom_template_arn}
	WisdomTemplateArn *string `field:"optional" json:"wisdomTemplateArn" yaml:"wisdomTemplateArn"`
}

